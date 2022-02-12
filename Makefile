GO_SOURCES := $(shell find . -name '*.go')

uploader.exe: ${GO_SOURCES}
	go build -o uploader.exe cmd/uploader/*.go

api.exe: ${GO_SOURCES}
	go build -o api.exe cmd/api/*.go

gcf.exe: ${GO_SOURCES}
	go build -o gcf.exe cmd/gcf/*.go

init:
	cp ~/.config/gcloud/application_default_credentials.json .
	docker-compose up

test:
	echo "FIXME"

clean:
	rm -f *.exe
