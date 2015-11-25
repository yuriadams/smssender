FROM golang:latest

MAINTAINER Yuri Adams

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/yuriadams/smssender

# Install our dependencies
RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/yuriadams/gocomtele

# Install api binary globally within container
RUN go install github.com/yuriadams/smssender

WORKDIR /go/src/github.com/yuriadams/smssender

# Set binary as entrypoint
ENTRYPOINT /go/bin/smssender

EXPOSE 8888
