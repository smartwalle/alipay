package alipay

import (
	"io/ioutil"
	"strings"
	"net/http"
	"net/url"
)

// TradeWapPay https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.stK0ff&docType=4&apiId=1046
func (this *AliPay) TradeWapPay(param AliPayTradeWapPay) (html string, url *url.URL, err error) {
	var buf = strings.NewReader(this.URLValues(param).Encode())

	req, err := http.NewRequest("POST", this.apiDomain, buf)
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	rep, err := this.client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer rep.Body.Close()

	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return "", nil, err
	}
	html = string(data)
	url = rep.Request.URL
	return html, url, err
}