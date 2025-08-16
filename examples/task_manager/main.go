package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/arnislvdev/go-guruui"
)

// Task represents a single task item
type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
	Priority    string
	Created     time.Time
	DueDate     *time.Time
	Category    string
	Tags        []string
}

// TaskManager handles task operations
type TaskManager struct {
	tasks  []Task
	nextID int
	ui     *guruui.UI
}

// NewTaskManager creates a new task manager instance
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks:  []Task{},
		nextID: 1,
		ui:     guruui.New(),
	}
}

// AddTask adds a new task with enhanced neo-brutalist UI
func (tm *TaskManager) AddTask() {
	tm.ui.BrutalistHeader("Add New Task", "Create a new task with enhanced details")

	title := tm.ui.Input("Task title")
	if title == "" {
		tm.ui.Error("Task title cannot be empty")
		return
	}

	description := tm.ui.Input("Task description (optional)")

	// Priority selection with geometric display
	priorities := []string{"Low", "Medium", "High", "Critical"}
	tm.ui.GeometricDivider("Priority Selection", guruui.PatternDiamond)
	priority := tm.ui.Select("Select priority:", priorities)

	// Category selection
	categories := []string{"Work", "Personal", "Study", "Health", "Finance", "Creative"}
	tm.ui.GeometricDivider("Category Selection", guruui.PatternHexagon)
	category := tm.ui.Select("Select category:", categories)

	// Due date input
	dueDateStr := tm.ui.Input("Due date (YYYY-MM-DD, optional)")
	var dueDate *time.Time
	if dueDateStr != "" {
		if parsed, err := time.Parse("2006-01-02", dueDateStr); err == nil {
			dueDate = &parsed
		} else {
			tm.ui.Warn("Invalid date format, skipping due date")
		}
	}

	// Tags input
	tagsInput := tm.ui.Input("Tags (comma-separated, optional)")
	var tags []string
	if tagsInput != "" {
		tagParts := strings.Split(tagsInput, ",")
		for _, tag := range tagParts {
			tags = append(tags, strings.TrimSpace(tag))
		}
	}

	// Create task
	task := Task{
		ID:          tm.nextID,
		Title:       title,
		Description: description,
		Status:      "Pending",
		Priority:    priority,
		Created:     time.Now(),
		DueDate:     dueDate,
		Category:    category,
		Tags:        tags,
	}

	tm.tasks = append(tm.tasks, task)
	tm.nextID++

	tm.ui.Success("Task added successfully!")

	// Show task in neo-brutalist panel
	tm.ui.BrutalistPanel("New Task Created", func() {
		tm.displayTaskDetails(task)
	}, guruui.BorderGeometric)
}

// ListTasks displays all tasks with enhanced visual layouts
func (tm *TaskManager) ListTasks() {
	if len(tm.tasks) == 0 {
		tm.ui.Info("No tasks found. Add some tasks first!")
		return
	}

	tm.ui.BrutalistHeader("Task List", "All your tasks in organized layouts")

	// Choose display style
	styles := []string{"Table View", "Hexagon Grid", "Diamond Matrix", "Fractal Tree"}
	choice := tm.ui.Select("Choose display style:", styles)

	switch choice {
	case "Table View":
		tm.displayTasksTable()
	case "Hexagon Grid":
		tm.displayTasksHexagon()
	case "Diamond Matrix":
		tm.displayTasksDiamond()
	case "Fractal Tree":
		tm.displayTasksFractal()
	}

	// Show summary with progress ring
	tm.ui.Space(1)
	tm.ui.ProgressRing(len(tm.tasks), len(tm.tasks), "Total Tasks")
	tm.ui.NpmStatus("info", fmt.Sprintf("Total tasks: %d", len(tm.tasks)))
}

// displayTasksTable shows tasks in traditional table format
func (tm *TaskManager) displayTasksTable() {
	headers := []string{"ID", "Title", "Status", "Priority", "Category", "Due Date"}
	rows := [][]string{}

	for _, task := range tm.tasks {
		dueDateStr := "No due date"
		if task.DueDate != nil {
			dueDateStr = task.DueDate.Format("2006-01-02")
		}

		rows = append(rows, []string{
			strconv.Itoa(task.ID),
			task.Title,
			task.Status,
			task.Priority,
			task.Category,
			dueDateStr,
		})
	}

	tm.ui.Table(headers, rows)
}

// displayTasksHexagon shows tasks in hexagonal grid layout
func (tm *TaskManager) displayTasksHexagon() {
	var taskStrings []string
	for _, task := range tm.tasks {
		taskStr := fmt.Sprintf("#%d: %s", task.ID, task.Title)
		taskStrings = append(taskStrings, taskStr)
	}

	tm.ui.Space(1)
	tm.ui.GeometricDivider("Hexagonal Layout", guruui.PatternHexagon)
	tm.ui.HexagonGrid(taskStrings, 3)
}

