.PHONY: deps docs sign

deps:
	go mod tidy

docs:
	go run docs/main.go

sign:
	gon -log-level=debug -log-json ./gon.json
