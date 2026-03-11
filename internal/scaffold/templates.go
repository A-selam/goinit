package scaffold

import "fmt"

func cmdMainTemplate() string {
	return `package main

import "log"

func main() {
	log.Println("API bootstrap pending")
}
`
}

func goModTemplate(moduleName string) string {
	return fmt.Sprintf("module %s\n\ngo 1.24.4\n", moduleName)
}

func gitignoreTemplate() string {
	return `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool
*.out

# Dependency directories
vendor/

# IDE files
.idea/
.vscode/

# Environment files
.env
`
}

func readmeTemplate(projectName string) string {
	return fmt.Sprintf(`# %s

Go project scaffolded with a clean architecture inspired structure.

## Folder Structure
- cmd/api: service entrypoint
- config: application configuration
- delivery: transport layer (controllers/routes)
- domain: entities and contracts
- usecase: business rules
- repository: data access implementations
- infrastructure: external adapters/middleware
- tests: unit and integration tests

## Quick Start
1. Run %s in your terminal.
2. Change into the generated folder.
3. Run go mod tidy.
4. Start coding your domain and use cases first.
`, projectName, "`go run . -name your-project`")
}

func dockerfileTemplate() string {
	return `FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o api ./cmd/api

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/api ./api
EXPOSE 8080
CMD ["./api"]
`
}

func ciTemplate() string {
	return `name: CI

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: "1.24.x"

      - name: Test
        run: go test ./...
`
}

func openAPITemplate() string {
	return `openapi: 3.0.3
info:
  title: Sample API
  version: 0.1.0
paths:
  /health:
    get:
      summary: Health check
      responses:
        "200":
          description: Healthy
`
}

func renderTemplate(projectName string) string {
	return fmt.Sprintf(`services:
  - type: web
    name: %s
    env: docker
    plan: free
    dockerfilePath: ./utils/Dockerfile
    healthCheckPath: /health
`, projectName)
}
