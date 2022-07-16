FROM golang:1.18.4-alpine as builder

RUN apk add build-base

WORKDIR /build

ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build -o server -a -ldflags '-linkmode external -extldflags "-static"'

FROM alpine:latest

RUN apk add chromium

WORKDIR /app

COPY --from=builder /build/yola.db /build/server ./

CMD ["./server"]
