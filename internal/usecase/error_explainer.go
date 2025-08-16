package usecase

import (
	"fmt"
	"strings"

	"github.com/arnislvdev/go-guru-ui/internal/domain"
	"github.com/arnislvdev/go-guru-ui/internal/infrastructure/ai"
	"github.com/arnislvdev/go-guru-ui/pkg/humor"
)

// ErrorExplainer handles the business logic for explaining errors
type ErrorExplainer struct {
	aiClient ai.Client
	humor    *humor.WTFMode
}

// NewErrorExplainer creates a new ErrorExplainer instance
func NewErrorExplainer() *ErrorExplainer {
	return &ErrorExplainer{
		aiClient: ai.NewOpenAIClient(),
		humor:    humor.NewWTFMode(),
	}
}

// Explain explains an error message in the specified mode
func (e *ErrorExplainer) Explain(errorMsg, file string, line int, mode string) (string, error) {
	// Parse the error to extract structured information
	parsedError := e.parseError(errorMsg, file, line)

	// Generate explanation using AI
	explanation, err := e.aiClient.ExplainError(parsedError)
	if err != nil {
		return "", fmt.Errorf("AI explanation failed: %w", err)
	}

	// Apply mode-specific formatting
	if mode == "wtf" {
		explanation = e.humor.Enhance(explanation, parsedError.Type)
	}

	return explanation, nil
}

// parseError extracts structured information from error messages
func (e *ErrorExplainer) parseError(errorMsg, file string, line int) *domain.Error {
	errorType := e.detectErrorType(errorMsg)

	return &domain.Error{
		Message:  errorMsg,
		Type:     errorType,
		File:     file,
		Line:     line,
		Severity: e.determineSeverity(errorType),
		Language: "go", // MVP: Go only, extendable later
	}
}

// detectErrorType identifies the type of error
func (e *ErrorExplainer) detectErrorType(errorMsg string) string {
	errorMsg = strings.ToLower(errorMsg)

	switch {
	case strings.Contains(errorMsg, "undefined"):
		return "undefined_symbol"
	case strings.Contains(errorMsg, "cannot use"):
		return "type_mismatch"
	case strings.Contains(errorMsg, "imported and not used"):
		return "unused_import"
	case strings.Contains(errorMsg, "declared and not used"):
		return "unused_variable"
	case strings.Contains(errorMsg, "missing return"):
		return "missing_return"
	case strings.Contains(errorMsg, "too many arguments"):
		return "argument_count"
	case strings.Contains(errorMsg, "not enough arguments"):
		return "argument_count"
	default:
		return "unknown"
	}
}

// determineSeverity assigns a severity level to the error
func (e *ErrorExplainer) determineSeverity(errorType string) string {
	switch errorType {
	case "undefined_symbol", "type_mismatch":
		return "error"
	case "unused_import", "unused_variable":
		return "warning"
	case "missing_return":
		return "error"
	case "argument_count":
		return "error"
	default:
		return "info"
	}
}
