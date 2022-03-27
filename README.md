[![codecov](https://codecov.io/gh/suzuito/blog1-go/branch/main/graph/badge.svg?token=YCIEOXPNO7)](https://codecov.io/gh/suzuito/blog1-go)

# Blog

http://blog.tach.dev

## Set up development environment

```bash
make init
```

```bash
make mockgen
```

Put test articles

```bash
docker-compose exec gcf /bin/bash -c './gcf.exe update-article -input-dir=./deployment/gcf/testdata'
docker-compose exec gcf /bin/bash -c './gcf.exe delete-article -input-dir=./deployment/gcf/testdata'
```

```bash
open http://localhost:8080/articles
```
