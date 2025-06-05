// internal/spans/spans.go
package spans

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/Knetic/govaluate"
)

// Span representa um evento auditável no LogLineOS
type Span struct {
	SpanID       string                 `json:"span_id"`        // ID único do span
	TaskID       string                 `json:"task_id"`        // ID do bloco no workflow
	Type         string                 `json:"type"`           // Tipo do bloco (e.g., "register", "mechanical")
	Status       string                 `json:"status"`         // Status da execução (success, error, warning, draft, skipped)
	Timestamp    string                 `json:"timestamp"`      // Timestamp da execução (ISO 8601)
	DurationMs   int64                  `json:"duration_ms"`    // Duração em milissegundos
	SourceWorkflow string               `json:"source_workflow,omitempty"` // Nome do workflow de origem
	InvokedBy    string                 `json:"invoked_by,omitempty"` // Quem invocou (usuário, sistema)
	Attempt      int                    `json:"attempt,omitempty"` // Número da tentativa (para retries)
	IsFallback   bool                   `json:"is_fallback,omitempty"` // Indica se este span é de um fallback
	Inputs       map[string]interface{} `json:"inputs,omitempty"`       // Entradas ou dependências do bloco
	Output       map[string]interface{} `json:"output,omitempty"`       // Saída do bloco
	Context      map[string]interface{} `json:"context,omitempty"`      // Contexto adicional (versão, ambiente)
	Message      string                 `json:"message,omitempty"`      // Mensagem resumida para o span
}

// JSONLWriter para escrever spans em formato JSON Lines
type JSONLWriter struct {
	file *os.File
	mu   sync.Mutex // Protege a escrita concorrente
}

// NewJSONLWriter cria um novo escritor de spans JSONL
func NewJSONLWriter(filePath string) (*JSONLWriter, error) {
	// Usar O_APPEND|O_CREATE|O_WRONLY para adicionar ao final ou criar se não existe
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("spans: falha ao abrir arquivo '%s' para escrita: %w", filePath, err)
	}
	return &JSONLWriter{file: file}, nil
}

// WriteSpan serializa um Span para JSON e o escreve no arquivo, seguido por uma nova linha.
func (w *JSONLWriter) WriteSpan(span Span) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	data, err := json.Marshal(span)
	if err != nil {
		return fmt.Errorf("spans: falha ao serializar span '%s': %w", span.SpanID, err)
	}

	if _, err := w.file.Write(data); err != nil {
		return fmt.Errorf("spans: falha ao escrever dados do span '%s': %w", span.SpanID, err)
	}
	if _, err := w.file.WriteString("\n"); err != nil { // JSONL requires newline
		return fmt.Errorf("spans: falha ao escrever newline para span '%s': %w", span.SpanID, err)
	}
	return nil
}

// Close fecha o arquivo de spans.
func (w *JSONLWriter) Close() error {
	if w.file == nil {
		return nil // Já fechado ou nunca aberto
	}
	err := w.file.Close()
	if err != nil {
		return fmt.Errorf("spans: falha ao fechar arquivo de spans: %w", err)
	}
	w.file = nil // Evita fechamento duplo
	return nil
}

// ReadAllSpans lê todos os spans de um arquivo JSONL
func ReadAllSpans(filePath string) ([]Span, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Span{}, nil // Retorna slice vazio se o arquivo não existe, sem erro
		}
		return nil, fmt.Errorf("spans: falha ao abrir arquivo '%s' para leitura: %w", filePath, err)
	}
	defer file.Close()

	var spans []Span
	decoder := json.NewDecoder(file)

	for {
		var span Span
		if err := decoder.Decode(&span); err == io.EOF {
			break // Fim do arquivo
		} else if err != nil {
			return nil, fmt.Errorf("spans: falha ao decodificar linha em '%s': %w", filePath, err)
		}
		spans = append(spans, span)
	}
	return spans, nil
}

// FilterSpans filtra uma lista de spans com base em uma expressão.
// Ex: "status == 'error' && duration_ms > 100"
func FilterSpans(spans []Span, filterExpr string) ([]Span, error) {
	if filterExpr == "" {
		return spans, nil // Sem filtro, retorna todos os spans
	}

	expression, err := govaluate.NewEvaluableExpression(filterExpr)
	if err != nil {
		return nil, fmt.Errorf("spans: expressão de filtro '%s' inválida: %w", filterExpr, err)
	}

	filtered := []Span{}
	for _, span := range spans {
		// Converte o span para um map[string]interface{} para avaliação
		spanMap, err := convertSpanToMap(span)
		if err != nil {
			return nil, fmt.Errorf("spans: erro ao converter span '%s' para map para avaliação: %w", span.SpanID, err)
		}

		result, err := expression.Evaluate(spanMap)
		if err != nil {
			// Erros de avaliação podem ocorrer se um campo não existir na expressão.
			// Consideramos que o span não corresponde ao filtro neste caso.
			continue
		}

		if boolResult, ok := result.(bool); ok && boolResult {
			filtered = append(filtered, span)
		}
	}
	return filtered, nil
}

