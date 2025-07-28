# vercheck

A lightweight Go library that helps CLI tools check for newer versions on GitHub and notify users about available updates.

## Features

- 🚀 Check GitHub releases for newer versions
- 📦 Automatic detection of installation method (Homebrew or go install)
- 🔔 User-friendly update notifications
- 🛠 Zero external dependencies (uses only Go standard library)
- ⚡ Lightweight and fast

## Installation

```bash
go get github.com/orangekame3/vercheck
```

## Usage

Add vercheck to your CLI tool's main function:

```go
import "github.com/orangekame3/vercheck"

func main() {
    vercheck.Check(vercheck.Options{
        CurrentVersion: "v1.2.3",
        RepoOwner: "orangekame3",
        RepoName: "lazypoetry",
    })
    
    // Your CLI tool logic here...
}
```

Example output when a newer version is available:
```
🔔 New version v1.3.0 is available! You're using v1.2.3.
👉 Update with: go install github.com/orangekame3/lazypoetry@latest
```

## API Reference

### Check(options Options)

The main function that checks for updates and displays notifications. Uses a 10-second timeout by default.

### CheckWithContext(ctx context.Context, options Options)

Context-aware version check that allows cancellation and custom timeouts.

### Options struct

```go
type Options struct {
    CurrentVersion string // Required: Current CLI version (e.g., "v1.2.3")
    RepoOwner      string // Required: GitHub repository owner
    RepoName       string // Required: GitHub repository name
    UpdateCommand  string // Optional: Custom update command (auto-detected if empty)
    Silent         bool   // Optional: Suppress error logging (default: false)
}
```

## Installation Method Detection

vercheck automatically detects how your CLI tool was installed and suggests the appropriate update command:

| Installation Method | Detection | Update Command |
|-------------------|-----------|----------------|
| Homebrew | Path contains `/Cellar/` | `brew upgrade <tool>` |
| go install | Default | `go install github.com/<owner>/<tool>@latest` |

## How It Works

1. Fetches the latest release tag from GitHub API (`/repos/{owner}/{repo}/releases/latest`)
2. Compares the latest version with the current version using semantic versioning
3. If a newer version exists, displays a notification with the appropriate update command
4. Gracefully handles errors without disrupting your CLI tool's operation

## Example

See the [example](example/main.go) directory for a complete example.

## Development

### Prerequisites

- Go 1.21 or higher
- golangci-lint (for linting)

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make coverage

# Run linting
make lint

# Run all CI checks
make ci
```

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## CI/CD

This project uses GitHub Actions for continuous integration:

- **Test**: Runs on multiple OS (Linux, macOS, Windows) and Go versions (1.21, 1.22, 1.23)
- **Lint**: Enforces code quality with golangci-lint
- **Coverage**: Reports test coverage to Codecov
- **Release**: Automatically creates GitHub releases for tags

## License

MIT License - see LICENSE file for details.