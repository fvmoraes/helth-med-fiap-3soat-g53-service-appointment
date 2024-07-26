# Use the official Golang image
FROM golang:1.19-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o helthmed-appointment ./cmd/helthmed-appointment

# Expose port 8082 to the outside world
EXPOSE 8082

# Command to run the executable
CMD ["./helthmed-appointment"]
