# tigolang
Utility docker image that randomly logs JSON-formatted data to stout to test container logs. Written in GoLang

## Requirements
* Go version 1.19+

## Run

### Go
```shell
$ go run main.go
```

### Docker build and run
```shell
$ docker build -t tigolang:1.0 -t tigolang:latest .
$ docker run tigolang:latest
```

### Docker
```shell
$ docker run devsareno/tigolang:latest
```
