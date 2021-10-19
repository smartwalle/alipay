package alipay_test

import (
	"testing"

	alipay "github.com/NeoclubTechnology/alipay/v3"
)

func TestClient_TradeAppPay(t *testing.T) {
	t.Log("========== TradeAppPay ==========")
	var p = alipay.TradeAppPay{}
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

func TestClient_TradePagePay(t *testing.T) {
	t.Log("========== TradePagePay ==========")
	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://220.112.233.229:3000/alipay"
	p.ReturnURL = "http://220.112.233.229:3000"
	p.Subject = "修正了中文的 Bug"
	p.OutTradeNo = "trade_no_201706230111212"
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	p.GoodsDetail = []*alipay.GoodsDetail{&alipay.GoodsDetail{
		GoodsId:   "123",
		GoodsName: "xxx",
		Quantity:  1,
		Price:     13,
	}}

	url, err := client.TradePagePay(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

func TestClient_TradePreCreate(t *testing.T) {
	t.Log("========== TradePreCreate ==========")
	var p = alipay.TradePreCreate{}
	p.OutTradeNo = "no_0001"
	p.Subject = "测试订单"
	p.TotalAmount = "10.10"

	rsp, err := client.TradePreCreate(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.CodeSuccess {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.QRCode)
}

func TestClient_TradePay(t *testing.T) {
	t.Log("========== Trade ==========")
	var p = alipay.TradePay{}
	p.OutTradeNo = "no_000111"
	p.Subject = "测试订单"
	p.TotalAmount = "10.10"
	p.Scene = "bar_code"
	p.AuthCode = "扫描用户的支付码"

	rsp, err := client.TradePay(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.CodeSuccess {
		t.Fatal(rsp.Content.Code, rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.Msg)
}

func TestClient_TradeQuery(t *testing.T) {
	t.Log("TradeQuery")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = "trade_no_20170623021124"
	p.QueryOptions = []string{"TRADE_SETTLE_INFO"}

	rsp, err := client.TradeQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.CodeSuccess {
		t.Fatal(rsp.Content.Code, rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.Msg)
}

func TestClient_TradeRefund(t *testing.T) {
	t.Log("========== TradeRefund ==========")
	var p = alipay.TradeRefund{}
	p.RefundAmount = "100"
	p.OutTradeNo = "trade_no_20170623021124"
	p.OutRequestNo = "111"
	rsp, err := client.TradeRefund(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp.Content)
}

func TestClient_TradeFastPayRefundQuery(t *testing.T) {
	t.Log("========== TradeFastPayRefundQuery ==========")
	var p = alipay.TradeFastPayRefundQuery{}
	p.OutTradeNo = "trade_no_20170623021124"
	p.OutRequestNo = "11"
	p.QueryOptions = []string{"refund_detail_item_list"}

	rsp, err := client.TradeFastPayRefundQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", rsp.Content)
}
