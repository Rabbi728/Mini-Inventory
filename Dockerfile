FROM golang:1.26-alpine

WORKDIR /app

RUN apk add --no-cache git bash

RUN go install github.com/air-verse/air@latest

COPY . .

CMD ["air"]