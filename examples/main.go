package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
)

var client *alipay.Client

const (
	kAppId      = "9021000122689420"
	kPrivateKey = "MIIEpQIBAAKCAQEA00M40KY4rqqxHHvk1InWLR8uTUtcsxfgoCPkjVEf9qgnrJv44asPn6xU9ry/FFdIlNRbymZS3aRhZs+PwmewjZQlRE23PigPwR5T3tffK82QDzfbunbyXxqEOmia3j16hmAt1J+EKJm2SwjkTfDApMj7oe9Wh3BU6mnul86RJjD5JifnIqbpq3qbJInCdkenB/m/kDizSD/qCBSw16BMSFktjVNVRr+76Sqer/UOCusqCc46yTynGD2Qc077FQ3gfG7WVpSOwETmdZpYzlt8Y3eMeLIFW5YLk8Dk1+gbELAzSknYA7Wz28HCtOub+R04Nl+MdeckpK3IBnNAlb45GQIDAQABAoIBAQC9024OlPzjfT5NMMnJa9zFiHnrO+cciTztx7KKhDVrRWb4wuEbrMAKIifp2Gj9FvyBtlqP/+c/fn+CiMhMzyyl2lKuEAKx1/9n8B8+YcwGqNtjwTYvUsevSr07Wliljqo0aeFkZryyWoOg4ml52vTOXEU2GT8vzXCPfQXE4/gqTj3vz/lxGrkSbKVn4f9vWE9GZ7kmZXp4Hhlao02z7y7M9Y5NV1wTzp5CmJqI7WbgUdaiVy8a/qaa3IvVcsZjWN/tATS3TredUEdi/HNAZ4VWFRomcfdCOohExEl5+BXWQKO58o3fFixQiXKqgfV7pVW5iSiTiu9PVf/qbqjfn4FxAoGBAPlVgd+u0Ewxr0p9KQim4PWcsj7wSUe2CZ6jg0X8EXUq4ao3TlZeDWTkq2s+rMY5+cw7w1LseuZoBV9HMVl/A20JP4AohO3JeEmhifYT9+Xi3oUZ5CzAe/t9J/gHQU/dIc0rzizGW4ATcxD0mB0B/xXXXToi3Q7i4Prph4iP/baVAoGBANjpJWHm+v87qecB71RIF5384jz3kMjB6aLnwztT1eF90+e/CxSEyt38z7nAFPqR7yZFhCIBvqwUc03DvRHJB35d1svUNGK2/WvQzv8qqMWoyWeWQgMI5frb61g6LPlrpt/VHd4oKu/EBaDTz/sv4Z48N9LMH25QnLOVbPnes+t1AoGBALwinCPXOY80skrnlA3WNcq+mPTKxNCaeDm+sbAeGmIpoCubHb4nq72kmgALQ70zQ9yqf6DTlYzDksIo6wDXyRL+Nm8So/L04ZHVlU8cFCLfg/404ioud570+jK57rs/ZWD2G9VHOwWyJ1H07c59kie94LecVOcFpLNPgRg3Zj7dAoGBALhwVrcOnrJUWz3waT6zJlHSa9UHdPcf27gHdfyGZsEcKwlNavCBkbZ8d20spSoC2OUkCxyXezf7E6g2Hhh3ZYXP6QrcX8bobkg0/y39ahDYWplYLL1D3boXMeCNZxyWqwh8wK7cXNYPSfHc6WZe6muQLK9zne5BYV1aW5WEBlzZAoGAT1HOGfkZoM3W6toiZF8yBtckl3Tm3rv45tV0Y/s9ymEwC/jaypyHuvIefnVAKtJshM7DkzKfZUs6Sm+naKJJuhgkkajQbuFngWBPUnA3obz3jWzWJPLm3jzXMF2xKvA0y8u3u+VFD8dUclviwT0QYfX7EMHY1uBz/ybWrSDB5f4="
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

	rsp, err := client.TradeQuery(context.Background(), p)
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
	if err = client.Request(context.Background(), p, &rsp); err != nil {
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
