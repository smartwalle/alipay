package alipay

import "encoding/json"

////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/doc2/detail.htm?treeId=270&articleId=105901&docType=1
type AliPayTradePagePay struct {
	NotifyURL string `json:"-"`
	ReturnURL string `json:"-"`

	// biz content，这四个参数是必须的
	Subject     string `json:"subject"`      // 订单标题
	OutTradeNo  string `json:"out_trade_no"` // 商户订单号，64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	TotalAmount string `json:"total_amount"` // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	ProductCode string `json:"product_code"` // 销售产品码，与支付宝签约的产品码名称。 注：目前仅支持FAST_INSTANT_TRADE_PAY

	Body               string `json:"body,omitempty"`                 // 订单描述
	GoodsDetail        string `json:"goods_detail,omitempty"`         // 订单包含的商品列表信息，Json格式，详见商品明细说明
	PassbackParams     string `json:"passback_params,omitempty"`      // 公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝只会在异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝
	ExtendParams       string `json:"extend_params,omitempty"`        // 业务扩展参数，详见业务扩展参数说明
	GoodsType          string `json:"goods_type,omitempty"`           // 商品主类型：0—虚拟类商品，1—实物类商品（默认） 注：虚拟类商品不支持使用花呗渠道
	TimeoutExpress     string `json:"timeout_express,omitempty"`      // 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m
	EnablePayChannels  string `json:"enable_pay_channels,omitempty"`  // 可用渠道，用户只能在指定渠道范围内支付 当有多个渠道时用“,”分隔   注：与disable_pay_channels互斥
	DisablePayChannels string `json:"disable_pay_channels,omitempty"` // 禁用渠道，用户不可用指定渠道支付 当有多个渠道时用“,”分隔  注：与enable_pay_channels互斥
	AuthToken          string `json:"auth_token,omitempty"`           // 针对用户授权接口，获取用户相关数据时，用于标识用户授权关系
	QRPayMode          string `json:"qr_pay_mode,omitempty"`          // PC扫码支付的方式，支持前置模式和跳转模式。
	QRCodeWidth        string `json:"qrcode_width,omitempty"`         // 商户自定义二维码宽度 注：qr_pay_mode=4时该参数生效
}

func (this AliPayTradePagePay) APIName() string {
	return "alipay.trade.page.pay"
}

func (this AliPayTradePagePay) Params() map[string]string {
	var m = make(map[string]string)
	m["notify_url"] = this.NotifyURL
	m["return_url"] = this.ReturnURL
	return m
}

func (this AliPayTradePagePay) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradePagePay) ExtJSONParamValue() string {
	var bytes, err = json.Marshal(this)
	if err != nil {
		return ""
	}
	return string(bytes)
}

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
		Code                string           `json:"code"`
		Msg                 string           `json:"msg"`
		SubCode             string           `json:"sub_code"`
		SubMsg              string           `json:"sub_msg"`
		BuyerLogonId        string           `json:"buyer_logon_id"`                // 买家支付宝账号
		BuyerPayAmount      float64          `json:"buyer_pay_amount,string"`       // 买家实付金额，单位为元，两位小数。
		BuyerUserId         string           `json:"buyer_user_id"`                 // 买家在支付宝的用户id
		InvoiceAmount       float64          `json:"invoice_amount,string"`         // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
		Openid              string           `json:"open_id"`                       // 买家支付宝用户号，该字段将废弃，不要使用
		OutTradeNo          string           `json:"out_trade_no"`                  // 商家订单号
		PointAmount         float64          `json:"point_amount,string"`           // 积分支付的金额，单位为元，两位小数。
		ReceiptAmount       float64          `json:"receipt_amount,string"`         // 实收金额，单位为元，两位小数
		SendPayDate         string           `json:"send_pay_date"`                 // 本次交易打款给卖家的时间
		TotalAmount         float64          `json:"total_amount,string"`           // 交易的订单金额
		TradeNo             string           `json:"trade_no"`                      // 支付宝交易号
		TradeStatus         string           `json:"trade_status"`                  // 交易状态
		AliPayStoreId       string           `json:"alipay_store_id"`               // 支付宝店铺编号
		StoreId             string           `json:"store_id"`                      // 商户门店编号
		TerminalId          string           `json:"terminal_id"`                   // 商户机具终端编号
		StoreName           string           `json:"store_name"`                    // 请求交易支付中的商户店铺的名称
		DiscountGoodsDetail string           `json:"discount_goods_detail"`         // 本次交易支付所使用的单品券优惠的商品优惠信息
		IndustrySepcDetail  string           `json:"industry_sepc_detail"`          // 行业特殊信息（例如在医保卡支付业务中，向用户返回医疗信息）。
		FundBillList        []*FundBill      `json:"fund_bill_list,omitempty"`      // 交易支付使用的资金渠道
		VoucherDetailList   []*VoucherDetail `json:"voucher_detail_list,omitempty"` // 本交易支付时使用的所有优惠券信息
	} `json:"alipay_trade_query_response"`
	Sign string `json:"sign"`
}

