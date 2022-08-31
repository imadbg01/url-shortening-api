# syntax=docker/dockerfile:1

FROM docker.io/golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

COPY store ./
COPY service ./
COPY client ./

RUN go build -o /shortening

CMD ["/shortening"]

