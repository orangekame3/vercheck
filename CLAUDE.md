# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

vercheck is a lightweight Go library that helps CLI tools check for newer versions on GitHub and notify users about available updates.

## Commands

```bash
# Run the example
go run example/main.go

# Run tests
make test

# Run tests with coverage
make coverage

# Run linter
make lint

# Format code
make fmt

# Run all CI checks
make ci

# Build the example
make build

# Update dependencies
make update

# Install development tools
make dev-deps

# Create a new release
make release VERSION=v1.0.0
```

## Architecture

The library consists of the following components:

1. **vercheck.go** - Main public API with the `Check()` function
2. **github.go** - GitHub API client for fetching latest releases
3. **detect.go** - Installation method detection and update command generation
4. **internal/semver.go** - Semantic version comparison logic

### Key Design Decisions

- Zero external dependencies - uses only Go standard library (no go.sum file needed)
- Graceful error handling - never disrupts the host CLI tool
- Simple API - single function call with options struct
- Automatic installation method detection (Homebrew vs go install)

## Development Guidelines

- Keep the API simple and minimal
- Maintain zero external dependencies
- All errors should be handled gracefully without panicking
- Version comparison should handle common semver formats (v1.2.3, 1.2.3, 1.2.3-beta)
- The library should never make the host CLI tool fail

## Testing

When testing changes:
1. Use the example program to verify basic functionality
2. Test with different version formats (v1.2.3, 1.2.3)
3. Test error cases (network failures, invalid GitHub repos)
4. Test both Homebrew and go install detection paths