package alipay

import (
	"testing"
	"fmt"
)

func TestAliPay_FundAuthOrderAppFreeze(t *testing.T) {
	fmt.Println("========== FundAuthOrderAppFreeze ==========")
	var p = AliPayFundAuthOrderAppFreeze{}
	p.OutOrderNo = "111"
	p.OutRequestNo = "xxxxx"
	p.OrderTitle = "test"
	p.Amount = "100"
	p.ProductCode = "PRE_AUTH_ONLINE"

	fmt.Println(client.FundAuthOrderAppFreeze(p))
}