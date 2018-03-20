# Faraway Api Punctuation

Micro service for the punctuation application

## Prerequisites

A working Go environment.

- [Go](https://golang.org/doc/install)
- [Dep](https://golang.github.io/dep/docs/introduction.html)
- [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

## Setup

Add this repo into your Go src folder

```sh
go get github.com/markbrownsword/faraway-api-punctuation
cd $GOPATH/src/github.com/markbrownsword/faraway-api-punctuation
dep ensure -vendor-only # installs dependencies
```

## Debug

```sh
go run *.go
```

## Build Docker Container

```sh
docker build -t faraway-api-punctuation .
```

## Run Docker Container

```sh
docker run -d -p 8080:8080 -e PORT=8080 faraway-api-punctuation
docker ps
docker container stop <imageid>

```

## Routes

- <http://localhost:8080/health/ping>