package alipay_test

import (
	"github.com/NeoclubTechnology/alipay/v3"
	"testing"
)

func TestClient_BillDownloadURLQuery(t *testing.T) {
	t.Log("========== BillDownloadURLQuery ==========")
	var p = alipay.BillDownloadURLQuery{}
	p.BillType = "trade"
	p.BillDate = "2019-01-01"
	rsp, err := client.BillDownloadURLQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.CodeSuccess {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.BillDownloadUrl)
}

func TestClient_BillBalanceQuery(t *testing.T) {
	t.Log("========== BillBalanceQuery ==========")
	var p = alipay.BillBalanceQuery{}
	rsp, err := client.BillBalanceQuery(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.CodeSuccess {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.TotalAmount, rsp.Content.FreezeAmount, rsp.Content.AvailableAmount)
}
