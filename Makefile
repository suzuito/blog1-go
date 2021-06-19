GO_SOURCES := $(shell find . -name '*.go')

uploader.exe: ${GO_SOURCES}
	go build -o uploader.exe main_uploader/main.go

api.exe: ${GO_SOURCES}
	go build -o api.exe main_api/main.go

start-api:
	source dev.sh && $(shell go env GOPATH)/bin/air -c .air-api.toml

clean:
	rm *.exe