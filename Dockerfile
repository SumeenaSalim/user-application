# Stage 1: Build the binary
FROM golang:1.21 AS builder

WORKDIR /app

# Copy only necessary files to build the binary
COPY go.mod go.sum .env ./
RUN go mod download

COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Stage 2: Create a minimal image to run the application
FROM alpine:latest

WORKDIR /app

# Copy only the necessary files from the builder stage
COPY --from=builder /app/app .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./app", "api"]
