FROM golang:1.18-alpine AS builder

LABEL stage=gobuilder


ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -tags wireinject -o service ./cmd/main.go ./cmd/wire_gen.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app

COPY --from=builder /build/service /app/service

COPY ./etc /app/etc

CMD ["./service", "-f", "./etc/config.yaml"]
