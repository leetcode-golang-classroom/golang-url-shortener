# golang-url-shorter

This repository is to implementation for shorten your input url with golang

for example:

1. Create shorten url
https://www.google.com -> 98sj1-293

2. Query with shorten url
http://localhost:8000/98sj1-293 -> https://www.google.com


## request lifecycle

repo <- service -> serializer -> http

## component diagram

![](./hexagonal-component%20diagram.drawio.png)


## setup storage environment

### start redis
```sh
cd docker

docker compose up -d redis
```
### start mongo
```sh
cd docker

docker compose up -d mongo
```

## run up service

### run up with redis

```sh
export URL_DB=redis
export REDIS_URL=redis://localhost:6379
make run
```
### run up with mongodb

```sh
export MONGO_URL=mongodb://localhost:27017/shortener
export MONGO_DB=shortener
export MONGO_TIMEOUT=30
export URL_DB=mongo
make run
```

## test with msgpack

```sh
make run-tool
```