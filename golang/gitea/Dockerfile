FROM golang:1.11
ARG http_proxy
ARG https_proxy

ENV http_proxy=${http_proxy}
ENV https_proxy=${https_proxy}

WORKDIR /go/src/app

RUN go get -d -u code.gitea.io/gitea
RUN cd ${GOPATH}/src/code.gitea.io/gitea && \
 git checkout v1.7.5 && \
 TAGS="bindata" make generate build