// internal/monitor/monitor.go
package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	errorsInternal "loglineos/internal/errors"
	"loglineos/internal/executor"
	"loglineos/internal/parser"
	"loglineos/internal/spans"
	"loglineos/internal/utils" // Para WebhookPayload
	"os"
	"regexp"
	"time"

	"github.com/google/uuid"
)

// Monitor representa um monitor contínuo para um bloco 'observe'
type Monitor struct {
	Block          parser.Block
	ExecutionContext *executor.ExecutionContext
	CancelFunc     context.CancelFunc
	LastObservedContent string // Para `interval`
	PatternRegex   *regexp.Regexp
}

// NewMonitor cria um novo monitor para um bloco 'observe'
func NewMonitor(block parser.Block, execCtx *executor.ExecutionContext) (*Monitor, error) {
	if block.Type != "observe" {
		return nil, fmt.Errorf("monitor: apenas blocos 'observe' podem ser monitorados continuamente")
	}

	var spec parser.ObserveSpec
	if err := json.Unmarshal(block.Spec, &spec); err != nil {
		return nil, fmt.Errorf("monitor: erro ao parsear spec de observe para monitor: %w", err)
	}

	if spec.Interval == "" && spec.OnEvent == "" {
		return nil, fmt.Errorf("monitor: bloco observe '%s' não tem 'interval' nem 'on_event' para monitoramento contínuo", block.ID)
	}

	re, err := regexp.Compile(spec.Pattern)
	if err != nil {
		return nil, fmt.Errorf("monitor: padrão regex inválido para bloco '%s': %w", block.ID, err)
	}

	if execCtx.OutputDir == "" {
		execCtx.OutputDir = "./data"
	}
	if execCtx.SpanWriter == nil {
		spanWriter, err := spans.NewJSONLWriter(execCtx.OutputDir + "/spans.jsonl")
		if err != nil {
			return nil, fmt.Errorf("monitor: falha ao criar span writer para monitor: %w", err)
		}
		execCtx.SpanWriter = spanWriter
	}


	return &Monitor{
		Block:          block,
		ExecutionContext: execCtx,
		PatternRegex:   re,
	}, nil
}

// Start inicia o monitoramento contínuo (polling) ou espera eventos
func (m *Monitor) Start(ctx context.Context) {
	var spec parser.ObserveSpec
	json.Unmarshal(m.Block.Spec, &spec)

	if spec.Interval != "" {
		duration, err := time.ParseDuration(spec.Interval)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Monitor '%s': erro ao parsear intervalo '%s': %v. Monitor não será iniciado.\n", m.Block.ID, spec.Interval, err)
			return
		}

		ticker := time.NewTicker(duration)
		defer ticker.Stop()

		fmt.Printf("✅ Monitor '%s' (Tipo: %s) iniciado com intervalo: %v\n", m.Block.ID, m.Block.Type, duration)

		for {
			select {
			case <-ctx.Done():
				fmt.Printf("⚪ Monitor '%s' parado (contexto cancelado).\n", m.Block.ID)
				return
			case <-ticker.C:
				m.performObservation(spec, nil) // Chamada de polling
			}
		}
	} else if spec.OnEvent != "" {
		fmt.Printf("✅ Monitor '%s' (Tipo: %s) esperando eventos de '%s'.\n", m.Block.ID, m.Block.Type, spec.OnEvent)
		// Esta goroutine apenas espera pelo cancelamento, os eventos são tratados via HandleWebhookEvent
		<-ctx.Done()
		fmt.Printf("⚪ Monitor '%s' parado (contexto cancelado).\n", m.Block.ID)
		return
	}
}

// HandleWebhookEvent é chamado pelo WebhookServer quando um evento relevante é recebido
func (m *Monitor) HandleWebhookEvent(payload utils.WebhookPayload) {
	var spec parser.ObserveSpec
	json.Unmarshal(m.Block.Spec, &spec) // Erro já tratado em NewMonitor

	fmt.Printf("   [MONITOR %s] Recebido evento de webhook em '%s'...\n", m.Block.ID, payload.Path)

	// Converte o corpo do webhook para string para a regex
	payloadContent := string(payload.Body)
	if payloadContent == "" && len(payload.Query) > 0 { // Se não tem corpo, pode ser via query params
		// Converte query params para string para a regex
		queryStr := ""
		for k, v := range payload.Query {
			queryStr += fmt.Sprintf("%s=%s&", k, strings.Join(v, ","))
		}
		payloadContent = queryStr
	}

	m.performObservation(spec, &payloadContent) // Passa o conteúdo do evento
}


