FROM golang:1.16 AS build
WORKDIR /arena
COPY . /arena
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o gin-rest-https-linux-amd64

FROM alpine:3.7
LABEL author=nluo@ptc.com
WORKDIR /arena

COPY --from=build /arena/gin-rest-https-linux-amd64 /arena/gin-rest-https-linux-amd64

ENTRYPOINT ["/arena/gin-rest-https-linux-amd64"]
