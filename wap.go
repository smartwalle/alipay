package alipay

// TradeWapPay https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
func (this *Client) TradeWapPay(param TradeWapPay) (url string, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return "", err
	}
	return this.apiDomain + "?" + p.Encode(), nil
}
