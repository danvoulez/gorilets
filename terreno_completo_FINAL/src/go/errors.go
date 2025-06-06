// internal/errors/errors.go
package errors

import "fmt"

// WorkflowError é uma interface para erros relacionados a workflows
type WorkflowError interface {
	error
	WorkflowError() bool
}

// BlockExecutionError representa um erro ocorrido durante a execução de um bloco.
type BlockExecutionError struct {
	BlockID string
	Type    string
	Cause   error // A causa original do erro
	Message string
}

func (e *BlockExecutionError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("erro no bloco '%s' (%s): %s, causa: %v", e.BlockID, e.Type, e.Message, e.Cause)
	}
	return fmt.Sprintf("erro no bloco '%s' (%s): %s", e.BlockID, e.Type, e.Message)
}

func (e *BlockExecutionError) Unwrap() error {
	return e.Cause
}

func (e *BlockExecutionError) WorkflowError() bool {
	return true
}

// ValidationError representa um erro de validação (schema ou semântica).
type ValidationError struct {
	FilePath string
	Errors   []string
	Cause    error // Causa original, se houver
}

func (e *ValidationError) Error() string {
	msg := fmt.Sprintf("erros de validação no arquivo '%s':\n", e.FilePath)
	for _, err := range e.Errors {
		msg += fmt.Sprintf("- %s\n", err)
	}
	if e.Cause != nil {
		msg += fmt.Sprintf("causa: %v\n", e.Cause)
	}
	return msg
}

func (e *ValidationError) WorkflowError() bool {
	return true
}

// NewBlockExecutionError cria uma nova BlockExecutionError
func NewBlockExecutionError(blockID, blockType, message string, cause error) *BlockExecutionError {
	return &BlockExecutionError{
		BlockID: blockID,
		Type:    blockType,
		Cause:   cause,
		Message: message,
	}
}

// NewValidationError cria uma nova ValidationError
func NewValidationError(filePath string, errors []string, cause error) *ValidationError {
	return &ValidationError{
		FilePath: filePath,
		Errors:   errors,
		Cause:    cause,
	}
}

// IsWorkflowError verifica se um erro é do tipo WorkflowError
func IsWorkflowError(err error) bool {
	_, ok := err.(WorkflowError)
	return ok
}
