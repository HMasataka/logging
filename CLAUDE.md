# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a lightweight Go logging library that extends Go's standard `log/slog` package with context-based attribute injection. The library allows adding contextual information to log records through Go's context mechanism, using a thread-safe implementation with `sync.Map`.

## Architecture

### Core Components

- **Context Management (`context.go`)**: Implements `WithValue()` function to add contextual logging information. Uses `sync.Map` for thread-safe value storage and creates immutable context copies.

- **Log Handler (`handler.go`)**: Custom `slog.Handler` implementation that wraps standard handlers and automatically injects context-based attributes into log records.

- **Example Usage (`_example/main.go`)**: Demonstrates the complete workflow of creating a logger with the custom handler and using context-based logging.

### Design Patterns

- **Decorator Pattern**: `LogHandler` wraps existing `slog.Handler` implementations
- **Context Propagation**: Extends Go's context with logging-specific enhancements
- **Immutable Context**: Creates new contexts instead of modifying existing ones
- **Thread-Safe Design**: Uses `sync.Map` for concurrent context value access

## Development Commands

This project uses [Task](https://taskfile.dev/) for build automation. All commands should be run from the repository root.

### Essential Commands

```bash
# List all available tasks
task

# Build the example binary
task build

# Run the example application
task run

# Run tests
task test

# Development workflow (build and run)
task dev
```

### Code Quality Commands

```bash
# Format Go code
task fmt

# Run go vet for static analysis
task vet

# Run all checks (format, vet, test)
task check

# Clean build artifacts
task clean
```

### Module Management

```bash
# Tidy go modules
task mod-tidy
```

### Multi-platform Builds

```bash
# Build for multiple platforms (Linux, Darwin, Windows)
task build-all
```

## Usage Pattern

The typical usage flow involves:

1. Create a logger with the custom handler wrapping a standard `slog.Handler`
2. Add contextual values using `logging.WithValue()`
3. Use standard `slog` functions with the enhanced context

Example:

```go
logger := slog.New(logging.NewHandler(slog.NewJSONHandler(os.Stdout, nil)))
ctx := logging.WithValue(context.Background(), "request_id", "abc123")
slog.ErrorContext(ctx, "Operation failed")
```

## Testing Notes

Currently, the project has minimal test coverage. When adding tests:

- Use standard Go testing patterns with `*_test.go` files
- Test context value propagation and immutability
- Verify thread-safety of context operations
- Test handler attribute injection functionality

## Dependencies

The project has zero external dependencies beyond Go's standard library:

- `context` - for Go context functionality
- `log/slog` - for structured logging
- `sync` - for thread-safe operations

This minimal dependency approach maintains simplicity and reduces potential compatibility issues.

