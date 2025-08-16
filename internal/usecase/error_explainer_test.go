package usecase

import (
	"testing"
)

func TestNewErrorExplainer(t *testing.T) {
	explainer := NewErrorExplainer()
	if explainer == nil {
		t.Error("NewErrorExplainer should not return nil")
	}
}

func TestDetectErrorType(t *testing.T) {
	explainer := &ErrorExplainer{}
	
	tests := []struct {
		input    string
		expected string
	}{
		{"undefined: fmt", "undefined_symbol"},
		{"cannot use nil as type string", "type_mismatch"},
		{"imported and not used", "unused_import"},
		{"declared and not used", "unused_variable"},
		{"missing return", "missing_return"},
		{"too many arguments", "argument_count"},
		{"not enough arguments", "argument_count"},
		{"unknown error", "unknown"},
	}
	
	for _, test := range tests {
		result := explainer.detectErrorType(test.input)
		if result != test.expected {
			t.Errorf("detectErrorType(%q) = %s, want %s", test.input, result, test.expected)
		}
	}
}

func TestDetermineSeverity(t *testing.T) {
	explainer := &ErrorExplainer{}
	
	tests := []struct {
		input    string
		expected string
	}{
		{"undefined_symbol", "error"},
		{"type_mismatch", "error"},
		{"unused_import", "warning"},
		{"unused_variable", "warning"},
		{"missing_return", "error"},
		{"argument_count", "error"},
		{"unknown", "info"},
	}
	
	for _, test := range tests {
		result := explainer.determineSeverity(test.input)
		if result != test.expected {
			t.Errorf("determineSeverity(%q) = %s, want %s", test.input, result, test.expected)
		}
	}
}
