# Use the Go official image
FROM golang:1.23.2-alpine

RUN apk update && apk add --no-cache curl

RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    chmod +x kubectl && \
    mv kubectl /usr/local/bin/

RUN curl -L "https://github.com/mikefarah/yq/releases/download/v4.44.5/yq_linux_amd64" -o yq && \
    chmod +x yq && \
    mv yq /usr/local/bin/

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

 