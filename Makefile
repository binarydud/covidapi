api.zip: api
	zip $@ $<

api: main.go
	go get .
	GOOS=linux GOARCH=amd64 go build -o $@