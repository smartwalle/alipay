package alipay

import (
	"testing"
	"time"
)

func TestBillDownloadurlQuery(t *testing.T) {
	type arg struct {
		param  BillDownloadurlQuery
		wanted string
	}

	testCases := []arg{
		{
			param: BillDownloadurlQuery{
				BillType: "trade",
				BillDate: "2017-02-24",
			},
			wanted: "10000",
		},
		{
			param: BillDownloadurlQuery{
				BillType: "trade",
				BillDate: time.Now().Format("2006-01-02"),
			},
			wanted: "40004",
		},
	}
	client := New(appID, "", publicKey, privateKey, false)
	for _, tc := range testCases {
		r, err := client.BillDownloadurlQuery(tc.param)
		t.Log(r, err)
		if err != nil {
			t.FailNow()
		}
		if r.AlipayDataDataserviceBillDownloadurlQueryResponse.Code != tc.wanted {
			t.FailNow()
		}
	}
}
