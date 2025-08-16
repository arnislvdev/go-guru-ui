package main

import (
	"time"

	"github.com/arnislvdev/go-guruui"
)

func main() {
	ui := guruui.New()

	// Header and basic text
	ui.Header("GuruUI Demo")
	ui.Info("Welcome to the GuruUI demonstration")
	ui.Space(1)

	// Text styling examples
	ui.Text("Primary text: ", guruui.StylePrimary)
	ui.TextLine("This is normal text", guruui.StylePrimary)

	ui.Text("Muted text: ", guruui.StylePrimary)
	ui.TextLine("This is secondary information", guruui.StyleMuted)

	ui.Text("Accent text: ", guruui.StylePrimary)
	ui.TextLine("This is highlighted text", guruui.StyleAccent)

	ui.Space(1)

	// Enhanced status messages with colors and icons
	ui.Info("Starting application...")
	ui.Success("Configuration loaded successfully")
	ui.Warn("Some features may be limited")
	ui.Error("Failed to connect to database")

	ui.Space(1)

	// Interactive selection with colored prompts
	ui.TextLine("Let's test the selection component:", guruui.StyleAccent)
	options := []string{"Option A", "Option B", "Option C", "Option D"}
	choice := ui.Select("Choose your preferred option:", options)
	ui.Success("You selected: " + choice)

	ui.Space(1)

	// Multi-selection with enhanced formatting
	ui.TextLine("Now let's test multi-selection:", guruui.StyleAccent)
	selected := ui.MultiSelect("Select multiple items:", options)
	ui.Info("You selected items: ")
	for _, index := range selected {
		ui.Text("  • ", guruui.StylePrimary)
		ui.TextLine(options[index], guruui.StylePrimary)
	}

	ui.Space(1)

	// Enhanced table with bold headers
	ui.TextLine("Here's a sample data table:", guruui.StyleAccent)
	headers := []string{"Name", "Age", "City", "Score"}
	rows := [][]string{
		{"Alice", "25", "New York", "95"},
		{"Bob", "30", "Los Angeles", "87"},
		{"Charlie", "28", "Chicago", "92"},
		{"Diana", "35", "Houston", "89"},
	}
	ui.Table(headers, rows)

	ui.Space(1)

	// NPM-style progress bar
	ui.TextLine("NPM-style progress demonstration:", guruui.StyleAccent)
	for i := 0; i <= 10; i++ {
		ui.ProgressBar(i, 10, "Processing")
		time.Sleep(200 * time.Millisecond)
		// Clear the line for next progress bar
		ui.Clear()
		ui.Header("GuruUI Demo")
		ui.Space(1)
	}

	ui.Space(1)

	// NPM-style installation progress
	ui.TextLine("NPM-style installation progress:", guruui.StyleAccent)
	ui.InstallProgress("guruui", "1.0.0", 5, 10)
	ui.InstallProgress("tcell", "2.6.0", 8, 10)
	ui.InstallProgress("tview", "0.0.0", 10, 10)

	ui.Space(1)

	// NPM-style package information
	ui.TextLine("NPM-style package info:", guruui.StyleAccent)
	ui.PackageInfo("guruui", "1.0.0", "A minimalist terminal UI library for Go")
	ui.PackageInfo("tcell", "2.6.0", "Terminal cell library")

	ui.Space(1)

	// NPM-style dependency tree
	ui.TextLine("NPM-style dependency tree:", guruui.StyleAccent)
	dependencies := map[string]string{
		"tcell": "2.6.0",
		"tview": "0.0.0",
	}
	ui.DependencyTree(dependencies, 0)

	ui.Space(1)

	// Enhanced panel demonstration
	ui.TextLine("Enhanced panel component:", guruui.StyleAccent)
	ui.Panel("Configuration", func() {
		ui.TextLine("Server: localhost:8080", guruui.StylePrimary)
		ui.TextLine("Database: postgres", guruui.StylePrimary)
		ui.TextLine("Cache: redis", guruui.StylePrimary)
	}, guruui.BorderThin)

	ui.Space(1)

	// NPM-style components
	ui.TextLine("NPM-style components:", guruui.StyleAccent)
	ui.NpmHeader("guruui-demo", "1.0.0", "Demo application for GuruUI")
	ui.NpmSection("Features")
	ui.NpmList([]string{"Color support", "Icons", "NPM-style formatting"}, guruui.IconStar)
	ui.NpmSection("Commands")
	ui.NpmCommand("go run main.go")
	ui.NpmOutput("Starting GuruUI demo...")
	ui.NpmStatus("success", "Demo completed successfully")

	ui.Space(1)

	// Layout utilities with enhanced colors
	ui.TextLine("Enhanced layout utilities:", guruui.StyleAccent)
	ui.PrintCentered("This text is centered")
	ui.PrintRightAligned("Right aligned text", 50)
	ui.Space(1)

	ui.TextLine("Enhanced grid layout:", guruui.StyleAccent)
	gridItems := []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5", "Item 6"}
	ui.Grid(gridItems, 3)

	ui.Space(1)

	// Enhanced box component
	ui.TextLine("Enhanced box component:", guruui.StyleAccent)
	ui.Box("This is some content\nthat spans multiple lines\nin a nice bordered box", "Content Box")

	ui.Space(1)

	// Enhanced divider
	ui.Divider("Section End")

	ui.Space(1)
	ui.Success("Demo completed successfully!")
	ui.TextLine("Thank you for trying GuruUI!", guruui.StyleAccent)
}