// performObservation executa a lógica de observação para um tick/evento
func (m *Monitor) performObservation(spec parser.ObserveSpec, eventContent *string) {
	span := spans.Span{
		SpanID:         uuid.New().String(),
		TaskID:         m.Block.ID,
		Type:           m.Block.Type,
		Timestamp:      time.Now().UTC().Format(time.RFC3339),
		SourceWorkflow: m.ExecutionContext.WorkflowName,
		InvokedBy:      "logline_monitor",
		Context:        map[string]interface{}{"interval": spec.Interval, "on_event": spec.OnEvent},
	}

	var currentContent string
	if eventContent != nil { // Se veio de um evento (webhook)
		currentContent = *eventContent
		span.Context["event_type"] = spec.OnEvent
		span.Context["event_source"] = spec.ListenPath
		span.Output["raw_event_content"] = currentContent
	} else { // Se veio de polling (interval)
		currentContent = readSimulatedSource(spec.Source) // Lê conteúdo simulado
	}
	
	isMatched := m.PatternRegex.MatchString(currentContent)
	contentChanged := (currentContent != m.LastObservedContent)
	m.LastObservedContent = currentContent // Atualiza último conteúdo (para polling)

	span.Output["source"] = spec.Source
	span.Output["pattern"] = spec.Pattern
	span.Output["matched_pattern"] = isMatched
	span.Output["content_changed"] = contentChanged
	span.Output["current_content_hash"] = fmt.Sprintf("%x", HashString(currentContent))


	if isMatched {
		span.Status = "success"
		span.Message = fmt.Sprintf("Padrão '%s' ENCONTRADO em '%s'. Conteúdo mudou: %t", spec.Pattern, spec.Source, contentChanged)
		fmt.Printf("   %s [MONITOR %s] %s\n", getStatusEmoji(span.Status), m.Block.ID, span.Message)
		// TODO: Disparar action (alert, trigger, commit) - Isso é complexo e precisa de um executor de actions
		// Por enquanto, apenas o span é gerado. A lógica de trigger block é no executor, não aqui.
	} else {
		span.Status = "warning"
		span.Message = fmt.Sprintf("Padrão '%s' NÃO ENCONTRADO em '%s'. Conteúdo mudou: %t", spec.Pattern, spec.Source, contentChanged)
		fmt.Printf("   %s [MONITOR %s] %s\n", getStatusEmoji(span.Status), m.Block.ID, span.Message)
	}

	if err := m.ExecutionContext.SpanWriter.WriteSpan(span); err != nil {
		fmt.Fprintf(os.Stderr, "Monitor '%s': erro ao escrever span: %v\n", m.Block.ID, err)
	}
	if m.ExecutionContext.SpanStream != nil {
		m.ExecutionContext.SpanStream <- span
	}
}

// readSimulatedSource simula a leitura de um arquivo/fonte externa
var simulatedContentCounter int
func readSimulatedSource(source string) string {
	simulatedContentCounter++
	if simulatedContentCounter%3 == 0 {
		return fmt.Sprintf("Conteúdo simulado de %s com PADRÃO_ALVO (mudou em %s)", source, time.Now().Format("15:04:05"))
	}
	return fmt.Sprintf("Conteúdo simulado de %s sem padrao (stable %s)", source, time.Now().Format("15:04:05"))
}

// HashString retorna um hash simples de uma string
func HashString(s string) uint32 {
	h := uint32(0)
	for i := 0; i < len(s); i++ {
		h = (h << 5) - h + uint32(s[i])
	}
	return h
}

// getStatusEmoji (copiado do executor, para uso independente aqui)
func getStatusEmoji(status string) string {
	switch status {
	case "success":
		return "🟢"
	case "error":
		return "❌"
	case "warning":
		return "🟡"
	case "draft":
		return "📝"
	case "skipped":
		return "⚪"
	default:
		return "⚪"
	}
}
