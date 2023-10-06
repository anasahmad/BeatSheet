# Use a lightweight base image with Go preinstalled
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the Go source code into the container
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN export PATH=$(go env GOPATH)/bin:$PATH

RUN swag init

# Build the Go application
RUN go build -o main .

# Expose a port (if your Go app listens on a specific port)
EXPOSE 8080

# Command to run your Go application
CMD ["./main"]