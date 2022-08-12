FROM golang:1.19-alpine AS builder

LABEL stage=gobuilder

ARG path=app/user/cmd/api
ARG buildpath=app/user/cmd
ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY ${buildpath} /build/${buildpath}
COPY common /build/common
COPY go.mod /build/go.mod
COPY go.sum /build/go.sum

RUN go mod download

RUN go build -ldflags="-s -w" ./${path}/main.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai
ENV APP_DEBUG false
WORKDIR /app

ARG path=app/user/cmd/api

COPY --from=builder /build/main /app/bin/main

COPY ${path}/etc /app/etc

CMD ["./bin/main", "api", "-f", "etc/config.yaml"]