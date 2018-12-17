FROM golang:1.11.3-alpine3.8

ADD ./go.mod /go/src/github.com/KageShiron/reqhack/server/
ADD ./go.sum /go/src/github.com/KageShiron/reqhack/server/
WORKDIR /go/src/github.com/KageShiron/reqhack/server
ENV GO111MODULE on
RUN go mod download
ADD . /go/src/github.com/KageShiron/reqhack/server/
RUN go install github.com/KageShiron/reqhack/server/

RUN apk add mariadb-client
COPY autodelete/ /etc/cron.daily/
RUN chmod 755 /etc/cron.daily/autodelete.sh


ENTRYPOINT entry.sh
EXPOSE 8081
