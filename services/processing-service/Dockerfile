# Use the official Golang image to create a build artifact.
FROM golang:1.18 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o processing-service .

######## Start a new stage from scratch #######
FROM alpine:latest

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/processing-service .

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./processing-service"]
