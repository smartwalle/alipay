package alipay

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_28/alipay.fund.trans.toaccount.transfer
// 单笔转账到支付宝账户接口请求结构
type FundTransToAccountTransfer struct {
	AppAuthToken  string `json:"-"`               // 可选
	OutBizNo      string `json:"out_biz_no"`      // 必选 商户转账唯一订单号
	PayeeType     string `json:"payee_type"`      // 必选 收款方账户类型,"ALIPAY_LOGONID":支付宝帐号
	PayeeAccount  string `json:"payee_account"`   // 必选 收款方账户。与payee_type配合使用
	Amount        string `json:"amount"`          // 必选 转账金额,元
	PayerShowName string `json:"payer_show_name"` // 可选 付款方显示姓名
	PayeeRealName string `json:"payee_real_name"` // 可选 收款方真实姓名,如果本参数不为空，则会校验该账户在支付宝登记的实名是否与收款方真实姓名一致。
	Remark        string `json:"remark"`          // 可选 转账备注,金额大于50000时必填
}

func (this FundTransToAccountTransfer) APIName() string {
	return "alipay.fund.trans.toaccount.transfer"
}

func (this FundTransToAccountTransfer) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this FundTransToAccountTransfer) ExtJSONParamName() string {
	return "biz_content"
}

func (this FundTransToAccountTransfer) ExtJSONParamValue() string {
	return marshal(this)
}

