# GuruUI API Reference

## Overview

GuruUI provides a clean, minimalist terminal user interface library for Go applications. All components use only black, white, and gray colors, with visual hierarchy achieved through typography, spacing, and borders.

## Core Types

### TextStyle
Controls how text is displayed:
- `StylePrimary`: Normal, readable text
- `StyleMuted`: Dimmed, secondary text  
- `StyleAccent`: Bold, emphasized text

### BorderStyle
Defines border appearance:
- `BorderNone`: No border
- `BorderThin`: Single-line borders
- `BorderThick`: Double-line borders

### Alignment
Controls horizontal positioning:
- `AlignLeft`: Left-aligned content
- `AlignCenter`: Centered content
- `AlignRight`: Right-aligned content

## Core Methods

### Initialization

#### `New() *UI`
Creates a new GuruUI instance.

```go
ui := guruui.New()
```

### Text and Styling

#### `Text(text string, style TextStyle)`
Prints styled text without a newline.

```go
ui.Text("Hello", guruui.StylePrimary)
ui.Text(" World", guruui.StyleAccent)
```

#### `TextLine(text string, style TextStyle)`
Prints styled text followed by a newline.

```go
ui.TextLine("This is a line", guruui.StylePrimary)
```

#### `Header(title string)`
Prints a prominent title with underlines.

```go
ui.Header("Application Name")
```

### Status Messages

#### `Info(msg string)`
Prints an informational message with [INFO] prefix.

```go
ui.Info("Starting application...")
```

#### `Success(msg string)`
Prints a success message with [SUCCESS] prefix.

```go
ui.Success("Operation completed successfully")
```

#### `Warn(msg string)`
Prints a warning message with [WARN] prefix.

```go
ui.Warn("Some features may be limited")
```

#### `Error(msg string)`
Prints an error message with [ERROR] prefix.

```go
ui.Error("Failed to connect to database")
```

### Layout and Spacing

#### `Space(lines int)`
Adds vertical spacing between components.

```go
ui.Space(2) // Adds 2 blank lines
```

#### `Separator()`
Prints a horizontal separator line.

```go
ui.Separator() // Prints 50 dashes
```

#### `Divider(text string)`
Creates a horizontal line with optional centered text.

```go
ui.Divider("Section End")
```

### Text Alignment

#### `Center(text string, width int) string`
Centers text within a specified width.

```go
centered := ui.Center("Title", 50)
```

#### `LeftAlign(text string, width int) string`
Left-aligns text with padding.

```go
aligned := ui.LeftAlign("Content", 30)
```

#### `RightAlign(text string, width int) string`
Right-aligns text with padding.

```go
aligned := ui.RightAlign("100", 20)
```

#### `Indent(text string, spaces int) string`
Adds left margin to multi-line text.

```go
indented := ui.Indent("Line 1\nLine 2", 4)
```

### Interactive Components

#### `Select(prompt string, options []string) string`
Displays a list of options and returns the user's choice.

```go
choice := ui.Select("Choose an option:", []string{"A", "B", "C"})
```

#### `MultiSelect(prompt string, options []string) []int`
Allows multiple selections, returns selected indices.

```go
selected := ui.MultiSelect("Pick items:", []string{"Item 1", "Item 2"})
```

#### `Input(prompt string) string`
Prompts for text input.

```go
name := ui.Input("Enter your name")
```

#### `Confirm(prompt string) bool`
Asks for yes/no confirmation.

```go
if ui.Confirm("Continue?") {
    // Proceed
}
```

### Data Display

#### `Table(headers []string, rows [][]string)`
Displays data in a clean, bordered table.

```go
headers := []string{"Name", "Age", "City"}
rows := [][]string{
    {"Alice", "25", "New York"},
    {"Bob", "30", "Los Angeles"},
}
ui.Table(headers, rows)
```

#### `ProgressBar(current, total int, label string)`
Shows progress with visual blocks.

```go
ui.ProgressBar(5, 10, "Processing")
```

#### `Grid(items []string, columns int)`
Organizes items in a grid layout.

```go
ui.Grid([]string{"A", "B", "C", "D"}, 2)
```

### Containers and Panels

#### `Panel(title string, content func(), style BorderStyle)`
Creates a bordered container for grouping content.

```go
ui.Panel("Configuration", func() {
    ui.TextLine("Server: localhost", guruui.StylePrimary)
    ui.TextLine("Port: 8080", guruui.StylePrimary)
}, guruui.BorderThin)
```

#### `Box(text string, title string)`
Creates a simple bordered box around text.

```go
ui.Box("Content here", "Title")
```

### Utility Methods

#### `Clear()`
Clears the terminal screen.

```go
ui.Clear()
```

#### `Exit(msg string, code int)`
Prints a message and exits the application.

```go
ui.Exit("Goodbye!", 0)
```

## Design Principles

### Color Usage
- **Primary**: Normal text (default terminal color)
- **Muted**: Dimmed text using ANSI dim attribute
- **Accent**: Bold text using ANSI bold attribute
- **No colors**: Only black, white, and gray

### Visual Hierarchy
- **Headers**: Uppercase with underlines
- **Borders**: Unicode box-drawing characters
- **Spacing**: Consistent margins and padding
- **Alignment**: Left, center, and right positioning

### Terminal Compatibility
- Uses standard ANSI escape sequences
- Unicode box-drawing characters for borders
- Fallback to ASCII characters if needed
- Works in most modern terminals

## Best Practices

### Component Organization
1. Use `Header()` for main sections
2. Group related content with `Panel()`
3. Use `Space()` for visual breathing room
4. Apply consistent text styling

### User Experience
1. Provide clear prompts and feedback
2. Use appropriate status message types
3. Validate user input gracefully
4. Show progress for long operations

### Code Style
1. Keep component usage simple and readable
2. Use descriptive variable names
3. Handle errors appropriately
4. Test components in various terminal sizes

## Examples

### Basic Application Structure
```go
func main() {
    ui := guruui.New()
    
    ui.Header("My CLI App")
    ui.Info("Starting up...")
    
    choice := ui.Select("Choose action:", []string{"Start", "Stop", "Status"})
    
    switch choice {
    case "Start":
        ui.Success("Application started")
    case "Stop":
        ui.Info("Stopping application...")
    case "Status":
        ui.Info("Application is running")
    }
}
```

### Data Display
```go
// Show user information
ui.Panel("User Profile", func() {
    ui.TextLine("Name: John Doe", guruui.StylePrimary)
    ui.TextLine("Email: john@example.com", guruui.StylePrimary)
    ui.TextLine("Status: Active", guruui.StyleAccent)
}, guruui.BorderThin)

// Display data table
headers := []string{"Metric", "Value", "Status"}
rows := [][]string{
    {"CPU Usage", "45%", "Normal"},
    {"Memory", "2.1GB", "Warning"},
    {"Disk", "75%", "Normal"},
}
ui.Table(headers, rows)
```

### Interactive Workflow
```go
// Multi-step process
ui.Header("Setup Wizard")
ui.Info("Let's configure your application")

name := ui.Input("Application name")
port := ui.Input("Port number")

if ui.Confirm("Start the application now?") {
    ui.Info("Starting " + name + " on port " + port)
    ui.ProgressBar(1, 1, "Initializing")
    ui.Success("Application is ready!")
}
```
