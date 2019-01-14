#!/bin/sh
/go/bin/migrate.linux-amd64 -path ./sql/ -database mysql://${DATA_SOURCE_NAME} force 1