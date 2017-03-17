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
		voucher_detail_list []VoucherDetail `json:"voucher_detail_list"` // 本交易支付时使用的所有优惠券信息
	} `json:"alipay_trade_query_response"`
	Sign string `json:"sign"`
}

type VoucherDetail struct {
	Id                 string `json:"id"`                  // 券id
	Name               string `json:"name"`                // 券名称
	Type               string `json:"type"`                // 当前有三种类型： ALIPAY_FIX_VOUCHER - 全场代金券, ALIPAY_DISCOUNT_VOUCHER - 折扣券, ALIPAY_ITEM_VOUCHER - 单品优惠
	Amount             string `json:"amount"`              // 优惠券面额，它应该会等于商家出资加上其他出资方出资
	MerchantContribute string `json:"merchant_contribute"` // 商家出资（特指发起交易的商家出资金额）
	OtherContribute    string `json:"other_contribute"`    // 其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
	Memo               string `json:"memo"`                // 优惠券备注信息
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
	NotifyURL    string `json:"-"`                      // 可选
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

type AliPayTradeCloseResponse struct {
	AliPayTradeClose struct {
		Code       string `json:"code"`
		Msg        string `json:"msg"`
		SubCode    string `json:"sub_code"`
		SubMsg     string `json:"sub_msg"`
		OutTradeNo string `json:"out_trade_no"`
		TradeNo    string `json:"trade_no"`
	} `json:"alipay_trade_close_response"`
	Sign string `json:"sign"`
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
		Code                 string `json:"code"`
		Msg                  string `json:"msg"`
		SubCode              string `json:"sub_code"`
		SubMsg               string `json:"sub_msg"`
		TradeNo              string `json:"trade_no"`       // 支付宝交易号
		OutTradeNo           string `json:"out_trade_no"`   // 商户订单号
		BuyerLogonId         string `json:"buyer_logon_id"` // 用户的登录id
		BuyerUserId          string `json:"buyer_user_id"`  // 买家在支付宝的用户id
		FundChange           string `json:"fund_change"`    // 本次退款是否发生了资金变化
		RefundFee            string `json:"refund_fee"`     // 退款总金额
		GmtRefundPay         string `json:"gmt_refund_pay"` // 退款支付时间
		StoreName            string `json:"store_name"`     // 交易在支付时候的门店名称
		RefundDetailItemList []struct {
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

//////////////////////////////////////////////////////////////////////////////////
//// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.CkYNiG&docType=4&apiId=1046
func (this AliPayTradeCreate) APIName() string {
	return "alipay.trade.create"
}

func (this AliPayTradeCreate) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayTradeCreate) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeCreate) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradeCreate struct {
	AppAuthToken         string             `json:"-"`                      // 可选
	OutTradeNo           string             `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	SellerId             string             `json:"seller_id,omitempty"`    // 卖家支付宝用户ID
	TotalAmount          string             `json:"total_amount"`           // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果同时传入了【打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【打折金额】+【不可打折金额】
	DiscountableAmount   string             `json:"discountable_amount"`    // 可打折金额. 参与优惠计算的金额，单位为元，精确到小数点后两位
	UndiscountableAmount string             `json:"undiscountable_amount"`
	BuyerLogonId         string             `json:"buyer_logon_id"`
	Subject              string             `json:"subject"`
	Body                 string             `json:"body"`
	BuyerId              string             `json:"buyer_id"`
	GoodsDetail          []*GoodsDetailItem `json:"goods_detail,omitempty"`
	OperatorId           string             `json:"operator_id"`
	StoreId              string             `json:"store_id"`
	TerminalId           string             `json:"terminal_id"`
	ExtendParams         *ExtendParamsItem  `json:"extend_params,omitempty"`
	TimeoutExpress       string             `json:"timeout_express"`
	RoyaltyInfo          *RoyaltyInfo       `json:"royalty_info,omitempty"`
	AliPayStoreId        string             `json:"alipay_store_id"`
	SubMerchant          []SubMerchantItem  `json:"sub_merchant"`
	MerchantOrderNo      string             `json:"merchant_order_no"`
}

type AliPayTradeCreateResponse struct {
	AliPayTradeCreateResponse struct {
		Code       string `json:"code"`
		Msg        string `json:"msg"`
		SubCode    string `json:"sub_code"`
		SubMsg     string `json:"sub_msg"`
		TradeNo    string `json:"trade_no"` // 支付宝交易号
		OutTradeNo string `json:"out_trade_no"`
	} `json:"alipay_trade_create_response"`
	Sign string `json:"sign"`
}

type ExtendParamsItem struct {
	SysServiceProviderId string `json:"sys_service_provider_id"`
	HbFqNum              string `json:"hb_fq_num"`
	HbFqSellerPercent    string `json:"hb_fq_seller_percent"`
	TimeoutExpress       string `json:"timeout_express"`
}

type RoyaltyInfo struct {
	RoyaltyType       string                   `json:"royalty_type"`
	RoyaltyDetailInfo []*RoyaltyDetailInfoItem `json:"royalty_detail_infos,omitempty"`
}

type RoyaltyDetailInfoItem struct {
	SerialNo         string `json:"serial_no"`
	TransInType      string `json:"trans_in_type"`
	BatchNo          string `json:"batch_no"`
	OutRelationId    string `json:"out_relation_id"`
	TransOutType     string `json:"trans_out_type"`
	TransOut         string `json:"trans_out"`
	TransIn          string `json:"trans_in"`
	Amount           string `json:"amount"`
	Desc             string `json:"desc"`
	AmountPercentage string `json:"amount_percentage"`
	AliPayStoreId    string `json:"alipay_store_id"`
}

type SubMerchantItem struct {
	MerchantId string `json:"merchant_id"`
}

type GoodsDetailItem struct {
	GoodsId       string `json:"goods_id"`
	AliPayGoodsId string `json:"alipay_goods_id"`
	GoodsName     string `json:"goods_name"`
	Quantity      string `json:"quantity"`
	Price         string `json:"price"`
	GoodsCategory string `json:"goods_category"`
	Body          string `json:"body"`
	ShowUrl       string `json:"show_url"`
}

//////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.IGVsS6&docType=4&apiId=850
type AliPayTradePay struct {
	AppAuthToken string `json:"-"` // 可选
	NotifyURL    string `json:"-"` // 可选

	OutTradeNo           string             `json:"out_trade_no"`           // 必须 商户订单号,64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	Scene                string             `json:"scene"`                  // 必须 支付场景 条码支付，取值：bar_code 声波支付，取值：wave_code	bar_code,wave_code
	AuthCode             string             `json:"auth_code"`              // 必须 支付授权码
	Subject              string             `json:"subject"`                // 必须 订单标题
	BuyerId              string             `json:"buyer_id"`               // 可选 家的支付宝用户id，如果为空，会从传入了码值信息中获取买家ID
	SellerId             string             `json:"seller_id"`              // 可选 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	TotalAmount          string             `json:"total_amount"`           // 可选 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果同时传入【可打折金额】和【不可打折金额】，该参数可以不用传入； 如果同时传入了【可打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【可打折金额】+【不可打折金额】
	DiscountableAmount   string             `json:"discountable_amount"`    // 可选 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。 如果该值未传入，但传入了【订单总金额】和【不可打折金额】，则该值默认为【订单总金额】-【不可打折金额】
	UnDiscountableAmount string             `json:"undiscountable_amount"`  // 可选 不参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]。如果该值未传入，但传入了【订单总金额】和【可打折金额】，则该值默认为【订单总金额】-【可打折金额】
	Body                 string             `json:"body"`                   // 可选 订单描述
	GoodsDetail          []*GoodsDetailItem `json:"goods_detail,omitempty"` // 可选 订单包含的商品列表信息，Json格式，其它说明详见商品明细说明
	OperatorId           string             `json:"operator_id"`            // 可选 商户操作员编号
	StoreId              string             `json:"store_id"`               // 可选 商户门店编号
	TerminalId           string             `json:"terminal_id"`            // 可选 商户机具终端编号
	AliPayStoreId        string             `json:"alipay_store_id"`        // 可选 支付宝的店铺编号
	TimeoutExpress       string             `json:"timeout_express"`        // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m
	AuthNo               string             `json:"auth_no"`                // 预授权号，预授权转交易请求中传入
}

func (this AliPayTradePay) APIName() string {
	return "alipay.trade.pay"
}

func (this AliPayTradePay) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this AliPayTradePay) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradePay) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradePayResponse struct {
	AliPayTradePay struct {
		Code                string `json:"code"`
		Msg                 string `json:"msg"`
		SubCode             string `json:"sub_code"`
		SubMsg              string `json:"sub_msg"`
		BuyerLogonId        string `json:"buyer_logon_id"`        // 买家支付宝账号
		BuyerPayAmount      string `json:"buyer_pay_amount"`      // 买家实付金额，单位为元，两位小数。
		BuyerUserId         string `json:"buyer_user_id"`         // 买家在支付宝的用户id
		CardBalance         string `json:"card_balance"`          // 支付宝卡余额
		DiscountGoodsDetail string `json:"discount_goods_detail"` // 本次交易支付所使用的单品券优惠的商品优惠信息
		FundBillList        []struct {
			FundChannel string `json:"fund_channel"` // 交易使用的资金渠道，详见 支付渠道列表
			Amount      string `json:"amount"`       // 该支付工具类型所使用的金额
			RealAmount  string `json:"real_amount"`  // 渠道实际付款金额
		} `json:"fund_bill_list"` // 交易支付使用的资金渠道
		GmtPayment        string          `json:"gmt_payment"`
		InvoiceAmount     string          `json:"invoice_amount"`      // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
		OutTradeNo        string          `json:"out_trade_no"`        // 创建交易传入的商户订单号
		TradeNo           string          `json:"trade_no"`            // 支付宝交易号
		PointAmount       string          `json:"point_amount"`        // 积分支付的金额，单位为元，两位小数。
		ReceiptAmount     string          `json:"receipt_amount"`      // 实收金额，单位为元，两位小数
		StoreName         string          `json:"store_name"`          // 发生支付交易的商户门店名称
		TotalAmount       string          `json:"total_amount"`        // 发该笔退款所对应的交易的订单金额
		VoucherDetailList []VoucherDetail `json:"voucher_detail_list"` // 本交易支付时使用的所有优惠券信息
	} `json:"alipay_trade_pay_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradePayResponse) IsSuccess() bool {
	if this.AliPayTradePay.Msg == "Success" {
		return true
	}
	return false
}
