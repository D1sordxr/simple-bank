FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o outbox-processor ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/outbox-processor .

COPY configs/app /app/configs/app

CMD ["./outbox-processor", "-config", "/app/configs/app/local.yaml"]
