FROM golang:latest
RUN mkdir /go/src/api

WORKDIR /go/src/api
ADD . /go/src/api

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]