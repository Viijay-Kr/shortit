# Build stage
FROM golang:1.24.2-alpine3.21 AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:3.21

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose the port your service runs on
EXPOSE 8001

# Run the binary
CMD ["./main"]
