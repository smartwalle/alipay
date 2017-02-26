package alipay

import (
	"fmt"
	"testing"
)

var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv8dXxi8wNAOqBNOh8Dv5
rh0BTb5KNgk62jDaS536Z1cDqq2JmpBYkBnzJXHAXEgBwPXgX8bGruMMjZKW8P4u
v3Rvj8Am9ewWwUK2U7m2ZB3Oo9MWtyYoiLGX1IA4FFenXzpPgm0WyzaeLX4yJ8j+
hVrRbgwbZzb9Aq0MyepnK5PVoSPLAPXxvWrIBTok1+liughxwD/7R+ldaQQCtWC7
nHBwOOChLkX6jenCOqi6LrTxJ4ycGTWTctngFMJO4YtMmq/2zrw+ovNqmxHJQAZw
uRFnKlZuFoEKPWyMGYtbvK9AWIfC8ubn30O5F9kfLMIHwAHCh0UipPSbKDwQ2BnW
swIDAQAB
-----END PUBLIC KEY-----`)

var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAv8dXxi8wNAOqBNOh8Dv5rh0BTb5KNgk62jDaS536Z1cDqq2J
mpBYkBnzJXHAXEgBwPXgX8bGruMMjZKW8P4uv3Rvj8Am9ewWwUK2U7m2ZB3Oo9MW
tyYoiLGX1IA4FFenXzpPgm0WyzaeLX4yJ8j+hVrRbgwbZzb9Aq0MyepnK5PVoSPL
APXxvWrIBTok1+liughxwD/7R+ldaQQCtWC7nHBwOOChLkX6jenCOqi6LrTxJ4yc
GTWTctngFMJO4YtMmq/2zrw+ovNqmxHJQAZwuRFnKlZuFoEKPWyMGYtbvK9AWIfC
8ubn30O5F9kfLMIHwAHCh0UipPSbKDwQ2BnWswIDAQABAoIBAH7QyfkSsTRkC+Sf
MaGTd1qscXVAVQCAf/tSfLeuIqx9PL57fNfJhdbcYg2rt8EOGKLJtHKBFlcFawKf
IdMAslcGHtOXA+xxDucDP2AEGVkA4OkyJ/46bGlfzn/Fvc+t2s6812Du1DjSyCxb
G711SuFSGdVEikZpdUt0tVU7/LcyKAEZd45Ct+F9MvrPECbSsfODvTOVDHO2k42f
iwSzLPVmM4wVUc2xA15O87jtDhRiAK/RveQ7J2TWcarkyCR8J+bf5GGA79LdE3vR
Kr/HAk7INVX4T6U9QuDF30mqNRsloQbNGdvqO65nafNHvuVzUiqPdSX7XQwg/cOO
mhSsUbkCgYEA8BQXaHn3psHUZx8zEwQFVyd6rzxb+7jmVlUT+jG1pSiZ4WAWxxqx
YVXhn2dbfatDxWoGOMsrDM/Qp8g81nMG01jtmJr2RKFhAbQl93ipGvvaCNoJ8Lx7
HpFSq7dETcCCAE7tYMk0LlcVwxeaIUWakDyBHgEy4Zp6lLwdwsh115UCgYEAzH8/
E5dTOcYdcxk7HLupEC9MCb+FshZT5UIN9I7zLNljQX2O/8m2THb+oZUoy30RVot+
kYjh5H8M5CYiP0Kkm0Ovq5KC0loyt5SfzWbgwHEldQUVp8woE0YdaJzGB/UnmI9m
dJBON1t3qbMWjlguXOD8bfriDRuefaZd9oVSQycCgYBcz+ecxEoxdY2fsDgWid9m
qiSLylHlJr4lcg6fEsieaOvUbUlg/7jDYGgxL8v28Vbp4us02ZZzBYQs2QRsA1wI
KMDx1jaOobTW68YhvcviWqsX8PMW1kbislu7dsY5KMsZQ2oRmLdLku8e1OkJI9d1
G27vIpeBEC+DgJYgz05/YQKBgQCStWNiQbkihKBSF7LR3Uvf4Z6yi6V16xDLM8Vh
Q0DwVxEfRd3WYjcXynLJJ4J54kMTDMaD0GkHDaMI9taw/bWr8jZQZ67VDILAM68l
o/3v8fyGZFxx4kSJ905X48kqolWC3LYLQA/tJQDHTUUMX/T7CynuGQQdlUfyKu3U
Uzd+FwKBgHW9Nur4eTxK1nIOZyGgCqL1duYsJQcPWyIcRMTSjOoQZ5ZUhQZTw1Hd
2CW0Iu2fXExESTIjwXJ0ZJXnCgFU8acQX5vtItC1BlMaucw9XTx1RBCVQdTZ7DSX
vTlWbWwZHVDP85dioLE9mfo5+Hh3SmHDi3TaVXjxeJsUgHkRgOX7
-----END RSA PRIVATE KEY-----
`)

