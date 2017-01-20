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
