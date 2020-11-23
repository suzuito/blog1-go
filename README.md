# Blog

BFF for http://blog.tach.dev

## Set up development environment

Environment variables1

```bash
source dev.sh
```

Google cloud platform credentials

```bash
export GOOGLE_APPLICATION_CREDENTIALS=./blog-minilla-276dbb450e29.json
```

## Usage

### Uploade `data/articles`

```bash
go build main_uploader/main.go

# Example
./main data/articles/hoge.md data/articles/fuga.md
```

### API

```bash
go build main_api/main.go
```


### Site map generator

```bash
go build main_sitemap_generator/main.go

# Example
./main https://hoge.com
```