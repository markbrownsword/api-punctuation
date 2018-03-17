# api-punctuation

Micro service for the punctuation application

## prerequisites
A working Go environment.
- [Go](https://golang.org/doc/install)
- [Dep](https://golang.github.io/dep/docs/introduction.html)
- [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

## setup
Add this repo into your Go src folder

```sh
$ go get github.com/markbrownsword/api-punctuation
$ cd $GOPATH/src/github.com/markbrownsword/api-punctuation
$ dep ensure # install dependencies
```

## run - using go

```sh
$ go run *.go
```

## build docker container

```sh
$ docker build -t faraway-api-punctuation .
$ docker image ls
```

## run - using docker container

```sh
$ docker run -d -p 8080:8080 -e PORT=8080 faraway-api-punctuation
```

## routes
- http://localhost:8080/health/ping