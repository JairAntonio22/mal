NAME=mal

BIN=$(shell go env GOPATH)/bin/$(NAME)

all: $(BIN)

$(BIN): cmd/* pkg/*
	go install ./cmd/$(NAME)

run:
	go run ./cmd/$(NAME)

tests:
	go test ./test

clean:
	go clean -i ./cmd/$(NAME)
