# Multi-stage build

FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/grpc-server/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app .

CMD ["./main"]