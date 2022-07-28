# mall-go

从开始工作至今一直想写一款功能全面(市面上热门APP的主要功能)的一款开源项目 (挖个坑)

既然是开源项目肯定所有的技术都得用最新的mixgo脚手架 微服务 容器部署 DTM分布式事务 等等 (不定时填坑
# 使用技术
- gin
- grpc
- redis
- mysql
- asynq(其他替代 go-queue)
- es(其他替代 gofound, zinc)
- prometheus
- dtm
- apisix
- wechat/alipay
- zap
- viper

# 业务功能
- 分期
- 众筹
- 秒杀
- 拼团
- 抽奖
- 配送
- 生鲜
- 优惠券
- 会员
- 多商户
- saas
- 二手交易
- 社区
  - 文字/图文/视频
  - 话题讨论
- IM
- 直播
- 悬赏

# 核心功能
- 登录
- 秒杀队列
- 使用dtm 高性能秒杀
- 使用dtm 分布式事务