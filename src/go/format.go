// internal/utils/format.go
package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"loglineos/internal/spans"
	"os"
	"text/tabwriter" // Para tabelas Markdown
)

// PrintSpansJSON imprime spans em formato JSON (indentado para leitura)
func PrintSpansJSON(spans []spans.Span, w io.Writer) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ") // Indenta para legibilidade
	return encoder.Encode(spans)
}

// PrintSpansMarkdown imprime spans em formato de tabela Markdown
func PrintSpansMarkdown(spans []spans.Span, w io.Writer) {
	// Definir as colunas que queremos exibir na tabela Markdown
	headers := []string{"SPAN_ID", "TASK_ID", "TYPE", "STATUS", "DURATION_MS", "ATTEMPT", "IS_FALLBACK", "TIMESTAMP", "MESSAGE"}
	
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', tabwriter.Debug) // Debug mostra os limites das colunas
	
	// Imprime cabeçalho
	for i, h := range headers {
		fmt.Fprint(tw, h)
		if i < len(headers)-1 {
			fmt.Fprint(tw, "\t")
		}
	}
	fmt.Fprintln(tw)

	// Imprime separador
	for i := range headers {
		for j := 0; j < len(headers[i]); j++ {
			fmt.Fprint(tw, "-")
		}
		if i < len(headers)-1 {
			fmt.Fprint(tw, "\t")
		}
	}
	fmt.Fprintln(tw)

	// Imprime dados
	for _, span := range spans {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%d\t%d\t%t\t%s\t%s\n",
			shortenID(span.SpanID),
			span.TaskID,
			span.Type,
			span.Status,
			span.DurationMs,
			span.Attempt,
			span.IsFallback,
			span.Timestamp,
			span.Message,
		)
	}
	tw.Flush() // Garante que tudo seja escrito
}

// PrintSpansCSV imprime spans em formato CSV
func PrintSpansCSV(spans []spans.Span, w io.Writer) error {
	writer := csv.NewWriter(w)

	// Escreve cabeçalho CSV
	headers := []string{"span_id", "task_id", "type", "status", "timestamp", "duration_ms", "source_workflow", "invoked_by", "attempt", "is_fallback", "message", "output_json", "context_json"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("falha ao escrever cabeçalho CSV: %w", err)
	}

	for _, span := range spans {
		outputJSON, _ := json.Marshal(span.Output) // Converte output para JSON string
		contextJSON, _ := json.Marshal(span.Context) // Converte context para JSON string
		row := []string{
			span.SpanID,
			span.TaskID,
			span.Type,
			span.Status,
			span.Timestamp,
			fmt.Sprintf("%d", span.DurationMs),
			span.SourceWorkflow,
			span.InvokedBy,
			fmt.Sprintf("%d", span.Attempt),
			fmt.Sprintf("%t", span.IsFallback),
			span.Message,
			string(outputJSON),
			string(contextJSON),
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("falha ao escrever linha CSV: %w", err)
		}
	}

	writer.Flush()
	return nil
}

// PrintAuditMetrics imprime as métricas coletadas
func PrintAuditMetrics(metrics spans.AuditMetrics, w io.Writer) {
	fmt.Fprintln(w, "\n--- Métricas de Auditoria ---")
	fmt.Fprintf(w, "Total de Spans: %d\n", metrics.TotalSpans)
	fmt.Fprintln(w, "Spans por Status:")
	for status, count := range metrics.SpansByStatus {
		fmt.Fprintf(w, "  - %s: %d\n", status, count)
	}
	fmt.Fprintln(w, "Spans por Tipo:")
	for typ, count := range metrics.SpansByType {
		fmt.Fprintf(w, "  - %s: %d\n", typ, count)
	}
	if len(metrics.AvgDurationMsByType) > 0 {
		fmt.Fprintln(w, "Duração Média (ms) por Tipo (Apenas Sucessos):")
		for typ, avg := range metrics.AvgDurationMsByType {
			fmt.Fprintf(w, "  - %s: %.2fms\n", typ, avg)
		}
	}
	fmt.Fprintf(w, "Duração Total (ms) de Todos os Spans: %dms\n", metrics.TotalDurationMs)
	fmt.Fprintln(w, "----------------------------")
}

// shortenID para exibir IDs mais curtos na tabela Markdown
func shortenID(id string) string {
	if len(id) > 8 {
		return id[:8] + "..."
	}
	return id
}
