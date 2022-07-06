FROM golang:1.18-alpine AS builder

RUN apk add --no-cache git

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod ./
COPY go.sum* ./

RUN go mod download
RUN go mod tidy
RUN go mod vendor

COPY . .
