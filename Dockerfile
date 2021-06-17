FROM golang:1.15-alpine as builder


RUN apk add --no-cache git make gcc musl-dev linux-headers bash

WORKDIR /

RUN git clone https://github.com/bisontrails/ette

RUN cd ette && make build

FROM alpine:latest

RUN apk add --no-cache ca-certificates 
COPY --from=builder /ette/ette /usr/local/bin/