APP_NAME := ecom
BIN_DIR := bin

.PHONY: build run test clean

build:
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/server

run: build
	go run ./cmd/server

test:
	go test ./...

clean:
	rm -f $(BIN_DIR)/$(APP_NAME)