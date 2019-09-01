FROM golang:alpine
ADD ./src/main.go /go/src/main.go
WORKDIR /go/src/
ENV PORT=8080
CMD ["go", "run", "main.go"]
