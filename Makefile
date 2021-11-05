.PHONY: deps docs

deps:
	go mod tidy

docs:
	go run docs/main.go
