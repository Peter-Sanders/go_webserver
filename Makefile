build:
	@templ generate
	@go build -o bin/server cmd/main.go

run: build migrate-up
	@./bin/server

test:
	@go test -v ./...

migrate-up:
	@bash migrate.sh up

migrate-down:
	@bash migrate.sh down
