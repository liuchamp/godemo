.PHONY: build test generate

generate:
	go generate ./...



build:
	go build -o main main.go



test:
	go test ./... -coverprofile=cover.out
	go tool cover -html=cover.out
