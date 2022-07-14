FROM alpine/git as Time
RUN git clone https://github.com/wolfcw/libfaketime /libfaketime \
 && apk -U add build-base
WORKDIR /libfaketime
RUN make \
 && make install

FROM golang:1.18.3-alpine

COPY --from=Time /usr/local/lib/faketime/libfaketimeMT.so.1 /lib/faketime.so
ENV LD_PRELOAD=/lib/faketime.so
ENV FAKETIME="-15d" 
ENV DONT_FAKE_MONOTONIC=1

RUN apk add build-base chromium

RUN mkdir /build
ADD ./ /build/
WORKDIR /build

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o server -a -ldflags '-linkmode external -extldflags "-static"'

CMD ["./server"]
