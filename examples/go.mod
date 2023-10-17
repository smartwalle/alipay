module github.com/smartwalle/alipay/examples

go 1.18

require (
	github.com/smartwalle/alipay/v3 v3.2.16
	github.com/smartwalle/ngx v1.0.9
	github.com/smartwalle/nsign v1.0.9
	github.com/smartwalle/xid v1.0.7
)

require github.com/smartwalle/ncrypto v1.0.4 // indirect

replace github.com/smartwalle/alipay/v3 => ../
