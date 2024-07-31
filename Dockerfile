FROM golang:latest
LABEL authors="Pop H2"
LABEL maintainer="Pop H2 <poph2@hiveops.io>"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]

