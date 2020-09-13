export ZOOKEEPER_URI= 127.0.0.1:2181

run:
	go run *.go
test:
	go test ./... -v
coverage:
	go test -v -coverpkg=./... -coverprofile=coverage.out ./...
.PHONY: all test clean