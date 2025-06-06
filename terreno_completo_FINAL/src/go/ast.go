// internal/parser/ast.go
package parser

import "encoding/json"

// Workflow representa a estrutura de alto nível de um arquivo .logline
type Workflow struct {
	Name        string  `yaml:"name" json:"name"`
	Description string  `yaml:"description" json:"description"`
	Tasks       []Block `yaml:"tasks" json:"tasks"`
}

// Block representa um bloco genérico da DSL LogLine
type Block struct {
	ID      string          `yaml:"id" json:"id"`
	Type    string          `yaml:"type" json:"type"`
	Spec    json.RawMessage `yaml:"spec" json:"spec"` // Conteúdo do 'spec' como JSON bruto para unmarshal posterior
	Fallback *FallbackConfig `yaml:"fallback,omitempty" json:"fallback,omitempty"`
	Retry    *RetryPolicy    `yaml:"retry,omitempty" json:"retry,omitempty"`
	When     string          `yaml:"when,omitempty" json:"when,omitempty"`
	Tasks    []Block         `yaml:"tasks,omitempty" json:"-"` // json:"-" para evitar Marshal recursivo em JSON
	Hooks    []string        `yaml:"hooks,omitempty" json:"hooks,omitempty"`
}

// FallbackConfig define a configuração para um bloco de fallback
type FallbackConfig struct {
    Type string          `yaml:"type" json:"type"` // "cognitive", "mechanical"
    Spec json.RawMessage `yaml:"spec" json:"spec"`
}

// RetryPolicy define a política de retry para um bloco
type RetryPolicy struct {
	MaxAttempts   int    `yaml:"max_attempts" json:"max_attempts"`
	InitialDelay  string `yaml:"initial_delay,omitempty" json:"initial_delay,omitempty"` // e.g., "1s", "500ms"
	BackoffFactor float64 `yaml:"backoff_factor,omitempty" json:"backoff_factor,omitempty"` // e.g., 2.0 for exponential
}

type RegisterSpec struct {
	Entity string                 `yaml:"entity" json:"entity"`
	ID     string                 `yaml:"id" json:"id"`
	Data   map[string]interface{} `yaml:"data" json:"data"`
	Schema string                 `yaml:"schema,omitempty" json:"schema,omitempty"`
}

type AffairSpec struct {
	Type     string `yaml:"type" json:"type"`
	From     string `yaml:"from" json:"from"`
	To       string `yaml:"to" json:"to"`
	Note     string `yaml:"note,omitempty" json:"note,omitempty"`
	Intensity int    `yaml:"intensity,omitempty" json:"intensity,omitempty"`
}

type ObserveSpec struct {
	Source  string          `yaml:"source" json:"source"`
	Pattern string          `yaml:"pattern" json:"pattern"`
	Action  ObserveAction   `yaml:"action" json:"action"`
	Interval string         `yaml:"interval,omitempty" json:"interval,omitempty"`
	OnEvent string          `yaml:"on_event,omitempty" json:"on_event,omitempty"`
	ListenPath string        `yaml:"listen_path,omitempty" json:"listen_path,omitempty"`
}

type ObserveAction struct {
	Type   string `yaml:"type" json:"type"`
	Target string `yaml:"target" json:"target"`
	Message string `yaml:"message,omitempty" json:"message,omitempty"`
}

type CommitSpec struct {
	Goal     string   `yaml:"goal" json:"goal"`
	Inputs   []string `yaml:"inputs" json:"inputs"`
	SignedBy []string `yaml:"signed_by,omitempty" json:"signed_by,omitempty"`
	Message  string   `yaml:"message,omitempty" json:"message,omitempty"`
}

type MechanicalSpec struct {
	Command string   `yaml:"command" json:"command"`
	Args    []string `yaml:"args,omitempty" json:"args,omitempty"`
	Timeout string   `yaml:"timeout,omitempty" json:"timeout,omitempty"`
	Env     map[string]string `yaml:"env,omitempty" json:"env,omitempty"`
}

type CognitiveSpec struct {
	LLM CognitiveLLMConfig `yaml:"llm" json:"llm"`
	ActionOnOutput *CognitiveActionOnOutput `yaml:"action_on_output,omitempty" json:"action_on_output,omitempty"`
}

type CognitiveLLMConfig struct {
	Goal  string `yaml:"goal" json:"goal"`
	Input string `yaml:"input" json:"input"`
	Trust string `yaml:"trust" json:"trust"`
	Model string `yaml:"model,omitempty" json:"model,omitempty"`
}

// CognitiveActionOnOutput define a ação a ser tomada com a saída do LLM
type CognitiveActionOnOutput struct {
    TriggerBlockID string `yaml:"trigger_block_id" json:"trigger_block_id"`
    Condition      string `yaml:"condition,omitempty" json:"condition,omitempty"`
    MapOutputToInput map[string]string `yaml:"map_output_to_input,omitempty" json:"map_output_to_input,omitempty"`
}

type LoopSpec struct {
	Count      *int   `yaml:"count,omitempty" json:"count,omitempty"`
	Until      string `yaml:"until,omitempty" json:"until,omitempty"`
	MaxDuration string `yaml:"max_duration,omitempty" json:"max_duration,omitempty"`
}

type AwaitSpec struct {
	SpanID    string `yaml:"span_id,omitempty" json:"span_id,omitempty"`
	TaskID    string `yaml:"task_id,omitempty" json:"task_id,omitempty"`
	Condition string `yaml:"condition,omitempty" json:"condition,omitempty"`
	Timeout   string `yaml:"timeout" json:"timeout"`
}