var appID = "2016073100129537"

func TestSign(t *testing.T) {

	//	var publickKey = []byte(`-----BEGIN PUBLIC KEY-----
	//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZ8lmnG4TMFqJty0FoqAHKxEIs
	//IcZaM2E8PXjXoQ+iszREsqi1B5Lmq7GeJ1/9N+OGDIjpDHnEfMMlHrj+5gYSTPab
	//bLGCvtcluPbI6R+uJz3uYGtPzqn4EKiNvC1ixANLmbhdqbb3KAkCcRltZOZYSerG
	//VE069nKjmSWRRlWJkQIDAQAB
	//-----END PUBLIC KEY-----`)

	//	var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
	//MIICXgIBAAKBgQDZ8lmnG4TMFqJty0FoqAHKxEIsIcZaM2E8PXjXoQ+iszREsqi1
	//B5Lmq7GeJ1/9N+OGDIjpDHnEfMMlHrj+5gYSTPabbLGCvtcluPbI6R+uJz3uYGtP
	//zqn4EKiNvC1ixANLmbhdqbb3KAkCcRltZOZYSerGVE069nKjmSWRRlWJkQIDAQAB
	//AoGBAKGAiRbfuYRSsYKSv6GB/fH3hOGXFZj5wfAVzVpcK23xRaYyjfm35w+v4yrD
	//GspVg/BtkXbAm+sSWLlFDuk0IwJGhqxmIJ6TvxRFDxep4Kz5LjPlTvG1/mvSKKW6
	//1uUZquT2Ll0qy2hwyui8K08+CWVxAO5NMskYHaztF8QArjgBAkEA8UtE8pawG/Cj
	//24kj//y10f36Yt9tQ00/Nu7hfLXJe292zWn0cCEZG2ukkt6kQQtNoUBpRMTj9cQR
	//hd+2hPgX8QJBAOc6z0sJUWwG6m13nSVlu6j2wmZTp6W9U52WNR364L1UDfn7MI/X
	//7roW1SdwQwdYNVgwvt2N1MNJheTxt5/ZK6ECQQCs3Ta08J2UNq69LZ+72ejMWz7R
	//LK3TZHjgOv0R4g5JPw6GlNzIo/2ftls92QEllBp2ZnXEDaYewOuo1B+nXTGRAkEA
	//h8qCp+dN+KnLDAQ9thObdCuNmHgyMOQRca8ffH6zcpwlJRP9vcuqd4AnJ2UHCA4m
	//Ladav1OmihToW74T/vyTYQJAAf+PKLRD+O2CuwcZJG0taEqL0RR+kXMEd0wLp4EN
	//pFwKUHMh+rm4/Asgy126+rS6Hr0QuNuoJuQbAr3Q28h7PQ==
	//-----END RSA PRIVATE KEY-----`)

	var client = New("2016073100129537", "2088102169227503", publicKey, privateKey, false)

	//var p = AliPayFastpayTradeRefundQuery{}
	//p.OutTradeNo = "1111111"
	//p.OutRequestNo = "1111111"
	//
	//var r, err = client.TradeFastpayRefundQuery(p)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(r.IsSuccess(), r)

	//var p = AliPayTradeRefund{}
	//p.RefundAmount = "10.00"
	//p.OutTradeNo = "1111111"
	//
	//var r, err = client.TradeRefund(p)
	//fmt.Println(r.IsSuccess(), err, r)

	var p = AliPayTradeWapPay{}
	p.NotifyURL = "http://203.86.24.181:3000/alipay"
	p.ReturnURL = "http://203.86.24.181:3000"
	p.Subject = "aa"
	p.OutTradeNo = "1111111"
	p.TotalAmount = "10.00"
	p.ProductCode = "eeeeee"

	fmt.Println(client.TradeWapPay(p))
}
