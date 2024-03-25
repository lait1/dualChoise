FROM golang:1.19 AS builder

WORKDIR /var/bestChoice
COPY go.mod go.sum ./
COPY vendor ./vendor/
COPY config ./config/
COPY cmd ./cmd/
COPY internal ./internal/

RUN CGO_ENABLED=1 GOOS=linux go build -ldflags "-linkmode external -extldflags '-static' -s -w" -o /go/bin/bestChoice ./cmd/main.go