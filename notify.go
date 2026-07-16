package alipay

import (
	"context"
	"errors"
	"io"
	"mime"
	"net/http"
	"net/url"
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
// Deprecated: use DecodeNotification or DecodeNotificationWithCharset instead.
func (c *Client) GetTradeNotification(req *http.Request) (notification *Notification, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	if err = req.ParseForm(); err != nil {
		return nil, err
	}
	return c.DecodeNotificationWithCharset(
		context.Background(),
		req.Form,
		charsetFromContentType(req.Header.Get("Content-Type")),
	)
}

func (c *Client) DecodeNotification(ctx context.Context, values url.Values) (notification *Notification, err error) {
	return c.DecodeNotificationWithCharset(ctx, values, "")
}

// DecodeNotificationWithCharset 解析并验签支付宝异步通知，并按声明的字符集
// 在验签成功后将通知字段转换为 UTF-8。values 中已签名的 charset
// 优先于显式传入的 charset，避免未签名的 HTTP 请求头改变解码方式。
func (c *Client) DecodeNotificationWithCharset(ctx context.Context, values url.Values, charset string) (notification *Notification, err error) {
	if err = c.VerifySign(ctx, values); err != nil {
		return nil, err
	}

	if signedCharset := values.Get("charset"); signedCharset != "" {
		charset = signedCharset
	}
	if values, err = decodeNotificationValues(values, charset); err != nil {
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

func decodeNotificationValues(values url.Values, charset string) (url.Values, error) {
	decoder := decoderForCharset(charset)
	if decoder == nil {
		return values, nil
	}

	decodedValues := make(url.Values, len(values))
	for key, valueList := range values {
		decodedList := make([]string, len(valueList))
		for index, value := range valueList {
			decoded, err := decoder.Bytes([]byte(value))
			if err != nil {
				return nil, err
			}
			decodedList[index] = string(decoded)
		}
		decodedValues[key] = decodedList
	}
	return decodedValues, nil
}

func charsetFromContentType(contentType string) string {
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return ""
	}
	return params["charset"]
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
