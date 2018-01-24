FROM golang:1.9.2-alpine3.7 AS build
LABEL maintainer="sudhanshu@go-jek.com"

# Just so you can login to it
RUN apk add --no-cache bash

ADD bin/samplecli_linux samplecli
ENV PORT 3000
EXPOSE 3000
ENTRYPOINT ["/samplecli"]