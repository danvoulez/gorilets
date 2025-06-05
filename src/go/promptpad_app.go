// internal/ui/promptpad_app.go
package ui

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"loglineos/cmd"
	errorsInternal "loglineos/internal/errors"
	"loglineos/internal/executor"
	"loglineos/internal/parser"
	"loglineos/internal/spans"
	"loglineos/internal/utils"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/rivo/tview"
)

// PromptPadApp contém a lógica da aplicação TUI
type PromptPadApp struct {
	app           *tview.Application
	pages         *tview.Pages
	inputField    *tview.InputField
	outputView    *tview.TextView
	workflowEditor *tview.TextArea
	graphView     *tview.TextView
	statusText    *tview.TextView
	
	currentWorkflowPath string
	loglineSchemaContent []byte
	spansOutputDir      string
	
	spanStream    chan spans.Span
}

// NewPromptPadApp cria e inicializa a aplicação PromptPad
func NewPromptPadApp(schemaContent []byte, spansOutputDir string) *PromptPadApp {
	app := tview.NewApplication()

	// Componentes da UI
	inputField := tview.NewInputField().
		SetLabel(">>> ").
		SetFieldWidth(0).
		SetBorder(true).
		SetTitle(" Comando ")

	outputView := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetWordWrap(true).
		SetMaxBuffer(1024 * 50)

	workflowEditor := tview.NewTextArea().
		SetPlaceholder("Cole seu workflow LogLine aqui ou use 'open <arquivo.logline>'").
		SetBorder(true).
		SetTitle(" Workflow Editor (F2 Salvar, F3 Validar, F4 Grafo) ")

	graphView := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetWordWrap(false).
		SetBorder(true).
		SetTitle(" Workflow Grafo ")

	statusText := tview.NewTextView().
		SetDynamicColors(true).
		SetText("Bem-vindo ao LogLineOS PromptPad!").SetTextAlign(tview.AlignCenter)

	// Layout principal
	mainLayout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(workflowEditor, 0, 3, false).
			AddItem(graphView, 0, 1, false),
			0, 4, false).
		AddItem(inputField, 1, 1, true).
		AddItem(statusText, 1, 1, false)

	pages := tview.NewPages().
		AddPage("main", mainLayout, true, true)

	p := &PromptPadApp{
		app:           app,
		pages:         pages,
		inputField:    inputField,
		outputView:    outputView,
		workflowEditor: workflowEditor,
		graphView:     graphView,
		statusText:    statusText,
		
		loglineSchemaContent: schemaContent,
		spansOutputDir:      spansOutputDir,
		
		spanStream:    make(chan spans.Span, 100),
	}

	inputField.SetDoneFunc(func(key tview.Key) {
		if key == tview.KeyEnter {
			command := strings.TrimSpace(inputField.GetText())
			inputField.SetText("")
			p.processCommand(command)
		}
	})

	app.SetInputCapture(func(event *tview.Event) *tview.Event {
		if p.app.GetFocus() == p.workflowEditor {
			switch event.Key() {
			case tview.KeyF2: // F2 para Salvar
				p.saveWorkflow()
				return nil
			case tview.KeyF3: // F3 para Validar
				p.validateCurrentWorkflow()
				return nil
			case tview.KeyF4: // F4 para Gerar Grafo
				p.generateWorkflowGraph()
				return nil
			}
		}
		return event
	})

	// Syntax highlighting e validação contínua (debounce)
	validationTimer := time.NewTimer(time.Millisecond * 200)
	validationTimer.Stop()
	go func() {
		for range validationTimer.C {
			p.syntaxHighlightAndValidate()
		}
	}()

	workflowEditor.SetChangedFunc(func() {
		validationTimer.Stop()
		validationTimer.Reset(time.Millisecond * 200)
	})

	go p.listenForSpans()

	return p
}

// Run inicia a aplicação TUI
func (p *PromptPadApp) Run() error {
	return p.app.SetRoot(p.pages, true).Run()
}

