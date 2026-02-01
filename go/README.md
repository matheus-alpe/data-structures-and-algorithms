# Data Structures & Algorithms - Go

## Prerequisites

- Go 1.21 or higher

## Running Tests

Run all tests:

```bash
go test ./...
# or
go test ./data_structures/...
```

Run tests with coverage:

```bash
go test -v -cover ./...
```

Run tests for a specific package:

```bash
go test ./...
go test ./data_structures/linked_list
```

Run tests with coverage report:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

Run benchmarks:

```bash
go test -bench=. ./...
```
