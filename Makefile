# Make and execute simple bank ledger applications.

# `make run` builds the go binary and runs it.
run: bin
	@./cli-app || true

# `make bin` builds the go binary for the cli app.
bin: clean
	@cd cmd/bank-ledger-cli && go build -o ../../cli-app

# `make test` tests all subdirectories in project.
test:
	@go test ./... -coverprofile=cover.out

# `make report` makes a test coverage report rendered in html.
report:
	@go tool cover -html=cover.out -o cover.html

# `make clean` removes any existing binary for the cli app.
clean:
	@rm cli-app || true
	