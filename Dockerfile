FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
COPY vendor ./vendor/
COPY cmd ./cmd/
COPY internal ./internal/

RUN CGO_ENABLED=0 GOOS=linux go build -o /bestChoice ./cmd/main.go

CMD ["/bestChoice"]