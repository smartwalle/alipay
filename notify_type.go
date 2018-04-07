package alipay

import (
	"errors"
	"net/http"
)

// https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.8AmJwg&treeId=203&articleId=105286&docType=1
type TradeNotification struct {
	AuthAppId         string `json:"auth_app_id"`         // App Id
	NotifyTime        string `json:"notify_time"`         // 通知时间
	NotifyType        string `json:"notify_type"`         // 通知类型
	NotifyId          string `json:"notify_id"`           // 通知校验ID
	AppId             string `json:"app_id"`              // 开发者的app_id
	Charset           string `json:"charset"`             // 编码格式
	Version           string `json:"version"`             // 接口版本
	SignType          string `json:"sign_type"`           // 签名类型
	Sign              string `json:"sign"`                // 签名
	TradeNo           string `json:"trade_no"`            // 支付宝交易号
	OutTradeNo        string `json:"out_trade_no"`        // 商户订单号
	OutBizNo          string `json:"out_biz_no"`          // 商户业务号
	BuyerId           string `json:"buyer_id"`            // 买家支付宝用户号
	BuyerLogonId      string `json:"buyer_logon_id"`      // 买家支付宝账号
	SellerId          string `json:"seller_id"`           // 卖家支付宝用户号
	SellerEmail       string `json:"seller_email"`        // 卖家支付宝账号
	TradeStatus       string `json:"trade_status"`        // 交易状态
	TotalAmount       string `json:"total_amount"`        // 订单金额
	ReceiptAmount     string `json:"receipt_amount"`      // 实收金额
	InvoiceAmount     string `json:"invoice_amount"`      // 开票金额
	BuyerPayAmount    string `json:"buyer_pay_amount"`    // 付款金额
	PointAmount       string `json:"point_amount"`        // 集分宝金额
	RefundFee         string `json:"refund_fee"`          // 总退款金额
	Subject           string `json:"subject"`             // 总退款金额
	Body              string `json:"body"`                // 商品描述
	GmtCreate         string `json:"gmt_create"`          // 交易创建时间
	GmtPayment        string `json:"gmt_payment"`         // 交易付款时间
	GmtRefund         string `json:"gmt_refund"`          // 交易退款时间
	GmtClose          string `json:"gmt_close"`           // 交易结束时间
	FundBillList      string `json:"fund_bill_list"`      // 支付金额信息
	PassbackParams    string `json:"passback_params"`     // 回传参数
	VoucherDetailList string `json:"voucher_detail_list"` // 优惠券信息
}

func (this *AliPay) GetTradeNotification(req *http.Request) (*TradeNotification, error) {
	return GetTradeNotification(req, this.AliPayPublicKey)
}

func GetTradeNotification(req *http.Request, aliPayPublicKey []byte) (noti *TradeNotification, err error) {
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}

	noti = &TradeNotification{}
	noti.AppId = req.FormValue("app_id")
	noti.AuthAppId = req.FormValue("auth_app_id")
	noti.NotifyId = req.FormValue("notify_id")
	noti.NotifyType = req.FormValue("notify_type")
	noti.NotifyTime = req.FormValue("notify_time")
	noti.TradeNo = req.FormValue("trade_no")
	noti.TradeStatus = req.FormValue("trade_status")
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

	if len(noti.NotifyId) == 0 {
		return nil, errors.New("不是有效的 Notify")
	}

	ok, err := verifySign(req.Form, aliPayPublicKey)
	if ok == false {
		return nil, err
	}
	return noti, err
}
