.PHONY: install build run-server run-client

# Install the dependencies.
install:
	go mod tidy

# Build GoRaft.
build:
	go build -o ./bin/goraft ./src/main.go

# Run on the terminal 1.
run-server:
	go run ./src/main.go

# Run on the terminal 2.
run-client:
	telnet localhost 7001
