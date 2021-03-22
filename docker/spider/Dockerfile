FROM golang:1.16.0-alpine3.13

ENV TZ=Asia/Tokyo

# https://youtu.be/Gcflde9fdPM
COPY ./docker/spider/root /var/spool/cron/crontabs/root
RUN chmod +x /bin/my_bash

COPY ./spider /go/src/spider

WORKDIR /go/src/spider

ENV GO111MODULE=on

RUN go mod download

RUN apk update \
    && apk add --no-cache git \
    && apk add mysql-client

CMD crond -l 2 -f
