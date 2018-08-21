package alipay

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_28/alipay.fund.trans.toaccount.transfer
// 单笔转账到支付宝账户接口请求结构
type AliPayFundTransToAccountTransfer struct {
	AppAuthToken  string `json:"-"`               // 可选
	OutBizNo      string `json:"out_biz_no"`      // 必选 商户转账唯一订单号
	PayeeType     string `json:"payee_type"`      // 必选 收款方账户类型,"ALIPAY_LOGONID":支付宝帐号
	PayeeAccount  string `json:"payee_account"`   // 必选 收款方账户。与payee_type配合使用
	Amount        string `json:"amount"`          // 必选 转账金额,元
	PayerShowName string `json:"payer_show_name"` // 可选 付款方显示姓名
	PayeeRealName string `json:"payee_real_name"` // 可选 收款方真实姓名,如果本参数不为空，则会校验该账户在支付宝登记的实名是否与收款方真实姓名一致。
	Remark        string `json:"remark"`          // 可选 转账备注,金额大于50000时必填
}

func (this AliPayFundTransToAccountTransfer) APIName() string {
	return "alipay.fund.trans.toaccount.transfer"
}

func (this AliPayFundTransToAccountTransfer) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayFundTransToAccountTransfer) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayFundTransToAccountTransfer) ExtJSONParamValue() string {
	return marshal(this)
}

// 单笔转账到支付宝账户接口响应参数
type AliPayFundTransToAccountTransferResponse struct {
	Body struct {
		Code     string `json:"code"`
		Msg      string `json:"msg"`
		SubCode  string `json:"sub_code"`
		SubMsg   string `json:"sub_msg"`
		OutBizNo string `json:"out_biz_no"` // 商户转账唯一订单号：发起转账来源方定义的转账单据号。请求时对应的参数，原样返回
		OrderId  string `json:"order_id"`   // 支付宝转账单据号，成功一定返回，失败可能不返回也可能返回
		PayDate  string `json:"pay_date"`   // 支付时间：格式为yyyy-MM-dd HH:mm:ss，仅转账成功返回
	} `json:"alipay_fund_trans_toaccount_transfer_response"`
	Sign string `json:"sign"`
}

