# go-sample


# startup

```shell
docker-compose up -d
```

# go run on docker

```shell
docker-compose exec go-dev go run cmd/hello/main.go
```

# build on docker

```shell
docker-compose exec go-dev go build -o dist/hello cmd/hello/main.go
```

# binary execute on docker

```shell
docker-compose exec go-deploy /go/hello
```
