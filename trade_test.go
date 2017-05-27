package alipay

import (
	"fmt"
	"testing"
)

func TestAliPay_TradeQuery(t *testing.T) {
	fmt.Println("========== TradeQuery ==========")
	type arg struct {
		outTradeNo string
		wanted     error
		name       string
	}

	testCaes := []arg{
		{"1111111", nil, "query success"},
		//TODO:add more test case
	}

	for _, tc := range testCaes {
		req := AliPayTradeQuery{
			OutTradeNo: tc.outTradeNo,
		}
		resp, err := client.TradeQuery(req)
		if err != tc.wanted {
			t.Errorf("%s input:%s wanted:%v get:%v", tc.name, tc.outTradeNo, tc.wanted, err)
		} else {
			t.Log(resp)
		}
	}
}

func TestAliPay_TradeAppPay(t *testing.T) {
	fmt.Println("========== TradeAppPay ==========")
	var p = AliPayTradeAppPay{}
	p.NotifyURL = "http://203.86.24.181:3000/alipay"
	p.Body = "body"
	p.Subject = "商品标题"
	p.OutTradeNo = "01010101"
	p.TotalAmount = "100.00"
	p.ProductCode = "p_1010101"
	fmt.Println(client.TradeAppPay(p))
}
