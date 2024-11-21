.PHONY: tidy run watch test lint

tidy:
	go mod tidy
	cp .env.example .env

run:
	go run cmd/web/main.go

watch:
	# go install github.com/cosmtrek/air@latest
	air

test:
	go test ./... -race

lint:
	golangci-lint run ./...