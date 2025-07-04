# Run the following command when creating a Docker Image:
# docker build -t payment-service .

# Run the following command when running the Docker Container, this injects the .env file at runtime:
# docker run --env-file .env -p 8084:8084 payment-service

# Step 1: Build the Go app
FROM golang:1.24-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o payment-service .

# Step 2: Create the final image (smaller image without the Go build tools)
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go binary from the build stage
COPY --from=build /app/payment-service /app/payment-service

# Expose the port the app will run on
EXPOSE 8084

# Command to run the application
CMD ["/app/payment-service"]