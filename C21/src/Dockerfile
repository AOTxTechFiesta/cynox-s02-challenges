# Stage 1: Build the Go application
FROM golang:1.24.1 AS builder

# Update and install necessary packages for SQLite3
RUN apt-get update
RUN apt-get install -y libsqlite3-dev gcc
RUN rm -rf /var/lib/apt/lists/*
RUN apt-get clean

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Enable CGO
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

# Build the Go application
RUN go build -o main .

# Stage 2: Run the Go application
FROM ubuntu:22.04

# Update and install necessary packages
RUN apt-get update
RUN apt-get install -y libsqlite3-0 ca-certificates
RUN rm -rf /var/lib/apt/lists/*
RUN apt-get clean

# Set environment variables
ENV PORT=8080

# Set the working directory in the new container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the template files to the new container
COPY --from=builder /app/templates ./templates

# Expose the port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