func (this *AliPayFundTransToAccountTransferResponse) IsSuccess() bool {
	if this.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_28/alipay.fund.trans.order.query/
// 查询转账订单接口请求参数
type AliPayFundTransOrderQuery struct {
	AppAuthToken string `json:"-"`                    // 可选
	OutBizNo     string `json:"out_biz_no,omitempty"` // 与 OrderId 二选一
	OrderId      string `json:"order_id,omitempty"`   // 与 OutBizNo 二选一
}

func (this AliPayFundTransOrderQuery) APIName() string {
	return "alipay.fund.trans.order.query"
}

func (this AliPayFundTransOrderQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayFundTransOrderQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayFundTransOrderQuery) ExtJSONParamValue() string {
	return marshal(this)
}

// 查询转账订单接口响应参数
type AliPayFundTransOrderQueryResponse struct {
	Body struct {
		Code           string `json:"code"`
		Msg            string `json:"msg"`
		SubCode        string `json:"sub_code"`
		SubMsg         string `json:"sub_msg"`
		OutBizNo       string `json:"out_biz_no"`       // 发起转账来源方定义的转账单据号。 该参数的赋值均以查询结果中 的 out_biz_no 为准。 如果查询失败，不返回该参数
		OrderId        string `json:"order_id"`         // 支付宝转账单据号，查询失败不返回。
		Status         string `json:"status"`           // 转账单据状态
		PayDate        string `json:"pay_date"`         // 支付时间
		ArrivalTimeEnd string `json:"arrival_time_end"` // 预计到账时间，转账到银行卡专用
		OrderFree      string `json:"order_fee"`        // 预计收费金额（元），转账到银行卡专用
		FailReason     string `json:"fail_reason"`      // 查询到的订单状态为FAIL失败或REFUND退票时，返回具体的原因。
		ErrorCode      string `json:"error_code"`       // 查询失败时，本参数为错误代 码。 查询成功不返回。 对于退票订单，不返回该参数。
	} `json:"alipay_fund_trans_order_query_response"`
	Sign string `json:"sign"`
}

func (this *AliPayFundTransOrderQueryResponse) IsSuccess() bool {
	if this.Body.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_28/alipay.fund.auth.order.app.freeze
// 线上资金授权冻结接口请求参数
type AliPayFundAuthOrderAppFreeze struct {
	AppAuthToken      string `json:"-"`                             // 可选
	OutOrderNo        string `json:"out_order_no"`                  // 必选, 商户授权资金订单号 ,不能包含除中文、英文、数字以外的字符，创建后不能修改，需要保证在商户端不重复。
	OutRequestNo      string `json:"out_request_no"`                // 必选, 商户本次资金操作的请求流水号，用于标示请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。
	OrderTitle        string `json:"order_title"`                   // 必选, 业务订单的简单描述，如商品名称等 长度不超过100个字母或50个汉字
	Amount            string `json:"amount"`                        // 必选, 需要冻结的金额，单位为：元（人民币），精确到小数点后两位 取值范围：[0.01,100000000.00]
	ProductCode       string `json:"product_code"`                  // 必选, 销售产品码，新接入线上预授权的业务，本字段取值固定为PRE_AUTH_ONLINE 。
	PayeeLogonId      string `json:"payee_logon_id,omitempty"`      // 收款方支付宝账号（Email或手机号），如果收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)同时传递，则以用户号(payee_user_id)为准，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayeeUserId       string `json:"payee_user_id,omitempty"`       // 收款方的支付宝唯一用户号,以2088开头的16位纯数字组成，如果非空则会在支付时校验交易的的收款方与此是否一致，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayTimeout        string `json:"pay_timeout,omitempty"`         // 该笔订单允许的最晚付款时间，逾期将关闭该笔订单 取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m 如果为空，默认15m
	ExtraParam        string `json:"extra_param,omitempty"`         // 业务扩展参数，用于商户的特定业务信息的传递，json格式。 1.授权业务对应的类目，key为category，value由支付宝分配，比如充电桩业务传 "CHARGE_PILE_CAR"； 2. 外部商户的门店编号，key为outStoreCode，可选； 3. 外部商户的门店简称，key为outStoreAlias，可选。
	EnablePayChannels string `json:"enable_pay_channels,omitempty"` // 商户可用该参数指定用户可使用的支付渠道，本期支持商户可支持三种支付渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。商户可设置一种支付渠道，也可设置多种支付渠道。
}

func (this AliPayFundAuthOrderAppFreeze) APIName() string {
	return "alipay.fund.auth.order.app.freeze"
}

func (this AliPayFundAuthOrderAppFreeze) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this AliPayFundAuthOrderAppFreeze) ExtJSONParamName() string {
	return "biz_content"
}

func (this AliPayFundAuthOrderAppFreeze) ExtJSONParamValue() string {
	return marshal(this)
}

type AliPayFundAuthOrderAppFreezeResponse struct {
	Body struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		AuthNo       string `json:"auth_no"`
		OutOrderNo   string `json:"out_order_no"`
		OperationId  string `json:"operation_id"`
		OutRequestNo string `json:"out_request_no"`
		Amount       string `json:"amount"`
		Status       string `json:"status"`
		PayerUserId  string `json:"payer_user_id"`
		GMTTrans     string `json:"gmt_trans"`
		PreAuthType  string `json:"pre_auth_type"`
		CreditAmount string `json:"credit_amount"`
		FundAmount   string `json:"fund_amount"`
	} `json:"alipay_fund_auth_order_app_freeze_response"`
	Sign string `json:"sign"`
}
