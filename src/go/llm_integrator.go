// internal/llm_integrator/llm_integrator.go
package llm_integrator

import (
	"context"
	"fmt"
	"loglineos/internal/config_loader"
	"time"

	"github.com/sashabaranov/go-openai" // Para OpenAI
	// Para uso futuro de JSON Schema, se a resposta do LLM precisar ser validada
	// "github.com/sashabaranov/go-openai/jsonschema"
	// "github.com/sashabaranov/go-openai/types" 

    "github.com/ollama/ollama-go/ollama" // Para Ollama
    "github.com/ollama/ollama-go/api"    // Para Ollama API

)

// LLMResponse representa a resposta de um LLM
type LLMResponse struct {
	Output      string                 `json:"output"`
	TrustLevel  string                 `json:"trust_level"`
	Provider    string                 `json:"provider"`
	Model       string                 `json:"model"`
	TokensUsed  int                    `json:"tokens_used,omitempty"`
	RawResponse map[string]interface{} `json:"raw_response,omitempty"` // Para inspeção
}

// LLMIntegrator interface para diferentes provedores de LLM
type LLMIntegrator interface {
	Generate(ctx context.Context, prompt string, trustLevel string, modelOverride string) (*LLMResponse, error)
}

// NewLLMIntegrator cria uma instância do integrador LLM com base na configuração
func NewLLMIntegrator(config *config_loader.LLMConfig) (LLMIntegrator, error) {
	switch config.Provider {
	case "openai":
		if config.OpenAI.APIKey == "" {
			return nil, fmt.Errorf("LLM: chave de API OpenAI não configurada")
		}
		return &OpenAIIntegrator{
			client: openai.NewClient(config.OpenAI.APIKey),
			defaultModel: config.OpenAI.Model,
		}, nil
	case "ollama":
        if config.Ollama.BaseURL == "" {
            return nil, fmt.Errorf("LLM: URL base do Ollama não configurada")
        }
        return &OllamaIntegrator{
            client: ollama.NewClient(config.Ollama.BaseURL),
            defaultModel: config.Ollama.Model,
        }, nil
	case "mock":
		return &MockLLMIntegrator{}, nil
	default:
		return nil, fmt.Errorf("LLM: provedor '%s' não suportado", config.Provider)
	}
}

// OpenAIIntegrator implementa LLMIntegrator para OpenAI
type OpenAIIntegrator struct {
	client       *openai.Client
	defaultModel string
}

func (o *OpenAIIntegrator) Generate(ctx context.Context, prompt string, trustLevel string, modelOverride string) (*LLMResponse, error) {
	model := o.defaultModel
	if modelOverride != "" {
		model = modelOverride
	}

	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: "Você é um assistente LogLineOS. Gero respostas concisas e úteis."},
		{Role: openai.ChatMessageRoleUser, Content: prompt},
	}

	req := openai.ChatCompletionRequest{
		Model:    model,
		Messages: messages,
		MaxTokens: 500, // Limite para evitar custos excessivos na demo
	}

	resp, err := o.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("LLM OpenAI: falha ao gerar resposta: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("LLM OpenAI: nenhuma escolha na resposta")
	}

	return &LLMResponse{
		Output:      resp.Choices[0].Message.Content,
		TrustLevel:  trustLevel,
		Provider:    "openai",
		Model:       model,
		TokensUsed:  resp.Usage.TotalTokens,
		RawResponse: map[string]interface{}{"usage": resp.Usage, "id": resp.ID},
	}, nil
}

// OllamaIntegrator implementa LLMIntegrator para Ollama
type OllamaIntegrator struct {
    client *ollama.Client
    defaultModel string
}

func (o *OllamaIntegrator) Generate(ctx context.Context, prompt string, trustLevel string, modelOverride string) (*LLMResponse, error) {
    model := o.defaultModel
    if modelOverride != "" {
        model = modelOverride
    }

    req := &api.GenerateRequest{
        Model:  model,
        Prompt: prompt,
        Stream: false, // Não usar stream para esta demo
    }

    resp, err := o.client.Generate(ctx, req)
    if err != nil {
        return nil, fmt.Errorf("LLM Ollama: falha ao gerar resposta: %w", err)
    }

    return &LLMResponse{
        Output:      resp.Response,
        TrustLevel:  trustLevel,
        Provider:    "ollama",
        Model:       model,
        TokensUsed:  0, // Ollama API response doesn't directly provide tokens
        RawResponse: map[string]interface{}{"done": resp.Done, "context": resp.Context},
    }, nil
}

// MockLLMIntegrator para simulação
type MockLLMIntegrator struct{}

func (m *MockLLMIntegrator) Generate(ctx context.Context, prompt string, trustLevel string, modelOverride string) (*LLMResponse, error) {
	// Simula um delay
	time.Sleep(500 * time.Millisecond)
	return &LLMResponse{
		Output:      fmt.Sprintf("MockLLM: Resposta simulada para '%s' (Trust: %s)", prompt, trustLevel),
		TrustLevel:  trustLevel,
		Provider:    "mock",
		Model:       "mock-model-v1",
		TokensUsed:  100,
		RawResponse: map[string]interface{}{"simulated": true},
	}, nil
}
