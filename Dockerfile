FROM golang:1.18.3-alpine as builder

RUN apk add build-base

RUN mkdir /build
ADD ./ /build/
WORKDIR /build

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o server -a -ldflags '-linkmode external -extldflags "-static"'

FROM alpine:latest

RUN apk add chromium

WORKDIR /app
COPY --from=builder /build/server .

CMD ["./server"]
