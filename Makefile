APP_NAME := ecom
BIN_DIR := bin

.PHONY: build run test govulncheck clean

build:
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/server

run: build
	go run ./cmd/server

test:
	go test ./...

govulncheck:
	go install golang.org/x/vuln/cmd/govulncheck@latest
	govulncheck ./...

clean:
	rm -f $(BIN_DIR)/$(APP_NAME)