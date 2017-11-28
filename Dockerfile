FROM alpine
RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*
COPY faktory-cron /faktory-cron
CMD /faktory-cron -config /crontab.yaml