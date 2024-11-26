# Step 1: Build the Go Application
FROM golang:1.23.0  as builder

# Ensure Go compiler builds a static linked binary
ENV CGO_ENABLED=0

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum  ./

# Download all dependencies.
RUN go mod download

# Install make and git
RUN apt update && apt install make

# Copy the rest of the application source code
COPY . .

# Install swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Build application
RUN make all

# Step 2: Run the Go Application
FROM alpine:latest 

RUN mkdir /config
COPY --from=builder /app/main /usr/local/bin/
COPY --from=builder /app/config/config.yml /config/config.yml

# Give Execution permission
RUN chmod +x /usr/local/bin/main

# Expose port 8080
EXPOSE  8080

# Entrypoint
CMD ["main"]
