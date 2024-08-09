# syntax=docker/dockerfile:1
ARG TARGET

FROM golang:1.21.1-alpine AS builder

ARG TARGET=$TARGET
ARG VERSION
ARG MODE
ARG MAIN
ARG LDFLAGS

WORKDIR /app

COPY . .

RUN go mod tidy

#RUN go build -ldflags="-s -w -X main.Version=${VERSION} -X main.Mode=${MODE}"  -o /build/$TARGET $MAIN
RUN go build -ldflags="${LDFLAGS}"  -o /build/$TARGET $MAIN

FROM alpine:3.20

ARG TARGET=$TARGET
ENV TARGET=$TARGET

ARG MODE=$MODE
ENV MODE=$MODE

ENV CONFIG_PATH=/config.yaml

COPY --from=builder /build/$TARGET /$TARGET
COPY ./config/config.yaml $CONFIG_PATH

CMD /$TARGET