// processCommand interpreta e executa comandos do PromptPad
func (p *PromptPadApp) processCommand(command string) {
	p.outputView.Clear()
	p.statusText.SetText(fmt.Sprintf("[blue]Executando: '%s'...", command)).SetTextAlign(tview.AlignCenter)
	p.app.Draw()

	parts := strings.Fields(command)
	if len(parts) == 0 {
		p.statusText.SetText("[red]Comando vazio.").SetTextAlign(tview.AlignCenter)
		return
	}

	ctx := context.Background()

	switch parts[0] {
	case "open":
		if len(parts) < 2 {
			p.outputView.SetText("[red]Erro: 'open' requer um caminho de arquivo.")
			p.statusText.SetText("[red]Uso: open <caminho_do_arquivo.logline>").SetTextAlign(tview.AlignCenter)
			return
		}
		filePath := parts[1]
		p.openWorkflow(filePath)
	case "run":
		workflowPath := p.currentWorkflowPath
		if len(parts) >= 2 {
			workflowPath = parts[1]
		}
		if workflowPath == "" {
			p.outputView.SetText("[red]Erro: Nenhum workflow carregado para 'run'. Use 'open' ou 'run <caminho>'.")
			p.statusText.SetText("[red]Falha ao executar.").SetTextAlign(tview.AlignCenter)
			return
		}
		p.saveAndExecute(ctx, workflowPath, false)
	case "simulate":
		workflowPath := p.currentWorkflowPath
		if len(parts) >= 2 {
			workflowPath = parts[1]
		}
		if workflowPath == "" {
			p.outputView.SetText("[red]Erro: Nenhum workflow carregado para 'simulate'. Use 'open' ou 'simulate <caminho>'.")
			p.statusText.SetText("[red]Falha ao simular.").SetTextAlign(tview.AlignCenter)
			return
		}
		p.saveAndExecute(ctx, workflowPath, true)
	case "audit":
		args := parts[1:]
		var buf bytes.Buffer
		err := cmd.RunAuditCommand("", strings.Join(args, " "), "", "markdown", "", false, &buf)
		p.outputView.WriteString(buf.String())
		if err != nil {
			p.outputView.WriteString(fmt.Sprintf("[red]❌ Erro na auditoria: %v\n", err))
			p.statusText.SetText("[red]Auditoria falhou.").SetTextAlign(tview.AlignCenter)
		} else {
			p.statusText.SetText("[green]Auditoria concluída.").SetTextAlign(tview.AlignCenter)
		}
		p.outputView.ScrollToEnd()
	case "commit":
		if p.currentWorkflowPath == "" {
			p.outputView.SetText("[red]Erro: Nenhum workflow aberto para commit.")
			p.statusText.SetText("[red]Falha ao commit.").SetTextAlign(tview.AlignCenter)
			return
		}
		goal := "Manual commit via PromptPad"
		message := "Commit manual do estado atual."
		inputs := []string{}
		
		go func() {
			err := cmd.RunCommitCommand(goal, inputs, []string{"PromptPad_User"}, message, p.spansOutputDir, false)
			p.app.QueueUpdateDraw(func() {
				if err != nil {
					p.outputView.WriteString(fmt.Sprintf("[red]❌ Erro no commit: %v\n", err))
					p.statusText.SetText("[red]Commit falhou.").SetTextAlign(tview.AlignCenter)
				} else {
					p.outputView.WriteString("[green]✅ Commit realizado com sucesso!\n")
					p.statusText.SetText("[green]Commit concluído.").SetTextAlign(tview.AlignCenter)
				}
				p.outputView.ScrollToEnd()
			})
		}()
	case "revert":
		p.outputView.Clear()
		p.outputView.WriteString("[yellow]Comando 'revert' invocado.\n")
		p.outputView.WriteString("Para reverter para um estado anterior, você precisará usar o Git.\n")
		p.outputView.WriteString("1. Use 'logline audit' para encontrar o 'snapshot_id' do ponto que deseja reverter.\n")
		p.outputView.WriteString("2. Use o hash do commit Git associado ao snapshot (se disponível no metadata.json).\n")
		p.outputView.WriteString("3. No seu terminal, fora do PromptPad, execute:\n")
		p.outputView.WriteString("   [blue]git checkout <hash_do_commit_do_snapshot>\n")
		p.outputView.WriteString("   [blue]OU: git checkout $(logline audit --filter 'snapshot_id == \"<id_do_snapshot>\"' --format json | jq -r '.[0].git_hash')\n")
		p.outputView.WriteString("[yellow]AVISO: Reverter um repositório Git pode sobrescrever arquivos. Salve suas mudanças antes!\n")
		p.statusText.SetText("[yellow]Instruções de reversão exibidas.").SetTextAlign(tview.AlignCenter)
		p.outputView.ScrollToEnd()
	case "replay":
		go p.replaySpansInTUI()
	case "?": // <-- NOVO: AI Assistant
		p.handleAIQuery(strings.Join(parts[1:], " "))
	case "clear":
		p.outputView.Clear()
		p.statusText.SetText("Saída limpa.").SetTextAlign(tview.AlignCenter)
	case "exit":
		p.app.Stop()
	default:
		p.outputView.SetText(fmt.Sprintf("[red]Comando desconhecido: '%s'", parts[0]))
		p.statusText.SetText("[red]Comando inválido.").SetTextAlign(tview.AlignCenter)
	}
}

