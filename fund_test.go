package alipay_test

import (
	"github.com/smartwalle/alipay"
	"testing"
)

func TestAliPay_FundTransToAccountTransfer(t *testing.T) {
	t.Log("========== AliPayFundTransToAccountTransfer ==========")
	var p = alipay.AliPayFundTransToAccountTransfer{}
	p.OutBizNo = "xxxx"
	p.PayeeType = "ALIPAY_LOGONID"
	p.PayeeAccount = "xwmkjn7612@sandbox.com"
	p.Amount = "100"
	rsp, err := client.FundTransToAccountTransfer(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Body.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Body.Msg, rsp.Body.SubMsg)
	}
	t.Log(rsp.Body.Msg)
}

func TestAliPay_FundAuthOrderVoucherCreate(t *testing.T) {
	t.Log("========== AliPayFundAuthOrderVoucherCreate ==========")
	var p = alipay.AliPayFundAuthOrderVoucherCreate{}
	p.OutOrderNo = "1111"
	p.OutRequestNo = "222"
	p.OrderTitle = "eee"
	p.Amount = "1001"
	rsp, err := client.FundAuthOrderVoucherCreate(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Body.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Body.Msg, rsp.Body.SubMsg)
	}
	t.Log(rsp.Body.Msg)
}

func TestAliPay_FundAuthOrderAppFreeze(t *testing.T) {
	t.Log("========== AliPayFundAuthOrderAppFreeze ==========")
	var p = alipay.AliPayFundAuthOrderAppFreeze{}
	p.OutOrderNo = "111"
	p.OutRequestNo = "xxxxx"
	p.OrderTitle = "test"
	p.Amount = "100"
	p.ProductCode = "PRE_AUTH_ONLINE"

	rsp, err := client.FundAuthOrderAppFreeze(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
}
