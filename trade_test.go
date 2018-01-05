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
		{"trade_no_20170623022111", nil, "query success"},
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

func TestAliPay_TradePagePay(t *testing.T) {
	fmt.Println("========== TradePagePay ==========")
	var p = AliPayTradePagePay{}
	p.NotifyURL = "http://220.112.233.229:3000/alipay"
	p.ReturnURL = "http://220.112.233.229:3000"
	p.Subject = "修正了中文的 Bug"
	p.OutTradeNo = "trade_no_20170623011112"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	fmt.Println(client.TradePagePay(p))
}

func TestAliPay_TradePreCreate(t *testing.T) {
	fmt.Println("========== TradePreCreate ==========")
	var p = AliPayTradePreCreate{}
	p.OutTradeNo = "no_0001"
	p.Subject = "测试订单"
	p.TotalAmount = "10.10"

	fmt.Println(client.TradePreCreate(p))
}
