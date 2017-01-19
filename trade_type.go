package alipay

type AliPayTradeQuery struct {
	AppAuthToken string `json:"-"`
	OutTradeNo   string `json:"out_trade_no"`
	TradeNo      string `json:"trade_no"`
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
		Code           string `json:"code"`
		Msg            string `json:"msg"`
		SubCode        string `json:"sub_code"`
		SubMsg         string `json:"sub_msg"`
		BuyerLogonId   string `json:"buyer_logon_id"`
		BuyerPayAmount string `json:"buyer_pay_amount"`
		BuyerUserId    string `json:"buyer_user_id"`
		InvoiceAmount  string `json:"invoice_amount"`
		Openid         string `json:"open_id"`
		OutTradeNo     string `json:"out_trade_no"`
		PointAmount    string `json:"point_amount"`
		ReceiptAmount  string `json:"receipt_amount"`
		SendPayDate    string `json:"send_pay_date"`
		TotalAmount    string `json:"total_amount"`
		TradeNo        string `json:"trade_no"`
		TradeStatus    string `json:"trade_status"`
	} `json:"alipay_trade_query_response"`
	Sign string `json:"sign"`
}

func (this *AliPayTradeQueryResponse) IsSuccess() (bool) {
	if this.AliPayTradeQuery.Msg == "Success" {
		return true
	}
	return false
}