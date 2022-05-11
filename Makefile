# up
up:
	docker compose build --no-cache
	docker compose up -d
# down
down:
	docker compose down
# build
build:
	go build -o main.go
# run
	go run main.go