package alipay

import "net/url"

// TradeWapPay 手机网站支付接口 https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
func (c *Client) TradeWapPay(param TradeWapPay) (result *url.URL, err error) {
	return c.BuildURL(param)
}

// TradeWapMergePay 无线Wap合并支付接口2.0 https://opendocs.alipay.com/open/028xra
// TODO TradeWapMergePay 接口未经测试
func (c *Client) TradeWapMergePay(param TradeWapMergePay) (result *url.URL, err error) {
	return c.BuildURL(param)
}
