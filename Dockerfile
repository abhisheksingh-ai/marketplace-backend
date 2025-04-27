# Use official Golang image
FROM golang:1.23.4

# Set environment variables
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first (to cache dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of your code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/api

# Expose port (same as your app runs)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
