package alipay_test

import (
	"testing"

	alipay "github.com/NeoclubTechnology/alipay/v3"
)

func TestClient_AgreementQuery(t *testing.T) {
	t.Log("========== AgreementQuery ==========")
	param := alipay.AgreementQuery{}
	t.Log(param)
	rsp, err := client.AgreementQuery(param)
	t.Log(rsp, err)
}

func TestClient_AgreementPageSign(t *testing.T) {
	t.Log("========== AgreementPageSign ==========")
	param := alipay.AgreementPageSign{
		ProductCode:         "train",
		ExternalAgreementNo: "sign1000",
	}
	t.Log(param)
	rsp, err := client.AgreementPageSign(param)
	t.Log(rsp, err)
}

func TestClient_AgreementUnsign(t *testing.T) {
	t.Log("========== AgreementUnsign ==========")
	param := alipay.AgreementUnsign{
		AgreementNo: "sign1000",
	}
	t.Log(param)
	rsp, err := client.AgreementUnsign(param)
	t.Log(rsp, err)
}
