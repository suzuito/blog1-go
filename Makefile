GO_SOURCES := $(shell find . -name '*.go')

uploader.exe: ${GO_SOURCES}
	go build -o uploader.exe main_uploader/main.go

api.exe: ${GO_SOURCES}
	go build -o api.exe main_api/main.go

init:
	docker-compose up -d
	until (docker-compose exec -T env /bin/bash -c 'curl http://localhost:8080') do echo "Wait for ready" && sleep 1; done

test:
	echo "FIXME"

clean:
	rm -f *.exe
