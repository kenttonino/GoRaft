.PHONY: install build-amd64 build-arm64 run-server run-client

install:
	go mod tidy

build: install
	go build -o ./bin/goraft ./src/main.go

run-server: build
	go run ./src/main.go

run-client:
	telnet localhost 7001
