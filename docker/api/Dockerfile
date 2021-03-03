FROM golang:1.16.0-alpine3.13

COPY ./api_v1 /go/src/api

WORKDIR /go/src/api

ENV GO111MODULE=on

RUN go mod download

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/pilu/fresh \
  && go get -u github.com/Watson-Sei/cloud-function-avoid-cli \
  && apk add mysql-client

ENV GOOGLE_APPLICATION_CREDENTIALS="/go/src/api/firebase_adminsdk.json"

EXPOSE 8080