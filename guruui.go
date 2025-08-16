package guruui

import (
	"fmt"
	"os"
	"strings"
)

// UI represents the main GuruUI interface.
// Provides methods for rendering various UI components in the terminal.
type UI struct {
	width int
}

// New creates a new GuruUI instance.
// Initializes the terminal and returns a ready-to-use UI.
func New() *UI {
	return &UI{}
}

// Header prints a prominent title at the top of the interface.
// Uses bold text and underlines for clear visual hierarchy.
func (ui *UI) Header(title string) {
	fmt.Printf("\n%s\n", strings.ToUpper(title))
	fmt.Println(strings.Repeat("=", len(title)))
	fmt.Println()
}

// Text prints text with the specified style.
// Styles include primary (normal), muted (dim), and accent (bold).
func (ui *UI) Text(text string, style TextStyle) {
	switch style {
	case StylePrimary:
		fmt.Print(text)
	case StyleMuted:
		fmt.Printf("\033[2m%s\033[0m", text)
	case StyleAccent:
		fmt.Printf("\033[1m%s\033[0m", text)
	default:
		fmt.Print(text)
	}
}

// TextLine prints styled text followed by a newline.
// Convenience method for common text output patterns.
func (ui *UI) TextLine(text string, style TextStyle) {
	ui.Text(text, style)
	fmt.Println()
}

// Info prints an informational message with icon and blue color.
// Used for general logs or status updates.
func (ui *UI) Info(msg string) {
	fmt.Printf("\033[34m%s\033[0m %s\n", IconInfo, msg)
}

// Success prints a success message with icon and green color.
// Used to confirm successful operations or positive outcomes.
func (ui *UI) Success(msg string) {
	fmt.Printf("\033[32m%s\033[0m %s\n", IconSuccess, msg)
}

// Warn prints a warning message with icon and yellow color.
// Used for non-critical issues or important notices.
func (ui *UI) Warn(msg string) {
	fmt.Printf("\033[33m%s\033[0m %s\n", IconWarning, msg)
}

// Error prints an error message with icon and red color.
// Used for errors, failures, or critical issues.
func (ui *UI) Error(msg string) {
	fmt.Printf("\033[31m%s\033[0m %s\n", IconError, msg)
}

// Separator prints a visual separator line.
// Useful for dividing sections of content.
func (ui *UI) Separator() {
	fmt.Println(strings.Repeat("-", 50))
}

// Space adds vertical spacing between components.
// Helps with readability and visual organization.
func (ui *UI) Space(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Println()
	}
}

// Clear clears the terminal screen.
// Useful for refreshing the interface or starting fresh.
func (ui *UI) Clear() {
	fmt.Print("\033[H\033[2J")
}

// Exit terminates the application with the given message.
// Prints the message and exits with the specified code.
func (ui *UI) Exit(msg string, code int) {
	if msg != "" {
		fmt.Println(msg)
	}
	os.Exit(code)
}
