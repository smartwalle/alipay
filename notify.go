package alipay

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	kSuccess = []byte("success")
)

func (c *Client) NotifyVerify(ctx context.Context, partnerId, notifyId string) bool {
	var values = url.Values{}
	values.Add("service", "notify_verify")
	values.Add("partner", partnerId)
	values.Add("notify_id", notifyId)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.notifyVerifyHost+"?"+values.Encode(), nil)
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
	return c.DecodeNotification(context.Background(), req.Form)
}

func (c *Client) DecodeNotification(ctx context.Context, values url.Values) (notification *Notification, err error) {
	if err = c.VerifySign(ctx, values); err != nil {
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
	notification.BuyerOpenId = values.Get("buyer_open_id")
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
	notification.SendBackFee = values.Get("send_back_fee")
	return notification, err
}

// GetTradeNotificationWithCharset parses notification and decodes GBK fields to UTF-8.
// It extracts charset from Content-Type header and converts GBK-encoded fields after signature verification.
func (c *Client) GetTradeNotificationWithCharset(req *http.Request) (notification *Notification, err error) {
	if req == nil {
		return nil, errors.New("request is nil")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}
	charset := parseCharsetFromContentType(req.Header.Get("Content-Type"))
	return c.DecodeNotificationWithCharset(req.Form, charset)
}

// DecodeNotificationWithCharset decodes notification and converts GBK fields to UTF-8 based on charset.
// If charset is GBK/GB2312/GB18030, it converts Subject, Body, BuyerLogonId, PassbackParams, RefundReason.
func (c *Client) DecodeNotificationWithCharset(values url.Values, charset string) (notification *Notification, err error) {
	notification, err = c.DecodeNotification(values)
	if err != nil {
		return nil, err
	}

	charset = strings.ToUpper(charset)
	if charset == "GBK" || charset == "GB2312" || charset == "GB18030" {
		notification.Subject = decodeGBK(notification.Subject)
		notification.Body = decodeGBK(notification.Body)
		notification.BuyerLogonId = decodeGBK(notification.BuyerLogonId)
		notification.PassbackParams = decodeGBK(notification.PassbackParams)
		notification.RefundReason = decodeGBK(notification.RefundReason)
	}

	return notification, nil
}

// parseCharsetFromContentType extracts charset from Content-Type header.
func parseCharsetFromContentType(contentType string) string {
	for _, part := range strings.Split(contentType, ";") {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(strings.ToLower(part), "charset=") {
			return strings.TrimPrefix(part, "charset=")
		}
	}
	return ""
}

// decodeGBK converts GBK-encoded string to UTF-8.
func decodeGBK(s string) string {
	if s == "" {
		return s
	}
	reader := transform.NewReader(strings.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	result, err := io.ReadAll(reader)
	if err != nil {
		return s
	}
	return string(result)
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
