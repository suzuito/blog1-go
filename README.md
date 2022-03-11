[![codecov](https://codecov.io/gh/suzuito/blog1-go/branch/main/graph/badge.svg?token=YCIEOXPNO7)](https://codecov.io/gh/suzuito/blog1-go)

# Blog

BFF for http://blog.tach.dev

## Set up development environment

Environment variables1

```bash
make init

docker-compose exec gcf /bin/bash -c './gcf.exe update-article -input-dir=./deployment/gcf/testdata'
docker-compose exec gcf /bin/bash -c './gcf.exe delete-article -input-dir=./deployment/gcf/testdata'
```

Gcloud set up

```bash
gcloud auth login
gcloud config set project suzuito-minilla
gcloud config set project suzuito-godzilla
```

Google cloud platform credentials

```bash
export GOOGLE_APPLICATION_CREDENTIALS=./suzuito-godzilla-276dbb450e29.json
```

## Usage

### Uploade `data/articles`

```bash
make uploader.exe
./uploader.exe data/articles/hoge.md data/articles/fuga.md
```

### API

```bash
make api.exe
./api.exe
```

### Site map generator

```bash
go build -o main.exe main_sitemap_generator/main.go

# Example
./main.exe https://blog.tach.dev
```

## Design

- https://github.com/suzuito/private/blob/main/hack/blog