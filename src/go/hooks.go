// internal/hooks/hooks.go
package hooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"loglineos/internal/config_loader"
	"loglineos/internal/spans" // Para SpanID e Writer
	"net/http"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid" // Para SpanID
)

// HookContext é o contexto passado para o template do hook e para o log
type HookContext struct {
	WorkflowName string                 `json:"workflow_name"`
	BlockID      string                 `json:"block_id"`
	Type         string                 `json:"type"`
	Status       string                 `json:"status"`
	SpanID       string                 `json:"span_id,omitempty"`
	Timestamp    string                 `json:"timestamp,omitempty"`
	DurationMs   int64                  `json:"duration_ms,omitempty"`
	Message      string                 `json:"message"`
	Error        string                 `json:"error,omitempty"` // Se for um hook de erro
	Output       map[string]interface{} `json:"output,omitempty"` // Output do span
	HookName     string                 `json:"hook_name"`       // Nome do hook que está sendo executado
}

// HooksManager gerencia a execução dos hooks
type HooksManager struct {
	HooksConfig config_loader.HooksConfig
	SpanWriter *spans.JSONLWriter // Para spans de hooks (injetado via context)
}

// NewHooksManager cria uma nova instância de HooksManager
func NewHooksManager(hooksConfig config_loader.HooksConfig, spanWriter *spans.JSONLWriter) *HooksManager {
	return &HooksManager{
		HooksConfig: hooksConfig,
		SpanWriter: spanWriter,
	}
}

// ExecuteHook executa um hook específico, se definido
func (hm *HooksManager) ExecuteHook(hookKey string, ctx *HookContext) {
	definitions, ok := hm.HooksConfig[hookKey]
	if !ok || len(definitions) == 0 {
		return // Nenhum hook definido para esta chave
	}

	fmt.Printf("   [HOOKS] Disparando hook '%s' para bloco '%s'...\n", hookKey, ctx.BlockID)

	for _, def := range definitions {
		hookSpan := spans.Span{
			SpanID:         uuid.New().String(),
			TaskID:         fmt.Sprintf("hook_%s_%s", ctx.BlockID, def.Name),
			Type:           "hook",
			Status:         "running",
			Timestamp:      time.Now().UTC().Format(time.RFC3339),
			SourceWorkflow: ctx.WorkflowName,
			InvokedBy:      "logline_hook_manager",
			Message:        fmt.Sprintf("Executando hook '%s' (tipo: %s)", def.Name, def.Type),
			Context:        map[string]interface{}{"hook_key": hookKey, "hook_name": def.Name, "hook_type": def.Type},
			Inputs:         map[string]interface{}{"block_context": ctx}, // Passa o contexto do bloco para o input do span
		}
        ctx.HookName = def.Name // Define o nome do hook no contexto para o template

		startTime := time.Now()
		var err error
		var hookOutput map[string]interface{}

		switch def.Type {
		case "shell":
			hookOutput, err = hm.executeShellHook(def, ctx)
		case "http":
			hookOutput, err = hm.executeHTTPHook(def, ctx)
		default:
			err = fmt.Errorf("tipo de hook '%s' não suportado", def.Type)
		}

		hookSpan.DurationMs = time.Since(startTime).Milliseconds()
		hookSpan.Output = hookOutput
		if err != nil {
			hookSpan.Status = "error"
			hookSpan.Message = fmt.Sprintf("Hook '%s' falhou: %v", def.Name, err)
			fmt.Printf("   ❌ [HOOK %s] Falha no hook '%s': %v\n", hookKey, def.Name, err)
		} else {
			hookSpan.Status = "success"
			hookSpan.Message = fmt.Sprintf("Hook '%s' executado com sucesso.", def.Name)
			fmt.Printf("   🟢 [HOOK %s] Sucesso no hook '%s'.\n", hookKey, def.Name)
		}
		
		if hm.SpanWriter != nil {
			hm.SpanWriter.WriteSpan(hookSpan)
		} else {
            fmt.Fprintf(os.Stderr, "Hooks: SpanWriter não inicializado para hook '%s'. Não será auditado.\n", def.Name)
        }
	}
}

// executeShellHook executa um comando shell definido no hook
func (hm *HooksManager) executeShellHook(def config_loader.HookDefinition, ctx *HookContext) (map[string]interface{}, error) {
	command, err := renderTemplate(def.Command, ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao renderizar comando do hook: %w", err)
	}

	var args []string
	for _, arg := range def.Args {
		renderedArg, err := renderTemplate(arg, ctx)
		if err != nil {
			return nil, fmt.Errorf("erro ao renderizar argumento '%s' do hook: %w", arg, err)
		}
		args = append(args, renderedArg)
	}

	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return map[string]interface{}{"stdout": string(output)}, fmt.Errorf("comando shell falhou: %w, output: %s", err, string(output))
	}
	return map[string]interface{}{"stdout": string(output)}, nil
}

// executeHTTPHook executa uma requisição HTTP definida no hook
func (hm *HooksManager) executeHTTPHook(def config_loader.HookDefinition, ctx *HookContext) (map[string]interface{}, error) {
	url, err := renderTemplate(def.URL, ctx)
	if err != nil {
		return nil, fmt.Errorf("erro ao renderizar URL do hook: %w", err)
	}
	method := def.Method
	if method == "" {
		method = "GET"
	}

	var bodyReader io.Reader
	if def.Body != "" {
		renderedBody, err := renderTemplate(def.Body, ctx)
		if err != nil {
			return nil, fmt.Errorf("erro ao renderizar corpo do hook: %w", err)
		}
		bodyReader = strings.NewReader(renderedBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar requisição HTTP: %w", err)
	}

	for k, v := range def.Headers {
		renderedValue, err := renderTemplate(v, ctx)
		if err != nil {
			return nil, fmt.Errorf("erro ao renderizar cabeçalho '%s' do hook: %w", k, err)
		}
		req.Header.Set(k, renderedValue)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("falha ao executar requisição HTTP: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("falha ao ler corpo da resposta HTTP: %w", err)
	}

	return map[string]interface{}{
		"status_code": resp.StatusCode,
		"response_body": string(respBody),
	}, nil
}

// renderTemplate processa uma string com dados do HookContext
func renderTemplate(text string, ctx *HookContext) (string, error) {
	// Usar Funções JSON para permitir acesso a campos de Output complexos
	funcMap := template.FuncMap{
		"json": func(v interface{}) (string, error) {
			b, err := json.Marshal(v)
			if err != nil {
				return "", err
			}
			return string(b), nil
		},
	}
	
	// Permite usar um template sem tags de ação se o texto não contiver '{{'
	// Isso evita erros de parsing para strings simples que não são templates.
	if !strings.Contains(text, "{{") && !strings.Contains(text, "}}") {
		return text, nil
	}

	tmpl, err := template.New("hook_template").Funcs(funcMap).Parse(text)
	if err != nil {
		return "", fmt.Errorf("erro ao parsear template: %w", err)
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, ctx)
	if err != nil {
		return "", fmt.Errorf("erro ao executar template: %w", err)
	}
	return buf.String(), nil
}
