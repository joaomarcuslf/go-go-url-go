# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY .env ./

RUN go mod download

COPY main.go ./
COPY configs ./configs
COPY encoders ./encoders
COPY store ./store
COPY handler ./handler

RUN go build -o /go-go-url-go

FROM ubuntu:18.04 AS run

WORKDIR /

RUN apt-get update && apt-get -y install redis-server

EXPOSE 80

COPY scripts/run.sh ./

COPY --from=build /go-go-url-go /go-go-url-go
COPY --from=build /app/.env /.env


CMD sh /run.sh
