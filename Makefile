.PHONY: test clean tf
test:
	go test -v ./...

api.zip: api
	zip $@ $<
	mv api.zip dist/.
	rm api

dist:
	mkdir dist

api: cmd/api/api.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $@ cmd/api/api.go

cache.zip: cache
	zip $@ $<
	mv cache.zip dist/.
	rm cache

cache: cmd/cache/cache.go
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $@ cmd/cache/cache.go

all: dist cache.zip api.zip
clean: 
	rm -rf dist
dev:
	go run server/server.go