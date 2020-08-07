package alipay

// BillDownloadURLQuery 查询对账单下载地址接口 https://docs.open.alipay.com/api_15/alipay.data.dataservice.bill.downloadurl.query
func (this *Client) BillDownloadURLQuery(param BillDownloadURLQuery) (result *BillDownloadURLQueryRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// BillBalanceQuery 支付宝商家账户当前余额查询接口 https://opendocs.alipay.com/apis/api_15/alipay.data.bill.balance.query
func (this *Client) BillBalanceQuery(param BillBalanceQuery) (result *BillBalanceQueryRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}
