

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

Sync all

```bash
go build main_uploader/main.go
```

### API

```bash
go build main_api/main.go
```