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

// FundAuthOrderFreeze https://docs.open.alipay.com/api_28/alipay.fund.auth.order.freeze/
// 资金授权冻结接口
func (this *AliPay) FundAuthOrderFreeze(param AliPayFundAuthOrderFreeze) (results *AliPayFundAuthOrderFreezeResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// FundAuthOrderUnfreeze https://docs.open.alipay.com/api_28/alipay.fund.auth.order.unfreeze/
// 资金授权解冻接口
func (this *AliPay) FundAuthOrderUnfreeze(param AliPayFundAuthOrderUnfreeze) (results *AliPayFundAuthOrderUnfreezeResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// FundAuthOperationCancel https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.cancel/
// 资金授权撤销接口
func (this *AliPay) FundAuthOperationCancel(param AliPayFundAuthOperationCancel) (results *AliPayFundAuthOperationCancelResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// FundAuthOperationDetailQuery https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.detail.query/
// 资金授权操作查询接口
func (this *AliPay) FundAuthOperationDetailQuery(param AliPayFundAuthOperationDetailQuery) (results *AliPayFundAuthOperationDetailQueryResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// FundAuthOrderAppFreeze https://docs.open.alipay.com/api_28/alipay.fund.auth.order.app.freeze
// 线上资金授权冻结接口
func (this *AliPay) FundAuthOrderAppFreeze(param AliPayFundAuthOrderAppFreeze) (results string, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return "", err
	}
	return p.Encode(), err
}
