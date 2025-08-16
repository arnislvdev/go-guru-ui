package guruui

import (
	"fmt"
	"strings"
)

// Center centers text within a given width.
// Useful for headers and titles that should be visually balanced.
func (ui *UI) Center(text string, width int) string {
	if len(text) >= width {
		return text
	}
	
	padding := (width - len(text)) / 2
	leftPad := padding
	rightPad := width - len(text) - leftPad
	
	return strings.Repeat(" ", leftPad) + text + strings.Repeat(" ", rightPad)
}

// RightAlign positions text to the right within a given width.
// Good for numbers, status indicators, and right-aligned content.
func (ui *UI) RightAlign(text string, width int) string {
	if len(text) >= width {
		return text
	}
	
	padding := width - len(text)
	return strings.Repeat(" ", padding) + text
}

// LeftAlign positions text to the left within a given width.
// Standard text alignment for most content.
func (ui *UI) LeftAlign(text string, width int) string {
	if len(text) >= width {
		return text[:width]
	}
	
	padding := width - len(text)
	return text + strings.Repeat(" ", padding)
}

// PrintCentered prints centered text with automatic width detection.
// Centers the text within the available terminal width.
func (ui *UI) PrintCentered(text string) {
	// Default width if we can't determine terminal size
	width := 80
	centered := ui.Center(text, width)
	fmt.Println(centered)
}

// PrintRightAligned prints right-aligned text.
// Useful for status information and right-side content.
func (ui *UI) PrintRightAligned(text string, width int) {
	aligned := ui.RightAlign(text, width)
	fmt.Print(aligned)
}

// PrintLeftAligned prints left-aligned text with padding.
// Ensures consistent left margin for content.
func (ui *UI) PrintLeftAligned(text string, width int) {
	aligned := ui.LeftAlign(text, width)
	fmt.Print(aligned)
}

// Indent adds left margin to text.
// Helps with visual hierarchy and content organization.
func (ui *UI) Indent(text string, spaces int) string {
	indent := strings.Repeat(" ", spaces)
	lines := strings.Split(text, "\n")
	
	for i, line := range lines {
		lines[i] = indent + line
	}
	
	return strings.Join(lines, "\n")
}

// PrintIndented prints text with left margin.
// Convenience method for indented output.
func (ui *UI) PrintIndented(text string, spaces int) {
	indented := ui.Indent(text, spaces)
	fmt.Print(indented)
}

// Box creates a simple bordered container around text.
// Uses basic characters for maximum terminal compatibility.
func (ui *UI) Box(text string, title string) {
	lines := strings.Split(text, "\n")
	maxWidth := 0
	
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}
	
	// Account for title if present
	if title != "" && len(title)+2 > maxWidth {
		maxWidth = len(title) + 2
	}
	
	// Top border with colored title
	if title != "" {
		fmt.Printf("┌─ \033[36m%s\033[0m %s\n", title, strings.Repeat("─", maxWidth-len(title)-2)) // Cyan title
	} else {
		fmt.Println("┌" + strings.Repeat("─", maxWidth))
	}
	
	// Content
	for _, line := range lines {
		padding := maxWidth - len(line)
		fmt.Printf("│ %s%s │\n", line, strings.Repeat(" ", padding))
	}
	
	// Bottom border
	fmt.Println("└" + strings.Repeat("─", maxWidth))
}

// Divider creates a horizontal line with optional text.
// Useful for separating sections of content.
func (ui *UI) Divider(text string) {
	if text == "" {
		fmt.Println(strings.Repeat("─", 50))
		return
	}
	
	lineWidth := 50
	textWidth := len(text) + 2
	sideWidth := (lineWidth - textWidth) / 2
	
	fmt.Print(strings.Repeat("─", sideWidth))
	fmt.Printf(" \033[36m%s\033[0m ", text) // Cyan text
	fmt.Println(strings.Repeat("─", lineWidth-sideWidth-textWidth))
}

// Grid displays content in a simple grid layout.
// Useful for organizing related information in columns.
func (ui *UI) Grid(items []string, columns int) {
	if columns <= 0 || len(items) == 0 {
		return
	}
	
	rows := (len(items) + columns - 1) / columns
	
	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			index := row*columns + col
			if index < len(items) {
				fmt.Printf("\033[90m%-20s\033[0m", items[index]) // Gray items
			}
		}
		fmt.Println()
	}
}

// NpmHeader creates an npm-style header with package info.
// Shows package name, version, and description in npm format.
func (ui *UI) NpmHeader(packageName, version, description string) {
	fmt.Printf("\033[32m%s\033[0m \033[90m@\033[0m\033[33m%s\033[0m\n", packageName, version) // Green package, gray @, yellow version
	if description != "" {
		fmt.Printf("  \033[90m%s\033[0m\n", description) // Gray description
	}
	fmt.Println()
}

// NpmSection creates a section header in npm style.
// Uses npm-style formatting for section titles.
func (ui *UI) NpmSection(title string) {
	fmt.Printf("\n\033[36m%s\033[0m\n", title) // Cyan section title
	fmt.Printf("\033[90m%s\033[0m\n", strings.Repeat("─", len(title))) // Gray underline
}

// NpmList displays a list of items in npm style.
// Shows items with npm-style formatting and icons.
func (ui *UI) NpmList(items []string, icon string) {
	for _, item := range items {
		fmt.Printf("  %s \033[90m%s\033[0m\n", icon, item) // Icon with gray item
	}
}

// NpmStatus shows status information in npm style.
// Displays status with appropriate colors and icons.
func (ui *UI) NpmStatus(status, message string) {
	var color, statusIcon string
	
	switch status {
	case "success":
		color = "\033[32m" // Green
		statusIcon = IconSuccess
	case "error":
		color = "\033[31m" // Red
		statusIcon = IconError
	case "warning":
		color = "\033[33m" // Yellow
		statusIcon = IconWarning
	case "info":
		color = "\033[34m" // Blue
		statusIcon = IconInfo
	default:
		color = "\033[90m" // Gray
		statusIcon = IconDot
	}
	
	fmt.Printf("%s%s\033[0m %s\n", color, statusIcon, message)
}

// NpmCommand shows a command in npm style.
// Displays commands with npm-style formatting.
func (ui *UI) NpmCommand(command string) {
	fmt.Printf("\033[90m$\033[0m \033[36m%s\033[0m\n", command) // Gray $, cyan command
}

// NpmOutput shows command output in npm style.
// Displays output with appropriate formatting.
func (ui *UI) NpmOutput(output string) {
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			fmt.Printf("  \033[90m%s\033[0m\n", line) // Gray output with indent
		}
	}
}
