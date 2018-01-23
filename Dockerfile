#FROM golang:1.9.2-alpine3.7 AS build
#TODO decide to use alpine or scratch
#More details here https://medium.com/@kelseyhightower/optimizing-docker-images-for-static-binaries-b5696e26eb07
FROM scratch
LABEL maintainer="sudhanshu@go-jek.com"

ADD bin/sample_linux sample
ENV PORT 80
EXPOSE 80
ENTRYPOINT ["/sample"]