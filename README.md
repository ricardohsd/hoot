# Twitter Scrapper

This is an example in how to fetch tweets for a given user using Golang and Postgres.

# How to use it

The following env vars must be set:

Database vars
```
DB_NAME
DB_USER
DB_PASSWORD
```

Twitter vars
```
CONSUMER_KEY
CONSUMER_SECRET
TOKEN
TOKEN_SECRET
```

First you must run `make setup` to create the database then you can build (make build) it.

```
hoot -user=username -amount=12
```

# Tasks

The following make tasks are available:

```
make setup
make setup_db
make setup_test
make migrate_up
make migrate_down
make test
```

Migrations are handled with [goose](https://github.com/pressly/goose):
```
go get -u github.com/pressly/goose/cmd/goose
```

you can create new migrations as in:
```
goose create MIGRATION_NAME sql
```