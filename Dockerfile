FROM golang:latest as builder
LABEL MAINTAINER="Muhammad Hafiedh"

WORKDIR /go/src/be-shop
COPY . .

RUN go mod download && \
    go mod verify


RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/be-shop ./cmd/be-shop


FROM alpine:latest
RUN apk update && \
    adduser -D appuser

COPY --from=builder /go/bin/be-shop /app/be-shop
COPY --from=builder /go/src/be-shop/.env /app/.env

USER appuser

WORKDIR /app

EXPOSE 8090

ENTRYPOINT ["/app/be-shop", "-env", "/app/.env"]