.PHONY: install build-amd64 build-arm64 run-server run-client

install:
	go mod tidy

build-amd64:
	GOARCH=amd64 go build -o ./bin/goraft-amd64 ./src/main.go

build-arm64:
	GOARCH=arm64 GOOS=darwin go build -o ./bin/goraft-arm64 ./src/main.go

run-test:


run-server:
	go run ./src/main.go

run-client:
	telnet localhost 7001
