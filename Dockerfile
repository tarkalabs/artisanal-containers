FROM golang:1.11-alpine AS build_base

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/github.com/tarkalabs/artisanal-containers

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -a -o artisanal-containers