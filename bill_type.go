package alipay

// BillDownloadURLQuery 查询对账单下载地址接口请求参数 https://docs.open.alipay.com/api_15/alipay.data.dataservice.bill.downloadurl.query
type BillDownloadURLQuery struct {
	AppAuthToken string `json:"-"` // 可选
	BillType     string `json:"bill_type"`
	BillDate     string `json:"bill_date"`
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
