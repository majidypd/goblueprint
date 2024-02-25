# Use the official Golang image as a base
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests into the container
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the port your app runs on
EXPOSE  8687

# Command to run the executable
ENTRYPOINT ["./main"]
CMD ["serv"]