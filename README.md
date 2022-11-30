# mall-go

[中文](./README-cn.md)

Always wanted to complete a fully functional open source project

Decided to develop a full-featured open source project in Go

# Use of technology
- gin、hertz
- grpc、kitex
- redis
- mysql
- mongodb
- asynq | go-queue
- amqp
- elasticsearch | gofound | zinc
- prometheus
- grafana
- jaeger
- dtm
- apisix
- wechat/alipay
- zap
- viper
- docker/docker-compose/kubernetes

# Service Functions
- [x] members
  - TODO recharge
- [x] balance
- [x] multi-merchant
  - TODO withdraw
- [ ] community⌛
- [ ] installments
- [ ] crowdfunding
- [ ] spike
- [ ] group buy
- [ ] lottery
- [ ] delivery
- [ ] fresh
- [ ] coupon
- [ ] second-hand transaction ? trade old things
- [ ] IM
- [ ] live streaming
- [ ] reward



# Project Description

The project builds scaffolding based on mixgo, which can realize flexible assembly of components. It is currently the client API interface, and admin related codes will not be implemented for the time being.

# Catalog introduction

- app:  Business code Include  api grpc mq job 
- common: common components error、middleware、interceptor、tool、ctxdata
- data: runtime data
- deployments: Deploy related configuration files
- docs: Project Series Documentation
- pkg: internal package

# Gateway

The front is slb followed by apisix

# Development mode

Use the microservice development pattern. api(http) --- rpc(grpc)

rpc provides basic service implementation.

api implements service aggregation business processing.


# Log

- logstash 
- filebeat

# Monitor

- prometheus

# Track

- jaeger

# pub/sub

- kafka
- mq

# Message queue、Delay queue、Timed task

- message queue
  - asynq
  - amqp
- delay queue
  - asnyq
  - amqp
- timed task
  - cron

# Distributed transaction

- dtm

# Deployment

develop use docker/docker-compose

deployment use kubernetes


# License

Apache License Version 2.0, http://www.apache.org/licenses/