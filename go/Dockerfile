FROM golang:latest
RUN mkdir /go/src/api

WORKDIR /go/src/api
ADD . /go/src/api

CMD ["go", "run", "cmd/main.go"]