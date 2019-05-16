package alipay

import (
	"net/http"
	"net/url"
	"strings"
)

// TradeWapPay https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
func (this *Client) TradeWapPay(param TradeWapPay) (url *url.URL, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return nil, err
	}
	var buf = strings.NewReader(p.Encode())

	req, err := http.NewRequest("POST", this.apiDomain, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", kContentType)

	rep, err := this.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer rep.Body.Close()

	if err != nil {
		return nil, err
	}
	url = rep.Request.URL
	return url, err
}
