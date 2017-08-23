package alipay

import (
	"net/http"
	"net/url"
	"strings"
)

// TradePagePay https://doc.open.alipay.com/doc2/detail.htm?treeId=270&articleId=105901&docType=1
func (this *AliPay) TradePagePay(param AliPayTradePagePay) (url *url.URL, err error) {
	var buf = strings.NewReader(this.URLValues(param).Encode())

	req, err := http.NewRequest("POST", this.apiDomain, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	rep, err := this.client.Do(req)
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

// TradeQuery https://doc.open.alipay.com/doc2/apiDetail.htm?apiId=757&docType=4
func (this *AliPay) TradeQuery(param AliPayTradeQuery) (results *AliPayTradeQueryResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeClose https://doc.open.alipay.com/doc2/apiDetail.htm?apiId=1058&docType=4
func (this *AliPay) TradeClose(param AliPayTradeClose) (results *AliPayTradeCloseResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeRefund https://doc.open.alipay.com/doc2/apiDetail.htm?apiId=759&docType=4
func (this *AliPay) TradeRefund(param AliPayTradeRefund) (results *AliPayTradeRefundResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeFastpayRefundQuery https://doc.open.alipay.com/doc2/apiDetail.htm?docType=4&apiId=1049
func (this *AliPay) TradeFastpayRefundQuery(param AliPayFastpayTradeRefundQuery) (results *AliPayFastpayTradeRefundQueryResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradePay https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.6jrv8J&docType=4&apiId=850
func (this *AliPay) TradePay(param AliPayTradePay) (results *AliPayTradePayResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradePreCreate https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.EnCSXC&docType=4&apiId=862
func (this *AliPay) TradePreCreate(param AliPayTradePreCreate) (results *AliPayTradePreCreateResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeCancel https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.UKvJeT&docType=4&apiId=866
func (this *AliPay) TradeCancel(param AliPayTradeCancel) (results *AliPayTradeCancelResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeCreate https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.x39G6l&docType=4&apiId=1046
func (this *AliPay) TradeCreate(param AliPayTradeCreate) (results *AliPayTradeCreateResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeAppPay https://doc.open.alipay.com/doc2/detail.htm?treeId=204&articleId=105462&docType=1
func (this *AliPay) TradeAppPay(param AliPayTradeAppPay) (results string, err error) {
	results = this.URLValues(param).Encode()
	return results, nil
}
