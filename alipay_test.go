package alipay_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/smartwalle/alipay/v3"
)

var (
	appID = "9021000122689420"

	privateKey = "MIIEpAIBAAKCAQEAtrCYn039gp/IjLZsNl64QjHjPfAklyYZyIPSqlPNmColRMOAb2rbnLASRpP1VgT7YzA4JgJ1f/fhGsSBexpIi2BOZdwDexBtmfe9dFGzbWhpqAxzWnZazLde+BGOs6BGk0v1B0cWUqlv6wgSSo57Xu7xL56a34gDsBi1qoXnu4f1CzrvviRsNCgDDurNsAtkjLzoPuzHri6sThsQ7P3amb3zyG5xVxSGZRFKgPNoiiZpBpXPoEwMrHQRE8rmsmgdz+E4YL1xuD+ICQxCCTBOJwUuDzPt5wr793Pxgloqh0p3yPvOShmbMAxUtLiGgcZNqxy49ddG89egVyGxRJpimQIDAQABAoIBAEfGUfAkn/j19cDy2sjxpcq79t+avYV0vqR8xgONMTUbOdEuTgN4JBgHRObdsoG9K1bo1uZ4CNnh9Vqi4YwP43h+uc5jBisPZUAciR5uCuRtJTWUzq032qybToB/xWTlD1VHflkBoM+RKhtY7HbGS8ocbj2bPpWbxnck/hqkyUpvkFkO0/ngHr6V44pxEK7sm53abiey4jtAQwJcLgS3wLcbSbVGsfRI1srq1I1s54EKzZfJxqRynuMiaKDPGHrUHEcNCQcWS3k4cU1sZF5jMk17f97SzwC8Iz0Kfd7zzw8IiGmvX7sYHJL0mPwqAF+rSlVwZs+Fj/DDcFcBHvpSlBECgYEA+BzskUzgQng3KoMY3Ho6bCxm2At00+Jmfz/bLknj/KB8qie0ionuvOkSFSc1rqPtijVN4L5EJRS1y8LWMGxJTsxAnqJdlLynFx7zi3n/C3Aywwhtz2ijyQNFLgZtnbhVIZYDnf8GwLsWPJjbEeE1JoDAO4wNtT4DuV5mHAdWqY8CgYEAvH9GuQgc+7un90kAFY9nyvlkD/D2cRcfF0Z+FPoj3/k3GS6pXWfnrORgkiGwyqm1e83Tx8RPYMtkRgAGy2mW4ibj7jbOFJYQNBZiML7DAtYuwDDdILm1d4F3840/QHYdtXdWHIryKmdthjmA/Bt1u0MSMmhaGVfHFdYUF2mNjVcCgYEA6mDAZN3fN0tCqakf0h6wk8E6AbqIySOkuW5ECa0JbnrYaRCK7xgva0sspsjcYDZAzX9fKv/xdanjtjE+jo2sjoBKRtCQYFH58dBuNoKvGEoL2ctbmEN7/QZW0oyF/ijEWq7Qie8AnQ3eiq3GvFQnFlEnxtidlmmXsQNop++SwScCgYBQCHJMyccUkx7D/fjNLrBRHAaCRjs81SZcSY/q9DIbPMNKK+e5Qw6499aQ9UENK3Vk9YWAAjf5zyHqHsTDxTdNGloYoKhrUTPcCczzCWvfXnVHIPgilvcXoJ7/h+9dPUlr7Rlg0RX1LyjvnqbHZBlewyGMyYXH0N80xEqPjj+NzQKBgQCyhmWVWUiZNNMV5aWcHy8XFoXuvLIFWbA2WvPoPG+Xia+5BO2ytTI0VJKBY5ACaEofWsy5R2/L6cJhYeSTGe9z0K6Wg96NsxV4BSawp+jseV7oi1HdpTMB4dGph0DJUFMJZ1Lm0s7r5aZ5pkZ6+JYGry1EGmjmR+xaVHodd2LNpw=="

	aliPublicKey = ""
)

var client *alipay.Client

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	var err error
	client, err = alipay.New(appID, privateKey, false)

	client.OnReceivedData(func(_ context.Context, method string, data []byte) {
		log.Println(method, string(data))
	})

	if err != nil {
		fmt.Println("初始化支付宝失败, 错误信息为", err)
		os.Exit(-1)
	}

	// 加载内容密钥（可选），详情查看 https://opendocs.alipay.com/common/02mse3
	fmt.Println("加载内容密钥", client.SetEncryptKey("iotxR/d99T9Awom/UaSqiQ=="))

	// 下面两种方式只能二选一
	var cert = true
	if cert {
		// 使用支付宝证书
		fmt.Println("加载证书", client.LoadAppCertPublicKeyFromFile("appPublicCert.crt"))
		fmt.Println("加载证书", client.LoadAliPayRootCertFromFile("alipayRootCert.crt"))
		fmt.Println("加载证书", client.LoadAlipayCertPublicKeyFromFile("alipayPublicCert.crt"))
	} else {
		// 使用支付宝公钥
		fmt.Println("加载公钥", client.LoadAliPayPublicKey(aliPublicKey))
	}
}

func TestClient_CertDownload(t *testing.T) {
	t.Log("========== CertDownload ==========")
	var p = alipay.CertDownload{}
	p.AliPayCertSN = ""
	rsp, err := client.CertDownload(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}
