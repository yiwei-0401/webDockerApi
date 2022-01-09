FROM golang:latest
MAINTAINER aw
WORKDIR /root

ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.io,direct"

RUN mkdir webDockerApi
COPY . webDockerApi/

RUN cd webDockerApi && go build -o webDockerExec && chmod +x webDockerExec
CMD ["./webDockerApi/webDockerExec"]