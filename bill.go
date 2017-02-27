package alipay

// BillDownloadURLQuery https://doc.open.alipay.com/docs/api.htm?spm=a219a.7386797.0.0.LwCBuJ&docType=4&apiId=1054
func (this *AliPay) BillDownloadURLQuery(param BillDownloadURLQuery) (results *BillDownloadURLQueryResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}
