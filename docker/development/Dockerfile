FROM golang:1.17.3-alpine3.14

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go get -u github.com/pilu/fresh

EXPOSE 80

ENTRYPOINT fresh