// saveAndExecute é uma função auxiliar para salvar e executar/simular
func (p *PromptPadApp) saveAndExecute(ctx context.Context, workflowPath string, dryRun bool) {
	if p.currentWorkflowPath != "" && p.workflowEditor.GetText() != "" {
		if err := ioutil.WriteFile(p.currentWorkflowPath, []byte(p.workflowEditor.GetText()), 0644); err != nil {
			p.outputView.WriteString(fmt.Sprintf("[red]Erro ao salvar antes de rodar: %v\n", err))
			p.statusText.SetText("[red]Falha ao salvar antes de rodar.").SetTextAlign(tview.AlignCenter)
			return
		}
	}

	p.outputView.Clear()
	action := "Executando"
	if dryRun { action = "Simulando" }
	p.outputView.WriteString(fmt.Sprintf("%s workflow: %s\n", action, workflowPath))
	p.statusText.SetText(fmt.Sprintf("[blue]%s '%s'...", action, workflowPath)).SetTextAlign(tview.AlignCenter)

	close(p.spanStream)
	p.spanStream = make(chan spans.Span, 100)
	go p.listenForSpans()

	workflowContent, err := ioutil.ReadFile(workflowPath)
	if err != nil {
		p.outputView.WriteString(fmt.Sprintf("[red]Erro ao ler arquivo: %v\n", err))
		p.statusText.SetText(fmt.Sprintf("[red]%s '%s' falhou.", action, workflowPath)).SetTextAlign(tview.AlignCenter)
		return
	}
	validationError := parser.ValidateContent(workflowPath, p.loglineSchemaContent, workflowContent)
	if validationError != nil {
		p.outputView.WriteString(fmt.Sprintf("[red]Workflow tem erros de validação:\n%s\n", validationError.Error()))
		p.statusText.SetText(fmt.Sprintf("[red]%s '%s' falhou devido a erros de validação.", action, workflowPath)).SetTextAlign(tview.AlignCenter)
		return
	}

	workflow, err := parser.ParseWorkflow(workflowContent)
	if err != nil {
		p.outputView.WriteString(fmt.Sprintf("[red]Erro ao parsear workflow: %v\n", err))
		p.statusText.SetText(fmt.Sprintf("[red]%s '%s' falhou.", action, workflowPath)).SetTextAlign(tview.AlignCenter)
		return
	}

	execContext := &executor.ExecutionContext{
		DryRun:     dryRun,
		SpanWriter: nil,
		SpanStream: p.spanStream,
		OutputDir:  p.spansOutputDir,
		EntityStore: make(map[string]executor.EntityRecord),
		WorkflowName: workflow.Name,
		HooksManager: cmd.GetHooksManager(), // Pega o HooksManager global
		LLMIntegrator: cmd.GetLLMIntegrator(), // Pega o LLMIntegrator global
	}

	// NOVO: Preenche o SpanWriter no HooksManager global para esta execução
	if execContext.HooksManager != nil {
		spanWriter, err := spans.NewJSONLWriter(execContext.OutputDir + "/spans.jsonl")
		if err != nil {
			p.outputView.WriteString(fmt.Sprintf("[red]Erro ao inicializar span writer para hooks: %v\n", err))
		} else {
			execContext.HooksManager.SpanWriter = spanWriter
			defer func() {
				if err := spanWriter.Close(); err != nil {
					fmt.Fprintf(os.Stderr, "Erro ao fechar SpanWriter de hooks: %v\n", err)
				}
			}()
		}
	}


	err = executor.ExecuteWorkflow(workflow, execContext)
	
	p.app.QueueUpdateDraw(func() {
		if err != nil {
			p.outputView.WriteString(fmt.Sprintf("[red]❌ Erro na %s: %v\n", strings.ToLower(action), err))
			p.statusText.SetText(fmt.Sprintf("[red]%s '%s' falhou.", action, workflowPath)).SetTextAlign(tview.AlignCenter)
		} else {
			p.outputView.WriteString(fmt.Sprintf("[green]✅ %s concluída.\n", action))
			p.statusText.SetText(fmt.Sprintf("[green]%s '%s' concluída com sucesso.", action, workflowPath)).SetTextAlign(tview.AlignCenter)
		}
		p.outputView.ScrollToEnd()
	})
}


