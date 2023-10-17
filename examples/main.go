package main

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
	"log"
	"net/http"
)

var client *alipay.Client

const (
	kAppId      = "9021000122689420"
	kPrivateKey = "MIIEpAIBAAKCAQEAtrCYn039gp/IjLZsNl64QjHjPfAklyYZyIPSqlPNmColRMOAb2rbnLASRpP1VgT7YzA4JgJ1f/fhGsSBexpIi2BOZdwDexBtmfe9dFGzbWhpqAxzWnZazLde+BGOs6BGk0v1B0cWUqlv6wgSSo57Xu7xL56a34gDsBi1qoXnu4f1CzrvviRsNCgDDurNsAtkjLzoPuzHri6sThsQ7P3amb3zyG5xVxSGZRFKgPNoiiZpBpXPoEwMrHQRE8rmsmgdz+E4YL1xuD+ICQxCCTBOJwUuDzPt5wr793Pxgloqh0p3yPvOShmbMAxUtLiGgcZNqxy49ddG89egVyGxRJpimQIDAQABAoIBAEfGUfAkn/j19cDy2sjxpcq79t+avYV0vqR8xgONMTUbOdEuTgN4JBgHRObdsoG9K1bo1uZ4CNnh9Vqi4YwP43h+uc5jBisPZUAciR5uCuRtJTWUzq032qybToB/xWTlD1VHflkBoM+RKhtY7HbGS8ocbj2bPpWbxnck/hqkyUpvkFkO0/ngHr6V44pxEK7sm53abiey4jtAQwJcLgS3wLcbSbVGsfRI1srq1I1s54EKzZfJxqRynuMiaKDPGHrUHEcNCQcWS3k4cU1sZF5jMk17f97SzwC8Iz0Kfd7zzw8IiGmvX7sYHJL0mPwqAF+rSlVwZs+Fj/DDcFcBHvpSlBECgYEA+BzskUzgQng3KoMY3Ho6bCxm2At00+Jmfz/bLknj/KB8qie0ionuvOkSFSc1rqPtijVN4L5EJRS1y8LWMGxJTsxAnqJdlLynFx7zi3n/C3Aywwhtz2ijyQNFLgZtnbhVIZYDnf8GwLsWPJjbEeE1JoDAO4wNtT4DuV5mHAdWqY8CgYEAvH9GuQgc+7un90kAFY9nyvlkD/D2cRcfF0Z+FPoj3/k3GS6pXWfnrORgkiGwyqm1e83Tx8RPYMtkRgAGy2mW4ibj7jbOFJYQNBZiML7DAtYuwDDdILm1d4F3840/QHYdtXdWHIryKmdthjmA/Bt1u0MSMmhaGVfHFdYUF2mNjVcCgYEA6mDAZN3fN0tCqakf0h6wk8E6AbqIySOkuW5ECa0JbnrYaRCK7xgva0sspsjcYDZAzX9fKv/xdanjtjE+jo2sjoBKRtCQYFH58dBuNoKvGEoL2ctbmEN7/QZW0oyF/ijEWq7Qie8AnQ3eiq3GvFQnFlEnxtidlmmXsQNop++SwScCgYBQCHJMyccUkx7D/fjNLrBRHAaCRjs81SZcSY/q9DIbPMNKK+e5Qw6499aQ9UENK3Vk9YWAAjf5zyHqHsTDxTdNGloYoKhrUTPcCczzCWvfXnVHIPgilvcXoJ7/h+9dPUlr7Rlg0RX1LyjvnqbHZBlewyGMyYXH0N80xEqPjj+NzQKBgQCyhmWVWUiZNNMV5aWcHy8XFoXuvLIFWbA2WvPoPG+Xia+5BO2ytTI0VJKBY5ACaEofWsy5R2/L6cJhYeSTGe9z0K6Wg96NsxV4BSawp+jseV7oi1HdpTMB4dGph0DJUFMJZ1Lm0s7r5aZ5pkZ6+JYGry1EGmjmR+xaVHodd2LNpw=="
	kServerPort = "9989"
	// TODO 设置回调地址域名
	kServerDomain = ""
)

func main() {
	var err error

	if client, err = alipay.New(kAppId, kPrivateKey, false); err != nil {
		log.Println("初始化支付宝失败", err)
		return
	}

	// 加载证书
	if err = client.LoadAppCertPublicKeyFromFile("appPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	if err = client.LoadAliPayRootCertFromFile("alipayRootCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	if err = client.LoadAlipayCertPublicKeyFromFile("alipayPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}

	if err = client.SetEncryptKey("iotxR/d99T9Awom/UaSqiQ=="); err != nil {
		log.Println("加载内容加密密钥发生错误", err)
		return
	}

	http.HandleFunc("/alipay/pay", pay)
	http.HandleFunc("/alipay/callback", callback)
	http.HandleFunc("/alipay/notify", notify)

	http.ListenAndServe(":"+kServerPort, nil)
}

func pay(writer http.ResponseWriter, request *http.Request) {
	var tradeNo = fmt.Sprintf("%d", xid.Next())

	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/alipay/notify"
	p.ReturnURL = kServerDomain + "/alipay/callback"
	p.Subject = "支付测试:" + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, _ := client.TradePagePay(p)
	http.Redirect(writer, request, url.String(), http.StatusTemporaryRedirect)
}

func callback(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	if err := client.VerifySign(request.Form); err != nil {
		log.Println("回调验证签名发生错误", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("回调验证签名发生错误"))
		return
	}

	log.Println("回调验证签名通过")

	// 示例一：使用已有接口进行查询
	var outTradeNo = request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo

	rsp, err := client.TradeQuery(p)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())))
		return
	}

	if rsp.IsFailure() {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf("订单 %s 支付成功", outTradeNo)))
}

func notify(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	var notification, err = client.DecodeNotification(request.Form)
	if err != nil {
		log.Println("解析异步通知发生错误", err)
		return
	}

	log.Println("解析异步通知成功:", notification.NotifyId)

	// 示例一：使用自定义请求进行查询
	var p = alipay.NewPayload("alipay.trade.query")
	p.AddBizField("out_trade_no", notification.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err = client.Request(p, &rsp); err != nil {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s \n", notification.OutTradeNo, err.Error())
		return
	}
	if rsp.IsFailure() {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s-%s \n", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)
		return
	}

	log.Printf("订单 %s 支付成功 \n", notification.OutTradeNo)

	client.ACKNotification(writer)
}
