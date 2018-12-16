FROM golang:latest

ADD ./go.mod /go/src/github.com/KageShiron/reqhack/server/
ADD ./go.sum /go/src/github.com/KageShiron/reqhack/server/
WORKDIR /go/src/github.com/KageShiron/reqhack/server
ENV GO111MODULE on
RUN go mod download
ADD . /go/src/github.com/KageShiron/reqhack/server/
RUN go install github.com/KageShiron/reqhack/server/

ENTRYPOINT /go/bin/server
EXPOSE 8081
