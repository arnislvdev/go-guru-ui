# GuruUI 🧘‍♂️

> A simple tool that explains coding errors in plain English and turns everyday words into computer commands

GuruUI is a simple tool written in Go that helps you understand programming errors and turns normal words into computer commands. It has two ways of talking to you: **Professional Mode** (clear and simple) and **WTF Mode** (fun and funny).

## What It Does

- **Explains Errors**: Makes coding errors easy to understand
- **Turns Words Into Commands**: Changes normal words into computer commands
- **Two Ways to Talk**: Professional (serious) and WTF (funny) modes
- **Uses AI**: Gets help from OpenAI to understand what you need
- **Works Everywhere**: Works on Windows, Mac, and Linux
- **Fast**: Built with Go so it's quick

## Installation

### What You Need

- Go 1.21 or newer
- An OpenAI API key (to get one, go to openai.com)

### Build from Source

```bash
# Get the code
git clone https://github.com/arnislvdev/go-guru-ui.git
cd go-guru-ui

# Get the parts it needs
go mod tidy

# Make the program
go build -o guruui cmd/guruui/main.go

# Put it where your computer can find it (optional)
sudo mv guruui /usr/local/bin/
```

### Setup

Tell the program your OpenAI API key:

```bash
guruui config set --ai-provider openai --api-key <your-api-key>
```

## How to Use

### Understanding Errors

```bash
# Explain a Go error
guruui explain "undefined: fmt"

# Explain with file info
guruui explain --file main.go --line 42 "cannot use nil as type string"

# Use funny mode
guruui --mode wtf explain "imported and not used"
```

### Turning Words Into Commands

```bash
# Turn normal words into computer commands
guruui translate "how do I check disk space"

# Add extra info for better results
guruui translate --context "I'm on Ubuntu" "install nginx"

# Use funny mode
guruui --mode wtf translate "check network connections"
```

### Changing Settings

```bash
# See current settings
guruui config show

# Set AI provider and API key
guruui config set --ai-provider openai --api-key <key>

# Go back to default settings
guruui config reset
```

## Examples

### Understanding Errors

**What you type:**
```bash
guruui explain "undefined: fmt"
```

**Serious mode answer:**
```
The error "undefined: fmt" means you're trying to use the 'fmt' package 
but haven't imported it. 

To fix this, add the following import statement at the top of your Go file:
import "fmt"
```

**Funny mode answer:**
```
🤦‍♂️ Oh boy, here we go again...

The error "undefined: fmt" means you're trying to use the 'fmt' package 
but haven't imported it. 

To fix this, add the following import statement at the top of your Go file:
import "fmt"

🎯 Pro tip: If you can't see it, it probably doesn't exist. Schrödinger's variable!
```

### Turning Words Into Commands

**What you type:**
```bash
guruui translate "check disk space on my system"
```

**What you get:**
```
Command: df -h

Explanation: The 'df' command displays disk space usage. The '-h' flag shows 
the output in human-readable format (GB, MB, etc.).
```

## How It's Built

```
cmd/guruui/           # Where the program starts
├── main.go          # Main program file
internal/             # Inside program code
├── cli/             # Command definitions
├── domain/          # Data structures
├── usecase/         # Main program logic
└── infrastructure/  # Connects to outside services (AI, storage)
pkg/                  # Public code
└── humor/           # Funny mode responses
```

## For Developers

### How the Code is Organized

- **Clean Structure**: Code is organized in clear, separate parts
- **Easy Testing**: Built so it's easy to test and fix
- **Flexible Design**: Can easily swap out AI providers and storage
- **Good Error Handling**: Gives helpful error messages

### Adding New Things

1. **New Error Types**: Add to `internal/domain/error.go`
2. **AI Prompts**: Change `internal/infrastructure/ai/openai.go`
3. **Funny Responses**: Add to `pkg/humor/wtf_mode.go`
4. **New Commands**: Create new files in `internal/cli/`

### Testing

```bash
# Test everything
go test ./...

# Test with coverage report
go test -cover ./...

# Test just one part
go test ./internal/usecase/
```

## Future Plans

### What's Done Now
- [x] Basic command structure
- [x] OpenAI connection for explaining errors
- [x] Simple error understanding for Go
- [x] Funny mode
- [x] Settings management

### Next Steps
- [ ] Support for Python, JavaScript, Rust errors
- [ ] Better error understanding
- [ ] Support for other AI services
- [ ] Plugin system for custom features

### Long Term
- [ ] Cloud service for teams
- [ ] Business features
- [ ] VS Code extension
- [ ] Web interface

## Want to Help?

1. Copy the project to your account
2. Make a new branch for your changes
3. Save your changes
4. Send your changes back
5. Ask to merge your changes

## License

This project is free to use under the MIT License - see the [LICENSE](LICENSE) file for details.

## Get Help

- 📖 [Documentation](https://docs.guruui.com)
- 🐛 [Report Problems](https://github.com/arnislvdev/go-guru-ui/issues)
- 💬 [Ask Questions](https://github.com/arnislvdev/go-guru-ui/discussions)

---

Made with ❤️ by Arnis
