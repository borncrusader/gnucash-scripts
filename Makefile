MIGRATIONS_DIR := ./db/migrations
DATABASE := ./db/matchers.db

makedb:
	if [ ! -f matchers.db ]; then
		sqlite3 matchers.db
	fi

mkmigration:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq -digits 3 ${MIGRATION}

migrate:
	migrate -source file://$(MIGRATIONS_DIR) -database sqlite3://$(DATABASE) up

build:
	go build ./...

vendor:
	go mod download; go mod verify; go mod tidy; go mod vendor
