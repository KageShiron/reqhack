#!/bin/sh
export $(cat .env | xargs)
if [ -z "$REQHACK_BASEHOST" ]; then
    echo "REQHACK_BASEHOST is empty. use localhost."
    export REQHACK_BASEHOST="localhost"
fi

if [ -z "$(grep REQHACK_RANDOM= .env)" ]; then
    echo ".env file has no REQHACK_RANDOM."
    echo REQHACK_RANDOM=$(od -vAn -tx8 -N8 < /dev/random | tr -d " ") >> .env
fi
echo xx $1

if [ "$1" = "build" ]; then
    docker-compose up --build
elif [ "$1" = "server" ]; then
    docker-compose build --no-cache reqhack-server
elif [ "$1" = "client" ]; then
    docker-compose build --no-cache nginx
else
    docker-compose up
fi
