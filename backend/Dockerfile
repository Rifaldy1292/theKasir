FROM golang:1.24-alpine AS builder

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /app

# Install dependencies first (for better cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .
RUN go build -o /app/api ./cmd/api/main.go

# Production stage
FROM alpine:latest

WORKDIR /app

# Copy the built binary and .env
COPY --from=builder /app/api /app/api
COPY .env .env

EXPOSE 8080

CMD ["/app/api"]
