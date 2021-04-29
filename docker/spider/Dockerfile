FROM golang:1.16.0-alpine3.13

USER root

COPY ./spider /go/src/spider

WORKDIR /go/src/spider

ENV GO111MODULE=auto

RUN go mod download

RUN apk update \
    && apk add --no-cache git \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
    && apk del tzdata \
    && rm -rf /var/cache/apk/* \
    && apk add at \
    && apk add mysql-client \
    && go get -u gorm.io/gorm \
    && go get -u gorm.io/driver/mysql \
    && apk add openrc \
    && openrc boot

RUN (crontab -l; echo "0       6       *       *       *       cd /go/src/spider && pwd && ls && go run main.go -mode 1") | crontab -

#COPY ./docker/spider/startup.sh /startup.sh
#RUN chmod 744 /startup.sh
#CMD ["/startup.sh"]