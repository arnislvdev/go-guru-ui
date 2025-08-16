package main

import (
	"time"

	"github.com/arnislvdev/go-guruui"
)

func main() {
	ui := guruui.New()

	// NPM-style header
	ui.NpmHeader("guruui-cli", "1.0.0", "Command line interface for GuruUI")

	// Welcome message
	ui.Info("Welcome to GuruUI CLI - Package Manager")
	ui.Space(1)

	// Main menu
	ui.NpmSection("Available Commands")
	options := []string{
		"Install packages",
		"List installed packages",
		"Update packages",
		"Remove packages",
		"Search packages",
		"Exit",
	}

	choice := ui.Select("Choose a command:", options)
	ui.Success("Selected: " + choice)
	ui.Space(1)

	switch choice {
	case "Install packages":
		// Package installation workflow
		ui.NpmSection("Package Installation")

		packageName := ui.Input("Package name")
		version := ui.Input("Version (or latest)")

		if version == "" {
			version = "latest"
		}

		ui.Info("Installing " + packageName + "@" + version)
		ui.Spinner("Resolving dependencies...")
		time.Sleep(2 * time.Second)

		// Show installation progress
		ui.Info("Installing packages...")
		ui.InstallProgress(packageName, version, 1, 5)
		time.Sleep(500 * time.Millisecond)
		ui.InstallProgress(packageName, version, 2, 5)
		time.Sleep(500 * time.Millisecond)
		ui.InstallProgress(packageName, version, 3, 5)
		time.Sleep(500 * time.Millisecond)
		ui.InstallProgress(packageName, version, 4, 5)
		time.Sleep(500 * time.Millisecond)
		ui.InstallProgress(packageName, version, 5, 5)

		ui.Success("Package installed successfully!")

		// Show package info
		ui.Space(1)
		ui.NpmSection("Installed Package")
		ui.PackageInfo(packageName, version, "Successfully installed package")

	case "List installed packages":
		// List packages
		ui.NpmSection("Installed Packages")

		packages := map[string]string{
			"guruui": "1.0.0",
			"tcell":  "2.6.0",
			"tview":  "0.0.0",
		}

		ui.DependencyTree(packages, 0)

	case "Update packages":
		// Update packages
		ui.NpmSection("Package Updates")

		ui.Info("Checking for updates...")
		ui.Spinner("Checking registry...")
		time.Sleep(2 * time.Second)

		// Show update progress
		ui.Info("Updating packages...")
		ui.InstallProgress("guruui", "1.1.0", 1, 3)
		time.Sleep(800 * time.Millisecond)
		ui.InstallProgress("tcell", "2.7.0", 2, 3)
		time.Sleep(800 * time.Millisecond)
		ui.InstallProgress("tview", "0.1.0", 3, 3)

		ui.Success("All packages updated successfully!")

	case "Remove packages":
		// Remove packages
		ui.NpmSection("Package Removal")

		packageToRemove := ui.Select("Select package to remove:", []string{"guruui", "tcell", "tview"})

		if ui.Confirm("Are you sure you want to remove " + packageToRemove + "?") {
			ui.Info("Removing " + packageToRemove)
			ui.Spinner("Uninstalling...")
			time.Sleep(1 * time.Second)
			ui.Success("Package removed successfully!")
		} else {
			ui.Info("Package removal cancelled")
		}

	case "Search packages":
		// Search packages
		ui.NpmSection("Package Search")

		searchTerm := ui.Input("Search term")
		ui.Info("Searching for: " + searchTerm)
		ui.Spinner("Searching registry...")
		time.Sleep(2 * time.Second)

		// Show search results
		ui.Success("Found 3 packages:")
		ui.Space(1)

		ui.PackageInfo("guruui", "1.0.0", "A minimalist terminal UI library for Go")
		ui.PackageInfo("guruui-cli", "1.0.0", "Command line interface for GuruUI")
		ui.PackageInfo("guruui-themes", "0.5.0", "Theme collection for GuruUI")

		// Show dependency tree for selected package
		if ui.Confirm("Show dependencies for guruui?") {
			ui.Space(1)
			ui.NpmSection("Dependencies")
			deps := map[string]string{
				"tcell": "2.6.0",
				"tview": "0.0.0",
			}
			ui.DependencyTree(deps, 0)
		}
	}

	ui.Space(1)
	ui.NpmSection("Session Summary")
	ui.NpmStatus("success", "CLI session completed successfully")
	ui.NpmStatus("info", "Thank you for using GuruUI CLI!")
}
