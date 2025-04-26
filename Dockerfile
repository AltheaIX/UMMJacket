# Stage 1: Build
FROM golang:1.23.6-alpine AS builder

# Install necessary packages
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum if you have them
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o app ./internal/cmd/main.go

# Stage 2: Run
FROM alpine:latest

# Install CA certificates
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/app .
COPY --from=builder /app/config.yaml .

# Command to run
CMD ["./app"]