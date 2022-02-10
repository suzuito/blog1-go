GO_SOURCES := $(shell find . -name '*.go')

uploader.exe: ${GO_SOURCES}
	go build -o uploader.exe cmd/uploader/main.go

updator-article.exe: ${GO_SOURCES}
	go build -o updator-article.exe cmd/updator-article/main.go

api.exe: ${GO_SOURCES}
	go build -o api.exe cmd/api/main.go

init:
	docker-compose up -d
	until (docker-compose exec -T env /bin/bash -c 'curl http://localhost:8080') do echo "Wait for ready" && sleep 1; done

test:
	echo "FIXME"

clean:
	rm -f *.exe
