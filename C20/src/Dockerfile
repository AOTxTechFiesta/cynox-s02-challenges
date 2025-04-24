# Stage 1: Build the Go application
FROM golang:1.24.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Stage 2: Run the Go application
FROM ubuntu:22.04

# Set environment variables
ENV PORT=8080

# Set the working directory in the new container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the static files to the new container
COPY --from=builder /app/static ./static

# Copy the templates to the new container
COPY --from=builder /app/templates ./templates

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
