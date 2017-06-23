package alipay

import (
	"fmt"
	"testing"
)

func TestAliPay_TradeWapPay(t *testing.T) {
	fmt.Println("========== TradeWapPay ==========")
	var p = AliPayTradeWapPay{}
	p.NotifyURL = "http://203.86.24.181:3000/alipay"
	p.ReturnURL = "http://203.86.24.181:3000"
	p.Subject = "修正了中文的 Bug"
	p.OutTradeNo = "trade_no_2017062302"
	p.TotalAmount = "10.00"
	p.ProductCode = "QUICK_WAP_WAY"

	var url, _ = client.TradeWapPay(p)
	fmt.Println(url)
}