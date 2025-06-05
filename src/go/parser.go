// internal/parser/parser.go
package parser

import (
	"fmt"
	errorsInternal "loglineos/internal/errors"

	"github.com/go-json-schema/go-json-schema-validator" // Para validação de schema
	"gopkg.in/yaml.v3" // Para parsing de YAML
)

// ParseWorkflow lê um arquivo YAML/JSON e o converte para a estrutura Workflow
func ParseWorkflow(content []byte) (*Workflow, error) {
	var workflow Workflow
	err := yaml.Unmarshal(content, &workflow)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer unmarshal do workflow: %w", err)
	}

	// TODO: Aqui entraria a validação semântica mais complexa
	// por exemplo, verificar se IDs referenciados em 'inputs' existem,
	// ou se condições 'when' são parsáveis (já feito no executor/planner).

	return &workflow, nil
}

// ValidateContent valida o conteúdo do workflow contra o schema JSON
// schemaContent: o conteúdo JSON do schema (logline.schema.json)
// workflowContent: o conteúdo YAML/JSON do workflow a ser validado
func ValidateContent(filePath string, schemaContent, workflowContent []byte) *errorsInternal.ValidationError {
	v := gojsonschema.NewValidator()
	schemaLoader := gojsonschema.NewBytesLoader(schemaContent)
	
	result, err := v.Validate(schemaLoader, gojsonschema.NewBytesLoader(workflowContent))
	if err != nil {
		return errorsInternal.NewValidationError(filePath, []string{"Erro interno ao processar schema ou documento"}, err)
	}

	if result.Valid() {
		return nil // Retorna nil se válido
	} else {
		errMsgs := make([]string, len(result.Errors()))
		for i, desc := range result.Errors() {
			errMsgs[i] = desc.String() // string() dá uma descrição mais legível
		}
		return errorsInternal.NewValidationError(filePath, errMsgs, nil)
	}
}
