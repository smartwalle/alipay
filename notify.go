package alipay

import (
	"errors"
	"io"
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

	data, err := io.ReadAll(rep.Body)
	if err != nil {
		return false
	}
	if string(data) == "true" {
		return true
	}
	return false
}

func (this *Client) GetTradeNotification(req *http.Request) (notification *TradeNotification, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}

	notification = &TradeNotification{}
	notification.AppId = req.FormValue("app_id")
	notification.AuthAppId = req.FormValue("auth_app_id")
	notification.NotifyId = req.FormValue("notify_id")
	notification.NotifyType = req.FormValue("notify_type")
	notification.NotifyTime = req.FormValue("notify_time")
	notification.TradeNo = req.FormValue("trade_no")
	notification.TradeStatus = TradeStatus(req.FormValue("trade_status"))
	notification.TotalAmount = req.FormValue("total_amount")
	notification.ReceiptAmount = req.FormValue("receipt_amount")
	notification.InvoiceAmount = req.FormValue("invoice_amount")
	notification.BuyerPayAmount = req.FormValue("buyer_pay_amount")
	notification.SellerId = req.FormValue("seller_id")
	notification.SellerEmail = req.FormValue("seller_email")
	notification.BuyerId = req.FormValue("buyer_id")
	notification.BuyerLogonId = req.FormValue("buyer_logon_id")
	notification.FundBillList = req.FormValue("fund_bill_list")
	notification.Charset = req.FormValue("charset")
	notification.PointAmount = req.FormValue("point_amount")
	notification.OutTradeNo = req.FormValue("out_trade_no")
	notification.OutBizNo = req.FormValue("out_biz_no")
	notification.GmtCreate = req.FormValue("gmt_create")
	notification.GmtPayment = req.FormValue("gmt_payment")
	notification.GmtRefund = req.FormValue("gmt_refund")
	notification.GmtClose = req.FormValue("gmt_close")
	notification.Subject = req.FormValue("subject")
	notification.Body = req.FormValue("body")
	notification.RefundFee = req.FormValue("refund_fee")
	notification.Version = req.FormValue("version")
	notification.SignType = req.FormValue("sign_type")
	notification.Sign = req.FormValue("sign")
	notification.PassbackParams = req.FormValue("passback_params")
	notification.VoucherDetailList = req.FormValue("voucher_detail_list")
	notification.AgreementNo = req.FormValue("agreement_no")
	notification.ExternalAgreementNo = req.FormValue("external_agreement_no")

	//if len(noti.NotifyId) == 0 {
	//	return nil, errors.New("不是有效的 Notify")
	//}

	ok, err := this.VerifySign(req.Form)
	if ok == false {
		return nil, err
	}
	return notification, err
}

func (this *Client) AckNotification(w http.ResponseWriter) {
	AckNotification(w)
}

func AckNotification(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(kSuccess)
}
