#!/usr/bin/env bash

## use docker to build the binary file for linux os
#docker run --rm \
#  -v "$PWD":/arena \
#  -w /arena \
#  -e GOOS=linux -e GOARCH=amd64 \
#  golang:1.16 go build -v -o gin-rest-linux-amd64
#

version=v1.0.1

docker build -t luolg/gin-rest-https:$version .
docker push luolg/gin-rest-https:$version
