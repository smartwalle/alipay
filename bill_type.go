// 查询对账单下载地址

package alipay

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

func (this BillDownloadURLQuery) ExtJSONParamName() string {
	return "biz_content"
}

func (this BillDownloadURLQuery) ExtJSONParamValue() string {
	return marshal(this)
}

type BillDownloadURLQueryResponse struct {
	AliPayDataServiceBillDownloadURLQueryResponse struct {
		Code            string `json:"code"`
		Msg             string `json:"msg"`
		SubCode         string `json:"sub_code"`
		SubMsg          string `json:"sub_msg"`
		BillDownloadUrl string `json:"bill_download_url"`
	} `json:"alipay_data_dataservice_bill_downloadurl_query_response"`
	Sign string `json:"sign"`
}
