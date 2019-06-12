package alipay

import (
	"net/url"
)

// TradePagePay https://docs.open.alipay.com/api_1/alipay.trade.page.pay
func (this *Client) TradePagePay(param TradePagePay) (results *url.URL, err error) {
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
func (this *Client) TradeAppPay(param TradeAppPay) (results string, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return "", err
	}
	return p.Encode(), err
}

// TradeFastPayRefundQuery https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query
func (this *Client) TradeFastPayRefundQuery(param TradeFastPayRefundQuery) (results *TradeFastPayRefundQueryRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeOrderSettle https://docs.open.alipay.com/api_1/alipay.trade.order.settle
func (this *Client) TradeOrderSettle(param TradeOrderSettle) (results *TradeOrderSettleRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeClose https://docs.open.alipay.com/api_1/alipay.trade.close/
func (this *Client) TradeClose(param TradeClose) (results *TradeCloseRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeCancel https://docs.open.alipay.com/api_1/alipay.trade.cancel/
func (this *Client) TradeCancel(param TradeCancel) (results *TradeCancelRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeRefund https://docs.open.alipay.com/api_1/alipay.trade.refund/
func (this *Client) TradeRefund(param TradeRefund) (results *TradeRefundRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradePreCreate https://docs.open.alipay.com/api_1/alipay.trade.precreate/
func (this *Client) TradePreCreate(param TradePreCreate) (results *TradePreCreateRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeQuery https://docs.open.alipay.com/api_1/alipay.trade.query/
func (this *Client) TradeQuery(param TradeQuery) (results *TradeQueryRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeCreate https://docs.open.alipay.com/api_1/alipay.trade.create/
func (this *Client) TradeCreate(param TradeCreate) (results *TradeCreateRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// Trade https://docs.open.alipay.com/api_1/alipay.trade.pay/
func (this *Client) TradePay(param TradePay) (results *TradePayRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}

// TradeOrderInfoSync https://docs.open.alipay.com/api_1/alipay.trade.orderinfo.sync/
func (this *Client) TradeOrderInfoSync(param TradeOrderInfoSync) (results *TradeOrderInfoSyncRsp, err error) {
	err = this.doRequest("POST", param, &results)
	return results, err
}
