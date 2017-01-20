package alipay

////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.t3A8tQ&docType=4&apiId=757
type AliPayTradeQuery struct {
	AppAuthToken string `json:"-"`                      // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	TradeNo      string `json:"trade_no,omitempty"`
}

func (this AliPayTradeQuery) APIName() string {
	return "alipay.trade.query"
}

func (this AliPayTradeQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayTradeQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeQuery) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradeQueryResponse struct {
	AliPayTradeQuery struct {
		Code                string `json:"code"`
		Msg                 string `json:"msg"`
		SubCode             string `json:"sub_code"`
		SubMsg              string `json:"sub_msg"`
		BuyerLogonId        string `json:"buyer_logon_id"`        // 买家支付宝账号
		BuyerPayAmount      string `json:"buyer_pay_amount"`      // 买家实付金额，单位为元，两位小数。
		BuyerUserId         string `json:"buyer_user_id"`         // 买家在支付宝的用户id
		InvoiceAmount       string `json:"invoice_amount"`        // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
		Openid              string `json:"open_id"`               // 买家支付宝用户号，该字段将废弃，不要使用
		OutTradeNo          string `json:"out_trade_no"`          // 商家订单号
		PointAmount         string `json:"point_amount"`          // 积分支付的金额，单位为元，两位小数。
		ReceiptAmount       string `json:"receipt_amount"`        // 实收金额，单位为元，两位小数
		SendPayDate         string `json:"send_pay_date"`         // 本次交易打款给卖家的时间
		TotalAmount         string `json:"total_amount"`          // 交易的订单金额
		TradeNo             string `json:"trade_no"`              // 支付宝交易号
		TradeStatus         string `json:"trade_status"`          // 交易状态
		AliPayStoreId       string `json:"alipay_store_id"`       // 支付宝店铺编号
		StoreId             string `json:"store_id"`              // 商户门店编号
		TerminalId          string `json:"terminal_id"`           // 商户机具终端编号
		StoreName           string `json:"store_name"`            // 请求交易支付中的商户店铺的名称
		DiscountGoodsDetail string `json:"discount_goods_detail"` // 本次交易支付所使用的单品券优惠的商品优惠信息
		IndustrySepcDetail  string `json:"industry_sepc_detail"`  // 行业特殊信息（例如在医保卡支付业务中，向用户返回医疗信息）。
		FundBillList        []struct {
			FundChannel string `json:"fund_channel"` // 交易使用的资金渠道，详见 支付渠道列表
			Amount      string `json:"amount"`       // 该支付工具类型所使用的金额
			RealAmount  string `json:"real_amount"`  // 渠道实际付款金额
		} `json:"fund_bill_list"` // 交易支付使用的资金渠道
		voucher_detail_list []struct {
			Id                 string `json:"id"`                  // 券id
			Name               string `json:"name"`                // 券名称
			Type               string `json:"type"`                // 当前有三种类型： ALIPAY_FIX_VOUCHER - 全场代金券, ALIPAY_DISCOUNT_VOUCHER - 折扣券, ALIPAY_ITEM_VOUCHER - 单品优惠
			Amount             string `json:"amount"`              // 优惠券面额，它应该会等于商家出资加上其他出资方出资
			MerchantContribute string `json:"merchant_contribute"` // 商家出资（特指发起交易的商家出资金额）
			OtherContribute    string `json:"other_contribute"`    // 其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
			Memo               string `json:"memo"`                // 优惠券备注信息

		} `json:"voucher_detail_list"` // 本交易支付时使用的所有优惠券信息
	} `json:"alipay_trade_query_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradeQueryResponse) IsSuccess() bool {
	if this.AliPayTradeQuery.Msg == "Success" {
		return true
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/doc2/apiDetail.htm?apiId=1058&docType=4
type AliPayTradeClose struct {
	AppAuthToken string `json:"-"`                      // 可选
	NotifyURL    string `json:"notify_url"`             // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	TradeNo      string `json:"trade_no,omitempty"`     // 与 OutTradeNo 二选一
	OperatorId   string `json:"operator_id,omitempty"`  // 可选
}

func (this AliPayTradeClose) APIName() string {
	return "alipay.trade.close"
}

func (this AliPayTradeClose) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this AliPayTradeClose) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeClose) ExtJSONParamValue() string {
	return marshal(this)
}

////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.5zkPUI&docType=4&apiId=759
type AliPayTradeRefund struct {
	AppAuthToken string `json:"-"`                      // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	TradeNo      string `json:"trade_no,omitempty"`     // 与 OutTradeNo 二选一
	RefundAmount string `json:"refund_amount"`          // 必须 需要退款的金额，该金额不能大于订单金额,单位为元，支持两位小数
	RefundReason string `json:"refund_reason"`          // 可选 退款的原因说明
	OutRequestNo string `json:"out_request_no"`         // 可选 标识一次退款请求，同一笔交易多次退款需要保证唯一，如需部分退款，则此参数必传。
	OperatorId   string `json:"operator_id"`            // 可选 商户的操作员编号
	StoreId      string `json:"store_id"`               // 可选 商户的门店编号
	TerminalId   string `json:"terminal_id"`            // 可选 商户的终端编号
}

func (this AliPayTradeRefund) APIName() string {
	return "alipay.trade.refund"
}

func (this AliPayTradeRefund) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayTradeRefund) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeRefund) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradeRefundResponse struct {
	AliPayTradeRefund struct {
		Code                    string `json:"code"`
		Msg                     string `json:"msg"`
		SubCode                 string `json:"sub_code"`
		SubMsg                  string `json:"sub_msg"`
		TradeNo                 string `json:"trade_no"`       // 支付宝交易号
		OutTradeNo              string `json:"out_trade_no"`   // 商户订单号
		BuyerLogonId            string `json:"buyer_logon_id"` // 用户的登录id
		BuyerUserId             string `json:"buyer_user_id"`  // 买家在支付宝的用户id
		FundChange              string `json:"fund_change"`    // 本次退款是否发生了资金变化
		RefundFee               string `json:"refund_fee"`     // 退款总金额
		GmtRefundPay            string `json:"gmt_refund_pay"` // 退款支付时间
		StoreName               string `json:"store_name"`     // 交易在支付时候的门店名称
		refund_detail_item_list []struct {
			FundChannel string `json:"fund_channel"` // 交易使用的资金渠道，详见 支付渠道列表
			Amount      string `json:"amount"`       // 该支付工具类型所使用的金额
			RealAmount  string `json:"real_amount"`  // 渠道实际付款金额
		} `json:"refund_detail_item_list"` // 退款使用的资金渠道
	} `json:"alipay_trade_refund_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradeRefundResponse) IsSuccess() bool {
	if this.AliPayTradeRefund.Msg == "Success" {
		return true
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.chC7PJ&docType=4&apiId=1049
type AliPayFastpayTradeRefundQuery struct {
	AppAuthToken string `json:"-"`                      // 可选
	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	TradeNo      string `json:"trade_no,omitempty"`     // 与 OutTradeNo 二选一
	OutRequestNo string `json:"out_request_no"`         // 必须 请求退款接口时，传入的退款请求号，如果在退款请求时未传入，则该值为创建交易时的外部交易号
}

func (this AliPayFastpayTradeRefundQuery) APIName() string {
	return "alipay.trade.fastpay.refund.query"
}

func (this AliPayFastpayTradeRefundQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayFastpayTradeRefundQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayFastpayTradeRefundQuery) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayFastpayTradeRefundQueryResponse struct {
	AliPayTradeFastpayRefundQueryResponse struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		OutRequestNo string `json:"out_request_no"` // 本笔退款对应的退款请求号
		OutTradeNo   string `json:"out_trade_no"`   // 创建交易传入的商户订单号
		RefundReason string `json:"refund_reason"`  // 发起退款时，传入的退款原因
		TotalAmount  string `json:"total_amount"`   // 发该笔退款所对应的交易的订单金额
		RefundAmount string `json:"refund_amount"`  // 本次退款请求，对应的退款金额
		TradeNo      string `json:"trade_no"`       // 支付宝交易号
	} `json:"alipay_trade_fastpay_refund_query_response"`
	Sign string `json:"sign"`
}

func (this *AliPayFastpayTradeRefundQueryResponse) IsSuccess() bool {
	if this.AliPayTradeFastpayRefundQueryResponse.Msg == "Success" {
		return true
	}
	return false
}
