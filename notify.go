package alipay

import (
	"errors"
	"io"
	"net/http"
	"net/url"
)

var (
	kSuccess = []byte("success")
)

func (c *Client) NotifyVerify(partnerId, notifyId string) bool {
	var values = url.Values{}
	values.Add("service", "notify_verify")
	values.Add("partner", partnerId)
	values.Add("notify_id", notifyId)
	req, err := http.NewRequest(http.MethodGet, c.notifyVerifyHost+"?"+values.Encode(), nil)
	if err != nil {
		return false
	}

	rsp, err := c.Client.Do(req)
	if err != nil {
		return false
	}
	defer rsp.Body.Close()

	data, err := io.ReadAll(rsp.Body)
	if err != nil {
		return false
	}
	if string(data) == "true" {
		return true
	}
	return false
}

// GetTradeNotification
// Deprecated: use DecodeNotification instead.
func (c *Client) GetTradeNotification(req *http.Request) (notification *Notification, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}
	return c.DecodeNotification(req.Form)
}

func (c *Client) DecodeNotification(values url.Values) (notification *Notification, err error) {
	if err = c.VerifySign(values); err != nil {
		return nil, err
	}

	notification = &Notification{}
	notification.AppId = values.Get("app_id")
	notification.AuthAppId = values.Get("auth_app_id")
	notification.NotifyId = values.Get("notify_id")
	notification.NotifyType = values.Get("notify_type")
	notification.NotifyTime = values.Get("notify_time")
	notification.TradeNo = values.Get("trade_no")
	notification.TradeStatus = TradeStatus(values.Get("trade_status"))
	notification.RefundStatus = values.Get("refund_status")
	notification.RefundReason = values.Get("refund_reason")
	notification.RefundAmount = values.Get("refund_amount")
	notification.TotalAmount = values.Get("total_amount")
	notification.ReceiptAmount = values.Get("receipt_amount")
	notification.InvoiceAmount = values.Get("invoice_amount")
	notification.BuyerPayAmount = values.Get("buyer_pay_amount")
	notification.SellerId = values.Get("seller_id")
	notification.SellerEmail = values.Get("seller_email")
	notification.BuyerId = values.Get("buyer_id")
	notification.BuyerLogonId = values.Get("buyer_logon_id")
	notification.FundBillList = values.Get("fund_bill_list")
	notification.Charset = values.Get("charset")
	notification.PointAmount = values.Get("point_amount")
	notification.OutTradeNo = values.Get("out_trade_no")
	notification.OutRequestNo = values.Get("out_request_no")
	notification.OutBizNo = values.Get("out_biz_no")
	notification.GmtCreate = values.Get("gmt_create")
	notification.GmtPayment = values.Get("gmt_payment")
	notification.GmtRefund = values.Get("gmt_refund")
	notification.GmtClose = values.Get("gmt_close")
	notification.Subject = values.Get("subject")
	notification.Body = values.Get("body")
	notification.RefundFee = values.Get("refund_fee")
	notification.Version = values.Get("version")
	notification.SignType = values.Get("sign_type")
	notification.Sign = values.Get("sign")
	notification.PassbackParams = values.Get("passback_params")
	notification.VoucherDetailList = values.Get("voucher_detail_list")
	notification.AgreementNo = values.Get("agreement_no")
	notification.ExternalAgreementNo = values.Get("external_agreement_no")
	notification.DBackStatus = values.Get("dback_status")
	notification.DBackAmount = values.Get("dback_amount")
	notification.BankAckTime = values.Get("bank_ack_time")
	return notification, err
}

// AckNotification
// Deprecated: use ACKNotification instead.
func (c *Client) AckNotification(w http.ResponseWriter) {
	AckNotification(w)
}

// ACKNotification 返回异步通知成功处理的消息给支付宝
func (c *Client) ACKNotification(w http.ResponseWriter) {
	ACKNotification(w)
}

// AckNotification
// Deprecated: use ACKNotification instead.
func AckNotification(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(kSuccess)
}

// ACKNotification 返回异步通知成功处理的消息给支付宝
func ACKNotification(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write(kSuccess)
}
