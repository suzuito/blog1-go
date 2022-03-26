GO_SOURCES := $(shell find . -name '*.go')

init:
	cp ~/.config/gcloud/application_default_credentials.json .
	docker-compose up

mockgen:
	sh mockgen.sh github.com/suzuito/blog1-go/pkg/usecase pkg/usecase/usecase.go
	sh mockgen.sh github.com/suzuito/blog1-go/pkg/usecase pkg/usecase/db.go
	sh mockgen.sh github.com/suzuito/blog1-go/pkg/usecase pkg/usecase/storage.go
	sh mockgen.sh github.com/suzuito/blog1-go/pkg/usecase pkg/usecase/bhtml.go
	sh mockgen.sh github.com/suzuito/blog1-go/pkg/usecase pkg/usecase/cmarkdown.go

test:
	go test -timeout 30s -coverprofile=coverage.txt -covermode=atomic ./...
	go tool cover -html=coverage.txt -o coverage.html

clean:
	rm -f *.exe
