package alipay

// FundTransToAccountTransfer https://docs.open.alipay.com/api_28/alipay.fund.trans.toaccount.transfer
// 单笔转账到支付宝账户接口
func (this *AliPay) FundTransToAccountTransfer(param AliPayFundTransToAccountTransfer) (results *AliPayFundTransToAccountTransferResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// FundTransOrderQuery https://docs.open.alipay.com/api_28/alipay.fund.trans.order.query/
// 查询转账订单接口
func (this *AliPay) FundTransOrderQuery(param AliPayFundTransOrderQuery) (results *AliPayFundTransOrderQueryResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}
