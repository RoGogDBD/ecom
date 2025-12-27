APP_NAME := ecom
BIN_DIR := bin
IMAGE_NAME := ecom-server

.PHONY: build run test clean docker-build docker-run docker-stop docker-clean docker-compose-up docker-compose-down

build:
	go build -o $(BIN_DIR)/$(APP_NAME) ./cmd/server

run: build
	go run ./cmd/server

test:
	go test ./...

clean:
	rm -f $(BIN_DIR)/$(APP_NAME)

# Docker команды
docker-build:
	docker build -t $(IMAGE_NAME):latest .

docker-run: docker-build
	docker run -d --name $(APP_NAME) -p 8080:8080 \
		-e SERVER_HOST=0.0.0.0 \
		-e SERVER_PORT=8080 \
		$(IMAGE_NAME):latest

docker-stop:
	docker stop $(APP_NAME) || true
	docker rm $(APP_NAME) || true

docker-clean: docker-stop
	docker rmi $(IMAGE_NAME):latest || true

# Docker Compose команды
docker-compose-up:
	docker-compose up -d --build

docker-compose-down:
	docker-compose down

docker-compose-logs:
	docker-compose logs -f
