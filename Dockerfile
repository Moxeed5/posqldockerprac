# Use the official Golang image as the base image
FROM golang:1.22-alpine

# Set the working directory
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Copy the .env file
COPY .env .env

# Build the Go app
RUN go build -o main ./cmd

# Run the executable
CMD ["./main"]
