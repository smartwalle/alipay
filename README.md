AliPay SDK for Golang

## 已实现接口

#### 手机网站支付API

* **手机网站支付接口**
	
	alipay.trade.wap.pay

* **交易查询接口**
	
	alipay.trade.query
	
#### 通知

* **验证是否是支付宝发来的通知**
	
	notify_verify
	
* **通知内容转换**
	
	将支付宝的通知内容转换为 Golang 的结构体
	
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

var html, _ = client.TradeWapPay(p)
// 将html输出到浏览器
```

#### 验证支付结果

有支付或者其它动作发生后，支付宝服务器会调用我们提供的 Notify URL，并向其传递会相关的信息。参考[手机网站支付结果异步通知](https://doc.open.alipay.com/docs/doc.htm?spm=a219a.7629140.0.0.XM5C4a&treeId=203&articleId=105286&docType=1)。

我们需要在提供的 Notify URL 服务中获取相关的参数并进行验证:

```Golang

http.HandleFunc("/alipay", func(rep http.ResponseWriter, req *http.Request) {
	var noti = alipay.GetTradeNotification(req)
	fmt.Println(noti)
	fmt.Println(client.NotifyVerify(noti.NotifyId))
})
```

如果 **client.NotifyVerify()** 方法返回的是 **true**，则表示是支付宝发送的通知，为了安全，切记这一步流程不可少。


此验证方法适用于支付宝所有情况下发送的 Notify，不管是手机 App 支付还是 Wap 支付。

