package alipay

import (
	"net/url"
)

// TradePagePay https://docs.open.alipay.com/270/alipay.trade.page.pay
func (this *AliPay) TradePagePay(param AliPayTradePagePay) (results *url.URL, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return nil, err
	}

	results, err = url.Parse(this.apiDomain + "?" + p.Encode())
	if err != nil {
		return nil, err
	}
	return results, err
}

// TradeAppPay https://docs.open.alipay.com/api_1/alipay.trade.app.pay
func (this *AliPay) TradeAppPay(param AliPayTradeAppPay) (results string, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return "", err
	}
	return p.Encode(), err
}

// TradeFastpayRefundQuery https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query
func (this *AliPay) TradeFastpayRefundQuery(param AliPayFastpayTradeRefundQuery) (results *AliPayFastpayTradeRefundQueryResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeOrderSettle https://docs.open.alipay.com/api_1/alipay.trade.order.settle
func (this *AliPay) TradeOrderSettle(param AliPayTradeOrderSettle) (results *AliPayTradeOrderSettleResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeClose https://docs.open.alipay.com/api_1/alipay.trade.close/
func (this *AliPay) TradeClose(param AliPayTradeClose) (results *AliPayTradeCloseResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeCancel https://docs.open.alipay.com/api_1/alipay.trade.cancel/
func (this *AliPay) TradeCancel(param AliPayTradeCancel) (results *AliPayTradeCancelResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeRefund https://docs.open.alipay.com/api_1/alipay.trade.refund/
func (this *AliPay) TradeRefund(param AliPayTradeRefund) (results *AliPayTradeRefundResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradePreCreate https://docs.open.alipay.com/api_1/alipay.trade.precreate/
func (this *AliPay) TradePreCreate(param AliPayTradePreCreate) (results *AliPayTradePreCreateResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeQuery https://docs.open.alipay.com/api_1/alipay.trade.query/
func (this *AliPay) TradeQuery(param AliPayTradeQuery) (results *AliPayTradeQueryResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeCreate https://docs.open.alipay.com/api_1/alipay.trade.create/
func (this *AliPay) TradeCreate(param AliPayTradeCreate) (results *AliPayTradeCreateResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradePay https://docs.open.alipay.com/api_1/alipay.trade.pay/
func (this *AliPay) TradePay(param AliPayTradePay) (results *AliPayTradePayResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeOrderInfoSync https://docs.open.alipay.com/api_1/alipay.trade.orderinfo.sync/
func (this *AliPay) TradeOrderInfoSync(param AliPayTradeOrderInfoSync) (results *AliPayTradeOrderInfoSyncResponse, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}
