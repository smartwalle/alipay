package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
	"log"
	"net/http"
)

var client *alipay.Client

const (
	kAppId      = "2016073100129537"
	kPrivateKey = "MIIEogIBAAKCAQEAmF/L4+mkeiFTaCeSy6DoOB3OGOKFmLElx/5pgOuPPGHLXjBsk6X+8Iq0VqRhug3hFndSPB4e8UxKftXmzjYswlhcEBG/huBlTQhiKbNAq7Iia1L3tIs+vV8lEuFJlI9lZwOfczHIDo3gYZFKnEumyLntWq5sebYkhazg9NG56D9cBeXPz7rPzKwzVmjCL8HkB4BkvEZYUqw0WFm6GFt0Pc0VOcCwoio9oRlZOgVI/kHCGbLfChgJAqs+Id1pYOvB0CoYv49uaXG4qqcUCG5tW9EE/IVUCXRjV7+v9ypd+PejNyFWtC+5S6I8xVa3LjmMpFr881n+QtKeYQq0hFijBQIDAQABAoIBABKqSXOVv0wmoOz1TAodn9Sf8gsiVHMr4BDrnUjpkhY3dI4JKIO9pckZdJXYdRAxew0heLVcizXLvqRi128TO9BiuoRNaETBYCdbi4rIJnfhzk2PUECRfhH8gbIaXsUP+7/uta2Kv5Lo1j+daKJUsg2MmQKusyMFqNunHbdfqYJFb5FVJieojngXhBeE/pOkCDKfVbbcYNg8MZSHqGBh6U15/H2qCewdfstnVJKMr8gs4xK7ITygev3jdE9PuxDpsiBexXNT4AXuhHpydYUKgPuAoFj2fnMIvdDY+bo5dpgN4sCFVjnPEp1HBkfviY1X2k3Ca/e0bGjfS1qxs2fDy1ECgYEA6bOPRx3AIO73aesBmVwRCxaZ2m1o9HsMmT0WeW/yaQ/NeEUFgn36BrKwkr8TQZwPrMUEMshFIowgSjOWSJbV243/xEi25i4IE76NeTPK6OJv1sOocc4dogpGmXq2U+lepPQTX8ZNXh6+MeT96IT//kIrc0u1tOe8v+Cfn6nUH48CgYEApum4o5VSHfHL3j3U6o9+eAsQXfoaPhEd9B2lkUU6btGLegR3lwycCDfP94dHKQe9JAL6UCvUcWvAEpU+gZT/vFrh0krs1XjP/i3S8PH5rd11NhRrYTn2K4ELDfcvhFYelHF7MLz9Yk6j5anze3fRGeAve45fqLl8CEaaVj3oSisCgYAeMHXnx+4T0wrfAd65Au2oswi48L1IJ8Ue3odStKVp8QKn8LKfgsqTpu2sZ0aDiTd1KBY8wSY9KkDZlQRq6CFENXm+z23hGj0s38bCy3AA2Y98/NV7rhah4hXwqat394OkZ2tBSgqgh/Ql2eD68oNnQwD96d/VOMJnPwsfwv6F/QKBgDcBHHyj74y4qwNRAwJNSVML6lfd3JoJkAJrZq1pz+jHGxyZrkNTv3Oh2OHsbZHi3/ynEpAq8XZzGLfHAPM5A9GxbWucj1GF350WwsXuJ+aY7VBmCEDhFfOeMeNnSvxkWO14PC2CiknEOpSrnfZZNMo/K8ae031JqssbYS78dblDAoGATqM5N4n8cS7NPzGLoj8+sZshJmtWZsefxV+QUBYKVeeDyjiPr/E1CA8zAewapnh5eeUF68A9uAqwXyuU8Cs8i6WaeBJf32llOgAgKfjaYtQJFfoc0xeGIPDltFQSL630pVwdIQTbYtj4WTm6UEPqGr/kBJ4/O7F4fWJOAyr0pT0="
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
	if err = client.LoadAppPublicCertFromFile("appPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	if err = client.LoadAliPayRootCertFromFile("alipayRootCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	if err = client.LoadAliPayPublicCertFromFile("alipayPublicCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}

	if err = client.SetEncryptKey("FtVd5SgrsUzYQRAPBmejHQ=="); err != nil {
		log.Println("加载内容加密密钥发生错误", err)
		return
	}

	var s = gin.Default()
	s.GET("/alipay/pay", pay)
	s.GET("/alipay/callback", callback)
	s.POST("/alipay/notify", notify)
	s.Run(":" + kServerPort)
}

func pay(c *gin.Context) {
	var tradeNo = fmt.Sprintf("%d", xid.Next())

	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/alipay/notify"
	p.ReturnURL = kServerDomain + "/alipay/callback"
	p.Subject = "支付测试:" + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, _ := client.TradePagePay(p)

	c.Redirect(http.StatusTemporaryRedirect, url.String())
}

func callback(c *gin.Context) {
	c.Request.ParseForm()

	err := client.VerifySign(c.Request.Form)
	if err != nil {
		log.Println("回调验证签名发生错误", err)
		c.String(http.StatusBadRequest, "回调验证签名发生错误")
		return
	}

	log.Println("回调验证签名通过")

	var outTradeNo = c.Request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo

	rsp, err := client.TradeQuery(p)
	if err != nil {
		c.String(http.StatusBadRequest, "验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())
		return
	}
	if rsp.Failed() {
		c.String(http.StatusBadRequest, "验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)
		return
	}

	c.String(http.StatusOK, "订单 %s 支付成功", outTradeNo)
}

func notify(c *gin.Context) {
	var noti, err = client.DecodeNotification(c.Request)
	if err != nil {
		log.Println("解析异步通知发生错误", err)
		return
	}

	log.Println("解析异步通知成功:", noti.NotifyId)

	var p = alipay.NewPayload("alipay.trade.query")
	p.AddField("out_trade_no", noti.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err = client.Request(p, &rsp); err != nil {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s \n", noti.OutTradeNo, err.Error())
		return
	}
	if rsp.Failed() {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s-%s \n", noti.OutTradeNo, rsp.Msg, rsp.SubMsg)
		return
	}

	log.Printf("订单 %s 支付成功 \n", noti.OutTradeNo)

	client.ACKNotification(c.Writer)
}
