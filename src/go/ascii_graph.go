// internal/utils/ascii_graph.go
package utils

import (
	"fmt"
	"loglineos/internal/parser"
	"strings"
)

// GenerateSimpleASCIIGraph gera uma representação pseudo-ASCII de um workflow.
// É uma representação linear com indentação e setas, não um grafo visual completo.
// Para um grafo DOT completo, use planner.BuildGraph.
func GenerateSimpleASCIIGraph(workflow *parser.Workflow) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Workflow: %s\n", workflow.Name))
	sb.WriteString(fmt.Sprintf("Descrição: %s\n", workflow.Description))
	sb.WriteString("----------------------------------------\n")

	var renderBlocks func(blocks []parser.Block, indent string)
	renderBlocks = func(blocks []parser.Block, indent string) {
		for i, block := range blocks {
			prefix := "├──"
			if i == len(blocks)-1 && len(block.Tasks) == 0 { // Último nó no nível atual
				prefix = "└──"
			}

			// Exibe o bloco principal
			sb.WriteString(fmt.Sprintf("%s%s %s (%s)\n", indent, prefix, block.ID, block.Type))

			// Adiciona anotações de when, retry, fallback para clareza
			detailsIndent := indent
			if i < len(blocks)-1 || len(block.Tasks) > 0 { // Se não for o último, usa linha de continuidade
				detailsIndent += "│   "
			} else {
				detailsIndent += "    "
			}

			if block.When != "" {
				sb.WriteString(fmt.Sprintf("%sWhen: %s\n", detailsIndent, block.When))
			}
			if block.Retry != nil {
				sb.WriteString(fmt.Sprintf("%sRetry: max_attempts=%d\n", detailsIndent, block.Retry.MaxAttempts))
			}
			if block.Fallback != nil {
				sb.WriteString(fmt.Sprintf("%sFallback: type=%s\n", detailsIndent, block.Fallback.Type))
			}
            if len(block.Hooks) > 0 {
                sb.WriteString(fmt.Sprintf("%sHooks: %s\n", detailsIndent, strings.Join(block.Hooks, ", ")))
            }


			// Se for um bloco 'loop', renderiza os blocos aninhados recursivamente
			if block.Type == "loop" && block.Tasks != nil && len(block.Tasks) > 0 {
				newIndent := indent
				if i < len(blocks)-1 {
					newIndent += "│   " // Linha vertical para continuidade
				} else {
					newIndent += "    " // Espaço para o último elemento
				}
				sb.WriteString(fmt.Sprintf("%s└── Tasks aninhadas:\n", indent))
				renderBlocks(block.Tasks, newIndent) // Chama recursivamente
			}
		}
	}

	renderBlocks(workflow.Tasks, "")
	sb.WriteString("----------------------------------------\n")
	return sb.String()
}
