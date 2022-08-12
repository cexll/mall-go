

default: user-api user-rpc balance-api balance-rpc coupon-api coupon-rpc delivery-api delivery-rpc fresh-api fresh-rpc goods-api goods-rpc im-api im-rpc installment-api installment-rpc lottery-api lottery-rpc merchants-api merchants-rpc order-api order-rpc payment-api payment-rpc raise-api raise-rpc seconds-api seconds-rpc spellgroup-api spellgroup-rpc

user-api:
	# 编译 user-api
	go build -ldflags="-s -w" -o build/user-api app/user/cmd/api/main.go
user-rpc:
	# 编译 user-rpc
	go build -ldflags="-s -w" -o build/user-rpc app/user/cmd/rpc/main.go

balance-api:
	# 编译 balance-api
	go build -ldflags="-s -w" -o build/balance-api app/balance/cmd/api/main.go
balance-rpc:
	# 编译 balance-rpc
	go build -ldflags="-s -w" -o build/balance-rpc app/balance/cmd/rpc/main.go

coupon-api:
	# 编译 coupon-api
	go build -ldflags="-s -w" -o build/coupon-api app/coupon/cmd/api/main.go
coupon-rpc:
	# 编译 coupon-rpc
	go build -ldflags="-s -w" -o build/coupon-rpc app/coupon/cmd/rpc/main.go

delivery-api:
	# 编译 delivery-api
	go build -ldflags="-s -w" -o build/delivery-api app/delivery/cmd/api/main.go
delivery-rpc:
	# 编译 delivery-rpc
	go build -ldflags="-s -w" -o build/delivery-rpc app/delivery/cmd/rpc/main.go

fresh-api:
	# 编译 fresh-api
	go build -ldflags="-s -w" -o build/fresh-api app/fresh/cmd/api/main.go
fresh-rpc:
	# 编译 fresh-rpc
	go build -ldflags="-s -w" -o build/fresh-rpc app/fresh/cmd/rpc/main.go

goods-api:
	# 编译 goods-api
	go build -ldflags="-s -w" -o build/goods-api app/goods/cmd/api/main.go
goods-rpc:
	# 编译 goods-rpc
	go build -ldflags="-s -w" -o build/goods-rpc app/goods/cmd/rpc/main.go

im-api:
	# 编译 im-api
	go build -ldflags="-s -w" -o build/im-api app/im/cmd/api/main.go
im-rpc:
	# 编译 im-rpc
	go build -ldflags="-s -w" -o build/im-rpc app/im/cmd/rpc/main.go

installment-api:
	# 编译 installment-api
	go build -ldflags="-s -w" -o build/installment-api app/installment/cmd/api/main.go
installment-rpc:
	# 编译 installment-rpc
	go build -ldflags="-s -w" -o build/installment-rpc app/installment/cmd/rpc/main.go

lottery-api:
	# 编译 lottery-api
	go build -ldflags="-s -w" -o build/lottery-api app/lottery/cmd/api/main.go
lottery-rpc:
	# 编译 lottery-rpc
	go build -ldflags="-s -w" -o build/lottery-rpc app/lottery/cmd/rpc/main.go

merchants-api:
	# 编译 merchants-api
	go build -ldflags="-s -w" -o build/merchants-api app/merchants/cmd/api/main.go
merchants-rpc:
	# 编译 merchants-rpc
	go build -ldflags="-s -w" -o build/merchants-rpc app/merchants/cmd/rpc/main.go

order-api:
	# 编译 order-api
	go build -ldflags="-s -w" -o build/order-api app/order/cmd/api/main.go
order-rpc:
	# 编译 order-rpc
	go build -ldflags="-s -w" -o build/order-rpc app/order/cmd/rpc/main.go

payment-api:
	# 编译 payment-api
	go build -ldflags="-s -w" -o build/payment-api app/payment/cmd/api/main.go
payment-rpc:
	# 编译 payment-rpc
	go build -ldflags="-s -w" -o build/payment-rpc app/payment/cmd/rpc/main.go

raise-api:
	# 编译 raise-api
	go build -ldflags="-s -w" -o build/raise-api app/raise/cmd/api/main.go
raise-rpc:
	# 编译 raise-rpc
	go build -ldflags="-s -w" -o build/raise-rpc app/raise/cmd/rpc/main.go

seconds-api:
	# 编译 seconds-api
	go build -ldflags="-s -w" -o build/seconds-api app/seconds/cmd/api/main.go
seconds-rpc:
	# 编译 seconds-rpc
	go build -ldflags="-s -w" -o build/seconds-rpc app/seconds/cmd/rpc/main.go

spellgroup-api:
	# 编译 spellgroup-api
	go build -ldflags="-s -w" -o build/spellgroup-api app/spellgroup/cmd/api/main.go
spellgroup-rpc:
	# 编译 spellgroup-rpc
	go build -ldflags="-s -w" -o build/spellgroup-rpc app/spellgroup/cmd/rpc/main.go



