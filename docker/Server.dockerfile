FROM golang:1.24.2

WORKDIR /app

COPY ./server .

RUN go mod download

RUN go build -o server .