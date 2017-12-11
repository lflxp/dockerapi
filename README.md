# DockerApi [![Travis](https://travis-ci.org/lflxp/dockerapi.svg?branch=master)](https://api.travis-ci.org/lflxp/dockerapi) [![GoDoc](https://godoc.org/github.com/lflxp/dockerapi?status.svg)](https://godoc.org/github.com/lflxp/dockerapi) [![Coverage Status](https://coveralls.io/repos/github/lflxp/dockerapi/badge.svg?branch=master)](https://coveralls.io/github/lflxp/dockerapi?branch=master)
# dockerapi
docker api 
类似 [Registrator](https://github.com/gliderlabs/registrator) 实现docker服务自动注册的功能

# sdk
https://docs.docker.com/develop/sdk/examples/#commit-a-container

# usage

## install

> go get github.com/lflxp/dockerapi

## build

> go build -o dockerapi dockerapi

## make docker images

### dockerfile

```
FROM alpine:latest
MAINTAINER "github.com/lflxp"

ADD dockerapi /bin/dockerapi
ENTRYPOINT ["/bin/dockerapi"]
```

### docker build

> docker build -t lxp/dockerapi:0.1 .

# Run

> docker run -d --net=host --volume=/var/run/docker.sock:/var/run/docker.sock --name register lxp/dockerapi:0.1 -host={{etcdHost}}:{{etcdPort}}