// displayTasksDiamond shows tasks in diamond matrix layout
func (tm *TaskManager) displayTasksDiamond() {
	var taskStrings []string
	for _, task := range tm.tasks {
		taskStr := fmt.Sprintf("#%d: %s", task.ID, task.Title)
		taskStrings = append(taskStrings, taskStr)
	}

	tm.ui.Space(1)
	tm.ui.GeometricDivider("Diamond Matrix", guruui.PatternDiamond)
	tm.ui.DiamondMatrix(taskStrings)
}

// displayTasksFractal shows tasks in fractal tree structure
func (tm *TaskManager) displayTasksFractal() {
	// Group tasks by category
	categoryMap := make(map[string]interface{})
	for _, task := range tm.tasks {
		if categoryMap[task.Category] == nil {
			categoryMap[task.Category] = make(map[string]interface{})
		}
		subMap := categoryMap[task.Category].(map[string]interface{})
		subMap[task.Title] = fmt.Sprintf("Status: %s, Priority: %s", task.Status, task.Priority)
	}

	tm.ui.Space(1)
	tm.ui.GeometricDivider("Fractal Tree by Category", guruui.PatternFractal)
	tm.ui.FractalTree(categoryMap, "", true)
}

// ViewTask shows detailed information about a specific task
func (tm *TaskManager) ViewTask() {
	if len(tm.tasks) == 0 {
		tm.ui.Info("No tasks to view")
		return
	}

	tm.ui.BrutalistHeader("View Task Details", "Comprehensive task information")

	// Create task options for selection
	var options []string
	for _, task := range tm.tasks {
		options = append(options, fmt.Sprintf("#%d - %s", task.ID, task.Title))
	}

	choice := tm.ui.Select("Select task to view:", options)

	// Extract task ID from choice
	parts := strings.Split(choice, " - ")
	if len(parts) != 2 {
		tm.ui.Error("Invalid selection")
		return
	}

	taskID, err := strconv.Atoi(strings.TrimPrefix(parts[0], "#"))
	if err != nil {
		tm.ui.Error("Invalid task ID")
		return
	}

	// Find and display task
	for _, task := range tm.tasks {
		if task.ID == taskID {
			tm.ui.BrutalistPanel("Task Details", func() {
				tm.displayTaskDetails(task)
			}, guruui.BorderLayered)
			return
		}
	}

	tm.ui.Error("Task not found")
}

// displayTaskDetails shows comprehensive task information with enhanced styling
func (tm *TaskManager) displayTaskDetails(task Task) {
	// Header with neo-brutalist style
	tm.ui.NpmHeader(fmt.Sprintf("Task #%d", task.ID), task.Status, task.Title)

	if task.Description != "" {
		tm.ui.TextLine("Description:", guruui.StyleAccent)
		tm.ui.TextLine(task.Description, guruui.StylePrimary)
		tm.ui.Space(1)
	}

	// Priority with color coding and geometric patterns
	priorityColor := "\033[32m" // Green for low
	switch task.Priority {
	case "Medium":
		priorityColor = "\033[33m" // Yellow
	case "High":
		priorityColor = "\033[31m" // Red
	case "Critical":
		priorityColor = "\033[35m" // Magenta
	}

	tm.ui.Text("Priority: ", guruui.StylePrimary)
	fmt.Printf("%s%s\033[0m %s\n", priorityColor, task.Priority, guruui.PatternDiamond[:2])

	// Category with hexagon pattern
	tm.ui.Text("Category: ", guruui.StylePrimary)
	fmt.Printf("%s %s\n", task.Category, guruui.PatternHexagon[:2])

	// Tags with star pattern
	if len(task.Tags) > 0 {
		tm.ui.Text("Tags: ", guruui.StylePrimary)
		fmt.Printf("%s %s\n", strings.Join(task.Tags, ", "), guruui.PatternStar[:2])
	}

	// Dates
	tm.ui.TextLine("Created: "+task.Created.Format("2006-01-02 15:04"), guruui.StylePrimary)

	if task.DueDate != nil {
		daysUntil := int(task.DueDate.Sub(time.Now()).Hours() / 24)
		dueDateStr := task.DueDate.Format("2006-01-02")

		if daysUntil < 0 {
			tm.ui.Text("Due date: ", guruui.StylePrimary)
			tm.ui.TextLine(dueDateStr+" (OVERDUE!)", guruui.StyleAccent)
		} else if daysUntil == 0 {
			tm.ui.Text("Due date: ", guruui.StylePrimary)
			tm.ui.TextLine(dueDateStr+" (Today)", guruui.StyleAccent)
		} else if daysUntil <= 3 {
			tm.ui.Text("Due date: ", guruui.StylePrimary)
			tm.ui.TextLine(fmt.Sprintf("%s (%d days left)", dueDateStr, daysUntil), guruui.StyleAccent)
		} else {
			tm.ui.TextLine("Due date: "+dueDateStr, guruui.StylePrimary)
		}
	}
}