type FundBill struct {
	FundChannel string  `json:"fund_channel"`       // 交易使用的资金渠道，详见 支付渠道列表
	Amount      string  `json:"amount"`             // 该支付工具类型所使用的金额
	RealAmount  float64 `json:"real_amount,string"` // 渠道实际付款金额
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
	if this.AliPayTradeQuery.Code == K_SUCCESS_CODE {
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
		Code                 string              `json:"code"`
		Msg                  string              `json:"msg"`
		SubCode              string              `json:"sub_code"`
		SubMsg               string              `json:"sub_msg"`
		TradeNo              string              `json:"trade_no"`                          // 支付宝交易号
		OutTradeNo           string              `json:"out_trade_no"`                      // 商户订单号
		BuyerLogonId         string              `json:"buyer_logon_id"`                    // 用户的登录id
		BuyerUserId          string              `json:"buyer_user_id"`                     // 买家在支付宝的用户id
		FundChange           string              `json:"fund_change"`                       // 本次退款是否发生了资金变化
		RefundFee            string              `json:"refund_fee"`                        // 退款总金额
		GmtRefundPay         string              `json:"gmt_refund_pay"`                    // 退款支付时间
		StoreName            string              `json:"store_name"`                        // 交易在支付时候的门店名称
		RefundDetailItemList []*RefundDetailItem `json:"refund_detail_item_list,omitempty"` // 退款使用的资金渠道
	} `json:"alipay_trade_refund_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradeRefundResponse) IsSuccess() bool {
	if this.AliPayTradeRefund.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

type RefundDetailItem struct {
	FundChannel string `json:"fund_channel"` // 交易使用的资金渠道，详见 支付渠道列表
	Amount      string `json:"amount"`       // 该支付工具类型所使用的金额
	RealAmount  string `json:"real_amount"`  // 渠道实际付款金额
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
	if this.AliPayTradeFastpayRefundQueryResponse.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.kqrPAp&docType=4&apiId=1147
type AliPayTradeOrderSettle struct {
	AppAuthToken      string              `json:"-"`                  // 可选
	OutRequestNo      string              `json:"out_request_no"`     // 必须 结算请求流水号 开发者自行生成并保证唯一性
	TradeNo           string              `json:"trade_no"`           // 必须 支付宝订单号
	RoyaltyParameters []*RoyaltyParameter `json:"royalty_parameters"` // 必须 分账明细信息
	OperatorId        string              `json:"operator_id"`        //可选 操作员id
}

func (this AliPayTradeOrderSettle) APIName() string {
	return "alipay.trade.order.settle"
}

func (this AliPayTradeOrderSettle) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayTradeOrderSettle) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeOrderSettle) ExtJSONParamValue() string {
	return marshal(this)
}

type RoyaltyParameter struct {
	TransOut         string  `json:"trans_out"`         // 可选 分账支出方账户，类型为userId，本参数为要分账的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。
	TransIn          string  `json:"trans_in"`          // 可选 分账收入方账户，类型为userId，本参数为要分账的支付宝账号对应的支付宝唯一用户号。以2088开头的纯16位数字。
	Amount           float64 `json:"amount"`            // 可选 分账的金额，单位为元
	AmountPercentage float64 `json:"amount_percentage"` // 可选 分账信息中分账百分比。取值范围为大于0，少于或等于100的整数。
	Desc             string  `json:"desc"`              // 可选 分账描述
}

//////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.CkYNiG&docType=4&apiId=1046
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
	SubMerchant          []*SubMerchantItem `json:"sub_merchant,omitempty"`
	MerchantOrderNo      string             `json:"merchant_order_no"`
}

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
		Code                string           `json:"code"`
		Msg                 string           `json:"msg"`
		SubCode             string           `json:"sub_code"`
		SubMsg              string           `json:"sub_msg"`
		BuyerLogonId        string           `json:"buyer_logon_id"`           // 买家支付宝账号
		BuyerPayAmount      string           `json:"buyer_pay_amount"`         // 买家实付金额，单位为元，两位小数。
		BuyerUserId         string           `json:"buyer_user_id"`            // 买家在支付宝的用户id
		CardBalance         string           `json:"card_balance"`             // 支付宝卡余额
		DiscountGoodsDetail string           `json:"discount_goods_detail"`    // 本次交易支付所使用的单品券优惠的商品优惠信息
		FundBillList        []*FundBill      `json:"fund_bill_list,omitempty"` // 交易支付使用的资金渠道
		GmtPayment          string           `json:"gmt_payment"`
		InvoiceAmount       string           `json:"invoice_amount"`                // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
		OutTradeNo          string           `json:"out_trade_no"`                  // 创建交易传入的商户订单号
		TradeNo             string           `json:"trade_no"`                      // 支付宝交易号
		PointAmount         string           `json:"point_amount"`                  // 积分支付的金额，单位为元，两位小数。
		ReceiptAmount       string           `json:"receipt_amount"`                // 实收金额，单位为元，两位小数
		StoreName           string           `json:"store_name"`                    // 发生支付交易的商户门店名称
		TotalAmount         string           `json:"total_amount"`                  // 发该笔退款所对应的交易的订单金额
		VoucherDetailList   []*VoucherDetail `json:"voucher_detail_list,omitempty"` // 本交易支付时使用的所有优惠券信息
	} `json:"alipay_trade_pay_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradePayResponse) IsSuccess() bool {
	if this.AliPayTradePay.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

//////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/doc2/detail.htm?treeId=204&articleId=105462&docType=1
type AliPayTradeAppPay struct {
	NotifyURL string `json:"-"` // 可选

	Body           string `json:"body"`            // 可选 对一笔交易的具体描述信息。如果是多种商品，请将商品描述字符串累加传给body。
	Subject        string `json:"subject"`         // 必须 商品的标题/交易标题/订单标题/订单关键字等。
	OutTradeNo     string `json:"out_trade_no"`    // 必须 商户网站唯一订单号
	TimeoutExpress string `json:"timeout_express"` // 可选 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
	TotalAmount    string `json:"total_amount"`    // 必须 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	SellerId       string `json:"seller_id"`       // 可选 收款支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	ProductCode    string `json:"product_code"`    // 必须 销售产品码，商家和支付宝签约的产品码
	//GoodsType          string `json:"goods_type"`           // 可选 商品主类型：0—虚拟类商品，1—实物类商品 注：虚拟类商品不支持使用花呗渠道
	//PassbackParams     string `json:"passback_params"`      // 可选 公用回传参数，如果请求时传递了该参数，则返回给商户时会回传该参数。支付宝会在异步通知时将该参数原样返回。本参数必须进行UrlEncode之后才可以发送给支付宝
	//PromoParams        string `json:"promo_params"`         // 可选 优惠参数 注：仅与支付宝协商后可用
	//ExtendParams       string `json:"extend_params"`        // 可选 业务扩展参数，详见下面的“业务扩展参数说明”
	//EnablePayChannels  string `json:"enable_pay_channels"`  // 可选 	可用渠道，用户只能在指定渠道范围内支付  当有多个渠道时用“,”分隔 注：与disable_pay_channels互斥
	//DisablePayChannels string `json:"disable_pay_channels"` // 可选 禁用渠道，用户不可用指定渠道支付  当有多个渠道时用“,”分隔 注：与enable_pay_channels互斥
	//StoreId            string `json:"store_id"`             // 可选 商户门店编号
}

func (this AliPayTradeAppPay) APIName() string {
	return "alipay.trade.app.pay"
}

func (this AliPayTradeAppPay) Params() map[string]string {
	var m = make(map[string]string)
	m["notify_url"] = this.NotifyURL
	return m
}

func (this AliPayTradeAppPay) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeAppPay) ExtJSONParamValue() string {
	return marshal(this)
}

