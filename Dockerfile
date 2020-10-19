FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/kainhuck/gold
COPY . $GOPATH/src/github.com/kainhuck/gold
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gold"]