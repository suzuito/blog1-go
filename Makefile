GO_SOURCES := $(shell find . -name '*.go')

init:
	cp ~/.config/gcloud/application_default_credentials.json .
	docker-compose up

mockgen:
	sh mockgen.sh pkg/usecase/usecase.go
	sh mockgen.sh pkg/usecase/db.go
	sh mockgen.sh pkg/usecase/storage.go

test:
	go test -timeout 30s -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html

clean:
	rm -f *.exe
