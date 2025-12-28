APP_NAME := ecom
BIN_DIR := bin
IMAGE_NAME := ecom-server

.PHONY: build run test lint coverage clean docker-build docker-up docker-down docker-logs docker-clean

build:
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/server

run: build
	go run ./cmd/server

test:
	go test ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

clean:
	rm -rf $(BIN_DIR)

# Docker
docker-build:
	docker build -t $(IMAGE_NAME):latest .

docker-up:
	docker compose up -d --build

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f

docker-clean:
	docker compose down -v
	docker rmi $(IMAGE_NAME):latest || true
	docker system prune -f
