package guruui

import (
	"testing"
)

func TestNew(t *testing.T) {
	ui := New()
	if ui == nil {
		t.Error("New() returned nil")
	}
}

func TestTextStyles(t *testing.T) {
	ui := New()
	
	// Test that Text method doesn't panic
	ui.Text("test", StylePrimary)
	ui.Text("test", StyleMuted)
	ui.Text("test", StyleAccent)
}

func TestStatusMessages(t *testing.T) {
	ui := New()
	
	// These should not panic
	ui.Info("test info")
	ui.Success("test success")
	ui.Warn("test warning")
	ui.Error("test error")
}

func TestHeader(t *testing.T) {
	ui := New()
	
	// Test header formatting
	ui.Header("Test Title")
	// Header should print uppercase title with underline
}

func TestSeparator(t *testing.T) {
	ui := New()
	
	// Test separator
	ui.Separator()
	// Should print 50 dashes
}

func TestSpace(t *testing.T) {
	ui := New()
	
	// Test spacing
	ui.Space(2)
	// Should add 2 newlines
}

func TestCenter(t *testing.T) {
	ui := New()
	
	result := ui.Center("test", 10)
	expected := "   test   " // 10 chars wide, centered
	
	if result != expected {
		t.Errorf("Center() returned '%s', expected '%s'", result, expected)
	}
}

func TestRightAlign(t *testing.T) {
	ui := New()
	
	result := ui.RightAlign("test", 10)
	expected := "      test" // 10 chars wide, right aligned
	
	if result != expected {
		t.Errorf("RightAlign() returned '%s', expected '%s'", result, expected)
	}
}

func TestLeftAlign(t *testing.T) {
	ui := New()
	
	result := ui.LeftAlign("test", 10)
	expected := "test      " // 10 chars wide, left aligned
	
	if result != expected {
		t.Errorf("LeftAlign() returned '%s', expected '%s'", result, expected)
	}
}

func TestIndent(t *testing.T) {
	ui := New()
	
	input := "line1\nline2"
	result := ui.Indent(input, 2)
	expected := "  line1\n  line2"
	
	if result != expected {
		t.Errorf("Indent() returned '%s', expected '%s'", result, expected)
	}
}

func TestDivider(t *testing.T) {
	ui := New()
	
	// Test divider without text
	ui.Divider("")
	
	// Test divider with text
	ui.Divider("Test")
}

func TestGrid(t *testing.T) {
	ui := New()
	
	items := []string{"a", "b", "c", "d"}
	ui.Grid(items, 2)
	// Should print items in 2 columns
}

func TestBox(t *testing.T) {
	ui := New()
	
	content := "test content"
	ui.Box(content, "Test Box")
	// Should create bordered box around content
}

func TestTable(t *testing.T) {
	ui := New()
	
	headers := []string{"Col1", "Col2"}
	rows := [][]string{
		{"row1col1", "row1col2"},
		{"row2col1", "row2col2"},
	}
	
	ui.Table(headers, rows)
	// Should create formatted table
}

func TestProgressBar(t *testing.T) {
	ui := New()
	
	ui.ProgressBar(5, 10, "Test Progress")
	// Should show 50% progress
}

func TestPanel(t *testing.T) {
	ui := New()
	
	ui.Panel("Test Panel", func() {
		ui.Text("Panel content", StylePrimary)
	}, BorderThin)
	// Should create bordered panel with content
}
