package alipay_test

import (
	"fmt"
	"os"

	alipay "github.com/smartwalle/alipay/v3"
)

var (
	appID = "2016073100129537"

	privateKey = "MIIEogIBAAKCAQEAmF/L4+mkeiFTaCeSy6DoOB3OGOKFmLElx/5pgOuPPGHLXjBsk6X+8Iq0VqRhug3hFndSPB4e8UxKftXmzjYswlhcEBG/huBlTQhiKbNAq7Iia1L3tIs+vV8lEuFJlI9lZwOfczHIDo3gYZFKnEumyLntWq5sebYkhazg9NG56D9cBeXPz7rPzKwzVmjCL8HkB4BkvEZYUqw0WFm6GFt0Pc0VOcCwoio9oRlZOgVI/kHCGbLfChgJAqs+Id1pYOvB0CoYv49uaXG4qqcUCG5tW9EE/IVUCXRjV7+v9ypd+PejNyFWtC+5S6I8xVa3LjmMpFr881n+QtKeYQq0hFijBQIDAQABAoIBABKqSXOVv0wmoOz1TAodn9Sf8gsiVHMr4BDrnUjpkhY3dI4JKIO9pckZdJXYdRAxew0heLVcizXLvqRi128TO9BiuoRNaETBYCdbi4rIJnfhzk2PUECRfhH8gbIaXsUP+7/uta2Kv5Lo1j+daKJUsg2MmQKusyMFqNunHbdfqYJFb5FVJieojngXhBeE/pOkCDKfVbbcYNg8MZSHqGBh6U15/H2qCewdfstnVJKMr8gs4xK7ITygev3jdE9PuxDpsiBexXNT4AXuhHpydYUKgPuAoFj2fnMIvdDY+bo5dpgN4sCFVjnPEp1HBkfviY1X2k3Ca/e0bGjfS1qxs2fDy1ECgYEA6bOPRx3AIO73aesBmVwRCxaZ2m1o9HsMmT0WeW/yaQ/NeEUFgn36BrKwkr8TQZwPrMUEMshFIowgSjOWSJbV243/xEi25i4IE76NeTPK6OJv1sOocc4dogpGmXq2U+lepPQTX8ZNXh6+MeT96IT//kIrc0u1tOe8v+Cfn6nUH48CgYEApum4o5VSHfHL3j3U6o9+eAsQXfoaPhEd9B2lkUU6btGLegR3lwycCDfP94dHKQe9JAL6UCvUcWvAEpU+gZT/vFrh0krs1XjP/i3S8PH5rd11NhRrYTn2K4ELDfcvhFYelHF7MLz9Yk6j5anze3fRGeAve45fqLl8CEaaVj3oSisCgYAeMHXnx+4T0wrfAd65Au2oswi48L1IJ8Ue3odStKVp8QKn8LKfgsqTpu2sZ0aDiTd1KBY8wSY9KkDZlQRq6CFENXm+z23hGj0s38bCy3AA2Y98/NV7rhah4hXwqat394OkZ2tBSgqgh/Ql2eD68oNnQwD96d/VOMJnPwsfwv6F/QKBgDcBHHyj74y4qwNRAwJNSVML6lfd3JoJkAJrZq1pz+jHGxyZrkNTv3Oh2OHsbZHi3/ynEpAq8XZzGLfHAPM5A9GxbWucj1GF350WwsXuJ+aY7VBmCEDhFfOeMeNnSvxkWO14PC2CiknEOpSrnfZZNMo/K8ae031JqssbYS78dblDAoGATqM5N4n8cS7NPzGLoj8+sZshJmtWZsefxV+QUBYKVeeDyjiPr/E1CA8zAewapnh5eeUF68A9uAqwXyuU8Cs8i6WaeBJf32llOgAgKfjaYtQJFfoc0xeGIPDltFQSL630pVwdIQTbYtj4WTm6UEPqGr/kBJ4/O7F4fWJOAyr0pT0="

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
		fmt.Println("加载证书", client.LoadAppPublicCertFromFile("appPublicCert.crt"))
		fmt.Println("加载证书", client.LoadAliPayRootCertFromFile("alipayRootCert.crt"))
		fmt.Println("加载证书", client.LoadAliPayPublicCertFromFile("alipayPublicCert.crt"))
	} else {
		// 使用支付宝公钥
		fmt.Println("加载公钥", client.LoadAliPayPublicKey(aliPublicKey))
	}
}
