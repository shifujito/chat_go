# go versino
FROM golang:1.16.3-alpine
# apudate git install
RUN apk add --update && apk add git
# app dir
RUN mkdir /go/src/app
# workgin dir
WORKDIR /go/src/app
# move work dir
ADD . /go/src/app/
