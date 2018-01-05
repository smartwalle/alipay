package alipay

import (
	"fmt"
	"testing"
	"time"
)

func TestAliPay_BillDownloadURLQuery(t *testing.T) {
	fmt.Println("========== BillDownloadURLQuery ==========")
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
