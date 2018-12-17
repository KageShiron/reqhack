#!/bin/sh
/go/bin/migrate /go/src/github.com/KageShiron/reqhack/server/
crond -f -d 8
/go/bin/server