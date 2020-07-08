FROM golang:1.13-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app/crud

RUN go build -o main .

CMD ["/app/main"]
