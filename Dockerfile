FROM golang

RUN mkdir /app

WORKDIR /app
COPY . .

RUN go get -d
RUN go build -o main

CMD ["./main"]