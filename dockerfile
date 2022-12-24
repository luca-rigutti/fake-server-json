# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . ./

RUN go build

EXPOSE 8080

ENTRYPOINT /app/start_server.sh