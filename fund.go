package alipay

// FundTransToAccountTransfer https://docs.open.alipay.com/api_28/alipay.fund.trans.toaccount.transfer
// 单笔转账到支付宝账户接口
func (this *Client) FundTransToAccountTransfer(param FundTransToAccountTransfer) (result *FundTransToAccountTransferRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// FundTransOrderQuery https://docs.open.alipay.com/api_28/alipay.fund.trans.order.query/
// 查询转账订单接口
func (this *Client) FundTransOrderQuery(param FundTransOrderQuery) (result *FundTransOrderQueryRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOrderVoucherCreate https://docs.open.alipay.com/api_28/alipay.fund.auth.order.voucher.create/
// 资金授权发码接口
func (this *Client) FundAuthOrderVoucherCreate(param FundAuthOrderVoucherCreate) (result *FundAuthOrderVoucherCreateRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOrderFreeze https://docs.open.alipay.com/api_28/alipay.fund.auth.order.freeze/
// 资金授权冻结接口
func (this *Client) FundAuthOrderFreeze(param FundAuthOrderFreeze) (result *FundAuthOrderFreezeRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOrderUnfreeze https://docs.open.alipay.com/api_28/alipay.fund.auth.order.unfreeze/
// 资金授权解冻接口
func (this *Client) FundAuthOrderUnfreeze(param FundAuthOrderUnfreeze) (result *FundAuthOrderUnfreezeRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOperationCancel https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.cancel/
// 资金授权撤销接口
func (this *Client) FundAuthOperationCancel(param FundAuthOperationCancel) (result *FundAuthOperationCancelRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOperationDetailQuery https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.detail.query/
// 资金授权操作查询接口
func (this *Client) FundAuthOperationDetailQuery(param FundAuthOperationDetailQuery) (result *FundAuthOperationDetailQueryRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOrderAppFreeze https://docs.open.alipay.com/api_28/alipay.fund.auth.order.app.freeze
// 线上资金授权冻结接口
func (this *Client) FundAuthOrderAppFreeze(param FundAuthOrderAppFreeze) (result string, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return "", err
	}
	return p.Encode(), err
}
