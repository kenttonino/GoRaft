.PHONY: install build run-server run-client

# Install the dependencies.
install:
	go mod tidy

# Build GoRaft for AMD64 architecture.
build-amd64:
	GOARCH=amd64 go build -o ./bin/goraft-amd64 ./src/main.go

# Build GoRaft for ARM64 architecture
build-arm64:
	GOARCH=arm64 GOOS=darwin go build -o ./bin/goraft-arm64 ./src/main.go

# Run on the terminal 1.
run-server:
	go run ./src/main.go

# Run on the terminal 2.
run-client:
	telnet localhost 7001
