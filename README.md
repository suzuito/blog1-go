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

### Front

```bash
go build -o main.exe main_front/main.go
./main.exe
```

### Uploade `data/articles`

```bash
go build main_uploader/main.go

# Example
./main data/articles/hoge.md data/articles/fuga.md
```

### API

```bash
go build -o main.exe main_api/main.go
./main.exe
```


### Site map generator

```bash
go build -o main.exe main_sitemap_generator/main.go

# Example
./main.exe https://blog.tach.dev

./main.exe -prerender https://blog.tach.dev
```