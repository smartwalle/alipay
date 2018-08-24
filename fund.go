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

// FundAuthOrderVoucherCreate https://docs.open.alipay.com/api_28/alipay.fund.auth.order.voucher.create/
// 资金授权发码接口
func (this *AliPay) FundAuthOrderVoucherCreate(param AliPayFundAuthOrderVoucherCreate) (results *AliPayFundAuthOrderVoucherCreateResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// FundTransOrderQuery https://docs.open.alipay.com/api_28/alipay.fund.auth.order.app.freeze
// 线上资金授权冻结接口
func (this *AliPay) FundAuthOrderAppFreeze(param AliPayFundAuthOrderAppFreeze) (results string, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return "", err
	}
	return p.Encode(), err
}
