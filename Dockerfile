FROM golang:cross

# The golang Docker sets the $GOPATH to be /go
# https://github.com/docker-library/golang/blob/c1baf037d71331eb0b8d4c70cff4c29cf124c5e0/1.4/Dockerfile
RUN mkdir -p /go/src/github.com/buildkite/golang-docker-example
WORKDIR /go/src/github.com/buildkite/golang-docker-example
