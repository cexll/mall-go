
.PHONY: user-api user-rpc

default: user-api

user-api:
	# 编译 user-api
	go build -o app/user/cmd/api/bin/user_api app/user/cmd/api/main.go
	# ./app/user/cmd/api/bin/user_api api

user-rpc:
	# 编译 user-rpc
	go build -o app/user/cmd/rpc/bin/user_rpc app/user/cmd/rpc/main.go
	# ./app/user/cmd/rpc/bin/user_rpc grpc:server


