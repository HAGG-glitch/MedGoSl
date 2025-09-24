# ---------- Build stage ----------
FROM golang:1.24.4 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd/medgosl

# ---------- Runtime stage ----------
FROM debian:bookworm-slim

WORKDIR /app

# copy binary only for production
COPY --from=builder /app/server .

# Default entrypoint (can be overridden by docker-compose)
CMD ["./server"]
