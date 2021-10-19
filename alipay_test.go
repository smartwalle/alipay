package alipay_test

import (
	"fmt"
	"os"

	"github.com/NeoclubTechnology/alipay/v3"
)

var (
	appID = "2016073100129537"

	privateKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC4UOTKtDstRrjyNPvek9eGqv1RYDmHtLw7olq4lBtqF/D+BfrPd04rMYoivqO5r+v1DQNUGs2yC48eQ4eWb0TSjl/kUy2jHzDcUGiZrGhgxw3e/TJ9w7ix6TrixVg5n542oC1zWstOl75gkiL7EkDfEKo/fZUbLt8aHgBW87NYE19obfAOxGn+YsH5Hpkl3GGby37Hq/mcX4tcWDYN55JPaAcSyjyeyl/uMA19StmQEuU992jqqYWd5Y3z58DotiD4dVAtjB5VogmCkLUcalvUB3N/z3y67GQT+gtOi42VfxnW/JoaXKurHK9Ukmw1GPu97iZAF3Sy19+mOdeR8I+RAgMBAAECggEAAXlEGwrN3lLOb8FUsjbkZkM/u0LVsuwTBTcLGqa0gWinmKBbnQULLvU6cYSssnNho5fzCt0b/+xvvII1t1I0bqqMwbqYhtFdBqXt8CycuQleZwYHPVIvS9zdh6qkRfGsxisJsf5r2bkE4KjKds9yjVYIxnEunAUH66GJxygzquSZQPxgYB3ASRkTzRRowe3ROqaGLF2ejEcvcASKAiIEaQ6Az0sMDtIUAcqFcN1mp9TQ+UyOgNsTw8qWHKv6z6XwjSfYcVfNhR5bcsf+3zr8CKK7cD6f6cjtIrrDSeDdbjdqXTMEYEOV4qs2PH6mNJwx2V1mAZKNZ3bIRGSiKme+sQKBgQD5UDqLBySXtqsCpMC31SvuHGOlkwEQAKslytJvE1kerh4VV9DCIyRsV9v57GZSc8HGo/AQq7dJAa3X9mumZ4vL8RIrtMkuU+scwT+AzPx8jeibNKLq1GI4GaPmdHJaJsBGtA2VUphWCw1HvHegJGgPUoymfpOc8iKwMdQZO3ZHxwKBgQC9QmWJ6zoBvkIOZdoNHpOMrHkIphzJagC/2dbdH6x58fEX787Nmd8yL3mU189EENmFJAdw9d5kPzi8Nxa84oIsbsQSTaYuF0VN6Kw+dZstoJ3U4pf5ReKjWiRNk6waeg6pzEYwH7mDbLcGBw+0+gW08KJLfsyl6aejJ75i9cnd5wKBgHa7UZYiabfi46BXq/wghlJYrNAOqWPgnaFa7Uq+0SN+Uo9hieba6565XOayQaykujUKn+qgjKI1LYB7N5tBFt+iSEAOUf1BM+g21DJX7Sq4Pn2j3K6vRLNo6ph2/nqWl91UJF/nvOrFSqbOR745eGFLs/Yas9v7qK92m4cEvXjDAoGAG2cOrp55YqE6jT0gCkBAGuEqER+EEYGgpCaVXqTkYy+tucqGBezejTSkhPGOWAucgxOJZEilL5ybyVyslSKyuF49U20cv5Ws+i/TKKP8mOmlkJpSaMw+mWpG0VitVZQQpXMnQnaFdMr74QqKsqh0xRMGXKn6VZd0J0Js5YUy+kcCgYEA5dxAzYHV2yH2/b/Uau99VUtp+xE3BssiEO5CiVLEvzLLQlqcqtFO3ptTMqfTeqe98iiYcr5EMSqLHte1qQeGziyKzPvMHkjrTPMDccVur0L2fcE+WBy2fNdBDzQQ5k3ra27/i5NcRzNAhoBXBNVRwnrLXpaUEjy7ERCoR6y1XkE="

	aliPublicKey = ""
)

var client *alipay.Client

func init() {
	var err error
	client, err = alipay.New(appID, privateKey, false)

	if err != nil {
		fmt.Println("初始化支付宝失败, 错误信息为", err)
		os.Exit(-1)
	}

	// 下面两种方式只能二选一
	var cert = true
	if cert {
		// 使用支付宝证书
		fmt.Println("加载证书", client.LoadAppPublicCertFromFile("appCertPublicKey_2016073100129537.crt"))
		fmt.Println("加载证书", client.LoadAliPayRootCertFromFile("alipayRootCert.crt"))
		fmt.Println("加载证书", client.LoadAliPayPublicCertFromFile("alipayCertPublicKey_RSA2.crt"))
	} else {
		// 使用支付宝公钥
		fmt.Println("加载公钥", client.LoadAliPayPublicKey(aliPublicKey))
	}
}
