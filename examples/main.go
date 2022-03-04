package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
	"log"
	"net/http"
)

var aliClient *alipay.Client

const (
	kAppId        = "2016073100129537"
	kPrivateKey   = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC4UOTKtDstRrjyNPvek9eGqv1RYDmHtLw7olq4lBtqF/D+BfrPd04rMYoivqO5r+v1DQNUGs2yC48eQ4eWb0TSjl/kUy2jHzDcUGiZrGhgxw3e/TJ9w7ix6TrixVg5n542oC1zWstOl75gkiL7EkDfEKo/fZUbLt8aHgBW87NYE19obfAOxGn+YsH5Hpkl3GGby37Hq/mcX4tcWDYN55JPaAcSyjyeyl/uMA19StmQEuU992jqqYWd5Y3z58DotiD4dVAtjB5VogmCkLUcalvUB3N/z3y67GQT+gtOi42VfxnW/JoaXKurHK9Ukmw1GPu97iZAF3Sy19+mOdeR8I+RAgMBAAECggEAAXlEGwrN3lLOb8FUsjbkZkM/u0LVsuwTBTcLGqa0gWinmKBbnQULLvU6cYSssnNho5fzCt0b/+xvvII1t1I0bqqMwbqYhtFdBqXt8CycuQleZwYHPVIvS9zdh6qkRfGsxisJsf5r2bkE4KjKds9yjVYIxnEunAUH66GJxygzquSZQPxgYB3ASRkTzRRowe3ROqaGLF2ejEcvcASKAiIEaQ6Az0sMDtIUAcqFcN1mp9TQ+UyOgNsTw8qWHKv6z6XwjSfYcVfNhR5bcsf+3zr8CKK7cD6f6cjtIrrDSeDdbjdqXTMEYEOV4qs2PH6mNJwx2V1mAZKNZ3bIRGSiKme+sQKBgQD5UDqLBySXtqsCpMC31SvuHGOlkwEQAKslytJvE1kerh4VV9DCIyRsV9v57GZSc8HGo/AQq7dJAa3X9mumZ4vL8RIrtMkuU+scwT+AzPx8jeibNKLq1GI4GaPmdHJaJsBGtA2VUphWCw1HvHegJGgPUoymfpOc8iKwMdQZO3ZHxwKBgQC9QmWJ6zoBvkIOZdoNHpOMrHkIphzJagC/2dbdH6x58fEX787Nmd8yL3mU189EENmFJAdw9d5kPzi8Nxa84oIsbsQSTaYuF0VN6Kw+dZstoJ3U4pf5ReKjWiRNk6waeg6pzEYwH7mDbLcGBw+0+gW08KJLfsyl6aejJ75i9cnd5wKBgHa7UZYiabfi46BXq/wghlJYrNAOqWPgnaFa7Uq+0SN+Uo9hieba6565XOayQaykujUKn+qgjKI1LYB7N5tBFt+iSEAOUf1BM+g21DJX7Sq4Pn2j3K6vRLNo6ph2/nqWl91UJF/nvOrFSqbOR745eGFLs/Yas9v7qK92m4cEvXjDAoGAG2cOrp55YqE6jT0gCkBAGuEqER+EEYGgpCaVXqTkYy+tucqGBezejTSkhPGOWAucgxOJZEilL5ybyVyslSKyuF49U20cv5Ws+i/TKKP8mOmlkJpSaMw+mWpG0VitVZQQpXMnQnaFdMr74QqKsqh0xRMGXKn6VZd0J0Js5YUy+kcCgYEA5dxAzYHV2yH2/b/Uau99VUtp+xE3BssiEO5CiVLEvzLLQlqcqtFO3ptTMqfTeqe98iiYcr5EMSqLHte1qQeGziyKzPvMHkjrTPMDccVur0L2fcE+WBy2fNdBDzQQ5k3ra27/i5NcRzNAhoBXBNVRwnrLXpaUEjy7ERCoR6y1XkE="
	kServerPort   = "9989"
	kServerDomain = "http://127.0.0.1" + ":" + kServerPort
)

func main() {
	var err error

	if aliClient, err = alipay.New(kAppId, kPrivateKey, false); err != nil {
		log.Println("初始化支付宝失败", err)
		return
	}

	// 使用支付宝证书
	if err = aliClient.LoadAppPublicCertFromFile("appCertPublicKey_2016073100129537.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}

	if err = aliClient.LoadAliPayRootCertFromFile("alipayRootCert.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	if err = aliClient.LoadAliPayPublicCertFromFile("alipayCertPublicKey_RSA2.crt"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}

	var s = gin.Default()
	s.GET("/alipay", pay)
	s.GET("/callback", callback)
	s.POST("/notify", notify)
	s.Run(":" + kServerPort)
}

func pay(c *gin.Context) {
	var tradeNo = fmt.Sprintf("%d", xid.Next())

	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/notify"
	p.ReturnURL = kServerDomain + "/callback"
	p.Subject = "支付测试:" + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, _ := aliClient.TradePagePay(p)

	c.Redirect(http.StatusTemporaryRedirect, url.String())
}

func callback(c *gin.Context) {
	c.Request.ParseForm()

	//ok, err := aliClient.VerifySign(c.Request.Form)
	//if err != nil {
	//	log.Println("回调验证签名发生错误", err)
	//	return
	//}
	//
	//if ok == false {
	//	log.Println("回调验证签名未通过")
	//	return
	//}

	var outTradeNo = c.Request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	rsp, err := aliClient.TradeQuery(p)
	if err != nil {
		c.String(http.StatusBadRequest, "验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())
		return
	}
	if rsp.IsSuccess() == false {
		c.String(http.StatusBadRequest, "验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Content.Msg, rsp.Content.SubMsg)
		return
	}

	c.String(http.StatusOK, "订单 %s 支付成功", outTradeNo)
}

func notify(c *gin.Context) {
	c.Request.ParseForm()

	ok, err := aliClient.VerifySign(c.Request.Form)
	if err != nil {
		log.Println("异步通知验证签名发生错误", err)
		return
	}

	if ok == false {
		log.Println("异步通知验证签名未通过")
		return
	}

	log.Println("异步通知验证签名通过")

	var outTradeNo = c.Request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo
	rsp, err := aliClient.TradeQuery(p)
	if err != nil {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s \n", outTradeNo, err.Error())
		return
	}
	if rsp.IsSuccess() == false {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s-%s \n", outTradeNo, rsp.Content.Msg, rsp.Content.SubMsg)
		return
	}

	log.Printf("订单 %s 支付成功 \n", outTradeNo)
}
