FROM golang:1.22.0-alpine AS builder
WORKDIR /builder
COPY ./ /builder

WORKDIR /builder
RUN go build -o ./deploy/bin/rssbot ./deploy/app

FROM alpine:3.15.0

WORKDIR /app
COPY --from=builder /builder/deploy/bin /app/
