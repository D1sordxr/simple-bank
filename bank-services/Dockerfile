FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bank-services ./cmd/command-query-service/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bank-services .

COPY configs/app /app/configs/app

CMD ["./bank-services", "-config", "/app/configs/app/local.yaml"]
