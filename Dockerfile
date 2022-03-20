FROM golang:alpine

RUN apk update && apk add librdkafka-dev pkgconf git curl gcc && mkdir /app

WORKDIR /app
COPY . .

RUN go get -d
RUN go build -o main

CMD ["./main"]