package alipay

// TradeWapPay 手机网站支付接口请求参数 https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
type TradeWapPay struct {
	Trade
	QuitURL    string `json:"quit_url,omitempty"`
	AuthToken  string `json:"auth_token,omitempty"`  // 针对用户授权接口，获取用户相关数据时，用于标识用户授权关系
	TimeExpire string `json:"time_expire,omitempty"` // 绝对超时时间，格式为yyyy-MM-dd HH:mm。
}

func (t TradeWapPay) APIName() string {
	return "alipay.trade.wap.pay"
}

func (t TradeWapPay) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = t.AppAuthToken
	m["notify_url"] = t.NotifyURL
	m["return_url"] = t.ReturnURL
	return m
}

// TradeWapMergePay 无线Wap合并支付接口2.0 https://opendocs.alipay.com/open/028xra
type TradeWapMergePay struct {
	AuxParam
	AppAuthToken string `json:"-"`            // 可选
	PreOrderNo   string `json:"pre_order_no"` // 必选 预下单号。通过 alipay.trade.merge.precreate(统一收单合并支付预创建接口)返回。
}

func (t TradeWapMergePay) APIName() string {
	return "alipay.trade.wap.merge.pay"
}

func (t TradeWapMergePay) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = t.AppAuthToken
	return m
}
