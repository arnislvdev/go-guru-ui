package usecase

import (
	"github.com/arnislvdev/go-guru-ui/internal/infrastructure/ai"
)

// CommandTranslator handles the business logic for translating natural language to CLI commands
type CommandTranslator struct {
	aiClient ai.Client
}

// NewCommandTranslator creates a new CommandTranslator instance
func NewCommandTranslator() *CommandTranslator {
	return &CommandTranslator{
		aiClient: ai.NewOpenAIClient(),
	}
}

// Translate converts a natural language query to a CLI command
func (c *CommandTranslator) Translate(query, context, mode string) (string, string, error) {
	// Use AI to translate the query
	command, err := c.aiClient.TranslateQuery(query, context)
	if err != nil {
		return "", "", err
	}

	// Format the output based on mode
	if mode == "wtf" {
		command.Explanation = c.addHumorToExplanation(command.Explanation)
	}

	return command.Command, command.Explanation, nil
}

// addHumorToExplanation adds some humor to the explanation
func (c *CommandTranslator) addHumorToExplanation(explanation string) string {
	// Simple humor addition - in production, you'd use the WTF mode package
	// For now, just return the explanation with a prefix
	return "🎯 " + explanation
}