//////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.L57zdH&docType=4&apiId=862
type AliPayTradePreCreate struct {
	AppAuthToken string `json:"-"` // 可选
	NotifyURL    string `json:"-"` // 可选

	OutTradeNo         string             `json:"out_trade_no"`            // 必须 商户订单号,64个字符以内、只能包含字母、数字、下划线；需保证在商户端不重复
	Subject            string             `json:"subject"`                 // 必须 订单标题
	SellerId           string             `json:"seller_id"`               // 可选 卖家支付宝用户ID。 如果该值为空，则默认为商户签约账号对应的支付宝用户ID
	TotalAmount        string             `json:"total_amount"`            // 必须 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果同时传入了【打折金额】，【不可打折金额】，【订单总金额】三者，则必须满足如下条件：【订单总金额】=【打折金额】+【不可打折金额】
	DiscountableAmount string             `json:"discountable_amount"`     // 可选 可打折金额. 参与优惠计算的金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000] 如果该值未传入，但传入了【订单总金额】，【不可打折金额】则该值默认为【订单总金额】-【不可打折金额】
	Body               string             `json:"body"`                    // 可选 对交易或商品的描述
	GoodsDetail        []*GoodsDetailItem `json:"goods_detail,omitempty"`  // 可选 订单包含的商品列表信息.Json格式. 其它说明详见：“商品明细说明”
	ExtendParams       string             `json:"extend_params,omitempty"` // 业务扩展参数，详见业务扩展参数说明
	OperatorId         string             `json:"operator_id"`             // 可选 商户操作员编号
	StoreId            string             `json:"store_id"`                // 可选 商户门店编号
	TerminalId         string             `json:"terminal_id"`             // 可选 商户机具终端编号
	TimeoutExpress     string             `json:"timeout_express"`         // 可选 该笔订单允许的最晚付款时间，逾期将关闭交易。取值范围：1m～15d。m-分钟，h-小时，d-天，1c-当天（1c-当天的情况下，无论交易何时创建，都在0点关闭）。 该参数数值不接受小数点， 如 1.5h，可转换为 90m。
}

