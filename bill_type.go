package alipay

// BillDownloadURLQuery 查询对账单下载地址接口请求参数 https://docs.open.alipay.com/api_15/alipay.data.dataservice.bill.downloadurl.query
type BillDownloadURLQuery struct {
	AppAuthToken string `json:"-"`         // 可选
	BillType     string `json:"bill_type"` // 必选 账单类型，商户通过接口或商户经开放平台授权后其所属服务商通过接口可以获取以下账单类型：trade、signcustomer；trade指商户基于支付宝交易收单的业务账单；signcustomer是指基于商户支付宝余额收入及支出等资金变动的帐务账单。
	BillDate     string `json:"bill_date"` // 必选 账单时间：日账单格式为yyyy-MM-dd，最早可下载2016年1月1日开始的日账单；月账单格式为yyyy-MM，最早可下载2016年1月开始的月账单。
}

func (this BillDownloadURLQuery) APIName() string {
	return "alipay.data.dataservice.bill.downloadurl.query"
}

func (this BillDownloadURLQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

// BillDownloadURLQueryRsp 查询对账单下载地址接口响应参数
type BillDownloadURLQueryRsp struct {
	Content struct {
		Code            Code   `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code"`
		SubMsg          string `json:"sub_msg"`
		BillDownloadUrl string `json:"bill_download_url"`
	} `json:"alipay_data_dataservice_bill_downloadurl_query_response"`
	Sign string `json:"sign"`
}

// BillBalanceQuery 支付宝商家账户当前余额查询 https://opendocs.alipay.com/apis/api_15/alipay.data.bill.balance.query
type BillBalanceQuery struct {
	AppAuthToken string `json:"-"` // 可选
}

func (this BillBalanceQuery) APIName() string {
	return "alipay.data.bill.balance.query"
}

func (this BillBalanceQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

// BillBalanceQueryRsp 支付宝商家账户当前余额查询响应参数
type BillBalanceQueryRsp struct {
	Content struct {
		Code            Code   `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code"`
		SubMsg          string `json:"sub_msg"`
		TotalAmount     string `json:"total_amount"`
		AvailableAmount string `json:"available_amount"`
		FreezeAmount    string `json:"freeze_amount"`
	} `json:"alipay_data_bill_balance_query_response"`
	Sign string `json:"sign"`
}
