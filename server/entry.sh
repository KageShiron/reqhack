#!/bin/sh
/go/src/github.com/KageShiron/reqhack/server/migrate.linux-amd64 -path ./sql/ -database mysql://${DATA_SOURCE_NAME} up
cron
/go/bin/server