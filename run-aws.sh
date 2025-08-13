cd app
sh build-aws.sh
docker compose build app server
docker compose up db server
