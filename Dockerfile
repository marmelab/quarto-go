FROM golang:1.8
RUN mkdir /code
WORKDIR /code

RUN go get github.com/golang/lint/golint