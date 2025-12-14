# claude-code-sandbox

[![Go CI](https://github.com/michael-freling/claude-code-sandbox/actions/workflows/go-ci.yml/badge.svg)](https://github.com/michael-freling/claude-code-sandbox/actions/workflows/go-ci.yml)

A sandbox repository demonstrating Go best practices with CI/CD automation.

## Features

- Simple calculator library demonstrating Go package structure
- Comprehensive unit tests with table-driven patterns
- GitHub Actions CI pipeline for automated testing and linting
- golangci-lint integration for code quality

## Prerequisites

- Go 1.25 or higher

## Installation

Clone the repository and download dependencies:

```bash
git clone https://github.com/michael-freling/claude-code-sandbox.git
cd claude-code-sandbox
go mod download
```

## Usage

### Running the Demo Application

```bash
go run ./cmd/calculator
```

### Running Tests

```bash
go test ./...
```

### Running Tests with Coverage

```bash
go test -cover ./...
```

### Running Linter

```bash
golangci-lint run
```

## Project Structure

```
.
├── cmd/
│   └── calculator/          # Main application entry point
│       └── main.go
├── internal/
│   └── calculator/          # Calculator package
│       ├── calculator.go
│       └── calculator_test.go
├── .github/
│   └── workflows/
│       └── go-ci.yml        # GitHub Actions CI pipeline
├── .golangci.yml            # golangci-lint configuration
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## CI/CD

The project uses GitHub Actions for continuous integration, which automatically:

- Runs all unit tests to ensure code correctness
- Executes golangci-lint to enforce code quality standards
- Validates the build across multiple Go versions

The CI pipeline is triggered on every push and pull request to ensure code quality is maintained.

## License

See the [LICENSE](LICENSE) file for details.
