package main

import (
	"time"

	"github.com/arnislvdev/go-guruui"
)

func main() {
	ui := guruui.New()

	ui.Header("Manga CLI")
	ui.Info("Welcome to the Manga Management System")
	ui.Space(1)

	// Simulate fetching data
	ui.Info("Fetching results from AniList...")
	ui.Spinner("Connecting to API")
	time.Sleep(1 * time.Second)
	ui.Success("Connected successfully")
	ui.Space(1)

	// Main menu
	ui.TextLine("What would you like to do?", guruui.StyleAccent)
	options := []string{
		"Search for manga",
		"View reading list",
		"Update progress",
		"Exit",
	}

	choice := ui.Select("Choose an option:", options)
	ui.Success("You selected: " + choice)
	ui.Space(1)

	if choice == "Search for manga" {
		// Search functionality
		ui.TextLine("Search for manga:", guruui.StyleAccent)
		searchTerm := ui.Input("Enter manga title")

		ui.Info("Searching for: " + searchTerm)
		ui.Spinner("Searching...")
		time.Sleep(2 * time.Second)

		// Display results
		ui.Success("Found 4 results:")
		ui.Space(1)

		headers := []string{"Title", "Author", "Status", "Score"}
		rows := [][]string{
			{"Bleach", "Tite Kubo", "Completed", "8.2"},
			{"Naruto", "Masashi Kishimoto", "Completed", "8.1"},
			{"One Piece", "Eiichiro Oda", "Ongoing", "9.1"},
			{"Dragon Ball", "Akira Toriyama", "Completed", "8.5"},
		}

		ui.Table(headers, rows)
		ui.Space(1)

		// Select from results
		mangaOptions := []string{"Bleach", "Naruto", "One Piece", "Dragon Ball"}
		selected := ui.Select("Which manga would you like to add?", mangaOptions)
		ui.Success("Added " + selected + " to your reading list!")

	} else if choice == "View reading list" {
		// Reading list
		ui.TextLine("Your Reading List:", guruui.StyleAccent)
		ui.Space(1)

		ui.Panel("Currently Reading", func() {
			ui.TextLine("One Piece - Chapter 1089", guruui.StylePrimary)
			ui.TextLine("Bleach - Chapter 686", guruui.StylePrimary)
		}, guruui.BorderThin)

		ui.Space(1)

		ui.Panel("Completed", func() {
			ui.TextLine("Naruto - All chapters", guruui.StylePrimary)
			ui.TextLine("Dragon Ball - All chapters", guruui.StylePrimary)
		}, guruui.BorderThin)

	} else if choice == "Update progress" {
		// Progress update
		ui.TextLine("Update Reading Progress:", guruui.StyleAccent)
		ui.Space(1)

		manga := ui.Select("Select manga:", []string{"One Piece", "Bleach"})
		chapter := ui.Input("Enter current chapter number")

		ui.Info("Updating progress for " + manga + " to chapter " + chapter)
		ui.ProgressBar(1, 1, "Saving progress")
		ui.Success("Progress updated successfully!")
	}

	ui.Space(1)
	ui.Divider("Session End")
	ui.Info("Thank you for using Manga CLI!")
}