// 单笔转账到支付宝账户接口响应参数
type FundTransToAccountTransferRsp struct {
	Content struct {
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

func (this *FundTransToAccountTransferRsp) IsSuccess() bool {
	if this.Content.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_28/alipay.fund.trans.order.query/
// 查询转账订单接口请求参数
type FundTransOrderQuery struct {
	AppAuthToken string `json:"-"`                    // 可选
	OutBizNo     string `json:"out_biz_no,omitempty"` // 与 OrderId 二选一
	OrderId      string `json:"order_id,omitempty"`   // 与 OutBizNo 二选一
}

func (this FundTransOrderQuery) APIName() string {
	return "alipay.fund.trans.order.query"
}

func (this FundTransOrderQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this FundTransOrderQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this FundTransOrderQuery) ExtJSONParamValue() string {
	return marshal(this)
}

// 查询转账订单接口响应参数
type FundTransOrderQueryRsp struct {
	Content struct {
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

func (this *FundTransOrderQueryRsp) IsSuccess() bool {
	if this.Content.Code == K_SUCCESS_CODE {
		return true
	}
	return false
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_28/alipay.fund.auth.order.voucher.create/
// 资金授权发码接口
type FundAuthOrderVoucherCreate struct {
	NotifyURL         string `json:"-"`
	AppAuthToken      string `json:"-"`                             // 可选
	OutOrderNo        string `json:"out_order_no"`                  // 必选, 商户授权资金订单号，创建后不能修改，需要保证在商户端不重复。
	OutRequestNo      string `json:"out_request_no"`                // 必选, 商户本次资金操作的请求流水号，用于标示请求流水的唯一性，需要保证在商户端不重复。
	ProductCode       string `json:"product_code,omitempty"`        // 必选, 销售产品码，后续新接入预授权当面付的业务，本字段取值固定为PRE_AUTH。
	OrderTitle        string `json:"order_title"`                   // 必选, 业务订单的简单描述，如商品名称等 长度不超过100个字母或50个汉字
	Amount            string `json:"amount"`                        // 必选, 需要冻结的金额，单位为：元（人民币），精确到小数点后两位 取值范围：[0.01,100000000.00]
	PayeeUserId       string `json:"payee_user_id,omitempty"`       // 可选, 收款方的支付宝唯一用户号,以2088开头的16位纯数字组成，如果非空则会在支付时校验交易的的收款方与此是否一致，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayeeLogonId      string `json:"payee_logon_id,omitempty"`      // 可选, 收款方支付宝账号（Email或手机号），如果收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)同时传递，则以用户号(payee_user_id)为准，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayTimeout        string `json:"pay_timeout,omitempty"`         // 可选, 该笔订单允许的最晚付款时间，逾期将关闭该笔订单 取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m 如果为空，默认15m
	ExtraParam        string `json:"extra_param,omitempty"`         // 可选, 业务扩展参数，用于商户的特定业务信息的传递，json格式。 1.授权业务对应的类目，key为category，value由支付宝分配，比如充电桩业务传 "CHARGE_PILE_CAR"； 2. 外部商户的门店编号，key为outStoreCode，可选； 3. 外部商户的门店简称，key为outStoreAlias，可选。
	TransCurrency     string `json:"trans_currency,omitempty"`      // 可选, 标价币种, amount 对应的币种单位。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP
	SettleCurrency    string `json:"settle_currency,omitempty"`     // 可选, 商户指定的结算币种。支持澳元：AUD, 新西兰元：NZD, 台币：TWD, 美元：USD, 欧元：EUR, 英镑：GBP
	EnablePayChannels string `json:"enable_pay_channels,omitempty"` // 可选, 商户可用该参数指定用户可使用的支付渠道，本期支持商户可支持三种支付渠道，余额宝（MONEY_FUND）、花呗（PCREDIT_PAY）以及芝麻信用（CREDITZHIMA）。商户可设置一种支付渠道，也可设置多种支付渠道。
}

func (this FundAuthOrderVoucherCreate) APIName() string {
	return "alipay.fund.auth.order.voucher.create"
}

func (this FundAuthOrderVoucherCreate) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this FundAuthOrderVoucherCreate) ExtJSONParamName() string {
	return "biz_content"
}

func (this FundAuthOrderVoucherCreate) ExtJSONParamValue() string {
	return marshal(this)
}

type FundAuthOrderVoucherCreateRsp struct {
	Content struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		OutOrderNo   string `json:"out_order_no"`
		OutRequestNo string `json:"out_request_no"`
		CodeType     string `json:"code_type"`
		CodeValue    string `json:"code_value"`
		CodeURL      string `json:"code_url"`
	} `json:"alipay_fund_auth_order_voucher_create_response"`
	Sign string `json:"sign"`
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_28/alipay.fund.auth.order.freeze/
// 资金授权冻结接口
type FundAuthOrderFreeze struct {
	NotifyURL    string `json:"-"`
	AppAuthToken string `json:"-"`                        // 可选
	AuthCode     string `json:"auth_code"`                // 必选, 支付授权码，25~30开头的长度为16~24位的数字，实际字符串长度以开发者获取的付款码长度为准
	AuthCodeType string `json:"auth_code_type"`           // 必选, 授权码类型 目前仅支持"bar_code"
	OutOrderNo   string `json:"out_order_no"`             // 必选, 商户授权资金订单号 ,不能包含除中文、英文、数字以外的字符，创建后不能修改，需要保证在商户端不重复。
	OutRequestNo string `json:"out_request_no"`           // 必选, 商户本次资金操作的请求流水号，用于标示请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。
	OrderTitle   string `json:"order_title"`              // 必选, 业务订单的简单描述，如商品名称等 长度不超过100个字母或50个汉字
	Amount       string `json:"amount"`                   // 必选, 需要冻结的金额，单位为：元（人民币），精确到小数点后两位 取值范围：[0.01,100000000.00]
	PayeeLogonId string `json:"payee_logon_id,omitempty"` // 可选, 收款方支付宝账号（Email或手机号），如果收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)同时传递，则以用户号(payee_user_id)为准，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayeeUserId  string `json:"payee_user_id,omitempty"`  // 可选, 收款方的支付宝唯一用户号,以2088开头的16位纯数字组成，如果非空则会在支付时校验交易的的收款方与此是否一致，如果商户有勾选花呗渠道，收款方支付宝登录号(payee_logon_id)和用户号(payee_user_id)不能同时为空。
	PayTimeout   string `json:"pay_timeout,omitempty"`    // 可选, 该笔订单允许的最晚付款时间，逾期将关闭该笔订单 取值范围：1m～15d。m-分钟，h-小时，d-天。 该参数数值不接受小数点， 如 1.5h，可转换为90m 如果为空，默认15m
	ExtraParam   string `json:"extra_param,omitempty"`    // 可选, 业务扩展参数，用于商户的特定业务信息的传递，json格式。 1.授权业务对应的类目，key为category，value由支付宝分配，比如充电桩业务传 "CHARGE_PILE_CAR"； 2. 外部商户的门店编号，key为outStoreCode，可选； 3. 外部商户的门店简称，key为outStoreAlias，可选。
	ProductCode  string `json:"product_code,omitempty"`   // 可选, 销售产品码，后续新接入预授权当面付的业务，本字段取值固定为PRE_AUTH。
}

func (this FundAuthOrderFreeze) APIName() string {
	return "alipay.fund.auth.order.voucher.create"
}

func (this FundAuthOrderFreeze) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this FundAuthOrderFreeze) ExtJSONParamName() string {
	return "biz_content"
}

func (this FundAuthOrderFreeze) ExtJSONParamValue() string {
	return marshal(this)
}

type FundAuthOrderFreezeRsp struct {
	Content struct {
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
	} `json:"alipay_fund_auth_order_freeze_response"`
	Sign string `json:"sign"`
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_28/alipay.fund.auth.order.unfreeze/
// 资金授权解冻接口
type FundAuthOrderUnfreeze struct {
	NotifyURL    string `json:"-"`
	AuthNo       string `json:"auth_no"`               // 必选,支付宝资金授权订单号,支付宝冻结时返回的交易号，数字格式 2016101210002001810258115912
	AppAuthToken string `json:"-"`                     // 可选
	OutRequestNo string `json:"out_request_no"`        // 必选, 商户本次资金操作的请求流水号，用于标示请求流水的唯一性，不能包含除中文、英文、数字以外的字符，需要保证在商户端不重复。
	Amount       string `json:"amount"`                // 必选, 本次操作解冻的金额，单位为：元（人民币），精确到小数点后两位，取值范围：[0.01,100000000.00]
	Remark       string `json:"remark"`                // 必选, 商户对本次解冻操作的附言描述，长度不超过100个字母或50个汉字
	ExtraParam   string `json:"extra_param,omitempty"` // 可选, 解冻扩展信息，json格式；unfreezeBizInfo 目前为芝麻消费字段，支持Key值如下： "bizComplete":"true" -- 选填：标识本次解冻用户是否履约，如果true信用单会完结为COMPLETE
}

func (this FundAuthOrderUnfreeze) APIName() string {
	return "alipay.fund.auth.order.unfreeze"
}

func (this FundAuthOrderUnfreeze) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this FundAuthOrderUnfreeze) ExtJSONParamName() string {
	return "biz_content"
}

func (this FundAuthOrderUnfreeze) ExtJSONParamValue() string {
	return marshal(this)
}

type FundAuthOrderUnfreezeRsp struct {
	Content struct {
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
		GMTTrans     string `json:"gmt_trans"`
		CreditAmount string `json:"credit_amount"`
		FundAmount   string `json:"fund_amount"`
	} `json:"alipay_fund_auth_order_unfreeze_response"`
	Sign string `json:"sign"`
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.cancel/
// 资金授权撤销接口
type FundAuthOperationCancel struct {
	NotifyURL    string `json:"-"`
	AppAuthToken string `json:"-"`                        // 可选
	AuthNo       string `json:"auth_no,omitempty"`        // 特殊可选, 支付宝授权资金订单号，与商户的授权资金订单号不能同时为空，二者都存在时，以支付宝资金授权订单号为准，该参数与支付宝授权资金操作流水号配对使用。
	OutOrderNo   string `json:"out_order_no,omitempty"`   // 特殊可选,  商户的授权资金订单号，与支付宝的授权资金订单号不能同时为空，二者都存在时，以支付宝的授权资金订单号为准，该参数与商户的授权资金操作流水号配对使用。
	OperationId  string `json:"operation_id,omitempty"`   // 特殊可选, 支付宝的授权资金操作流水号，与商户的授权资金操作流水号不能同时为空，二者都存在时，以支付宝的授权资金操作流水号为准，该参数与支付宝授权资金订单号配对使用。
	OutRequestNo string `json:"out_request_no,omitempty"` // 特殊可选, 商户的授权资金操作流水号，与支付宝的授权资金操作流水号不能同时为空，二者都存在时，以支付宝的授权资金操作流水号为准，该参数与商户的授权资金订单号配对使用。
	Remark       string `json:"remark"`                   // 必选, 商户对本次撤销操作的附言描述，长度不超过100个字母或50个汉字
}

func (this FundAuthOperationCancel) APIName() string {
	return "alipay.fund.auth.operation.cancel"
}

func (this FundAuthOperationCancel) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	return m
}

func (this FundAuthOperationCancel) ExtJSONParamName() string {
	return "biz_content"
}

func (this FundAuthOperationCancel) ExtJSONParamValue() string {
	return marshal(this)
}

type FundAuthOperationCancelRsp struct {
	Content struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		AuthNo       string `json:"auth_no"`
		OutOrderNo   string `json:"out_order_no"`
		OperationId  string `json:"operation_id"`
		OutRequestNo string `json:"out_request_no"`
		Action       string `json:"action"`
	} `json:"alipay_fund_auth_operation_cancel_response"`
	Sign string `json:"sign"`
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_28/alipay.fund.auth.operation.detail.query/
// 资金授权操作查询接口
type FundAuthOperationDetailQuery struct {
	AppAuthToken string `json:"-"`              // 可选
	AuthNo       string `json:"auth_no"`        // 特殊可选, 支付宝授权资金订单号，与商户的授权资金订单号不能同时为空，二者都存在时，以支付宝资金授权订单号为准，该参数与支付宝授权资金操作流水号配对使用。
	OutOrderNo   string `json:"out_order_no"`   // 特殊可选, 商户的授权资金订单号，与支付宝的授权资金订单号不能同时为空，二者都存在时，以支付宝的授权资金订单号为准，该参数与商户的授权资金操作流水号配对使用。
	OperationId  string `json:"operation_id"`   // 特殊可选, 支付宝的授权资金操作流水号，与商户的授权资金操作流水号不能同时为空，二者都存在时，以支付宝的授权资金操作流水号为准，该参数与支付宝授权资金订单号配对使用。
	OutRequestNo string `json:"out_request_no"` // 特殊可选, 商户的授权资金操作流水号，与支付宝的授权资金操作流水号不能同时为空，二者都存在时，以支付宝的授权资金操作流水号为准，该参数与商户的授权资金订单号配对使用。
}

func (this FundAuthOperationDetailQuery) APIName() string {
	return "alipay.fund.auth.operation.detail.query"
}

func (this FundAuthOperationDetailQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this FundAuthOperationDetailQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this FundAuthOperationDetailQuery) ExtJSONParamValue() string {
	return marshal(this)
}

type FundAuthOperationDetailQueryRsp struct {
	Content struct {
		AuthNo                  string `json:"auth_no"`
		OutOrderNo              string `json:"out_order_no"`
		TotalFreezeAmount       string `json:"total_freeze_amount"`
		RestAmount              string `json:"rest_amount"`
		TotalPayAmount          string `json:"total_pay_amount"`
		OrderTitle              string `json:"order_title"`
		PayerLogonId            string `json:"payer_logon_id"`
		PayerUserId             string `json:"payer_user_id"`
		ExtraParam              string `json:"extra_param"`
		OperationId             string `json:"operation_id"`
		OutRequestNo            string `json:"out_request_no"`
		Amount                  string `json:"amount"`
		OperationType           string `json:"operation_type"`
		Status                  string `json:"status"`
		Remark                  string `json:"remark"`
		GMTCreate               string `json:"gmt_create"`
		GMTTrans                string `json:"gmt_trans"`
		PreAuthType             string `json:"pre_auth_type"`
		TransCurrency           string `json:"trans_currency"`
		TotalFreezeCreditAmount string `json:"total_freeze_credit_amount"`
		TotalFreezeFundAmount   string `json:"total_freeze_fund_amount"`
		TotalPayCreditAmount    string `json:"total_pay_credit_amount"`
		TotalPayFundAmount      string `json:"total_pay_fund_amount"`
		RestCreditAmount        string `json:"rest_credit_amount"`
		RestFundAmount          string `json:"rest_fund_amount"`
		CreditAmount            string `json:"credit_amount"`
		FundAmount              string `json:"fund_amount"`
	} `json:"alipay_fund_auth_operation_detail_query_response"`
	Sign string `json:"sign"`
}

////////////////////////////////////////////////////////////////////////////////
// https://docs.open.alipay.com/api_28/alipay.fund.auth.order.app.freeze
// 线上资金授权冻结接口请求参数
type FundAuthOrderAppFreeze struct {
	NotifyURL         string `json:"-"`
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

func (this FundAuthOrderAppFreeze) APIName() string {
	return "alipay.fund.auth.order.app.freeze"
}

func (this FundAuthOrderAppFreeze) Params() map[string]string {
	var m = make(map[string]string)
	if this.AppAuthToken != "" {
		m["app_auth_token"] = this.AppAuthToken
	}
	m["notify_url"] = this.NotifyURL
	return m
}

func (this FundAuthOrderAppFreeze) ExtJSONParamName() string {
	return "biz_content"
}

func (this FundAuthOrderAppFreeze) ExtJSONParamValue() string {
	return marshal(this)
}

type FundAuthOrderAppFreezeRsp struct {
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
