go install github.com/swaggo/swag/cmd/swag@v1.16.4
go mod tidy
swag init --output ./docs

docker compose down -v
docker compose up --build
