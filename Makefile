build:
	go build ./...

vendor:
	go mod download; go mod verify; go mod tidy; go mod vendor
