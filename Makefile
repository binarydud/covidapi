.PHONY: test
test:
	go test -v ./...
api.zip: api
	zip $@ $<

api: cmd/api/api.go
	GOOS=linux GOARCH=amd64 go build -o $@ cmd/api/api.go

cache.zip: cache
	zip $@ $<

cache: cmd/cache/cache.go
	GOOS=linux GOARCH=amd64 go build -o $@ cmd/cache/cache.go

all: cache.zip api.zip