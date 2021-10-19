package alipay_test

import (
	"testing"

	alipay "github.com/NeoclubTechnology/alipay/v3"
)

func TestClient_UserCertifyOpenInitialize(t *testing.T) {
	t.Log("========== UserCertifyOpenInitialize ==========")
	var p = alipay.UserCertifyOpenInitialize{}
	p.OuterOrderNo = "xxxx"
	p.BizCode = alipay.CertifyBizCodeFace
	p.IdentityParam.IdentityType = "CERT_INFO"
	p.IdentityParam.CertType = "IDENTITY_CARD"
	p.IdentityParam.CertName = "沙箱环境"
	p.IdentityParam.CertNo = "829297191402263571"
	p.MerchantConfig.ReturnURL = "http://127.0.0.1"
	rsp, err := client.UserCertifyOpenInitialize(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.CodeSuccess {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.CertifyId)
}

func TestClient_UserCertifyOpenCertify(t *testing.T) {
	t.Log("========== UserCertifyOpenCertify ==========")
	var p = alipay.UserCertifyOpenCertify{}
	p.CertifyId = "xxxx"
	rsp, err := client.UserCertifyOpenCertify(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
}

func TestClient_UserCertifyOpenQuery(t *testing.T) {
	t.Log("========== UserCertifyOpenQuery ==========")
	var p = alipay.UserCertifyOpenQuery{}
	p.CertifyId = "xxxx"
	rsp, err := client.UserCertifyOpenQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.CodeSuccess {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.Msg)
}
