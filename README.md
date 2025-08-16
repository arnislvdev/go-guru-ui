# GuruUI

A minimalist terminal user interface library for Go. Clean, black-and-white design inspired by modern web aesthetics, optimized for terminal use.

## Features

- **Minimalist Design**: Clean, professional appearance without ASCII art or unnecessary colors
- **Black & White Only**: Uses whitespace, borders, and typography for visual hierarchy
- **Ready-to-Use Components**: Drop-in UI primitives that work out of the box
- **Opinionated Defaults**: Sensible styling that looks good without configuration

## Components

- **Text Styling**: Primary, muted, and accent text variants
- **Status Messages**: Info, Success, Warning, and Error with clear prefixes
- **Interactive Selectors**: Single and multi-choice selection widgets
- **Tables**: Clean borders with alternating row backgrounds
- **Progress Bars**: Simple block-based progress indicators
- **Spinners**: Loading indicators for async operations
- **Panels**: Content grouping with clean borders

## Quick Start

```go
package main

import "github.com/arnislvdev/go-guruui"

func main() {
    ui := guruui.New()
    
    ui.Header("My CLI App")
    ui.Info("Starting up...")
    
    choice := ui.Select("Choose an option:", []string{"Option 1", "Option 2", "Option 3"})
    ui.Success("Selected: " + choice)
    
    ui.Table([]string{"Name", "Value"}, [][]string{
        {"Item 1", "100"},
        {"Item 2", "200"},
    })
}
```

## Installation

```bash
go get github.com/arnislvdev/go-guruui
```

## Design Philosophy

GuruUI follows a strict black-and-white aesthetic where visual hierarchy is achieved through:
- Whitespace and alignment
- Border styles and thickness
- Text weight and emphasis
- Alternating backgrounds

No colors, no emojis, no fluff. Just clean, professional terminal interfaces that developers can rely on.

## License

MIT
