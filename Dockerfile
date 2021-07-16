# Builder
FROM golang:1.15.3-alpine3.12 as builder

RUN apk update && apk upgrade && \
  apk --update add git make

WORKDIR /parking_lot

COPY . .

RUN make env

RUN make engine


# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
  apk --update --no-cache add tzdata && \
  mkdir /app 

WORKDIR /parking_lot 

EXPOSE 8080

COPY --from=builder /parking_lot/engine /parking_lot

CMD /parking_lot/engine
