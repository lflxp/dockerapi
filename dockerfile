FROM alpine:latest
MAINTAINER "github.com/lflxp"

ADD dockerapi /bin/dockerapi
ENTRYPOINT ["/bin/dockerapi"]
