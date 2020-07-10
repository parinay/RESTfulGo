FROM golang:1.13-alpine

ARG VERSION
RUN mkdir /app

ADD . /app

WORKDIR /app/crud

RUN CGO_ENABLED=0  GOOS=linux  go build -o main ./cmd/*.go

CMD ["/app/crud/main"]
