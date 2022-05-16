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
tests:
	go test -v ./...
# normal benthos
benthos-run:
	benthos -c ./benthos/config.yaml
# benthos docker
benthos-run-docker:
	docker run --rm -v ./benthos/config.yaml jeffail/benthos
