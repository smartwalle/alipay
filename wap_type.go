package alipay

// TradeWapPay 手机网站支付接口请求参数 https://docs.open.alipay.com/api_1/alipay.trade.wap.pay/
type TradeWapPay struct {
	Trade
	QuitURL    string `json:"quit_url,omitempty"`
	AuthToken  string `json:"auth_token,omitempty"`  // 针对用户授权接口，获取用户相关数据时，用于标识用户授权关系
	TimeExpire string `json:"time_expire,omitempty"` // 绝对超时时间，格式为yyyy-MM-dd HH:mm。
}

func (this TradeWapPay) APIName() string {
	return "alipay.trade.wap.pay"
}

func (this TradeWapPay) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["notify_url"] = this.NotifyURL
	m["return_url"] = this.ReturnURL
	return m
}
