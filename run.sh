export REQHACK_RANDOM=$(od -vAn -tx8 -N8 < /dev/random | tr -d " ");
docker-compose up
