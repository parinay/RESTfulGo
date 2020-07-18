FROM golang:1.13-alpine AS builder

ARG VERSION
RUN mkdir /app

ADD . /app

WORKDIR /app/crud

RUN CGO_ENABLED=0  GOOS=linux  go build -o main ./server/*.go


FROM alpine:latest AS production
COPY --from=builder  /app .
CMD ["/app/crud/main"]
