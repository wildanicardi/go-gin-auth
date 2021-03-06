ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api/github.com/wildanicardi/go-gin-auth/
WORKDIR /api/github.com/wildanicardi/go-gin-auth/

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api/github.com/wildanicardi/go-gin-auth/
WORKDIR /api/github.com/wildanicardi/go-gin-auth/
COPY --from=builder /api/github.com/wildanicardi/go-gin-auth/main .

EXPOSE 3000

ENTRYPOINT ["./main"]