// openWorkflow carrega e exibe o conteúdo de um workflow no editor, e o valida
func (p *PromptPadApp) openWorkflow(filePath string) {
	content, err := cmd.LoadFileContent(filePath)
	if err != nil {
		p.outputView.SetText(fmt.Sprintf("[red]Erro ao abrir workflow: %v", err))
		p.statusText.SetText("[red]Falha ao carregar workflow.").SetTextAlign(tview.AlignCenter)
		return
	}
	p.workflowEditor.SetText(string(content), false)
	p.currentWorkflowPath = filePath
	p.statusText.SetText(fmt.Sprintf("[green]Workflow '%s' carregado.", filePath)).SetTextAlign(tview.AlignCenter)
	p.syntaxHighlightAndValidate()
	p.generateWorkflowGraph()
}

// saveWorkflow salva o conteúdo atual do editor no arquivo
func (p *PromptPadApp) saveWorkflow() {
	if p.currentWorkflowPath == "" {
		p.statusText.SetText("[red]Nenhum workflow aberto para salvar.").SetTextAlign(tview.AlignCenter)
		return
	}
	content := []byte(p.workflowEditor.GetText())
	err := ioutil.WriteFile(p.currentWorkflowPath, content, 0644)
	if err != nil {
		p.statusText.SetText(fmt.Sprintf("[red]Erro ao salvar: %v", err)).SetTextAlign(tview.AlignCenter)
		return
	}
	p.statusText.SetText(fmt.Sprintf("[green]Workflow '%s' salvo com sucesso.", p.currentWorkflowPath)).SetTextAlign(tview.AlignCenter)
	p.syntaxHighlightAndValidate()
	p.generateWorkflowGraph()
}

// validateCurrentWorkflow valida o conteúdo atual do editor em tempo real
func (p *PromptPadApp) validateCurrentWorkflow() {
	if p.currentWorkflowPath == "" {
		p.statusText.SetText("[yellow]Nenhum workflow carregado para validação.").SetTextAlign(tview.AlignCenter)
		return
	}

	content := []byte(p.workflowEditor.GetText())
	validationError := parser.ValidateContent(p.currentWorkflowPath, p.loglineSchemaContent, content)

	p.app.QueueUpdateDraw(func() {
		p.outputView.Clear()
		if validationError != nil {
			p.outputView.SetTitle("[red] Erros de Validação! ")
			p.outputView.WriteString(fmt.Sprintf("[red]❌ %s\n", validationError.Error()))
			p.statusText.SetText(fmt.Sprintf("[red]Workflow inválido.").SetTextAlign(tview.AlignCenter))
		} else {
			p.outputView.SetTitle("[green] Validação ")
			p.outputView.WriteString("[green]✅ Workflow válido de acordo com o schema.\n")
			p.statusText.SetText("[green]Workflow válido.").SetTextAlign(tview.AlignCenter)
		}
		p.outputView.ScrollToBeginning()
	})
}


