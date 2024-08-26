FROM golang:1.22-alpine

ENV GO111MODULE=on
ENV APP_ENV development
ENV GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore

RUN apk add bash ca-certificates curl git gcc g++ libc-dev unzip autoconf automake libtool make
RUN mkdir -p /app/services/cart-bff

WORKDIR /app/services/cart-bff

ADD ./services/cart-bff .

RUN cd ../../ && go work init && go work use -r

RUN go install github.com/air-verse/air@latest
RUN go work sync

CMD ["air"]

