FROM golang:1.11

ENV http_proxy http://fwlit0020.lit.hamburg.de:80
ENV https_proxy http://fwlit0020.lit.hamburg.de:80

WORKDIR /go/src/app

COPY dupover.go .

RUN go get ./...
RUN go build dupover.go