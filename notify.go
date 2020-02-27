package alipay

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var (
	kSuccess = []byte("success")
)

func NewRequest(method, url string, params url.Values) (*http.Request, error) {
	var m = strings.ToUpper(method)
	var body io.Reader
	if m == "GET" || m == "HEAD" {
		if len(params) > 0 {
			if strings.Contains(url, "?") {
				url = url + "&" + params.Encode()
			} else {
				url = url + "?" + params.Encode()
			}
		}
	} else {
		body = strings.NewReader(params.Encode())
	}
	return http.NewRequest(m, url, body)
}

func (this *Client) NotifyVerify(partnerId, notifyId string) bool {
	var param = url.Values{}
	param.Add("service", "notify_verify")
	param.Add("partner", partnerId)
	param.Add("notify_id", notifyId)
	req, err := NewRequest("GET", this.notifyVerifyDomain, param)
	if err != nil {
		return false
	}

	rep, err := this.Client.Do(req)
	if err != nil {
		return false
	}
	defer rep.Body.Close()

	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return false
	}
	if string(data) == "true" {
		return true
	}
	return false
}

func (this *Client) GetTradeNotification(req *http.Request) (noti *TradeNotification, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	noti = &TradeNotification{}
	noti.AppId = req.FormValue("app_id")
	noti.AuthAppId = req.FormValue("auth_app_id")
	noti.NotifyId = req.FormValue("notify_id")
	noti.NotifyType = req.FormValue("notify_type")
	noti.NotifyTime = req.FormValue("notify_time")
	noti.TradeNo = req.FormValue("trade_no")
	noti.TradeStatus = TradeStatus(req.FormValue("trade_status"))
	noti.TotalAmount = req.FormValue("total_amount")
	noti.ReceiptAmount = req.FormValue("receipt_amount")
	noti.InvoiceAmount = req.FormValue("invoice_amount")
	noti.BuyerPayAmount = req.FormValue("buyer_pay_amount")
	noti.SellerId = req.FormValue("seller_id")
	noti.SellerEmail = req.FormValue("seller_email")
	noti.BuyerId = req.FormValue("buyer_id")
	noti.BuyerLogonId = req.FormValue("buyer_logon_id")
	noti.FundBillList = req.FormValue("fund_bill_list")
	noti.Charset = req.FormValue("charset")
	noti.PointAmount = req.FormValue("point_amount")
	noti.OutTradeNo = req.FormValue("out_trade_no")
	noti.OutBizNo = req.FormValue("out_biz_no")
	noti.GmtCreate = req.FormValue("gmt_create")
	noti.GmtPayment = req.FormValue("gmt_payment")
	noti.GmtRefund = req.FormValue("gmt_refund")
	noti.GmtClose = req.FormValue("gmt_close")
	noti.Subject = req.FormValue("subject")
	noti.Body = req.FormValue("body")
	noti.RefundFee = req.FormValue("refund_fee")
	noti.Version = req.FormValue("version")
	noti.SignType = req.FormValue("sign_type")
	noti.Sign = req.FormValue("sign")
	noti.PassbackParams = req.FormValue("passback_params")
	noti.VoucherDetailList = req.FormValue("voucher_detail_list")

	//if len(noti.NotifyId) == 0 {
	//	return nil, errors.New("不是有效的 Notify")
	//}

	ok, err := this.VerifySign(req.Form)
	if ok == false {
		return nil, err
	}
	return noti, err
}

func (this *Client) AckNotification(w http.ResponseWriter) {
	AckNotification(w)
}

func AckNotification(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(kSuccess)
}
