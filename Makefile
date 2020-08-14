makedb:
	if [ ! -f matchers.db ]; then
		sqlite3 matchers.db
	fi

build:
	go build ./...

vendor:
	go mod download; go mod verify; go mod tidy; go mod vendor
