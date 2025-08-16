package domain

// Error represents a programming error with structured information
type Error struct {
	Message  string `json:"message"`
	Type     string `json:"type"`
	File     string `json:"file,omitempty"`
	Line     int    `json:"line,omitempty"`
	Severity string `json:"severity"`
	Language string `json:"language"`
}

// ErrorType constants
const (
	ErrorTypeUndefinedSymbol = "undefined_symbol"
	ErrorTypeTypeMismatch    = "type_mismatch"
	ErrorTypeUnusedImport    = "unused_import"
	ErrorTypeUnusedVariable  = "unused_variable"
	ErrorTypeMissingReturn   = "missing_return"
	ErrorTypeArgumentCount   = "argument_count"
	ErrorTypeUnknown         = "unknown"
)

// Severity constants
const (
	SeverityInfo    = "info"
	SeverityWarning = "warning"
	SeverityError   = "error"
	SeverityFatal   = "fatal"
)

// Language constants
const (
	LanguageGo      = "go"
	LanguagePython  = "python"
	LanguageJS      = "javascript"
	LanguageRust    = "rust"
	LanguageUnknown = "unknown"
)
