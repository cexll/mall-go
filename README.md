# mall-go

从开始工作至今一直想写一款功能全面(市面上热门APP的主要功能)的一款开源项目 (挖个坑)

既然是开源项目肯定所有的技术都得用最新的mixgo脚手架 微服务 容器部署 DTM分布式事务 等等 (不定时填坑

# 使用技术
- gin、hertz
- grpc、kitex
- redis
- mysql
- mongodb
- asynq(其他替代 go-queue)
- es(其他替代 gofound， zinc)
- prometheus
- grafana
- jaeger
- dtm
- apisix
- wechat/alipay
- zap
- viper
- docker/docker-compose/k8s

# 业务功能
- [x] 会员
  - TODO 充值
- [x] 余额
- [x] 多商户
  - TODO 提现
- [ ] 社区⌛
- [ ] 分期
- [ ] 众筹
- [ ] 秒杀
- [ ] 拼团
- [ ] 抽奖
- [ ] 配送
- [ ] 生鲜
- [ ] 优惠券
- [ ] 二手交易
- [ ] IM
- [ ] 直播
- [ ] 悬赏



# 项目简介

项目基于mixgo搭建脚手架，可实现灵活拼装组件，现目前为客户端api接口，暂不会实现admin相关代码

# 目录介绍

- app:  所有业务代码包含api、rpc以及mq（消息队列、延迟队列、定时任务）
- common: 通用组件 error、middleware、interceptor、tool、ctxdata等
- data: 产生的数据
- deployments: 项目系列配置文件
- docs: 项目系列文档
- pkg: 内部package

# 网关

apisix做外网关 网关前面是slb

# 开发模式

本项目使用的是微服务开发，api （http） + rpc（grpc） ， api充当聚合服务，复杂、涉及到其他业务调用的统一写在rpc中，如果一些不会被其他服务依赖使用的简单业务，可以直接写在api的logic中

# 日志

暂时还在考虑是用logstash还是filebeat

# 监控

监控采用prometheus

# 链路追踪
jaeger

# 发布订阅

- kafka 
- mq

# 消息队列、延迟队列、定时任务

- 消息队列
  - asynq
  - amqp
- 延迟队列
  - asnyq
  - amqp
- 定时任务
  - cron

# 分布式事务

分布式事务使用dtm

# 部署

本项目开发环境推荐docker-compose，使用直链方式，放弃服务注册发现中间件（etcd、nacos、consul等）带来的麻烦

测试、线上部署使用k8s（也不需要etcd、nacos、consul等）

# License

Apache License Version 2.0, http://www.apache.org/licenses/