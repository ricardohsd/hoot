setup: setup_db
	go get -u github.com/pressly/goose/cmd/goose

setup_db:
	psql -c "CREATE DATABASE $(DB_NAME)"

build:
	go build -o hoot cmd/*

setup_test:
	psql -c "DROP DATABASE IF EXISTS $(DB_NAME)_test"
	psql -c "CREATE DATABASE $(DB_NAME)_test"

migrate_up:
	goose -dir migrations/ postgres "postgres://$(DB_USER):$(DB_PASSWORD)@127.0.0.1:5432/$(DB_NAME)" up

migrate_down:
	goose -dir migrations/ postgres "postgres://$(DB_USER):$(DB_PASSWORD)@127.0.0.1:5432/$(DB_NAME)" down

test: setup_test
	goose -dir migrations/ postgres "postgres://$(DB_USER):$(DB_PASSWORD)@127.0.0.1:5432/$(DB_NAME)_test" up
	go test -v ./...