// syntaxHighlightAndValidate combina highlight de sintaxe e validação (refatorado para ser mais rápido)
func (p *PromptPadApp) syntaxHighlightAndValidate() {
	if p.currentWorkflowPath == "" {
		return
	}

	content := p.workflowEditor.GetText() // Obtém o texto UMA VEZ

	// Aplicar highlight
	highlightedText := p.applySyntaxHighlight(content)

	// Validar (agora usa o novo método de parser.ValidateContent)
	validationError := parser.ValidateContent(p.currentWorkflowPath, p.loglineSchemaContent, []byte(content))

	p.app.QueueUpdateDraw(func() {
		// Atualiza o texto do editor com o highlight ANTES de exibir erros
		p.workflowEditor.SetText(highlightedText, false)

		// Exibir erros/status da validação
		p.outputView.Clear()
		if validationError != nil {
			p.outputView.SetTitle("[red] Erros de Validação! ")
			p.outputView.WriteString(fmt.Sprintf("[red]❌ %s\n", validationError.Error()))
			p.statusText.SetText(fmt.Sprintf("[red]Workflow inválido.").SetTextAlign(tview.AlignCenter))
		} else {
			p.outputView.SetTitle("[green] Validação ")
			p.outputView.WriteString("[green]✅ Workflow válido de acordo com o schema.\n")
			p.statusText.SetText("[green]Workflow válido.").SetTextAlign(tview.AlignCenter)
		}
		p.outputView.ScrollToBeginning()
	})
}


// applySyntaxHighlight aplica um highlight de sintaxe básico ao texto do workflow
func (p *PromptPadApp) applySyntaxHighlight(text string) string {
	var sb strings.Builder
	lines := strings.Split(text, "\n")
	
	// Regexes para tipos e palavras-chave
	keywordRegex := regexp.MustCompile(`\b(name|description|tasks|id|type|spec|fallback|retry|when|inputs|command|args|entity|data|goal|from|to|note|pattern|source|action|message|signed_by|max_attempts|interval|on_event|count|until|timeout|llm|model|trust|input|listen_path|trigger_block_id|condition|map_output_to_input|hooks)\b:`)
	typeRegex := regexp.MustCompile(`type:\s*["']?(register|affair|observe|commit|mechanical|cognitive|loop|await|alert|trigger)["']?`)
	stringValRegex := regexp.MustCompile(`(".*?")`)
	numberValRegex := regexp.MustCompile(`\b(\d+(\.\d+)?)\b`)

	for _, line := range lines {
		// Aplicar cores em ordem, do mais geral para o mais específico
		// Primeiro, a cor padrão (branca ou do tema)
		coloredLine := "[white]" + line + "[white]" // Reset cor no final da linha

		// Tipos de blocos (amarelo)
		coloredLine = typeRegex.ReplaceAllString(coloredLine, "type: [yellow]$1[white]")

		// Palavras-chave (laranja)
		coloredLine = keywordRegex.ReplaceAllString(coloredLine, "[orange]$1[white]:")

		// Valores de string (verde)
		coloredLine = stringValRegex.ReplaceAllString(coloredLine, "[green]$1[white]")
		
		// Valores numéricos (azul claro) - pode conflitar com strings, ordem importa
		coloredLine = numberValRegex.ReplaceAllString(coloredLine, "[aqua]$1[white]")

		sb.WriteString(coloredLine)
		sb.WriteString("\n")
	}
	return sb.String()
}


// generateWorkflowGraph gera e exibe a representação ASCII do grafo no painel `graphView`
func (p *PromptPadApp) generateWorkflowGraph() {
	if p.currentWorkflowPath == "" {
		p.graphView.SetText("[yellow]Nenhum workflow aberto para gerar grafo.").SetTextAlign(tview.AlignCenter)
		return
	}

	content := []byte(p.workflowEditor.GetText())
	workflow, err := parser.ParseWorkflow(content)
	if err != nil {
		p.app.QueueUpdateDraw(func() {
			p.graphView.SetText(fmt.Sprintf("[red]Erro ao parsear workflow para grafo: %v", err)).SetTextAlign(tview.AlignCenter)
			p.statusText.SetText("[red]Erro ao gerar grafo.").SetTextAlign(tview.AlignCenter)
		})
		return
	}

	asciiGraph := utils.GenerateSimpleASCIIGraph(workflow)
	p.app.QueueUpdateDraw(func() {
		p.graphView.SetText(asciiGraph)
		p.graphView.ScrollToBeginning()
		p.graphView.SetTitle(fmt.Sprintf(" Workflow Grafo (%s) ", filepath.Base(p.currentWorkflowPath)))
		p.statusText.SetText("[green]Grafo do workflow gerado.").SetTextAlign(tview.AlignCenter)
	})
}


// replaySpansInTUI reproduz os spans em câmera lenta
func (p *PromptPadApp) replaySpansInTUI() {
	p.outputView.Clear()
	p.outputView.WriteString("[yellow]Iniciando Replay de Spans... (Pressione Ctrl+C para parar fora do PromptPad se for stuck)\n")
	p.statusText.SetText("[blue]Replay ativo.').").SetTextAlign(tview.AlignCenter)

	spanFilePath := filepath.Join(p.spansOutputDir, "spans.jsonl")
	spansToReplay, err := spans.ReadAllSpans(spanFilePath)
	if err != nil {
		p.outputView.WriteString(fmt.Sprintf("[red]Erro ao ler spans para replay: %v\n", err))
		p.statusText.SetText("[red]Replay falhou.").SetTextAlign(tview.AlignCenter)
		return
	}
	if len(spansToReplay) == 0 {
		p.outputView.WriteString("[yellow]Nenhum span encontrado para replay. Execute um workflow primeiro.\n")
		p.statusText.SetText("[yellow]Replay sem spans.").SetTextAlign(tview.AlignCenter)
		return
	}

	spansToReplay, _ = spans.SortSpans(spansToReplay, "timestamp", "asc")

	go func() {
		for i, span := range spansToReplay {
			p.app.QueueUpdateDraw(func() {
				statusEmoji := getStatusEmoji(span.Status)
				outputLine := fmt.Sprintf("[%d/%d] %s [%s] %s: %s (Duração: %dms",
					i+1, len(spansToReplay), statusEmoji, span.Status, span.TaskID, span.Message, span.DurationMs)
				if span.Attempt > 1 {
					outputLine += fmt.Sprintf(", Tentativa: %d", span.Attempt)
				}
				if span.IsFallback {
					outputLine += ", Fallback!"
				}
				outputLine += ")\n"
				p.outputView.WriteString(outputLine)
				if span.Status == "error" {
					p.outputView.WriteString(fmt.Sprintf("  Detalhes do Erro: %v\n", span.Output))
				}
				p.outputView.ScrollToEnd()
			})
			time.Sleep(500 * time.Millisecond) // Pausa para cada span
		}
		p.app.QueueUpdateDraw(func() {
			p.outputView.WriteString("[green]Replay concluído.\n")
			p.statusText.SetText("[green]Replay concluído.").SetTextAlign(tview.AlignCenter)
		})
	}()
}

