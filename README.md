# goinit

`goinit` is a tiny Go CLI that scaffolds a clean architecture inspired project structure.

It helps you start faster with a consistent layout for:

- `domain` (entities and contracts)
- `usecase` (business rules)
- `repository` and `infrastructure` (external adapters)
- `delivery` (controllers and routes)

## Why this is useful

When you create many side projects, setup friction slows you down. This script standardizes the first 5-10 minutes so you can focus on actual business logic.

## Features

- Interactive or flag-based usage
- Generates a clean architecture folder layout
- Supports `-dry-run` for previewing output
- Supports `-force` for existing non-empty target directories

## Project Structure (this repo)

- `main.go`: CLI entrypoint and flag handling
- `internal/scaffold/plan.go`: plan building logic for dirs/files
- `internal/scaffold/generator.go`: filesystem write operations
- `internal/scaffold/templates.go`: starter file templates
- `internal/scaffold/scaffold_test.go`: unit tests for planning logic
- `docs/blog-outline-clean-architecture.md`: blog writing blueprint

## Usage

### 1. Run interactively

```bash
go run .
```

You will be prompted for a project name.

### 2. Run with flags

```bash
go run . -name my-go-app
```

### 3. Set a custom module path

```bash
go run . -name my-go-app -module github.com/yourname/my-go-app
```

### 4. Generate in another directory

```bash
go run . -name my-go-app -path ../sandbox
```

### 5. Preview only (no writes)

```bash
go run . -name my-go-app -dry-run
```

### 6. Allow writing into an existing non-empty project folder

```bash
go run . -name my-go-app -force
```

## Generated Template Layout

```text
my-go-app/
  cmd/api/
  config/
  delivery/controller/
  delivery/route/
  domain/
  infrastructure/middleware/
  repository/
  usecase/
  utils/
  tests/
  .github/workflows/
```

## Development

Run tests:

```bash
go test ./...
```
