#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
docker stop register
docker rm register
docker build -t lxp/register .
docker run -d --net=host --volume=/var/run/docker.sock:/var/run/docker.sock --name register lxp/register -host=10.6.200.8:2379
