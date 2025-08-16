# GuruUI 🧘‍♂️

> AI-powered error explanation and command translation for developers

GuruUI is a Go-based developer tool that explains programming errors in plain English and translates natural language into CLI commands. It features two modes: **Professional Mode** for clear, concise explanations, and **WTF Mode** for the same explanations with humor, sarcasm, and memes.

## Features

- 🚀 **Error Explanation**: Understand compilation and runtime errors in plain English
- 🔄 **Command Translation**: Convert natural language to executable CLI commands
- 🎭 **Dual Modes**: Professional (clear) and WTF (humorous) output modes
- 🤖 **AI-Powered**: Uses OpenAI GPT-4 for intelligent error analysis
- 🎯 **Multi-Platform**: Supports Linux, macOS, and Windows
- ⚡ **Fast**: Written in Go for optimal performance

## Installation

### Prerequisites

- Go 1.21 or later
- OpenAI API key

### Build from Source

```bash
# Clone the repository
git clone https://github.com/guruui/guruui.git
cd guruui

# Install dependencies
go mod tidy

# Build the binary
go build -o guruui cmd/guruui/main.go

# Move to your PATH (optional)
sudo mv guruui /usr/local/bin/
```

### Configuration

Set up your OpenAI API key:

```bash
guruui config set --ai-provider openai --api-key <your-api-key>
```

## Usage

### Error Explanation

```bash
# Explain a Go compilation error
guruui explain "undefined: fmt"

# Explain with file context
guruui explain --file main.go --line 42 "cannot use nil as type string"

# Use WTF mode for humorous explanations
guruui --mode wtf explain "imported and not used"
```

### Command Translation

```bash
# Translate natural language to CLI commands
guruui translate "how do I check disk space"

# Add context for better results
guruui translate --context "I'm on Ubuntu" "install nginx"

# WTF mode for fun translations
guruui --mode wtf translate "check network connections"
```

### Configuration Management

```bash
# Show current configuration
guruui config show

# Set AI provider and API key
guruui config set --ai-provider openai --api-key <key>

# Reset to defaults
guruui config reset
```

## Examples

### Error Explanation

**Input:**
```bash
guruui explain "undefined: fmt"
```

**Professional Mode Output:**
```
The error "undefined: fmt" means you're trying to use the 'fmt' package 
but haven't imported it. 

To fix this, add the following import statement at the top of your Go file:
import "fmt"
```

**WTF Mode Output:**
```
🤦‍♂️ Oh boy, here we go again...

The error "undefined: fmt" means you're trying to use the 'fmt' package 
but haven't imported it. 

To fix this, add the following import statement at the top of your Go file:
import "fmt"

🎯 Pro tip: If you can't see it, it probably doesn't exist. Schrödinger's variable!
```

### Command Translation

**Input:**
```bash
guruui translate "check disk space on my system"
```

**Output:**
```
Command: df -h

Explanation: The 'df' command displays disk space usage. The '-h' flag shows 
the output in human-readable format (GB, MB, etc.).
```

## Architecture

```
cmd/guruui/           # CLI entry point
├── main.go          # Main application
internal/             # Internal application code
├── cli/             # CLI command definitions
├── domain/          # Domain models and entities
├── usecase/         # Business logic and use cases
└── infrastructure/  # External integrations (AI, storage)
pkg/                  # Public packages
└── humor/           # WTF mode humor enhancement
```

## Development

### Project Structure

- **Clean Architecture**: Separation of concerns with clear layers
- **Dependency Injection**: Easy to test and maintain
- **Interface-based Design**: Pluggable AI providers and storage
- **Error Handling**: Comprehensive error handling with user-friendly messages

### Adding New Features

1. **New Error Types**: Add to `internal/domain/error.go`
2. **AI Prompts**: Modify `internal/infrastructure/ai/openai.go`
3. **Humor Responses**: Extend `pkg/humor/wtf_mode.go`
4. **CLI Commands**: Create new files in `internal/cli/`

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific package tests
go test ./internal/usecase/
```

## Roadmap

### MVP (Current)
- [x] Basic CLI structure with Cobra
- [x] OpenAI integration for error explanation
- [x] Simple error parsing for Go
- [x] WTF mode with humor
- [x] Configuration management

### Phase 2
- [ ] Support for Python, JavaScript, Rust errors
- [ ] Advanced error parsing with AST
- [ ] Multiple AI provider support (Anthropic, local models)
- [ ] Plugin system for custom integrations

### Phase 3
- [ ] Cloud service with team collaboration
- [ ] Enterprise features (SSO, compliance)
- [ ] VS Code extension
- [ ] Web dashboard

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📖 [Documentation](https://docs.guruui.com)
- 🐛 [Issue Tracker](https://github.com/guruui/guruui/issues)
- 💬 [Discussions](https://github.com/guruui/guruui/discussions)
- 🐦 [Twitter](https://twitter.com/guruui)

---

Made with ❤️ by the GuruUI team
