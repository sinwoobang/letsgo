FROM ubuntu:18.04
FROM golang:1.11
MAINTAINER Sin-Woo Bang <sinwoobang@gmail.com>

RUN apt-get update

WORKDIR /go/src/letsgo/
COPY . .

# It should be changed to "go build".
CMD ["go", "run", "main.go"]
EXPOSE 3333