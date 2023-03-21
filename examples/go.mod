module github.com/smartwalle/alipay/examples

go 1.12

require (
	github.com/gin-gonic/gin v1.9.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/smartwalle/alipay/v3 v3.2.0
	github.com/smartwalle/xid v1.0.6
)

replace github.com/smartwalle/alipay/v3 => ../
