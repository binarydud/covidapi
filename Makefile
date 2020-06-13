api.zip: api
	zip $@ $<

api: api.go
	GOOS=linux GOARCH=amd64 go build -o $@ api.go

processor.zip: processor
	zip $@ $<

processor: processor.go
	GOOS=linux GOARCH=amd64 go build -o $@ processor.go

all: processor.zip api.zip