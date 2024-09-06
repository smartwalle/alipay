package alipay_test

import (
	"context"
	"github.com/smartwalle/alipay/v3"
	"testing"
)

func TestClient_SecurityRiskComplaintInfoBatchQuery(t *testing.T) {
	t.Log("========== SecurityRiskComplaintInfoBatchQuery ==========")
	payload := alipay.SecurityRiskComplaintInfoBatchQueryReq{}
	payload.CurrentPageNum = 1
	payload.PageSize = 20
	payload.GmtComplaintStart = "2024-09-05 00:00:00"
	payload.GmtComplaintEnd = "2024-09-06 00:00:00"
	t.Log(payload)
	result, err := client.SecurityRiskComplaintInfoBatchQuery(context.Background(), payload)
	if err != nil {
		t.Fatal(err)
	}

	if result.IsFailure() {
		t.Fatal(result.Msg, result.SubMsg)
	}
	t.Logf("%v", result)
}

func TestClient_SecurityRiskComplaintInfoQuery(t *testing.T) {
	t.Log("========== TestClient_SecurityRiskComplaintInfoQuery ==========")
	payload := alipay.SecurityRiskComplaintInfoQueryReq{}
	payload.ComplainId = 158237746
	t.Log(payload)
	result, err := client.SecurityRiskComplaintInfoQuery(context.Background(), payload)
	if err != nil {
		t.Fatal(err)
	}

	if result.IsFailure() {
		t.Fatal(result.Msg, result.SubMsg)
	}
	t.Logf("%v", result)
}

func TestClient_SecurityRiskComplaintProcessFinish(t *testing.T) {
	t.Log("========== TestClient_SecurityRiskComplaintProcessFinish ==========")
	payload := alipay.SecurityRiskComplaintProcessFinishReq{}
	payload.IdList = []int64{158237746}
	payload.ProcessCode = "REFUND"
	payload.Remark = "公司内部测试用户"
	t.Log(payload)
	result, err := client.SecurityRiskComplaintProcessFinish(context.Background(), payload)
	if err != nil {
		t.Fatal(err)
	}

	if result.IsFailure() {
		t.Fatal(result.Msg, result.SubMsg)
	}
	t.Logf("%v", result)
}
