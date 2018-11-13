package alipay_test

import (
	"github.com/smartwalle/alipay"
	"testing"
)

func TestAliPay_TradeAppPay(t *testing.T) {
	t.Log("========== TradeAppPay ==========")
	var p = alipay.AliPayTradeAppPay{}
	p.NotifyURL = "http://203.86.24.181:3000/alipay"
	p.Body = "body"
	p.Subject = "商品标题"
	p.OutTradeNo = "01010101"
	p.TotalAmount = "100.00"
	p.ProductCode = "p_1010101"
	param, err := client.TradeAppPay(p)
	if err != nil {
		t.FailNow()
	}
	t.Log(param)
}

func TestAliPay_TradePagePay(t *testing.T) {
	t.Log("========== TradePagePay ==========")
	var p = alipay.AliPayTradePagePay{}
	p.NotifyURL = "http://220.112.233.229:3000/alipay"
	p.ReturnURL = "http://220.112.233.229:3000"
	p.Subject = "修正了中文的 Bug"
	p.OutTradeNo = "trade_no_20170623011121"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"
	url, err := client.TradePagePay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

func TestAliPay_TradePreCreate(t *testing.T) {
	t.Log("========== TradePreCreate ==========")
	var p = alipay.AliPayTradePreCreate{}
	p.OutTradeNo = "no_0001"
	p.Subject = "测试订单"
	p.TotalAmount = "10.10"

	rsp, err := client.TradePreCreate(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.AliPayPreCreateResponse.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.AliPayPreCreateResponse.Msg, rsp.AliPayPreCreateResponse.SubMsg)
	}
	t.Log(rsp.AliPayPreCreateResponse.QRCode)
}

func TestAliPay_TradePay(t *testing.T) {
	t.Log("========== TradePay ==========")
	var p = alipay.AliPayTradePay{}
	p.OutTradeNo = "no_000111"
	p.Subject = "测试订单"
	p.TotalAmount = "10.10"
	p.Scene = "bar_code"
	p.AuthCode = "扫描用户的支付码"

	rsp, err := client.TradePay(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.AliPayTradePay.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.AliPayTradePay.Msg, rsp.AliPayTradePay.SubMsg)
	}
	t.Log(rsp.AliPayTradePay.Msg)
}

func TestAliPay_TradeRefund(t *testing.T) {
	t.Log("========== TradeRefund ==========")
	var p = alipay.AliPayTradeRefund{}
	p.RefundAmount = "10"
	p.OutTradeNo = "trade_no_20170623011121"
	rsp, err := client.TradeRefund(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp.AliPayTradeRefund)
}
