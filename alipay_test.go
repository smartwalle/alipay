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

	privateKey = "MIIEpQIBAAKCAQEA00M40KY4rqqxHHvk1InWLR8uTUtcsxfgoCPkjVEf9qgnrJv44asPn6xU9ry/FFdIlNRbymZS3aRhZs+PwmewjZQlRE23PigPwR5T3tffK82QDzfbunbyXxqEOmia3j16hmAt1J+EKJm2SwjkTfDApMj7oe9Wh3BU6mnul86RJjD5JifnIqbpq3qbJInCdkenB/m/kDizSD/qCBSw16BMSFktjVNVRr+76Sqer/UOCusqCc46yTynGD2Qc077FQ3gfG7WVpSOwETmdZpYzlt8Y3eMeLIFW5YLk8Dk1+gbELAzSknYA7Wz28HCtOub+R04Nl+MdeckpK3IBnNAlb45GQIDAQABAoIBAQC9024OlPzjfT5NMMnJa9zFiHnrO+cciTztx7KKhDVrRWb4wuEbrMAKIifp2Gj9FvyBtlqP/+c/fn+CiMhMzyyl2lKuEAKx1/9n8B8+YcwGqNtjwTYvUsevSr07Wliljqo0aeFkZryyWoOg4ml52vTOXEU2GT8vzXCPfQXE4/gqTj3vz/lxGrkSbKVn4f9vWE9GZ7kmZXp4Hhlao02z7y7M9Y5NV1wTzp5CmJqI7WbgUdaiVy8a/qaa3IvVcsZjWN/tATS3TredUEdi/HNAZ4VWFRomcfdCOohExEl5+BXWQKO58o3fFixQiXKqgfV7pVW5iSiTiu9PVf/qbqjfn4FxAoGBAPlVgd+u0Ewxr0p9KQim4PWcsj7wSUe2CZ6jg0X8EXUq4ao3TlZeDWTkq2s+rMY5+cw7w1LseuZoBV9HMVl/A20JP4AohO3JeEmhifYT9+Xi3oUZ5CzAe/t9J/gHQU/dIc0rzizGW4ATcxD0mB0B/xXXXToi3Q7i4Prph4iP/baVAoGBANjpJWHm+v87qecB71RIF5384jz3kMjB6aLnwztT1eF90+e/CxSEyt38z7nAFPqR7yZFhCIBvqwUc03DvRHJB35d1svUNGK2/WvQzv8qqMWoyWeWQgMI5frb61g6LPlrpt/VHd4oKu/EBaDTz/sv4Z48N9LMH25QnLOVbPnes+t1AoGBALwinCPXOY80skrnlA3WNcq+mPTKxNCaeDm+sbAeGmIpoCubHb4nq72kmgALQ70zQ9yqf6DTlYzDksIo6wDXyRL+Nm8So/L04ZHVlU8cFCLfg/404ioud570+jK57rs/ZWD2G9VHOwWyJ1H07c59kie94LecVOcFpLNPgRg3Zj7dAoGBALhwVrcOnrJUWz3waT6zJlHSa9UHdPcf27gHdfyGZsEcKwlNavCBkbZ8d20spSoC2OUkCxyXezf7E6g2Hhh3ZYXP6QrcX8bobkg0/y39ahDYWplYLL1D3boXMeCNZxyWqwh8wK7cXNYPSfHc6WZe6muQLK9zne5BYV1aW5WEBlzZAoGAT1HOGfkZoM3W6toiZF8yBtckl3Tm3rv45tV0Y/s9ymEwC/jaypyHuvIefnVAKtJshM7DkzKfZUs6Sm+naKJJuhgkkajQbuFngWBPUnA3obz3jWzWJPLm3jzXMF2xKvA0y8u3u+VFD8dUclviwT0QYfX7EMHY1uBz/ybWrSDB5f4="

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
