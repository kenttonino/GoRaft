.PHONY: install build run-test run-server run-client

install:
	go mod tidy

build: install
	go build -o ./bin/goraft ./src/main.go

run-test:
	go test -v -race ./src/...

run-server: build
	go run ./src/main.go

run-client:
	telnet localhost 7001
