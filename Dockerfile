# BINARY BUILD
FROM golang@sha256:2293e952c79b8b3a987e1e09d48b6aa403d703cef9a8fa316d30ba2918d37367 as builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/faktory-cron


# IMAGE BUILD
FROM alpine:latest

COPY --from=builder /go/bin/faktory-cron /faktory-cron

ENTRYPOINT ["/faktory-cron", "-config", "crontab.yaml"]
