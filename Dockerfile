FROM alpine:3.20 as root-certificates
RUN apk add -U --no-cache ca-certificates
RUN addgroup -g 1001 app
RUN adduser app -u 1001 -D -G app /home/app

FROM golang:1.22.5 as builder
WORKDIR /app
COPY --from=root-certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -o main .

FROM scratch as final
COPY --from=root-certificates /etc/passwd /etc/passwd
COPY --from=root-certificates /etc/group /etc/group
COPY --chown=1001:1001 --from=root-certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=1001:1001 --from=builder  /app/main /main
USER app
ENTRYPOINT ["/main"]

