// internal/config_loader/config_loader.go
package config_loader

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config representa a estrutura do arquivo logline.yml
type Config struct {
	LLM   LLMConfig   `mapstructure:"llm"`
	Hooks HooksConfig `mapstructure:"hooks"`
}

// LLMConfig para configurações de LLM
type LLMConfig struct {
	Provider string       `mapstructure:"provider"`
	OpenAI   OpenAIConfig `mapstructure:"openai"`
	Ollama   OllamaConfig `mapstructure:"ollama"`
}

// OpenAIConfig para configurações da OpenAI
type OpenAIConfig struct {
	APIKey string `mapstructure:"api_key"`
	Model  string `mapstructure:"model"`
}

// OllamaConfig para configurações do Ollama
type OllamaConfig struct {
    BaseURL string `mapstructure:"base_url"`
    Model   string `mapstructure:"model"`
}

// HooksConfig para todas as definições de hooks
type HooksConfig map[string][]HookDefinition

// HookDefinition define um hook individual
type HookDefinition struct {
	Name    string            `mapstructure:"name"`
	Type    string            `mapstructure:"type"` // "shell" ou "http"
	Command string            `mapstructure:"command"`
	Args    []string          `mapstructure:"args"`
	URL     string            `mapstructure:"url"`
	Method  string            `mapstructure:"method"`
	Headers map[string]string `mapstructure:"headers"`
	Body    string            `mapstructure:"body"`
}

// LoadConfig carrega a configuração do arquivo logline.yml
func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()
	v.SetConfigName("logline") // Nome do arquivo sem extensão
	v.SetConfigType("yaml")    // Tipo de arquivo
	
    // Caminhos de busca padrão
	v.AddConfigPath(configPath) // Caminho passado (ex: ./config)
	v.AddConfigPath("./config") // Caminho padrão relativo ao executável
	v.AddConfigPath("$HOME/.loglineos") // Caminho no diretório home do usuário
	v.AddConfigPath(".") // Diretório atual

	v.SetDefault("llm.provider", "mock") // Default para mock se não configurado
	v.SetDefault("llm.openai.model", "gpt-3.5-turbo")
    v.SetDefault("llm.ollama.base_url", "http://localhost:11434")
    v.SetDefault("llm.ollama.model", "llama3")

	if err := v.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            fmt.Fprintf(os.Stderr, "Configuração: arquivo logline.yml não encontrado. Usando defaults ou variáveis de ambiente.\n")
        } else {
            return nil, fmt.Errorf("config: erro ao ler arquivo de configuração: %w", err)
        }
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("config: erro ao fazer unmarshal da configuração: %w", err)
	}

    // Validações básicas
    if config.LLM.Provider != "mock" && config.LLM.Provider != "openai" && config.LLM.Provider != "ollama" {
        return nil, fmt.Errorf("config: provedor LLM '%s' inválido. Use 'openai', 'ollama' ou 'mock'.", config.LLM.Provider)
    }
    if config.LLM.Provider == "openai" && config.LLM.OpenAI.APIKey == "" {
        fmt.Fprintf(os.Stderr, "ATENÇÃO: Provedor OpenAI selecionado, mas 'llm.openai.api_key' não configurado em logline.yml. Chamadas falharão.\n")
    }

	return &config, nil
}

// GetDefaultConfigPath retorna o caminho padrão para a configuração
func GetDefaultConfigPath() string {
    return filepath.Join(".", "config")
}
