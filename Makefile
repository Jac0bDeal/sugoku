all: clean lint test goNES

clean:
	@echo "Cleaning bin/..."
	@rm -rf bin/*

dependencies:
	@echo "Installing project dependencies..."
	@go install honnef.co/go/tools/cmd/staticcheck@latest

sugoku:
	@echo "Building sugoku binary for use on local system..."
	@go build -o bin/sugoku ./cmd/sugoku

lint:
	@echo "Running linters..."
	@go vet ./...
	@staticcheck ./...

test:
	@echo "Running tests..."
	@go test ./...
