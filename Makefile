GO_SOURCES := $(shell find . -name '*.go')

uploader.exe: ${GO_SOURCES}
	go build -o uploader.exe main_uploader/main.go

api.exe: ${GO_SOURCES}
	go build -o api.exe main_api/main.go

sitemap_generator.exe: ${GO_SOURCES}
	go build -o sitemap_generator.exe main_sitemap_generator/main.go

start-api:
	source dev.sh && $(shell go env GOPATH)/bin/air -c .air-api.toml

clean:
	rm *.exe