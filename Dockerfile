FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/wuyueCreator/gin_blog
COPY . $GOPATH/src/github.com/wuyueCreator/gin_blog
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gin_blog"]