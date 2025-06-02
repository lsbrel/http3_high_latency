FROM golang:1.24.2-alpine

WORKDIR /app

COPY ./client .

RUN go mod download

RUN go build -o client .