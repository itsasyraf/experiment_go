sudo docker run --rm -v $(pwd):/app -w /app golang:1.22 go get github.com/go-chi/chi/v5
sudo docker run --rm -v $(pwd):/app -w /app golang:1.22 go mod tidy
