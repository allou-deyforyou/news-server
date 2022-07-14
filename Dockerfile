FROM golang:1.18.3-alpine

RUN apk add build-base chromium

RUN mkdir /build
ADD ./ /build/
WORKDIR /build

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o server -a -ldflags '-linkmode external -extldflags "-static"'

RUN git clone https://github.com/wolfcw/libfaketime.git 
RUN cd /libfaketime/src && make install

CMD ["./server"]
