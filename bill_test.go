package alipay

import (
	"testing"
	"time"
)

func TestBillDownloadURLQuery(t *testing.T) {
	type arg struct {
		param  BillDownloadURLQuery
		wanted string
	}

	testCases := []arg{
		{
			param: BillDownloadURLQuery{
				BillType: "trade",
				BillDate: "2017-02-24",
			},
			wanted: "10000",
		},
		{
			param: BillDownloadURLQuery{
				BillType: "trade",
				BillDate: time.Now().Format("2006-01-02"),
			},
			wanted: "40004",
		},
	}
	client := New(appID, "", publicKey, privateKey, false)
	for _, tc := range testCases {
		r, err := client.BillDownloadURLQuery(tc.param)
		t.Log(r, err)
		if err != nil {
			t.FailNow()
		}
		if r.AliPayDataServiceBillDownloadURLQueryResponse.Code != tc.wanted {
			t.FailNow()
		}
	}
}
