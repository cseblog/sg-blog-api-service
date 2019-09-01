FROM golang:alpine
ADD ./src/hello /go/src/hello
WORKDIR /go/src/hello
ENV PORT=3001
CMD ["go", "run", "hello.go"]
