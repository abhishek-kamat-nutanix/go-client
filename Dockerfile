# Use the Go official image
FROM golang:1.23.2-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY . ./

# Download Go modules
RUN go mod download

# Build the server
RUN go build -o server ./reader/server

# Run the server
CMD ["./server"]

 