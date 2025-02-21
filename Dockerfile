# Use a lightweight base image
FROM golang:1.21-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Use a smaller image for runtime
FROM alpine:3.15

# Copy the built binary
COPY --from=builder /app/main /app/main
COPY --from=builder /app/dev.json .
COPY .env ./.env
# Expose the port your app listens on
EXPOSE 8080

# Command to run your app
CMD ["/app/main"]