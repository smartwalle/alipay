package alipay_test

import (
	"context"
	"testing"

	"github.com/smartwalle/alipay/v3"
)

func TestClient_BillDownloadURLQuery(t *testing.T) {
	t.Log("========== BillDownloadURLQuery ==========")
	var p = alipay.BillDownloadURLQuery{}
	p.BillType = "trade"
	p.BillDate = "2019-01-01"
	rsp, err := client.BillDownloadURLQuery(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_BillBalanceQuery(t *testing.T) {
	t.Log("========== BillBalanceQuery ==========")
	var p = alipay.BillBalanceQuery{}
	rsp, err := client.BillBalanceQuery(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_BillAccountLogQuery(t *testing.T) {
	t.Log("========== BillAccountLogQuery ==========")
	var p = alipay.BillAccountLogQuery{}
	rsp, err := client.BillAccountLogQuery(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}
