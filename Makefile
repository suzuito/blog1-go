GO_SOURCES := $(shell find . -name '*.go')

uploader.exe: ${GO_SOURCES}
	go build -o uploader.exe main_uploader/main.go
