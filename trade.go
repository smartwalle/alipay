package alipay

import (
	"context"
	"net/url"
)

// TradePagePay 统一收单下单并支付页面接口 https://docs.open.alipay.com/api_1/alipay.trade.page.pay
func (c *Client) TradePagePay(param TradePagePay) (result *url.URL, err error) {
	return c.BuildURL(param)
}

// TradeAppPay App支付接口 https://docs.open.alipay.com/api_1/alipay.trade.app.pay
func (c *Client) TradeAppPay(param TradeAppPay) (result string, err error) {
	return c.EncodeParam(param)
}

// TradeFastPayRefundQuery 统一收单交易退款查询接口 https://docs.open.alipay.com/api_1/alipay.trade.fastpay.refund.query
func (c *Client) TradeFastPayRefundQuery(ctx context.Context, param TradeFastPayRefundQuery) (result *TradeFastPayRefundQueryRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradeOrderSettle 统一收单交易结算接口 https://docs.open.alipay.com/api_1/alipay.trade.order.settle
func (c *Client) TradeOrderSettle(ctx context.Context, param TradeOrderSettle) (result *TradeOrderSettleRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradeClose 统一收单交易关闭接口 https://docs.open.alipay.com/api_1/alipay.trade.close/
func (c *Client) TradeClose(ctx context.Context, param TradeClose) (result *TradeCloseRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradeCancel 统一收单交易撤销接口 https://docs.open.alipay.com/api_1/alipay.trade.cancel/
func (c *Client) TradeCancel(ctx context.Context, param TradeCancel) (result *TradeCancelRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradeRefund 统一收单交易退款接口 https://docs.open.alipay.com/api_1/alipay.trade.refund/
func (c *Client) TradeRefund(ctx context.Context, param TradeRefund) (result *TradeRefundRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradePreCreate 统一收单线下交易预创建接口 https://docs.open.alipay.com/api_1/alipay.trade.precreate/
func (c *Client) TradePreCreate(ctx context.Context, param TradePreCreate) (result *TradePreCreateRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradeQuery 统一收单线下交易查询接口 https://docs.open.alipay.com/api_1/alipay.trade.query/
func (c *Client) TradeQuery(ctx context.Context, param TradeQuery) (result *TradeQueryRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradeCreate 统一收单交易创建接口 https://docs.open.alipay.com/api_1/alipay.trade.create/
func (c *Client) TradeCreate(ctx context.Context, param TradeCreate) (result *TradeCreateRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradePay 统一收单交易支付接口 https://docs.open.alipay.com/api_1/alipay.trade.pay/
func (c *Client) TradePay(ctx context.Context, param TradePay) (result *TradePayRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradeOrderInfoSync 支付宝订单信息同步接口 https://docs.open.alipay.com/api_1/alipay.trade.orderinfo.sync/
func (c *Client) TradeOrderInfoSync(ctx context.Context, param TradeOrderInfoSync) (result *TradeOrderInfoSyncRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradeRefundAsync 统一收单交易退款(异步)接口 https://opendocs.alipay.com/pre-apis/api_pre/alipay.trade.refund.apply
func (c *Client) TradeRefundAsync(ctx context.Context, param TradeRefundAsync) (result *TradeRefundAsyncRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradeMergePreCreate 统一收单合并支付预创建接口请求参数 https://opendocs.alipay.com/open/028xr9
// TODO TradeMergePreCreate 接口未经测试
func (c *Client) TradeMergePreCreate(ctx context.Context, param TradeMergePreCreate) (result *TradeMergePreCreateRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// TradeAppMergePay App合并支付接口 https://opendocs.alipay.com/open/028py8
// TODO TradeAppMergePay 接口未经测试
func (c *Client) TradeAppMergePay(param TradeAppPay) (result string, err error) {
	return c.EncodeParam(param)
}

// OpenMiniOrderCreate 小程序交易组件业务单创建 https://opendocs.alipay.com/mini/54f80876_alipay.open.mini.order.create
func (c *Client) OpenMiniOrderCreate(ctx context.Context, param OpenMiniOrderCreate) (result *OpenMiniOrderCreateResponse, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}
