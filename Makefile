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
run:
	go run main.go
# test
test:
	go test ./...
# benthos
benthos-run:
	docker run --rm -v ./benthos/config.yaml jeffail/benthos
