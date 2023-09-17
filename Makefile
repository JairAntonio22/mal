NAME=mal

BIN=$(shell go env GOPATH)/bin/$(NAME)

all: $(BIN)

$(BIN): cmd/* pkg/*
	go install ./cmd/$(NAME)

run:
	go run ./cmd/$(NAME)

tests:
	go test ./test

coverage:
	go test -cover -coverpkg=./pkg/... ./test

clean:
	go clean -i ./cmd/$(NAME)
