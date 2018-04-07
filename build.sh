#!/bin/sh

# build in docker
docker run --rm -it -v "$PWD":/go/src/github.com/ququzone/http-trash -w /go/src/github.com/ququzone/http-trash golang:1.10 go build -v

docker build -t http-trash .