// handleAIQuery processa uma pergunta para o assistente de IA
func (p *PromptPadApp) handleAIQuery(query string) {
	p.outputView.Clear()
	p.outputView.WriteString(fmt.Sprintf("[blue]🧠 Assistente IA: '%s'\n", query))
	p.statusText.SetText("[blue]Consultando assistente IA...").SetTextAlign(tview.AlignCenter)

	llmIntegrator := cmd.GetLLMIntegrator()
	if llmIntegrator == nil {
		p.outputView.WriteString("[red]Erro: LLM Integrator não está disponível. Verifique sua configuração de logline.yml.\n")
		p.statusText.SetText("[red]Assistente IA indisponível.").SetTextAlign(tview.AlignCenter)
		return
	}

	// Constrói um prompt baseado na query e no contexto atual (spans recentes, workflow)
	contextualPrompt := p.buildContextualPromptForAI(query)

	go func() {
		resp, err := llmIntegrator.Generate(context.Background(), contextualPrompt, "medium", "")
		p.app.QueueUpdateDraw(func() {
			if err != nil {
				p.outputView.WriteString(fmt.Sprintf("[red]❌ Erro ao consultar IA: %v\n", err))
				p.statusText.SetText("[red]Consulta IA falhou.").SetTextAlign(tview.AlignCenter)
			} else {
				p.outputView.WriteString(fmt.Sprintf("[green]✅ Resposta da IA:\n%s\n", resp.Output))
				p.statusView.SetText("[green]Resposta da IA recebida.").SetTextAlign(tview.AlignCenter)
			}
			p.outputView.ScrollToEnd()
		})
	}()
}

// buildContextualPromptForAI constrói um prompt com base na query e no contexto do PromptPad
func (p *PromptPadApp) buildContextualPromptForAI(query string) string {
	var sb strings.Builder
	sb.WriteString("Você é um assistente LogLineOS no PromptPad. Responda à pergunta do usuário. Inclua informações sobre o workflow atual e spans recentes, se relevantes.\n\n")
	sb.WriteString("PERGUNTA DO USUÁRIO: " + query + "\n\n")

	// Adicionar o workflow atual, se houver
	if p.currentWorkflowPath != "" {
		sb.WriteString("WORKFLOW ATUAL:\n")
		sb.WriteString("Caminho: " + p.currentWorkflowPath + "\n")
		sb.WriteString("Conteúdo:\n```yaml\n" + p.workflowEditor.GetText() + "\n```\n\n")
	}

	// Adicionar spans recentes (últimos 10, por exemplo)
	spanFilePath := filepath.Join(p.spansOutputDir, "spans.jsonl")
	recentSpans, err := spans.ReadAllSpans(spanFilePath)
	if err == nil && len(recentSpans) > 0 {
		sb.WriteString("SPANS RECENTES (Últimos 10):\n")
		if len(recentSpans) > 10 {
			recentSpans = recentSpans[len(recentSpans)-10:] // Pega os 10 últimos
		}
		for _, s := range recentSpans {
			spanJSON, _ := json.Marshal(s) // Converte span para JSON
			sb.WriteString(fmt.Sprintf("- %s\n", string(spanJSON)))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("Por favor, seja conciso e direto ao ponto. Use o contexto fornecido para sua resposta.")
	return sb.String()
}


// listenForSpans escuta o canal de spans e os adiciona ao outputView
func (p *PromptPadApp) listenForSpans() {
	for span := range p.spanStream {
		p.app.QueueUpdateDraw(func() {
			statusEmoji := getStatusEmoji(span.Status)
			outputLine := fmt.Sprintf("%s [%s] %s: %s (Duração: %dms",
				statusEmoji, span.Status, span.TaskID, span.Message, span.DurationMs)
			if span.Attempt > 1 {
				outputLine += fmt.Sprintf(", Tentativa: %d", span.Attempt)
			}
			if span.IsFallback {
				outputLine += ", Fallback!"
			}
			outputLine += ")\n"
			p.outputView.WriteString(outputLine)
			if span.Status == "error" {
				p.outputView.WriteString(fmt.Sprintf("  Detalhes do Erro: %v\n", span.Output))
			}
			p.outputView.ScrollToEnd()
		})
	}
}


// Funções utilitárias (repetidas do executor para este pacote, idealmente centralizadas)
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
