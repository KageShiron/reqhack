export REQHACK_RANDOM=$(od -vAn -tx4 -N4 < /dev/random)
docker-compose up