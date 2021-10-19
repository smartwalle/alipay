package alipay_test

import (
	"testing"

	alipay "github.com/NeoclubTechnology/alipay/v3"
)

func TestClient_FundTransToAccountTransfer(t *testing.T) {
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
	if rsp.Content.Code != alipay.CodeSuccess {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.Msg)
}

func TestClient_FundAuthOrderVoucherCreate(t *testing.T) {
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
	if rsp.Content.Code != alipay.CodeSuccess {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.Msg)
}

func TestClient_FundAuthOrderAppFreeze(t *testing.T) {
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

func TestClient_FundTransUniTransfer(t *testing.T) {
	t.Log("========== FundTransUniTransfer ==========")
	var param = alipay.FundTransUniTransfer{
		OutBizNo:    "1111",
		TransAmount: "10.00",
		ProductCode: "TRANS_ACCOUNT_NO_PWD",
		BizScene:    "DIRECT_TRANSFER",
		OrderTitle:  "remark",
		PayeeInfo: &alipay.PayeeInfo{
			Identity:     "xwmkjn7612@sandbox.com",
			IdentityType: "ALIPAY_LOGON_ID",
			Name:         "沙箱环境",
		},
		Remark: "remark",
	}
	rsp, err := client.FundTransUniTransfer(param)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
}

func TestClient_FundTransCommonQuery(t *testing.T) {
	t.Log("========== FundTransCommonQuery ==========")
	var param = alipay.FundTransCommonQuery{
		ProductCode: "TRANS_ACCOUNT_NO_PWD",
		BizScene:    "DIRECT_TRANSFER",
		OutBizNo:    "1111",
	}
	rsp, err := client.FundTransCommonQuery(param)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
}

func TestClient_FundAccountQuery(t *testing.T) {
	t.Log("========== FundAccountQuery ==========")
	var param = alipay.FundAccountQuery{
		AliPayUserId: "2088102169227503",
	}
	rsp, err := client.FundAccountQuery(param)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
}
