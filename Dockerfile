FROM golang:1.8
RUN mkdir /code
WORKDIR /code

RUN go get github.com/gorilla/mux
# RUN apt-get update && apt-get upgrade --yes