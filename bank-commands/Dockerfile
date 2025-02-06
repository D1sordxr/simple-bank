FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bank-commands ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bank-commands .

COPY configs/app /app/configs/app

CMD ["./bank-commands", "-config", "/app/configs/app/local.yaml"]
