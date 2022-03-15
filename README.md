[![codecov](https://codecov.io/gh/suzuito/blog1-go/branch/main/graph/badge.svg?token=YCIEOXPNO7)](https://codecov.io/gh/suzuito/blog1-go)

# Blog

BFF for http://blog.tach.dev

## Set up development environment

Gcloud set up

```bash
gcloud auth login
gcloud auth application-default login
gcloud config set project suzuito-minilla
```

```bash
make init
```

Put test articles

```bash
docker-compose exec gcf /bin/bash -c './gcf.exe update-article -input-dir=./deployment/gcf/testdata'
docker-compose exec gcf /bin/bash -c './gcf.exe delete-article -input-dir=./deployment/gcf/testdata'
```

```bash
curl http://localhost:8080/articles
```

## Design

- https://github.com/suzuito/private/blob/main/hack/blog