// convertSpanToMap converte uma struct Span para um map[string]interface{} para govaluate
func convertSpanToMap(span Span) (map[string]interface{}, error) {
	// A maneira mais fácil de fazer isso é serializar e desserializar
	// No futuro, considere usar reflexão para evitar o overhead de JSON
	data, err := json.Marshal(span)
	if err != nil {
		return nil, fmt.Errorf("spans: falha ao serializar span para map de avaliação: %w", err)
	}
	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		return nil, fmt.Errorf("spans: falha ao desserializar span para map de avaliação: %w", err)
	}
	return m, nil
}


// SortSpans ordena uma lista de spans por um campo e ordem (asc/desc)
func SortSpans(spans []Span, sortBy string, order string) ([]Span, error) {
	if sortBy == "" {
		return spans, nil
	}

	sortedSpans := make([]Span, len(spans))
	copy(sortedSpans, spans)

	isAsc := (order == "asc" || order == "")

	sort.Slice(sortedSpans, func(i, j int) bool {
		s1 := sortedSpans[i]
		s2 := sortedSpans[j]

		var less bool
		switch sortBy {
		case "timestamp":
			t1, err1 := time.Parse(time.RFC3339, s1.Timestamp)
			t2, err2 := time.Parse(time.RFC3339, s2.Timestamp)
			if err1 == nil && err2 == nil {
				less = t1.Before(t2)
			} else {
				// Fallback para comparação de string se parse falhar
				less = s1.Timestamp < s2.Timestamp
			}
		case "duration_ms":
			less = s1.DurationMs < s2.DurationMs
		case "type":
			less = s1.Type < s2.Type
		case "status":
			less = s1.Status < s2.Status
		case "task_id":
			less = s1.TaskID < s2.TaskID
		case "attempt":
			less = s1.Attempt < s2.Attempt
		case "is_fallback":
			// false (não é fallback) vem antes de true (é fallback) em ordem ascendente
			less = !s1.IsFallback && s2.IsFallback
		case "span_id":
			less = s1.SpanID < s2.SpanID
		default:
			// No caso de campos aninhados ou complexos, seria necessário mais lógica (reflexão, etc.)
			// Para simplicidade, vamos usar fmt.Sprintf para tentar comparar, o que não é ideal para todos os tipos.
			// Um sistema premium teria um `GetField(span, fieldPath string) (interface{}, error)` robusto.
			less = fmt.Sprintf("%v", s1) < fmt.Sprintf("%v", s2)
		}

		if !isAsc {
			return !less // Inverte para ordem descendente
		}
		return less
	})

	return sortedSpans, nil
}

// AuditMetrics guarda métricas agregadas dos spans
type AuditMetrics struct {
	TotalSpans    int                       `json:"total_spans"`
	SpansByStatus map[string]int            `json:"spans_by_status"`
	SpansByType   map[string]int            `json:"spans_by_type"`
	AvgDurationMsByType map[string]float64    `json:"avg_duration_ms_by_type"`
	TotalDurationMs int64                     `json:"total_duration_ms"`
}

// CollectMetrics coleta métricas básicas de uma lista de spans
func CollectMetrics(spans []Span) AuditMetrics {
	metrics := AuditMetrics{
		TotalSpans:    len(spans),
		SpansByStatus: make(map[string]int),
		SpansByType:   make(map[string]int),
		AvgDurationMsByType: make(map[string]float64),
		TotalDurationMs: 0,
	}

	typeDurations := make(map[string][]int64)

	for _, span := range spans {
		metrics.SpansByStatus[span.Status]++
		metrics.SpansByType[span.Type]++
		metrics.TotalDurationMs += span.DurationMs

		if span.Status == "success" {
			typeDurations[span.Type] = append(typeDurations[span.Type], span.DurationMs)
		}
	}

	for spanType, durations := range typeDurations {
		if len(durations) > 0 {
			var sum int64
			for _, d := range durations {
				sum += d
			}
			metrics.AvgDurationMsByType[spanType] = float64(sum) / float64(len(durations))
		}
	}

	return metrics
}
