.PHONY: build test

FILES := $(shell find ./ -name "*_test.go")
test:
	go test ${FILES}

build:
	dep ensure && \
	go build -o dist/handler cmd/handler/handler.go && \
	go build -o dist/cli cmd/cli/cli.go && \
	go build -o dist/api cmd/api/api.go && \
	zip -j dist/package.zip dist/*
