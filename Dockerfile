FROM golang:1.26-alpine

WORKDIR /app

RUN apk add --no-cache git bash

RUN go install github.com/air-verse/air@latest

COPY . .

RUN [ -f main.go ] && \
    [ ! -f go.mod ] && go mod init basic-inventory-app

RUN go mod tidy

CMD ["air"]