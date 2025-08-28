# Unit Tests

This directory contains unit tests for the Lyrics REST API project. The tests are organized to mirror the project structure.

## Structure

-   `api/`: Contains tests for API components
    -   `services/`: Tests for service layer components
    -   `repositories/`: Tests for repository layer components
    -   `handlers/`: Tests for handler layer components

## Running Tests

To run all unit tests:

```bash
go test ./tests/unit/...
```

To run specific tests:

```bash
go test ./tests/unit/api/services/...
```

To run tests with coverage:

```bash
go test ./tests/unit/... -cover
```

To generate a coverage report:

```bash
go test ./tests/unit/... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```
