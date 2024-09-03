.PHONY: all

all: run

run:
	@go mod tidy
	@go run starscanner
