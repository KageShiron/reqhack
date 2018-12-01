#! /bin/sh
if [ -z "$REQHACK_BASEHOST" ]; then
    echo "REQHACK_BASEHOST is empty. use localhost."
    export REQHACK_BASEHOST="localhost"
fi
export REQHACK_RANDOM=$(od -vAn -tx8 -N8 < /dev/random | tr -d " ");
docker-compose up