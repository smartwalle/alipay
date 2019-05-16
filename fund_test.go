package alipay_test

import (
	"github.com/smartwalle/alipay"
	"testing"
)

func TestAliPay_FundTransToAccountTransfer(t *testing.T) {
	t.Log("========== FundTransToAccountTransfer ==========")
	var p = alipay.FundTransToAccountTransfer{}
	p.OutBizNo = "xxxx"
	p.PayeeType = "ALIPAY_LOGONID"
	p.PayeeAccount = "xwmkjn7612@sandbox.com"
	p.Amount = "100"
	rsp, err := client.FundTransToAccountTransfer(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.Msg)
}

func TestAliPay_FundAuthOrderVoucherCreate(t *testing.T) {
	t.Log("========== FundAuthOrderVoucherCreate ==========")
	var p = alipay.FundAuthOrderVoucherCreate{}
	p.OutOrderNo = "1111"
	p.OutRequestNo = "222"
	p.OrderTitle = "eee"
	p.Amount = "1001"
	rsp, err := client.FundAuthOrderVoucherCreate(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.Msg)
}

func TestAliPay_FundAuthOrderAppFreeze(t *testing.T) {
	t.Log("========== FundAuthOrderAppFreeze ==========")
	var p = alipay.FundAuthOrderAppFreeze{}
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
