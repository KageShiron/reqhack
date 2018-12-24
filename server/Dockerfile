FROM golang:1.11.3
RUN apt update && apt upgrade -y && apt install -y mariadb-client cron

RUN mkdir -p /go/bin/ && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.1.0/migrate.linux-amd64.tar.gz | tar xvz

ADD ./go.mod /go/src/github.com/KageShiron/reqhack/server/
ADD ./go.sum /go/src/github.com/KageShiron/reqhack/server/
WORKDIR /go/src/github.com/KageShiron/reqhack/server

ENV GO111MODULE on
RUN go mod download

ADD . /go/src/github.com/KageShiron/reqhack/server/
RUN go install github.com/KageShiron/reqhack/server/


COPY autodelete/ /etc/cron.daily/
RUN chmod 755 /etc/cron.daily/autodelete.sh

ENTRYPOINT entry.sh
EXPOSE 8081
