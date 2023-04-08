FROM golang:1.19.5

WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod ./
COPY go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Set the working directory to /app/cmd/shortner-go
WORKDIR /app/cmd/shortner-go

# Copy .env file to the container
COPY cmd/shortner-go/.env ./

# Build the Go application
RUN go build -o shortner-go .

# Expose port 8080
EXPOSE 8080

# Run the Go application
CMD ["./shortner-go"]

