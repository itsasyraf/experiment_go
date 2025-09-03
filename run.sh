go install github.com/swaggo/swag/cmd/swag@v1.16.4
go mod tidy
# swag init --output ./docs

docker compose down -v
docker compose -f ./build/docker/docker-compose.yml --env-file ./.env up --build

rm -fr ./docs
rm -fr ./go.sum
