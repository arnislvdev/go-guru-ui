package ai

import (
	"context"
	"fmt"
	"strings"

	"github.com/arnislvdev/go-guru-ui/internal/domain"
	"github.com/sashabaranov/go-openai"
)

// OpenAIClient implements the AI Client interface using OpenAI
type OpenAIClient struct {
	client *openai.Client
	config *Config
}

// NewOpenAIClient creates a new OpenAI client
func NewOpenAIClient() *OpenAIClient {
	// TODO: Get API key from config/environment
	apiKey := "your-openai-api-key" // This should come from config

	client := openai.NewClient(apiKey)

	return &OpenAIClient{
		client: client,
		config: DefaultConfig(),
	}
}

// ExplainError explains a programming error using OpenAI
func (c *OpenAIClient) ExplainError(err *domain.Error) (string, error) {
	prompt := c.buildErrorExplanationPrompt(err)

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: c.config.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful programming mentor who explains errors in clear, beginner-friendly terms.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: c.config.MaxTokens,
		},
	)

	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}

// TranslateQuery converts natural language to CLI commands
func (c *OpenAIClient) TranslateQuery(query, context string) (*domain.Command, error) {
	prompt := c.buildTranslationPrompt(query, context)

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: c.config.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a CLI expert who translates natural language into executable commands. Return only the command and a brief explanation.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			MaxTokens: c.config.MaxTokens,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("OpenAI API error: %w", err)
	}

	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no response from OpenAI")
	}

	// Parse the response to extract command and explanation
	// This is a simplified parser - in production, you'd want more robust parsing
	response := resp.Choices[0].Message.Content
	lines := strings.Split(response, "\n")

	var command string
	var explanation string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Command:") {
			command = strings.TrimPrefix(line, "Command:")
			command = strings.TrimSpace(command)
		} else if strings.HasPrefix(line, "Explanation:") {
			explanation = strings.TrimPrefix(line, "Explanation:")
			explanation = strings.TrimSpace(explanation)
		}
	}

	return &domain.Command{
		Command:     command,
		Explanation: explanation,
		Platform:    c.detectPlatform(context),
		Context:     context,
	}, nil
}

// GetProvider returns the provider name
func (c *OpenAIClient) GetProvider() string {
	return "openai"
}

// buildErrorExplanationPrompt creates a prompt for error explanation
func (c *OpenAIClient) buildErrorExplanationPrompt(err *domain.Error) string {
	prompt := fmt.Sprintf(`Explain this %s programming error in clear, beginner-friendly terms:

Error: %s
Type: %s
Severity: %s
Language: %s`, err.Severity, err.Message, err.Type, err.Severity, err.Language)

	if err.File != "" {
		prompt += fmt.Sprintf("\nFile: %s", err.File)
	}
	if err.Line > 0 {
		prompt += fmt.Sprintf("\nLine: %d", err.Line)
	}

	prompt += "\n\nProvide a clear explanation and suggest how to fix it."
	return prompt
}

// buildTranslationPrompt creates a prompt for command translation
func (c *OpenAIClient) buildTranslationPrompt(query, context string) string {
	prompt := fmt.Sprintf(`Translate this natural language request into a CLI command:

Request: %s`, query)

	if context != "" {
		prompt += fmt.Sprintf("\nContext: %s", context)
	}

	prompt += "\n\nRespond with:\nCommand: <the actual command>\nExplanation: <brief explanation of what it does>"
	return prompt
}

// detectPlatform attempts to detect the platform from context
func (c *OpenAIClient) detectPlatform(context string) string {
	context = strings.ToLower(context)

	switch {
	case strings.Contains(context, "linux") || strings.Contains(context, "ubuntu") || strings.Contains(context, "debian"):
		return domain.PlatformLinux
	case strings.Contains(context, "macos") || strings.Contains(context, "mac") || strings.Contains(context, "darwin"):
		return domain.PlatformMacOS
	case strings.Contains(context, "windows") || strings.Contains(context, "win"):
		return domain.PlatformWindows
	default:
		return domain.PlatformUnknown
	}
}
