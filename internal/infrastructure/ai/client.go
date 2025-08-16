package ai

import (
	"github.com/arnislvdev/go-guru-ui/internal/domain"
)

// Client defines the interface for AI providers
type Client interface {
	// ExplainError explains a programming error in plain English
	ExplainError(err *domain.Error) (string, error)

	// TranslateQuery converts natural language to CLI commands
	TranslateQuery(query, context string) (*domain.Command, error)

	// GetProvider returns the name of the AI provider
	GetProvider() string
}

// Config holds configuration for AI clients
type Config struct {
	Provider  string `json:"provider"`
	APIKey    string `json:"api_key"`
	Model     string `json:"model"`
	MaxTokens int    `json:"max_tokens"`
}

// DefaultConfig returns default AI configuration
func DefaultConfig() *Config {
	return &Config{
		Provider:  "openai",
		Model:     "gpt-4",
		MaxTokens: 1000,
	}
}
