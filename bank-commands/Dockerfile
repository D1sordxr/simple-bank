FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bank-commands ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/bank-commands .

COPY configs/app/local.yaml /config.yaml

CMD ["./bank-commands", "-config", "/config.yaml"]
