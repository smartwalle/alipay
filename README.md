AliPay SDK for Golang

## 已实现接口

#### 手机网站支付API

* **手机网站支付接口**
	
	alipay.trade.wap.pay

* **统一收单线下交易查询**
	
	alipay.trade.query
	
* **统一收单交易支付接口**
	
	alipay.trade.pay
	
* **统一收单交易关闭接口**

	alipay.trade.close

* **统一收单交易退款接口**

	alipay.trade.refund

* **统一收单交易退款查询**

	alipay.trade.fastpay.refund.query
	
#### 通知
	
* **通知内容转换及签名验证**
	
	将支付宝的通知内容转换为 Golang 的结构体，并且验证其合法性。
	
## 集成流程

从[支付宝开放平台](https://open.alipay.com/)申请创建相关的应用，使用自己的支付宝账号登录即可。

#### 沙箱环境

支付宝开放平台为每一个应用提供了沙箱环境，供开发人员开发测试使用。

沙箱环境是独立的，每一个应用都会有一个商家账号和买家账号。

#### 应用信息配置

参考[官网文档](https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.5pgfxp&treeId=200&articleId=105894&docType=1) 进行应用的配置。

本 SDK 中的签名方法为 RSA2，所以请注意配置 **RSA2(SHA256)密钥**。

请参考 [如何生成 RSA 密钥](https://doc.open.alipay.com/docs/doc.htm?treeId=291&articleId=105971&docType=1)。

#### 创建 Wap 支付

``` Golang
var client = alipay.New(appId, partnerId, publickKey, privateKey, false)

var p = AliPayTradeWapPay{}
p.NotifyURL = "xxx"
p.Subject = "标题"
p.OutTradeNo = "传递一个唯一单号"
p.TotalAmount = "10.00"
p.ProductCode = "商品编码"

var html, url, _ = client.TradeWapPay(p)
// 直接访问该 URL 就可以了
```

#### 验证支付结果

有支付或者其它动作发生后，支付宝服务器会调用我们提供的 Notify URL，并向其传递会相关的信息。参考[手机网站支付结果异步通知](https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.XM5C4a&treeId=203&articleId=105286&docType=1)。

我们需要在提供的 Notify URL 服务中获取相关的参数并进行验证:

```Golang

var client = alipay.New(appId, partnerId, publickKey, privateKey, false)
client.AliPayPublicKey = xxx // 从支付宝管理后台获取支付宝提供的公钥
 
http.HandleFunc("/alipay", func(rep http.ResponseWriter, req *http.Request) {
	var noti, _ = client.GetTradeNotification(req)
	if noti != nil {
		fmt.Println("支付成功")
	} else {
		fmt.Println("支付失败")
	}
})
```

此验证方法适用于支付宝所有情况下发送的 Notify，不管是手机 App 支付还是 Wap 支付。

#### 鸣谢
感谢 [@wusphinx](https://github.com/wusphinx) 对本项目的支持。
