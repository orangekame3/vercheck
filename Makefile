.PHONY: all test lint fmt vet clean coverage build

# Default target
all: test lint

# Run tests
test:
	go test -v -race ./...

# Run tests with coverage
coverage:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html

# Run linter
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...
	gofumpt -w .

# Run go vet
vet:
	go vet ./...

# Build example
build:
	cd example && go build -v .

# Clean build artifacts
clean:
	rm -f coverage.txt coverage.html
	rm -f example/example

# Install development dependencies
dev-deps:
	go install mvdan.cc/gofumpt@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run all checks (used in CI)
ci: fmt vet test lint fmt-check

# Check if code is formatted
fmt-check:
	@if [ -n "$$(go fmt ./...)" ]; then \
		echo "Code is not formatted. Run 'make fmt' to fix."; \
		exit 1; \
	fi

# Update dependencies
update:
	go get -u ./...
	go mod tidy

# Verify go.mod is clean
mod-verify:
	go mod tidy
	@if [ -f go.sum ]; then \
		git diff --exit-code go.mod go.sum; \
	else \
		git diff --exit-code go.mod; \
	fi

# Create a new release tag
release:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make release VERSION=v1.0.0"; \
		exit 1; \
	fi
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)