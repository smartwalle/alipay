package alipay_test

import (
	"context"
	"testing"

	"github.com/smartwalle/alipay/v3"
)

func TestClient_AgreementQuery(t *testing.T) {
	t.Log("========== AgreementQuery ==========")
	param := alipay.AgreementQuery{}
	t.Log(param)
	rsp, err := client.AgreementQuery(context.Background(), param)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_AgreementPageSign(t *testing.T) {
	t.Log("========== AgreementPageSign ==========")
	param := alipay.AgreementPageSign{
		ProductCode:         "train",
		ExternalAgreementNo: "sign1000",
	}
	rsp, err := client.AgreementPageSign(param)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(rsp)
}

func TestClient_AgreementUnsign(t *testing.T) {
	t.Log("========== AgreementUnsign ==========")
	param := alipay.AgreementUnsign{
		AgreementNo: "sign1000",
	}
	rsp, err := client.AgreementUnsign(context.Background(), param)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_DecodePhoneNumber(t *testing.T) {
	var payload = `{"response":"3UO4ElDwNFuhOg+3UVTlCaiHYgtPFis0EOEy/gBPRuCw1ZX4P5gNsbbATblw5LWfKsEpa5FuFwNUrz14E2Dr8Q==","sign":"N8T+WmCZk5ulYd32bS9uR9KFYkTojqKaYKVAi3HFEp/6RrFrSMoRzD5J29OHfghlxoOe6ZG5Hf2ZKaxdtxURgY04YhmPbygS/l1ECJDmB2yGsPPaTbFP/o20QdpPmFriAVOiWwQLbnVUzm0uwb33zx38YUhyg8L9Rw/q0ts7ZQbIooJXg4JpLQ7cpAScxilRg1JnsTzmClz+UQbOTatnl8gz9NqyTStuAUXecVEOGR/nHMkO53WFkJw3TgAYEdWMclORBUoylQ3O+n8Dkq0uBrWoNQnm7921GUZ9QbrZjcE+zbkoxBfR1jxA7Pp78Pyy6eTVcfyZVO/xQfliBBm5nQ=="}`
	var mobile, err = client.DecodePhoneNumber(context.Background(), payload)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(mobile)
}
