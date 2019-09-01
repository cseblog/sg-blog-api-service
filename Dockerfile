FROM golang:alpine
ADD ./src/ /go/src/
WORKDIR /go/src/
ENV PORT=8080
CMD ["go", "run", "main.go"]
