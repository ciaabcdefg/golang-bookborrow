run: build
	@./bin/server

build:
	@go build -o bin/server cmd/server/server.go

test:
	@go test -v ./...

migration:
	@migrate create -ext sql -dir db/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/migrate.go up

migrate-down:
	@go run cmd/migrate/migrate.go down

migrate-fix:
	@go run cmd/migrate/migrate.go fix $(filter-out $@,$(MAKECMDGOALS))

migrate-roll:
	@go run cmd/migrate/migrate.go rollback

migrate-ver:
	@go run cmd/migrate/migrate.go version

%:
	@: