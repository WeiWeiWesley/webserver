FROM golang:1.11-alpine

COPY . /go/src/webserver
WORKDIR /go/src/webserver

RUN go build -o app main.go

ENTRYPOINT [ "./app" ]