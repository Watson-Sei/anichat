FROM golang:1.15.6-alpine3.12

COPY ./api_v1 /go/src/api

WORKDIR /go/src/api

ENV GO111MODULE=on

RUN go mod download

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/pilu/fresh \
  && apk add mysql-client

EXPOSE 8080