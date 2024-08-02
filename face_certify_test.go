package alipay_test

import (
	"context"
	"github.com/smartwalle/alipay/v3"
	"testing"
)

func TestClient_FaceVerificationInitialize(t *testing.T) {
	t.Log("========== FaceVerificationInitialize ==========")
	var p = alipay.FaceVerificationInitialize{}

	p.OuterOrderNo = "xxxxxx"
	p.BizCode = "DATA_DIGITAL_BIZ_CODE_FACE_VERIFICATION"
	p.IdentityType = "CERT_INFO"
	p.CertType = "IDENTITY_CARD"
	p.CertName = "张三"
	p.CertNo = "131128199004234511"

	rsp, err := client.FaceVerificationInitialize(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_FaceVerificationQuery(t *testing.T) {
	t.Log("========== FaceVerificationQuery ==========")
	var p = alipay.FaceVerificationQuery{}

	p.CertifyId = "xxxxxx"

	rsp, err := client.FaceVerificationQuery(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}