// UpdateTaskStatus allows changing task status with enhanced UI
func (tm *TaskManager) UpdateTaskStatus() {
	if len(tm.tasks) == 0 {
		tm.ui.Info("No tasks to update")
		return
	}

	tm.ui.BrutalistHeader("Update Task Status", "Change task workflow status")

	// Create task options
	var options []string
	for _, task := range tm.tasks {
		options = append(options, fmt.Sprintf("#%d - %s (%s)", task.ID, task.Title, task.Status))
	}

	choice := tm.ui.Select("Select task to update:", options)

	// Extract task ID
	parts := strings.Split(choice, " - ")
	if len(parts) != 2 {
		tm.ui.Error("Invalid selection")
		return
	}

	taskID, err := strconv.Atoi(strings.TrimPrefix(parts[0], "#"))
	if err != nil {
		tm.ui.Error("Invalid task ID")
		return
	}

	// Find task and update status
	for i := range tm.tasks {
		if tm.tasks[i].ID == taskID {
			statuses := []string{"Pending", "In Progress", "Completed", "Cancelled"}
			newStatus := tm.ui.Select("Select new status:", statuses)

			oldStatus := tm.tasks[i].Status
			tm.tasks[i].Status = newStatus

			tm.ui.Success(fmt.Sprintf("Task #%d status updated: %s → %s", taskID, oldStatus, newStatus))

			// Show progress ring for completion
			if newStatus == "Completed" {
				completed := 0
				for _, t := range tm.tasks {
					if t.Status == "Completed" {
						completed++
					}
				}
				tm.ui.ProgressRing(completed, len(tm.tasks), "Completion Progress")
			}
			return
		}
	}

	tm.ui.Error("Task not found")
}

// DeleteTask removes a task with enhanced confirmation
func (tm *TaskManager) DeleteTask() {
	if len(tm.tasks) == 0 {
		tm.ui.Info("No tasks to delete")
		return
	}

	tm.ui.BrutalistHeader("Delete Task", "Remove tasks with confirmation")

	// Create task options
	var options []string
	for _, task := range tm.tasks {
		options = append(options, fmt.Sprintf("#%d - %s", task.ID, task.Title))
	}

	choice := tm.ui.Select("Select task to delete:", options)

	// Extract task ID
	parts := strings.Split(choice, " - ")
	if len(parts) != 2 {
		tm.ui.Error("Invalid selection")
		return
	}

	taskID, err := strconv.Atoi(strings.TrimPrefix(parts[0], "#"))
	if err != nil {
		tm.ui.Error("Invalid task ID")
		return
	}

	// Find task and confirm deletion
	for i, task := range tm.tasks {
		if task.ID == taskID {
			if tm.ui.Confirm(fmt.Sprintf("Are you sure you want to delete task #%d: %s?", taskID, task.Title)) {
				// Remove task
				tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
				tm.ui.Success(fmt.Sprintf("Task #%d deleted successfully", taskID))
			} else {
				tm.ui.Info("Task deletion cancelled")
			}
			return
		}
	}

	tm.ui.Error("Task not found")
}

