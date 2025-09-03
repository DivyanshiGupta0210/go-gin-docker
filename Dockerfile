# Use Golang image
FROM golang:1.24.5

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod ./
COPY go.sum ./

# Download dependencies
RUN go mod download

# Copy all source files
COPY . .

# Build the Go app
RUN go build -o main .

EXPOSE 8080

# Run the binary
CMD ["./main"]
