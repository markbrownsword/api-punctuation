# api-punctuation

Micro service for the punctuation application

## prerequisites
- [Go](https://golang.org/doc/install)
- [Dep](https://golang.github.io/dep/docs/introduction.html)
- [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

## setup
Clone this repo into your Go src folder

## run - using go
From the project root, start the web server

    $ go run *.go

## build docker container

```sh
$ docker build -t faraway-api-punctuation .
$ docker image ls
$ docker run -p 5000:5000 -e PORT=5000 faraway-api-punctuation
```

## run - using docker container

```sh
$ docker run -d -p 8080:8080 -e PORT=8080 faraway-api-punctuation
```

## routes
- http://localhost:8080/health/ping