// ShowDashboard displays task statistics with neo-brutalist styling
func (tm *TaskManager) ShowDashboard() {
	tm.ui.BrutalistHeader("Task Dashboard", "Comprehensive overview and statistics")

	if len(tm.tasks) == 0 {
		tm.ui.Info("No tasks yet. Create your first task to get started!")
		return
	}

	// Calculate statistics
	total := len(tm.tasks)
	pending := 0
	inProgress := 0
	completed := 0
	cancelled := 0
	overdue := 0

	// Category breakdown
	categoryCount := make(map[string]int)

	for _, task := range tm.tasks {
		switch task.Status {
		case "Pending":
			pending++
		case "In Progress":
			inProgress++
		case "Completed":
			completed++
		case "Cancelled":
			cancelled++
		}

		categoryCount[task.Category]++

		// Check for overdue tasks
		if task.DueDate != nil && task.Status != "Completed" && task.Status != "Cancelled" {
			if task.DueDate.Before(time.Now()) {
				overdue++
			}
		}
	}

	// Display statistics in neo-brutalist panels
	tm.ui.BrutalistPanel("Task Statistics", func() {
		tm.ui.NpmList([]string{
			fmt.Sprintf("Total Tasks: %d", total),
			fmt.Sprintf("Pending: %d", pending),
			fmt.Sprintf("In Progress: %d", inProgress),
			fmt.Sprintf("Completed: %d", completed),
			fmt.Sprintf("Cancelled: %d", cancelled),
		}, guruui.IconDot)
	}, guruui.BorderGeometric)

	// Category breakdown
	tm.ui.BrutalistPanel("Category Breakdown", func() {
		for category, count := range categoryCount {
			fmt.Printf("\033[36m%s\033[0m: %d tasks\n", category, count)
		}
	}, guruui.BorderLayered)

	if overdue > 0 {
		tm.ui.Space(1)
		tm.ui.NpmStatus("warning", fmt.Sprintf("You have %d overdue tasks!", overdue))
	}

	// Progress indicators
	if total > 0 {
		tm.ui.Space(1)
		tm.ui.GeometricDivider("Progress Tracking", guruui.PatternWave)

		// Multiple progress styles
		tm.ui.ProgressBar(completed, total, "Overall Progress")
		tm.ui.ProgressWave(completed, total, "Wave Progress")
		tm.ui.ProgressRing(completed, total, "Ring Progress")

		completionRate := float64(completed) / float64(total) * 100
		tm.ui.NpmStatus("info", fmt.Sprintf("Completion rate: %.1f%%", completionRate))
	}
}

// ShowAdvancedFeatures demonstrates the new neo-brutalist components
func (tm *TaskManager) ShowAdvancedFeatures() {
	tm.ui.BrutalistHeader("Advanced UI Features", "Neo-brutalist terminal design showcase")

	// Geometric patterns demonstration
	tm.ui.BrutalistPanel("Geometric Patterns", func() {
		fmt.Printf("Hexagon: %s\n", guruui.PatternHexagon[:10])
		fmt.Printf("Diamond: %s\n", guruui.PatternDiamond[:8])
		fmt.Printf("Triangle: %s\n", guruui.PatternTriangle[:8])
		fmt.Printf("Square: %s\n", guruui.PatternSquare[:8])
		fmt.Printf("Circle: %s\n", guruui.PatternCircle[:8])
	}, guruui.BorderGeometric)

	// Progress variations
	tm.ui.BrutalistPanel("Progress Variations", func() {
		tm.ui.ProgressBar(7, 10, "Standard Bar")
		tm.ui.ProgressWave(7, 10, "Wave Style")
		tm.ui.ProgressRing(7, 10, "Ring Style")
	}, guruui.BorderLayered)

	// Layout demonstrations
	tm.ui.BrutalistPanel("Layout Demonstrations", func() {
		items := []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5", "Item 6"}

		tm.ui.TextLine("Hexagon Grid:", guruui.StyleAccent)
		tm.ui.HexagonGrid(items, 3)

		tm.ui.Space(1)
		tm.ui.TextLine("Diamond Matrix:", guruui.StyleAccent)
		tm.ui.DiamondMatrix(items)
	}, guruui.BorderFractal)
}

// Run starts the enhanced task manager application
func (tm *TaskManager) Run() {
	tm.ui.BrutalistHeader("Neo-Brutalist Task Manager", "Advanced terminal UI with unique designs")
	tm.ui.Info("Welcome to the future of terminal task management!")
	tm.ui.Space(1)

	for {
		tm.ui.NpmSection("Main Menu")
		options := []string{
			"Add new task",
			"List all tasks",
			"View task details",
			"Update task status",
			"Delete task",
			"Show dashboard",
			"Advanced features",
			"Exit",
		}

		choice := tm.ui.Select("Choose an action:", options)
		tm.ui.Space(1)

		switch choice {
		case "Add new task":
			tm.AddTask()
		case "List all tasks":
			tm.ListTasks()
		case "View task details":
			tm.ViewTask()
		case "Update task status":
			tm.UpdateTaskStatus()
		case "Delete task":
			tm.DeleteTask()
		case "Show dashboard":
			tm.ShowDashboard()
		case "Advanced features":
			tm.ShowAdvancedFeatures()
		case "Exit":
			tm.ui.BrutalistHeader("Goodbye", "Session summary and farewell")
			tm.ui.Success("Thank you for using Neo-Brutalist Task Manager!")
			tm.ui.NpmStatus("info", fmt.Sprintf("You managed %d tasks in this session", len(tm.tasks)))
			return
		}

		tm.ui.Space(1)
		tm.ui.GeometricDivider("", guruui.PatternFractal)
		tm.ui.Space(1)
	}
}

func main() {
	taskManager := NewTaskManager()
	taskManager.Run()
}
