FROM golang:1.18-alpine AS builder

LABEL stage=gobuilder

ARG path=/app/user/cmd/api

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -ldflags="-s -w" -o /build .${path}/main.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app

ARG path=/app/user/cmd/api

COPY --from=builder /build/main /app/bin/main

COPY ${path}/etc /app/etc

CMD ["./bin/main", "-f", "./etc/config.yaml"]