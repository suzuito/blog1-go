[![codecov](https://codecov.io/gh/suzuito/blog1-go/branch/main/graph/badge.svg?token=YCIEOXPNO7)](https://codecov.io/gh/suzuito/blog1-go)

# Blog

http://blog.tach.dev

## Set up development environment

```bash
make init
```

### Run server

```bash
make run-server
open http://localhost:8080/articles
```

### Put test articles

```bash
make run-gcf
docker-compose exec gcf /bin/bash -c './gcf.exe update-article -input-dir=./deployment/gcf/testdata'
docker-compose exec gcf /bin/bash -c './gcf.exe delete-article -input-dir=./deployment/gcf/testdata'
```

# Unit test

Regenerate gomock

```bash
make mockgen
```

Run test

```bash
# Run test not using GCP resource
go test -timeout 30s ./pkg/usecase
# Run test using GCP resource
make run-ut
docker-compose exec server go test -timeout 30s ./internal/bgcp/fdb
# Run all test and coverage
make run-ut
docker-compose exec server make test
```
