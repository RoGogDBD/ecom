FROM golang:1.25-alpine AS build

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./cmd/server

FROM alpine:3.20

WORKDIR /app
COPY --from=build /out/server /app/server
COPY config.json /app/config.json

ENTRYPOINT ["/app/server", "-config", "/app/config.json"]
