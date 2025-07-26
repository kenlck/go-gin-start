# ---- Build Stage ----
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build -o server main.go

# ---- Run Stage ----
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .
COPY .env ./
COPY migrations ./migrations

EXPOSE 8080

CMD ["./server"]
