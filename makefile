.PHONY: run build clean help

help:
	@echo "Available commands:"
	@echo "  make run       - Run the application"
	@echo "  make build     - Build the binary"
	@echo "  make clean     - Remove build artifacts"
	@echo "  make test      - Run tests"

run:
	go run .

build:
	go build -o linkedin-cli

clean:
	rm -f linkedin-cli

test:
	go test ./...