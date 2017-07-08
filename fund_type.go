package alipay

////////////////////////////////////////////////////////////////////////////////
// https://doc.open.alipay.com/docs/api.htm?apiId=1321&docType=4
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
// https://doc.open.alipay.com/docs/api.htm?spm=a219a.7395905.0.0.SIkNrH&docType=4&apiId=1322
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
