FROM golang:latest as builder


ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the application source code
COPY . .

COPY .env ./app/be-shop/.env

# Build the Go binary and output it to /app/be-shop
RUN go build -o /app/be-shop ./cmd/be-shop

# Stage 2: Create a minimal image for running the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/be-shop /app/be-shop

# Ensure the binary is executable
RUN chmod +x /app/be-shop

# Expose the port that the application will run on
EXPOSE 8090

# Run the compiled binary
CMD ["/app/be-shop"]