func (this AliPayTradePreCreate) APIName() string {
	return "alipay.trade.precreate"
}

func (this AliPayTradePreCreate) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this AliPayTradePreCreate) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradePreCreate) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradePreCreateResponse struct {
	AliPayPreCreateResponse struct {
		Code       string `json:"code"`
		Msg        string `json:"msg"`
		SubCode    string `json:"sub_code"`
		SubMsg     string `json:"sub_msg"`
		OutTradeNo string `json:"out_trade_no"` // 创建交易传入的商户订单号
		QRCode     string `json:"qr_code"`      // 当前预下单请求生成的二维码码串，可以用二维码生成工具根据该码串值生成对应的二维码
	} `json:"alipay_trade_precreate_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradePreCreateResponse) IsSuccess() bool {
	if this.AliPayPreCreateResponse.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

//////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.UKvJeT&docType=4&apiId=866
type AliPayTradeCancel struct {
	AppAuthToken string `json:"-"` // 可选
	NotifyURL    string `json:"-"` // 可选

	OutTradeNo string `json:"out_trade_no"` // 原支付请求的商户订单号,和支付宝交易号不能同时为空
	TradeNo    string `json:"trade_no"`     // 支付宝交易号，和商户订单号不能同时为空
}

func (this AliPayTradeCancel) APIName() string {
	return "alipay.trade.cancel"
}

func (this AliPayTradeCancel) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this AliPayTradeCancel) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayTradeCancel) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayTradeCancelResponse struct {
	AliPayTradeCancelResponse struct {
		Code       string `json:"code"`
		Msg        string `json:"msg"`
		SubCode    string `json:"sub_code"`
		SubMsg     string `json:"sub_msg"`
		TradeNo    string `json:"trade_no"`     // 支付宝交易号
		OutTradeNo string `json:"out_trade_no"` // 创建交易传入的商户订单号
		RetryFlag  string `json:"retry_flag"`   // 是否需要重试
		Action     string `json:"action"`       // 本次撤销触发的交易动作 close：关闭交易，无退款 refund：产生了退款
	} `json:"alipay_trade_cancel_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradeCancelResponse) IsSuccess() bool {
	if this.AliPayTradeCancelResponse.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}
