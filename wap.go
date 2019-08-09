package alipay

import "net/url"

// TradeWapPay 手机网站支付接口 https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
func (this *Client) TradeWapPay(param TradeWapPay) (result *url.URL, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return nil, err
	}

	result, err = url.Parse(this.apiDomain + "?" + p.Encode())
	if err != nil {
		return nil, err
	}

	return result, err
}
