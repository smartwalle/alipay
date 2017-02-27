package alipay

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

// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.x39G6l&docType=4&apiId=1046
func (this *AliPay) TradeCreate(param AliPayTradeCreate) (results *AliPayTradeCreateResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}