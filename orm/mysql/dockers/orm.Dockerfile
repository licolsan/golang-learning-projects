FROM golang:1.13.4-alpine3.10

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
