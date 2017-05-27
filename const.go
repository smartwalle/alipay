package alipay

const (
	TimeForamt = "2006-01-02 15:04:05"

	K_ALI_PAY_TRADE_STATUS_WAIT_BUYER_PAY = "WAIT_BUYER_PAY" // 交易创建，等待买家付款
	K_ALI_PAY_TRADE_STATUS_TRADE_CLOSED   = "TRADE_CLOSED"   // 未付款交易超时关闭，或支付完成后全额退款
	K_ALI_PAY_TRADE_STATUS_TRADE_SUCCESS  = "TRADE_SUCCESS"  // 交易支付成功
	K_ALI_PAY_TRADE_STATUS_TRADE_FINISHED = "TRADE_FINISHED" // 交易结束，不可退款

	K_ALI_PAY_SANDBOX_API_URL    = "https://openapi.alipaydev.com/gateway.do"
	K_ALI_PAY_PRODUCTION_API_URL = "https://openapi.alipay.com/gateway.do"

	FixFormat   = "JSON"
	FixCharset  = "utf-8"
	FixVersion  = "1.0"
	FixSignType = "RSA2"

	// https://doc.open.alipay.com/docs/doc.htm?treeId=291&articleId=105806&docType=1
	// TODO:返回码根据文档统一处理,供外部调用解析
	SUCCESS_CODE = "10000"
)
