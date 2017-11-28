FROM golang:1.9-alpine AS builder
WORKDIR /go/src/github.com/cdrx/faktory_cron
COPY . /go/src/github.com/cdrx/faktory_cron
RUN apk add --update make
RUN make build


FROM alpine:3.6
COPY --from=builder /go/src/github.com/cdrx/faktory_cron/faktory-cron /faktory-cron
CMD /faktory-cron -config /crontab.yaml
