FROM alpine:3.20 as root-certificates
RUN apk add -U --no-cache ca-certificates
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app

FROM golang:latest as builder
LABEL authors="Pop H2"
LABEL maintainer="Pop H2 <poph2@hiveops.io>"

WORKDIR /app

COPY --from=root-certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -o main .


FROM scratch as final

COPY --from=root-certificates /etc/passwd /etc/passwd
COPY --from=root-certificates /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=1001:1001 --from=builder  /app/main /main
COPY --chown=1001:1001 --from=builder  /app/apiary.yml /apiary.yml
USER app

CMD ["./main"]
