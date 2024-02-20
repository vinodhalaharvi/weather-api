# Start by using the Go 1.21 base image for building the application.
# This image is compatible with multiple architectures, including ARM.
FROM golang:1.21-alpine as builder

# Create and set the working directory inside the container.
WORKDIR /app

# Copy the Go module files and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application's source code.
COPY . .

# Build the application. Disable CGO and force the Linux OS and ARM architecture.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -v -o weather-api

# Use the Alpine Linux image for the final stage to keep it lightweight.
# The Alpine image supports multiple architectures, including ARM.
FROM alpine:latest

# Install ca-certificates for HTTPS requests (if your app needs them).
RUN apk --no-cache add ca-certificates

# Set the working directory in the container.
WORKDIR /root/

# Copy the statically-linked binary from the builder stage.
COPY --from=builder /app/weather-api .

# Command to run the application.
CMD ["./weather-api"]
