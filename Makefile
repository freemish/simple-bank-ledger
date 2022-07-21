# Make and execute simple bank ledger applications.

# `make cli-app-exec` builds the go binary and runs it.
cli-app-exec: cli-app
	@./cli-app || true

# `make cli-app` builds the go binary for the cli app.
cli-app:
	@cd cmd/bank-ledger-cli && go build -o ../../cli-app

# `make test-all` tests all subdirectories in project.
test-all:
	@go test ./...
