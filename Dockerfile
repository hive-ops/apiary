FROM golang:alpine3.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

## Build the Go app
RUN go build -o bin/main .

EXPOSE 2468

# Command to run the executable
CMD ["./bin/main"]
