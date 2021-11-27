# go versino
FROM golang:1.16.3-alpine
# apudate git install
RUN apk add --update && apk add git  && apk add curl
# app dir
RUN mkdir /go/src/app
# workgin dir
WORKDIR /go/src/app/server
# move work dir
COPY . /go/src/app
RUN go install github.com/cosmtrek/air@latest
RUN go mod download
# run comman
CMD ["air"]
