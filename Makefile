.PHONY: test
test:
	go test -v ./...

api.zip: dist api
	zip dist/$@ $<
	rm api

dist:
	mkdir dist

api: cmd/api/api.go
	GOOS=linux GOARCH=amd64 go build -o $@ cmd/api/api.go

cache.zip: dist cache
	zip dist/$@ $<
	rm cache

cache: cmd/cache/cache.go
	GOOS=linux GOARCH=amd64 go build -o $@ cmd/cache/cache.go

all: cache.zip api.zip