package alipay

// FundTransToAccountTransfer 单笔转账到支付宝账户接口 https://docs.open.alipay.com/api_28/alipay.fund.trans.toaccount.transfer
func (c *Client) FundTransToAccountTransfer(param FundTransToAccountTransfer) (result *FundTransToAccountTransferRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// FundTransOrderQuery 查询转账订单接口 https://docs.open.alipay.com/api_28/alipay.fund.trans.order.query/
func (c *Client) FundTransOrderQuery(param FundTransOrderQuery) (result *FundTransOrderQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOrderVoucherCreate 资金授权发码接口 https://docs.open.alipay.com/api_28/alipay.fund.auth.order.voucher.create/
func (c *Client) FundAuthOrderVoucherCreate(param FundAuthOrderVoucherCreate) (result *FundAuthOrderVoucherCreateRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOrderFreeze 资金授权冻结接口 https://docs.open.alipay.com/api_28/alipay.fund.auth.order.freeze/
func (c *Client) FundAuthOrderFreeze(param FundAuthOrderFreeze) (result *FundAuthOrderFreezeRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOrderUnfreeze 资金授权解冻接口 https://docs.open.alipay.com/api_28/alipay.fund.auth.order.unfreeze/
func (c *Client) FundAuthOrderUnfreeze(param FundAuthOrderUnfreeze) (result *FundAuthOrderUnfreezeRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOperationCancel 资金授权撤销接口 https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.cancel/
func (c *Client) FundAuthOperationCancel(param FundAuthOperationCancel) (result *FundAuthOperationCancelRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOperationDetailQuery 资金授权操作查询接口 https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.detail.query/
func (c *Client) FundAuthOperationDetailQuery(param FundAuthOperationDetailQuery) (result *FundAuthOperationDetailQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// FundAuthOrderAppFreeze 线上资金授权冻结接口 https://docs.open.alipay.com/api_28/alipay.fund.auth.order.app.freeze
func (c *Client) FundAuthOrderAppFreeze(param FundAuthOrderAppFreeze) (result string, err error) {
	return c.EncodeParam(param)
}

// FundTransUniTransfer 单笔转账接口 https://docs.open.alipay.com/api_28/alipay.fund.trans.uni.transfer/
func (c *Client) FundTransUniTransfer(param FundTransUniTransfer) (result *FundTransUniTransferRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// FundTransCommonQuery 转账业务单据查询接口 https://docs.open.alipay.com/api_28/alipay.fund.trans.common.query/
func (c *Client) FundTransCommonQuery(param FundTransCommonQuery) (result *FundTransCommonQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// FundAccountQuery 支付宝资金账户资产查询接口  https://docs.open.alipay.com/api_28/alipay.fund.account.query
func (c *Client) FundAccountQuery(param FundAccountQuery) (result *FundAccountQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// FundTransAppPay 现金红包无线支付接口 https://opendocs.alipay.com/open/03rbyf
func (c *Client) FundTransAppPay(param FundTransAppPay) (result string, err error) {
	return c.EncodeParam(param)
}
