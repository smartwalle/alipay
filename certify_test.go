package alipay_test

import (
	"context"
	"testing"

	"github.com/smartwalle/alipay/v3"
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
	rsp, err := client.UserCertifyOpenInitialize(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
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
	rsp, err := client.UserCertifyOpenQuery(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_UserCertdocCertverifyPreconsult(t *testing.T) {
	t.Log("========== UserCertdocCertverifyPreconsult ==========")
	var p = alipay.UserCertDocCertVerifyPreConsult{}
	p.UserName = "xxxx"
	p.CertType = "IDENTITY_CARD"
	p.CertNo = "xxxx"
	p.Mobile = "xxxx"
	p.LogonId = "xxxx"
	rsp, err := client.UserCertDocCertVerifyPreConsult(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_UserCertdocCertverifyConsult(t *testing.T) {
	t.Log("========== UserCertdocCertverifyConsult ==========")
	var p = alipay.UserCertDocCertVerifyConsult{}
	p.VerifyId = "xxxx"
	rsp, err := client.UserCertDocCertVerifyConsult(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}
