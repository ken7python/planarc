cd app
sh build-aws.sh
docker compose build app db
docker compose up db server