# Repository Guidelines

## Project Structure & Module Organization
- Place the HTTP entrypoint in `cmd/api/main.go`; keep framework wiring and startup code there.
- Keep application code under `internal/` (e.g., `internal/handlers`, `internal/services`, `internal/repositories`) to avoid leaking APIs.
- Shared, exportable utilities go in `pkg/` only when they are meant for reuse by other modules.
- Store configuration defaults in `configs/` (YAML/JSON) and app docs or diagrams in `docs/`. Keep scripts used in CI or local setup under `scripts/`.
- Tests live alongside code in `*_test.go` files. Add integration or contract tests under `test/` if they span multiple packages.

## Build, Test, and Development Commands
- `go mod tidy` to sync dependencies with `go.mod`/`go.sum` before committing.
- `go fmt ./...` followed by `go vet ./...` to format and catch common issues.
- `go build ./cmd/api` to compile the Echo server binary.
- `go run ./cmd/api` to run the server locally during development.
- `go test ./...` for the full test suite; add `-race` when diagnosing concurrency issues.

## Coding Style & Naming Conventions
- Follow standard Go style (tabs for indentation). Always run `gofmt` or configure your editor to do so on save.
- Keep functions short and focused. Use dependency injection for services to simplify testing.
- Name handlers by resource and action (e.g., `CreateUser`, `ListOrders`). Export only what needs to be shared.
- Use structured logging with clear context keys; prefer request-scoped loggers in handlers.

## Testing Guidelines
- Use Go’s `testing` package with table-driven tests for handlers, services, and middleware.
- Mock external dependencies (DB, HTTP clients) via interfaces; prefer in-memory fakes for Echo contexts where possible.
- Cover both success and failure paths; include regression tests when fixing bugs.
- Run `go test ./...` before pushing; target high coverage on business-critical packages.

## Commit & Pull Request Guidelines
- Write imperative, concise commit messages; prefer Conventional Commit prefixes when they help (e.g., `feat: add order handler`).
- Keep PRs small and focused; include a summary of changes, testing performed, and any new configuration or migrations.
- Link related issues and add screenshots or curl examples for new endpoints or behavior changes.
- Update docs/config samples when introducing new environment variables or routes.

## Security & Configuration Tips
- Keep secrets out of source control; load runtime settings from environment variables or a `.env` file ignored by Git.
- Validate all incoming request data in handlers or middleware; return clear error responses without leaking internals.
- Enable the Go race detector (`-race`) and consider static checks (e.g., `staticcheck`, `golangci-lint`) in CI as the project grows.
