package guruui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Select displays a list of options and returns the user's choice.
// Shows numbered options with clean formatting for easy selection.
func (ui *UI) Select(prompt string, options []string) string {
	fmt.Printf("\033[36m%s\033[0m\n", prompt) // Cyan prompt

	for i, option := range options {
		fmt.Printf("  \033[33m%d\033[0m. %s\n", i+1, option) // Yellow numbers
	}

	fmt.Printf("\nEnter your choice (1-%d): ", len(options))

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(options) {
		ui.Error("Invalid choice. Please try again.")
		return ui.Select(prompt, options)
	}

	return options[choice-1]
}

// MultiSelect allows users to choose multiple options.
// Returns a slice of selected option indices.
func (ui *UI) MultiSelect(prompt string, options []string) []int {
	fmt.Printf("\033[36m%s\033[0m\n", prompt)                                      // Cyan prompt
	fmt.Printf("\033[90mEnter numbers separated by spaces (e.g., 1 3 4)\033[0m\n") // Gray instruction

	for i, option := range options {
		fmt.Printf("  \033[33m%d\033[0m. %s\n", i+1, option) // Yellow numbers
	}

	fmt.Print("\nYour choices: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var selected []int
	parts := strings.Fields(input)

	for _, part := range parts {
		if choice, err := strconv.Atoi(part); err == nil && choice >= 1 && choice <= len(options) {
			selected = append(selected, choice-1)
		}
	}

	if len(selected) == 0 {
		ui.Error("No valid choices selected. Please try again.")
		return ui.MultiSelect(prompt, options)
	}

	return selected
}

// Table displays data in a clean, bordered table format.
// Uses alternating row backgrounds for better readability.
func (ui *UI) Table(headers []string, rows [][]string) {
	if len(headers) == 0 {
		return
	}

	// Calculate column widths
	colWidths := make([]int, len(headers))
	for i, header := range headers {
		colWidths[i] = len(header)
	}

	for _, row := range rows {
		for i, cell := range row {
			if i < len(colWidths) && len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	// Print top border
	ui.printTableBorder(colWidths, "top")

	// Print headers with bold styling
	ui.printTableRow(headers, colWidths, true)

	// Print separator
	ui.printTableBorder(colWidths, "separator")

	// Print rows with alternating background hints
	for _, row := range rows {
		ui.printTableRow(row, colWidths, false)
	}

	// Print bottom border
	ui.printTableBorder(colWidths, "bottom")
}

// printTableRow prints a single row of the table.
// Handles padding and alignment for consistent appearance.
func (ui *UI) printTableRow(cells []string, widths []int, isHeader bool) {
	fmt.Print("│")

	for i, cell := range cells {
		if i >= len(widths) {
			break
		}

		padding := widths[i] - len(cell)
		if isHeader {
			fmt.Printf(" \033[1m%s\033[0m%s │", cell, strings.Repeat(" ", padding)) // Bold headers
		} else {
			fmt.Printf(" %s%s │", cell, strings.Repeat(" ", padding))
		}
	}

	fmt.Println()
}

// printTableBorder prints table borders with proper corner characters.
// Creates clean, professional-looking table edges.
func (ui *UI) printTableBorder(widths []int, borderType string) {
	switch borderType {
	case "top":
		fmt.Print("┌")
		for i, width := range widths {
			if i > 0 {
				fmt.Print("┬")
			}
			fmt.Print(strings.Repeat("─", width+2))
		}
		fmt.Println("┐")
	case "separator":
		fmt.Print("├")
		for i, width := range widths {
			if i > 0 {
				fmt.Print("┼")
			}
			fmt.Print(strings.Repeat("─", width+2))
		}
		fmt.Println("┤")
	case "bottom":
		fmt.Print("└")
		for i, width := range widths {
			if i > 0 {
				fmt.Print("┴")
			}
			fmt.Print(strings.Repeat("─", width+2))
		}
		fmt.Println("┘")
	}
}

// ProgressBar displays a progress indicator with npm-style blocks.
// Shows completion percentage and visual progress representation.
func (ui *UI) ProgressBar(current, total int, label string) {
	if total <= 0 {
		return
	}

	percentage := float64(current) / float64(total)
	barWidth := 30
	filled := int(float64(barWidth) * percentage)

	// NPM-style progress bar with colors
	fmt.Printf("\033[36m%s\033[0m [", label) // Cyan label

	for i := 0; i < barWidth; i++ {
		if i < filled {
			fmt.Print("\033[32m█\033[0m") // Green filled blocks
		} else {
			fmt.Print("\033[90m░\033[0m") // Gray empty blocks
		}
	}

	fmt.Printf("] \033[33m%d%%\033[0m\n", int(percentage*100)) // Yellow percentage
}

// Spinner displays a loading animation with npm-style characters.
// Shows rotating characters to indicate ongoing operations.
func (ui *UI) Spinner(msg string) {
	spinnerChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

	fmt.Printf("\033[36m%s\033[0m %s", msg, spinnerChars[0]) // Cyan message with spinner

	// Note: This is a basic implementation.
	// For real-time animation, you'd need to handle terminal control sequences.
	fmt.Println()
}

// Panel creates a bordered container for grouping content.
// Useful for organizing related information visually.
func (ui *UI) Panel(title string, content func(), style BorderStyle) {
	if style == BorderNone {
		if title != "" {
			fmt.Printf("\033[90m--- %s ---\033[0m\n", title) // Gray title
		}
		content()
		return
	}

	// Top border with colored title
	if title != "" {
		fmt.Printf("┌─ \033[36m%s\033[0m %s\n", title, strings.Repeat("─", 50-len(title)-4)) // Cyan title
	} else {
		fmt.Println("┌" + strings.Repeat("─", 50))
	}

	// Content
	content()

	// Bottom border
	fmt.Println("└" + strings.Repeat("─", 50))
}

// Input prompts the user for text input with colored prompt.
// Returns the entered string for further processing.
func (ui *UI) Input(prompt string) string {
	fmt.Printf("\033[36m%s\033[0m: ", prompt) // Cyan prompt

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Confirm asks for yes/no confirmation with colored prompt.
// Returns true for yes, false for no.
func (ui *UI) Confirm(prompt string) bool {
	fmt.Printf("\033[33m%s\033[0m (y/n): ", prompt) // Yellow prompt

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	return input == "y" || input == "yes"
}

// InstallProgress shows npm-style installation progress.
// Useful for package managers or dependency installation.
func (ui *UI) InstallProgress(packageName, version string, current, total int) {
	if total <= 0 {
		return
	}

	percentage := float64(current) / float64(total)
	barWidth := 20
	filled := int(float64(barWidth) * percentage)

	// NPM-style install progress
	fmt.Printf("\033[32m%s\033[0m \033[90m@\033[0m\033[33m%s\033[0m ", packageName, version) // Green package, gray @, yellow version

	// Progress bar
	fmt.Print("[")
	for i := 0; i < barWidth; i++ {
		if i < filled {
			fmt.Print("\033[32m█\033[0m") // Green filled
		} else {
			fmt.Print("\033[90m░\033[0m") // Gray empty
		}
	}
	fmt.Printf("] \033[33m%d%%\033[0m\n", int(percentage*100)) // Yellow percentage
}

// PackageInfo displays package information in npm style.
// Shows package details with icons and colors.
func (ui *UI) PackageInfo(packageName, version, description string) {
	fmt.Printf("\033[32m%s\033[0m \033[90m@\033[0m\033[33m%s\033[0m\n", packageName, version)
	if description != "" {
		fmt.Printf("  \033[90m%s\033[0m\n", description) // Gray description
	}
}

// DependencyTree shows dependency relationships in tree format.
// Useful for displaying package dependencies.
func (ui *UI) DependencyTree(dependencies map[string]string, level int) {
	indent := strings.Repeat("  ", level)

	for pkg, version := range dependencies {
		if level == 0 {
			fmt.Printf("%s\033[32m%s\033[0m \033[90m@\033[0m\033[33m%s\033[0m\n", indent, pkg, version)
		} else {
			fmt.Printf("%s\033[90m├─\033[0m \033[32m%s\033[0m \033[90m@\033[0m\033[33m%s\033[0m\n", indent, pkg, version)
		}
	}
}

// === NEW NEO-BRUTALIST COMPONENTS ===

// BrutalistPanel creates a neo-brutalist styled container with layered borders
func (ui *UI) BrutalistPanel(title string, content func(), style BorderStyle) {
	switch style {
	case BorderGeometric:
		ui.printGeometricBorder(title, content)
	case BorderLayered:
		ui.printLayeredBorder(title, content)
	case BorderFractal:
		ui.printFractalBorder(title, content)
	default:
		ui.Panel(title, content, style)
	}
}

// printGeometricBorder creates a geometric pattern border
func (ui *UI) printGeometricBorder(title string, content func()) {
	width := 60
	titleLen := len(title)

	// Top border with geometric pattern
	fmt.Print(CornerTopLeft)
	fmt.Print(strings.Repeat("─", (width-titleLen-6)/2))
	fmt.Printf(" \033[36m%s\033[0m ", title)
	fmt.Print(strings.Repeat("─", width-titleLen-6-(width-titleLen-6)/2))
	fmt.Println(CornerTopRight)

	// Content
	content()

	// Bottom border
	fmt.Print(CornerBottomLeft)
	fmt.Print(strings.Repeat("─", width-2))
	fmt.Println(CornerBottomRight)
}

// printLayeredBorder creates a layered depth effect
func (ui *UI) printLayeredBorder(title string, content func()) {
	width := 60
	titleLen := len(title)

	// Multiple layered borders
	fmt.Print("╔")
	fmt.Print(strings.Repeat("═", (width-titleLen-6)/2))
	fmt.Printf(" \033[36m%s\033[0m ", title)
	fmt.Print(strings.Repeat("═", width-titleLen-6-(width-titleLen-6)/2))
	fmt.Println("╗")

	fmt.Print("║")
	fmt.Print(strings.Repeat(" ", width-2))
	fmt.Println("║")

	// Content
	content()

	fmt.Print("║")
	fmt.Print(strings.Repeat(" ", width-2))
	fmt.Println("║")

	fmt.Print("╚")
	fmt.Print(strings.Repeat("═", width-2))
	fmt.Println("╝")
}

// printFractalBorder creates a fractal-like border pattern
func (ui *UI) printFractalBorder(title string, content func()) {
	width := 60
	titleLen := len(title)

	// Fractal top border
	fmt.Print("◢")
	fmt.Print(strings.Repeat("─", (width-titleLen-6)/2))
	fmt.Printf(" \033[36m%s\033[0m ", title)
	fmt.Print(strings.Repeat("─", width-titleLen-6-(width-titleLen-6)/2))
	fmt.Println("◣")

	// Content
	content()

	// Fractal bottom border
	fmt.Print("◤")
	fmt.Print(strings.Repeat("─", width-2))
	fmt.Println("◥")
}

// ProgressRing displays a circular progress indicator
func (ui *UI) ProgressRing(current, total int, label string) {
	if total <= 0 {
		return
	}

	percentage := float64(current) / float64(total)
	ringFrames := []string{"◐", "◑", "◒", "◓"}
	frameIndex := int(percentage*float64(len(ringFrames))) % len(ringFrames)

	fmt.Printf("\033[36m%s\033[0m %s \033[33m%.1f%%\033[0m\n",
		label, ringFrames[frameIndex], percentage*100)
}

// ProgressWave displays a wave-like progress indicator
func (ui *UI) ProgressWave(current, total int, label string) {
	if total <= 0 {
		return
	}

	percentage := float64(current) / float64(total)
	waveWidth := 20
	filled := int(float64(waveWidth) * percentage)

	fmt.Printf("\033[36m%s\033[0m [", label)

	for i := 0; i < waveWidth; i++ {
		if i < filled {
			fmt.Print("\033[32m" + string(ProgressWave[i%len(ProgressWave)]) + "\033[0m")
		} else {
			fmt.Print("\033[90m" + string(ProgressWave[i%len(ProgressWave)]) + "\033[0m")
		}
	}

	fmt.Printf("] \033[33m%.1f%%\033[0m\n", percentage*100)
}

// HexagonGrid displays items in a hexagonal grid pattern
func (ui *UI) HexagonGrid(items []string, columns int) {
	if len(items) == 0 || columns <= 0 {
		return
	}

	rows := (len(items) + columns - 1) / columns

	for row := 0; row < rows; row++ {
		// Print hexagon row
		for col := 0; col < columns; col++ {
			index := row*columns + col
			if index >= len(items) {
				break
			}

			// Hexagon pattern
			if col > 0 {
				fmt.Print("  ") // Spacing between hexagons
			}

			fmt.Printf("\033[36m⬡\033[0m %s", items[index])
		}
		fmt.Println()

		// Print connecting lines (except for last row)
		if row < rows-1 {
			for col := 0; col < columns; col++ {
				index := row*columns + col
				if index >= len(items) {
					break
				}

				if col > 0 {
					fmt.Print("  ")
				}
				fmt.Print("  ╱╲")
			}
			fmt.Println()
		}
	}
}

// DiamondMatrix displays items in a diamond-shaped matrix
func (ui *UI) DiamondMatrix(items []string) {
	if len(items) == 0 {
		return
	}

	// Calculate diamond dimensions
	size := 1
	for size*size < len(items) {
		size += 2
	}

	center := size / 2

	for row := 0; row < size; row++ {
		// Calculate how many items to show in this row
		itemsInRow := 1 + 2*abs(row-center)
		startCol := center - (itemsInRow-1)/2

		// Print leading spaces
		fmt.Print(strings.Repeat("  ", abs(row-center)))

		// Print items in this row
		for col := startCol; col < startCol+itemsInRow; col++ {
			index := row*size + col
			if index < len(items) {
				fmt.Printf("\033[36m◆\033[0m %s  ", items[index])
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Println()
	}
}

// FractalTree displays hierarchical data in a fractal tree structure
func (ui *UI) FractalTree(data map[string]interface{}, prefix string, isLast bool) {
	for key, value := range data {
		// Print current node
		if prefix == "" {
			fmt.Printf("\033[36m%s\033[0m\n", key)
		} else {
			if isLast {
				fmt.Printf("%s└─ \033[36m%s\033[0m\n", prefix, key)
			} else {
				fmt.Printf("%s├─ \033[36m%s\033[0m\n", prefix, key)
			}
		}

		// Recursively print children
		if subMap, ok := value.(map[string]interface{}); ok {
			newPrefix := prefix
			if prefix != "" {
				if isLast {
					newPrefix += "   "
				} else {
					newPrefix += "│  "
				}
			}

			// Find if this is the last child
			keys := make([]string, 0, len(subMap))
			for k := range subMap {
				keys = append(keys, k)
			}

			for i, subKey := range keys {
				isLastChild := i == len(keys)-1
				ui.FractalTree(map[string]interface{}{subKey: subMap[subKey]}, newPrefix, isLastChild)
			}
		} else {
			// Print leaf value
			if prefix != "" {
				if isLast {
					fmt.Printf("%s   \033[90m%s\033[0m\n", prefix, fmt.Sprint(value))
				} else {
					fmt.Printf("%s│  \033[90m%s\033[0m\n", prefix, fmt.Sprint(value))
				}
			}
		}
	}
}

// AnimatedSpinner displays a rotating spinner with custom patterns
func (ui *UI) AnimatedSpinner(msg string, pattern []string, frame int) {
	if len(pattern) == 0 {
		return
	}

	frameIndex := frame % len(pattern)
	fmt.Printf("\033[36m%s\033[0m %s", msg, pattern[frameIndex])
}

// BrutalistHeader creates a neo-brutalist styled header
func (ui *UI) BrutalistHeader(title string, subtitle string) {
	width := 70
	titleLen := len(title)
	subtitleLen := len(subtitle)

	// Top border with brutalist style
	fmt.Print("╔")
	fmt.Print(strings.Repeat("═", (width-titleLen-2)/2))
	fmt.Printf(" \033[1m\033[36m%s\033[0m ", title)
	fmt.Print(strings.Repeat("═", width-titleLen-2-(width-titleLen-2)/2))
	fmt.Println("╗")

	// Subtitle line
	if subtitle != "" {
		fmt.Print("║")
		fmt.Print(strings.Repeat(" ", (width-subtitleLen-2)/2))
		fmt.Printf(" \033[90m%s\033[0m ", subtitle)
		fmt.Print(strings.Repeat(" ", width-subtitleLen-2-(width-subtitleLen-2)/2))
		fmt.Println("║")
	}

	// Bottom border
	fmt.Print("╚")
	fmt.Print(strings.Repeat("═", width-2))
	fmt.Println("╝")
}

// GeometricDivider creates a divider with geometric patterns
func (ui *UI) GeometricDivider(text string, pattern string) {
	if text == "" {
		fmt.Println(strings.Repeat(pattern, 20))
		return
	}

	textLen := len(text)
	patternLen := len(pattern)
	repeatCount := (40 - textLen) / patternLen

	fmt.Print(strings.Repeat(pattern, repeatCount))
	fmt.Printf(" \033[36m%s\033[0m ", text)
	fmt.Println(strings.Repeat(pattern, repeatCount))
}

// Utility function for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
