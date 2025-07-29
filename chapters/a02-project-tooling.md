# Project and Tooling Quick Reference

## Purpose and Usage

This appendix provides a comprehensive reference for Go's module system, project organization, essential toolchain commands, and testing framework. Use this as your go-to guide for setting up projects, managing dependencies, and using Go's built-in tools effectively.

**Target Audience:** Developers working through the curriculum who need quick access to Go tooling, project setup, and best practices.

**How to Use:** Reference specific sections when setting up new projects, managing dependencies, or when you need to recall specific command syntax and options.

## Table of Contents

1. [Module System](#module-system)
2. [Project Organization](#project-organization)
3. [Package System and Visibility](#package-system-and-visibility)
4. [Essential Go Commands](#essential-go-commands)
5. [Testing Framework](#testing-framework)
6. [Build and Deployment](#build-and-deployment)
7. [Code Quality Tools](#code-quality-tools)
8. [Development Workflow](#development-workflow)

---

## Module System

Go modules, introduced in Go 1.11, are the standard way to manage dependencies and versioning.

### Module Initialization

```bash
# Initialize a new module
go mod init example.com/myproject

# Initialize with custom module path
go mod init github.com/username/myproject

# Initialize in existing directory with inferred name
go mod init
```

### Module Files

#### go.mod

```go
module github.com/username/myproject

go 1.24

require (
    github.com/gorilla/mux v1.8.0
    github.com/lib/pq v1.10.9
    golang.org/x/crypto v0.14.0
)

require (
    github.com/gorilla/context v1.1.1 // indirect
    golang.org/x/sys v0.13.0 // indirect
)

replace github.com/old/package => github.com/new/package v1.2.3

exclude github.com/problematic/package v1.0.0

retract v1.0.1 // Published accidentally
```

#### go.sum

```
github.com/gorilla/mux v1.8.0 h1:i40aqfkR1h2SlN9hojwV5ZA91wcXFOvtP6o09YNMjJjU=
github.com/gorilla/mux v1.8.0/go.mod h1:DVb2xZr1YK5FkZOFEhHzHlxD2h5IfmOx2gJOUvjvjKo=
```

### Dependency Management

```bash
# Add dependency (automatically updates go.mod)
go get github.com/gorilla/mux

# Add specific version
go get github.com/gorilla/mux@v1.8.0

# Add latest pre-release
go get github.com/gorilla/mux@latest

# Add from specific commit
go get github.com/gorilla/mux@commit-hash

# Add from specific branch
go get github.com/gorilla/mux@branch-name

# Update all dependencies
go get -u ./...

# Update direct dependencies only
go get -u

# Remove unused dependencies
go mod tidy

# Download dependencies without building
go mod download

# Vendor dependencies (copy to vendor/)
go mod vendor

# Verify dependencies
go mod verify

# Show module graph
go mod graph

# Show why a dependency is needed
go mod why github.com/gorilla/mux

# Edit go.mod programmatically
go mod edit -require github.com/new/dep@v1.0.0
go mod edit -replace github.com/old/dep=github.com/new/dep@v1.0.0
go mod edit -exclude github.com/bad/dep@v1.0.0
```

### Version Selection

```bash
# List available versions
go list -m -versions github.com/gorilla/mux

# Show current version info
go list -m github.com/gorilla/mux

# Show all dependencies with versions
go list -m all

# Upgrade to latest patch version
go get github.com/gorilla/mux@patch

# Upgrade to latest minor version
go get github.com/gorilla/mux@minor

# Downgrade to specific version
go get github.com/gorilla/mux@v1.7.4
```

---

## Project Organization

### Standard Project Layout

```
myproject/
├── go.mod                      # Module definition
├── go.sum                      # Dependency checksums
├── README.md                   # Project documentation
├── LICENSE                     # License file
├── Makefile                    # Build automation
├── .gitignore                  # Git ignore patterns
├── .dockerignore               # Docker ignore patterns
├── Dockerfile                  # Container definition
├── docker-compose.yml          # Multi-container setup
│
├── cmd/                        # Main applications
│   ├── server/                 # Server application
│   │   └── main.go
│   ├── cli/                    # CLI application
│   │   └── main.go
│   └── worker/                 # Background worker
│       └── main.go
│
├── internal/                   # Private packages
│   ├── config/                 # Configuration
│   │   ├── config.go
│   │   └── config_test.go
│   ├── database/               # Database layer
│   │   ├── database.go
│   │   ├── migrations/
│   │   └── queries/
│   ├── handlers/               # HTTP handlers
│   │   ├── auth.go
│   │   ├── users.go
│   │   └── handlers_test.go
│   └── models/                 # Data models
│       ├── user.go
│       └── models_test.go
│
├── pkg/                        # Public packages (reusable)
│   ├── api/                    # API client
│   │   ├── client.go
│   │   └── client_test.go
│   └── utils/                  # Utility functions
│       ├── crypto.go
│       └── crypto_test.go
│
├── api/                        # API definitions
│   ├── openapi.yaml            # OpenAPI specification
│   └── proto/                  # Protocol buffer definitions
│       └── service.proto
│
├── web/                        # Web assets
│   ├── static/                 # Static files
│   │   ├── css/
│   │   ├── js/
│   │   └── images/
│   └── templates/              # HTML templates
│       ├── base.html
│       └── index.html
│
├── scripts/                    # Build and deployment scripts
│   ├── build.sh
│   ├── deploy.sh
│   └── setup.sh
│
├── deployments/                # Deployment configurations
│   ├── docker/
│   ├── kubernetes/
│   │   ├── deployment.yaml
│   │   └── service.yaml
│   └── terraform/
│
├── docs/                       # Documentation
│   ├── architecture.md
│   ├── api.md
│   └── deployment.md
│
├── examples/                   # Example code
│   └── simple/
│       └── main.go
│
├── test/                       # Integration tests
│   ├── integration/
│   └── fixtures/
│       └── sample_data.json
│
└── vendor/                     # Vendored dependencies (optional)
    └── github.com/
```

### Package Naming Conventions

```go
// Good package names (short, clear, lowercase)
package user
package http
package crypto
package json

// Avoid these patterns
package userManager    // camelCase
package user_manager   // snake_case
package UserManager    // PascalCase
package users          // plural (usually)

// Package with multiple words
package httputil       // No separator
package nethttp        // Abbreviate when clear

// Avoid stuttering
package user
func (u User) UserName() string  // Bad: user.UserName()
func (u User) Name() string      // Good: user.Name()
```

---

## Package System and Visibility

### Package Declaration and Imports

```go
// Package declaration (first non-comment line)
package main      // Executable package
package user      // Library package
package userutil  // Utility package

// Import statements
import "fmt"                    // Standard library
import "net/http"              // Standard library subpackage

// Grouped imports (preferred style)
import (
    // Standard library first
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"

    // Third-party packages second
    "github.com/gorilla/mux"
    "github.com/lib/pq"
    "golang.org/x/crypto/bcrypt"

    // Local packages last
    "github.com/mycompany/myproject/internal/config"
    "github.com/mycompany/myproject/internal/database"
)

// Import variations
import (
    "fmt"                                    // Standard import
    nethttp "net/http"                      // Aliased import
    . "strings"                             // Dot import (discouraged)
    _ "github.com/lib/pq"                   // Blank import (side effects)
)
```

### Visibility Rules

Go uses capitalization to determine visibility:

```go
package user

// Exported (public) - starts with capital letter
type User struct {
    ID    int    // Exported field
    Name  string // Exported field
    email string // Unexported field (private)
}

// Exported function
func NewUser(name, email string) *User {
    return &User{
        Name:  name,
        email: email,
    }
}

// Exported method
func (u *User) GetEmail() string {
    return u.email
}

// Unexported function (package-private)
func validateEmail(email string) bool {
    // Validation logic
    return strings.Contains(email, "@")
}

// Exported constant
const MaxNameLength = 100

// Unexported constant
const defaultTimeout = 30 * time.Second

// Exported variable (rare, usually avoid)
var DefaultUser = &User{Name: "Anonymous"}

// Unexported variable
var userCount int
```

### Package Initialization

```go
package database

import (
    "database/sql"
    "log"
    _ "github.com/lib/pq"  // Driver registration
)

// Package-level variables
var (
    db        *sql.DB
    connected bool
)

// init functions run automatically before main()
func init() {
    log.Println("Initializing database package")
    
    var err error
    db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    
    connected = true
}

// Multiple init functions allowed (run in declaration order)
func init() {
    if connected {
        log.Println("Database connection established")
    }
}

// Exported function using package-level variables
func GetDB() *sql.DB {
    if !connected {
        log.Fatal("Database not initialized")
    }
    return db
}
```

---

## Essential Go Commands

### Building and Running

```bash
# Run Go program directly
go run main.go
go run *.go
go run ./cmd/server

# Build executable
go build                        # Build in current directory
go build -o myapp              # Build with custom name
go build ./cmd/server          # Build specific package
go build -o bin/server ./cmd/server

# Install executable to $GOPATH/bin
go install
go install ./cmd/server

# Cross-compilation
GOOS=linux GOARCH=amd64 go build -o myapp-linux
GOOS=windows GOARCH=amd64 go build -o myapp.exe
GOOS=darwin GOARCH=arm64 go build -o myapp-mac

# Build with flags
go build -ldflags "-X main.version=1.0.0"
go build -ldflags "-s -w"      # Strip debug info
go build -tags production       # Build with tags
```

### Testing

```bash
# Run tests
go test                         # Current package
go test ./...                   # All packages recursively
go test -v                      # Verbose output
go test -short                  # Skip long-running tests
go test -timeout 30s            # Set timeout

# Test specific functions/packages
go test -run TestUserCreate     # Run specific test
go test -run TestUser           # Run tests matching pattern
go test ./internal/user         # Test specific package

# Coverage
go test -cover                  # Show coverage percentage
go test -coverprofile=cover.out # Generate coverage profile
go tool cover -html=cover.out   # View coverage in browser
go test -covermode=atomic       # Thread-safe coverage

# Benchmarks
go test -bench=.                # Run all benchmarks
go test -bench=BenchmarkSort    # Run specific benchmark
go test -benchmem               # Include memory stats
go test -cpuprofile=cpu.prof    # Generate CPU profile

# Race detection
go test -race                   # Enable race detector
go run -race main.go           # Run with race detection

# Test caching
go clean -testcache            # Clear test cache
go test -count=1               # Disable test caching
```

### Code Quality and Analysis

```bash
# Format code
go fmt ./...                   # Format all packages
gofmt -w .                     # Format with gofmt directly
gofmt -s -w .                  # Simplify code while formatting

# Vet (examine code for issues)
go vet ./...                   # Check all packages
go vet -shadow ./...           # Check for shadowed variables

# Documentation
go doc fmt.Println             # Show documentation
go doc -all fmt                # Show all documentation for package
godoc -http=:6060              # Start documentation server

# List packages and dependencies
go list ./...                  # List all packages
go list -m all                 # List all modules
go list -deps ./...            # List dependencies
go list -json ./...            # Output in JSON format
```

### Module and Dependency Commands

```bash
# Module management
go mod init myproject          # Initialize module
go mod tidy                    # Add missing, remove unused
go mod download                # Download dependencies
go mod vendor                  # Copy dependencies to vendor/
go mod verify                  # Verify dependencies
go mod graph                   # Print module requirement graph
go mod why github.com/pkg/errors # Explain why module is needed

# Environment and version info
go version                     # Show Go version
go env                         # Show Go environment
go env GOPATH                  # Show specific env var
go env -w GOPROXY=direct       # Set environment variable
```

### Cleaning and Maintenance

```bash
# Clean build artifacts
go clean                       # Clean current package
go clean -cache                # Clean build cache
go clean -testcache            # Clean test cache
go clean -modcache             # Clean module cache
go clean -i                    # Clean installed packages

# Generate code
go generate ./...              # Run go:generate directives

# Get help
go help                        # General help
go help build                  # Help for specific command
go help buildmode              # Help for build modes
```

---

## Testing Framework

### Basic Test Structure

```go
// user_test.go
package user

import (
    "testing"
)

// Test function must start with Test
func TestCreateUser(t *testing.T) {
    user := NewUser("Alice", "alice@example.com")
    
    if user.Name != "Alice" {
        t.Errorf("Expected name 'Alice', got '%s'", user.Name)
    }
    
    if user.GetEmail() != "alice@example.com" {
        t.Errorf("Expected email 'alice@example.com', got '%s'", user.GetEmail())
    }
}

// Table-driven tests (idiomatic Go pattern)
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name  string
        email string
        want  bool
    }{
        {"valid email", "user@example.com", true},
        {"empty email", "", false},
        {"no @ symbol", "userexample.com", false},
        {"no domain", "user@", false},
        {"no username", "@example.com", false},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := validateEmail(tt.email)
            if got != tt.want {
                t.Errorf("validateEmail(%q) = %v, want %v", tt.email, got, tt.want)
            }
        })
    }
}

// Subtest example
func TestUserOperations(t *testing.T) {
    user := NewUser("Bob", "bob@example.com")
    
    t.Run("GetName", func(t *testing.T) {
        if got := user.Name; got != "Bob" {
            t.Errorf("Name = %q, want %q", got, "Bob")
        }
    })
    
    t.Run("GetEmail", func(t *testing.T) {
        if got := user.GetEmail(); got != "bob@example.com" {
            t.Errorf("GetEmail() = %q, want %q", got, "bob@example.com")
        }
    })
}
```

### Test Helpers and Setup

```go
// Test helper functions
func TestMain(m *testing.M) {
    // Setup before all tests
    setupDatabase()
    
    // Run tests
    code := m.Run()
    
    // Cleanup after all tests
    teardownDatabase()
    
    os.Exit(code)
}

func setupTest(t *testing.T) *Database {
    t.Helper()  // Mark as helper function
    
    db, err := NewTestDatabase()
    if err != nil {
        t.Fatalf("Failed to setup test database: %v", err)
    }
    
    // Cleanup when test finishes
    t.Cleanup(func() {
        db.Close()
    })
    
    return db
}

func TestWithDatabase(t *testing.T) {
    db := setupTest(t)
    
    // Use db in test
    user := &User{Name: "Alice"}
    err := db.SaveUser(user)
    if err != nil {
        t.Fatalf("Failed to save user: %v", err)
    }
}

// Skip tests conditionally
func TestExpensiveOperation(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping expensive test in short mode")
    }
    
    // Long-running test
}

// Parallel tests
func TestParallel1(t *testing.T) {
    t.Parallel()
    // This test runs in parallel with other parallel tests
}

func TestParallel2(t *testing.T) {
    t.Parallel()
    // This test runs in parallel with other parallel tests
}
```

### Benchmarks

```go
// Benchmark function must start with Benchmark
func BenchmarkStringConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result := "hello" + "world"
        _ = result  // Prevent optimization
    }
}

// Benchmark with setup
func BenchmarkMapLookup(b *testing.B) {
    m := make(map[string]int)
    for i := 0; i < 1000; i++ {
        m[fmt.Sprintf("key%d", i)] = i
    }
    
    b.ResetTimer()  // Don't count setup time
    
    for i := 0; i < b.N; i++ {
        _ = m["key500"]
    }
}

// Benchmark with different inputs
func BenchmarkSort(b *testing.B) {
    sizes := []int{100, 1000, 10000}
    
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size_%d", size), func(b *testing.B) {
            for i := 0; i < b.N; i++ {
                b.StopTimer()
                data := generateRandomData(size)
                b.StartTimer()
                
                sort.Ints(data)
            }
        })
    }
}

// Memory allocation tracking
func BenchmarkMemAlloc(b *testing.B) {
    b.ReportAllocs()  // Report memory allocations
    
    for i := 0; i < b.N; i++ {
        slice := make([]int, 1000)
        _ = slice
    }
}
```

### Example Tests

```go
// Example function for documentation
func ExampleUser_GetEmail() {
    user := NewUser("Alice", "alice@example.com")
    fmt.Println(user.GetEmail())
    // Output: alice@example.com
}

// Example with multiple outputs
func ExampleValidateEmail() {
    fmt.Println(validateEmail("user@example.com"))
    fmt.Println(validateEmail("invalid-email"))
    // Output:
    // true
    // false
}

// Example with unordered output
func ExampleMap() {
    m := map[string]int{"b": 2, "a": 1}
    for k, v := range m {
        fmt.Printf("%s: %d\n", k, v)
    }
    // Unordered output:
    // a: 1
    // b: 2
}
```

### Test Utilities

```go
// Testing HTTP handlers
func TestHTTPHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/users/123", nil)
    rec := httptest.NewRecorder()
    
    handler := http.HandlerFunc(getUserHandler)
    handler.ServeHTTP(rec, req)
    
    if rec.Code != http.StatusOK {
        t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
    }
    
    expected := `{"name":"Alice"}`
    if rec.Body.String() != expected {
        t.Errorf("Expected body %q, got %q", expected, rec.Body.String())
    }
}

// Temporary files and directories
func TestFileOperations(t *testing.T) {
    // Create temporary directory
    tmpDir := t.TempDir()
    
    // Create temporary file
    tmpFile, err := os.CreateTemp(tmpDir, "test-*.txt")
    if err != nil {
        t.Fatalf("Failed to create temp file: %v", err)
    }
    defer tmpFile.Close()
    
    // Use temporary file in test
    content := "test content"
    if _, err := tmpFile.WriteString(content); err != nil {
        t.Fatalf("Failed to write to temp file: %v", err)
    }
}

// Golden files (expected output comparison)
func TestGoldenFile(t *testing.T) {
    output := generateOutput()
    
    goldenFile := "testdata/expected_output.golden"
    
    if *update {
        // Update golden file
        os.WriteFile(goldenFile, output, 0644)
        return
    }
    
    expected, err := os.ReadFile(goldenFile)
    if err != nil {
        t.Fatalf("Failed to read golden file: %v", err)
    }
    
    if !bytes.Equal(output, expected) {
        t.Errorf("Output doesn't match golden file")
        t.Logf("Expected:\n%s", expected)
        t.Logf("Got:\n%s", output)
    }
}
```

---

## Build and Deployment

### Build Options

```bash
# Basic build
go build -o myapp

# Optimized build (smaller binary)
go build -ldflags="-s -w" -o myapp

# Static linking (no external dependencies)
CGO_ENABLED=0 go build -a -ldflags="-s -w" -o myapp

# Inject version information
go build -ldflags="-X main.version=1.0.0 -X main.buildTime=$(date +%Y-%m-%d_%H:%M:%S)"

# Cross-platform builds
GOOS=linux GOARCH=amd64 go build -o myapp-linux
GOOS=windows GOARCH=amd64 go build -o myapp.exe
GOOS=darwin GOARCH=arm64 go build -o myapp-mac-arm64

# Build with specific tags
go build -tags "production,mysql" -o myapp
```

### Build Tags

```go
// +build linux
// build constraint (old syntax)

//go:build linux
// build constraint (new syntax, Go 1.17+)

//go:build linux && amd64

//go:build !windows

//go:build production || development

// File: config_dev.go
//go:build !production

package config

const DatabaseURL = "localhost:5432"

// File: config_prod.go
//go:build production

package config

const DatabaseURL = "prod-db:5432"
```

### Docker Integration

```dockerfile
# Multi-stage Docker build
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

### Makefile Example

```makefile
.PHONY: build test clean run docker

# Variables
BINARY_NAME=myapp
VERSION=$(shell git describe --tags --always --dirty)
BUILD_TIME=$(shell date +%Y-%m-%d_%H:%M:%S)
LDFLAGS=-ldflags="-X main.version=${VERSION} -X main.buildTime=${BUILD_TIME}"

# Default target
all: test build

# Build the application
build:
	go build ${LDFLAGS} -o bin/${BINARY_NAME} .

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run the application
run:
	go run ${LDFLAGS} .

# Clean build artifacts
clean:
	go clean
	rm -rf bin/
	rm -f coverage.out coverage.html

# Build for multiple platforms
build-all:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-windows-amd64.exe .
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY_NAME}-darwin-amd64 .

# Docker build
docker:
	docker build -t ${BINARY_NAME}:${VERSION} .

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run linters
lint:
	go fmt ./...
	go vet ./...
	golint ./...

# Development server with auto-reload
dev:
	air -c .air.toml
```

---

## Code Quality Tools

### Formatting and Linting

```bash
# Format code
go fmt ./...
gofmt -w .
goimports -w .  # Also manages imports

# Vet (static analysis)
go vet ./...
go vet -shadow ./...

# External linters (install separately)
golint ./...                    # Style checker
golangci-lint run              # Meta-linter (recommended)
staticcheck ./...              # Advanced static analysis

# golangci-lint configuration (.golangci.yml)
```

```yaml
# .golangci.yml
run:
  timeout: 5m
  tests: true

linters-settings:
  govet:
    check-shadowing: true
  golint:
    min-confidence: 0.8
  gocyclo:
    min-complexity: 10

linters:
  enable:
    - errcheck
    - gosimple
    - govet
    - ineffassign
    - staticcheck
    - typecheck
    - unused
    - varcheck
    - deadcode
    - gosec
    - misspell

issues:
  exclude-rules:
    - path: _test\.go
      linters:
        - gosec
```

### Security Analysis

```bash
# Security scanner
gosec ./...

# Check for vulnerabilities
go list -json -m all | nancy sleuth

# Dependency scanning
govulncheck ./...
```

---

## Development Workflow

### Git Hooks Integration

```bash
# .git/hooks/pre-commit
#!/bin/sh
go fmt ./...
go vet ./...
go test -short ./...
```

### Air (Live Reload)

```toml
# .air.toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
args_bin = []
bin = "tmp/main"
cmd = "go build -o ./tmp/main ."
delay = 1000
exclude_dir = ["assets", "tmp", "vendor", "testdata"]
exclude_file = []
exclude_regex = ["_test.go"]
exclude_unchanged = false
follow_symlink = false
full_bin = ""
include_dir = []
include_ext = ["go", "tpl", "tmpl", "html"]
include_file = []
kill_delay = 0
log = "build-errors.log"
send_interrupt = false
stop_on_root = false

[color]
app = ""
build = "yellow"
main = "magenta"
runner = "green"
watcher = "cyan"

[log]
time = false

[misc]
clean_on_exit = false
```

### VSCode Configuration

```json
// .vscode/settings.json
{
    "go.formatTool": "goimports",
    "go.lintTool": "golangci-lint",
    "go.lintOnSave": "package",
    "go.vetOnSave": "package",
    "go.buildOnSave": "off",
    "go.testOnSave": false,
    "go.coverOnSave": false,
    "go.useLanguageServer": true,
    "gopls": {
        "usePlaceholders": true,
        "staticcheck": true
    }
}

// .vscode/launch.json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",
            "env": {},
            "args": []
        }
    ]
}
```

### Environment Setup

```bash
# .env file for development
DATABASE_URL=postgres://user:password@localhost/myapp_dev
REDIS_URL=redis://localhost:6379
LOG_LEVEL=debug
PORT=8080

# Load environment variables
export $(cat .env | xargs)
```

This comprehensive reference covers all the essential aspects of Go project organization, tooling, and development workflow. Use it as your go-to resource for setting up projects, managing dependencies, and maintaining code quality.