package alipay

// FundTransToAccountTransfer https://doc.open.alipay.com/docs/api.htm?apiId=1321&docType=4
// 单笔转账到支付宝账户接口
func (this *AliPay) FundTransToAccountTransfer(param AliPayFundTransToAccountTransfer) (results *AliPayFundTransToAccountTransferResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// FundTransOrderQuery https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.SIkNrH&docType=4&apiId=1322
// 查询转账订单接口fund.trans.order.query
func (this *AliPay) FundTransOrderQuery(param AliPayFundTransOrderQuery) (results *AliPayFundTransOrderQueryResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}
