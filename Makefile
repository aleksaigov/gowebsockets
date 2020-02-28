IMAGE_NAME=wss_server

build:
	go build -o ./bin/$(IMAGE_NAME) ./cmd/worker

run: build
	./bin/$(IMAGE_NAME)
