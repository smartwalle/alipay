package alipay

import (
	"encoding/json"
)

const (
	kSandboxURL        = "https://openapi.alipaydev.com/gateway.do"
	kProductionURL     = "https://openapi.alipay.com/gateway.do"
	kProductionMAPIURL = "https://mapi.alipay.com/gateway.do"

	kFormat      = "JSON"
	kCharset     = "utf-8"
	kVersion     = "1.0"
	kContentType = "application/x-www-form-urlencoded;charset=utf-8"
	kTimeFormat  = "2006-01-02 15:04:05"
)

const (
	kResponseSuffix = "_response"
	kErrorResponse  = "error_response"
	kSignNodeName   = "sign"
)

const (
	K_SIGN_TYPE_RSA2 = "RSA2"
	K_SIGN_TYPE_RSA  = "RSA"
)

const (
	// https://doc.open.alipay.com/docs/doc.htm?treeId=291&articleId=105806&docType=1
	K_SUCCESS_CODE = "10000"
)

type Param interface {
	// 用于提供访问的 method
	APIName() string

	// 返回参数列表
	Params() map[string]string

	// 返回扩展 JSON 参数的字段名称
	ExtJSONParamName() string

	// 返回扩展 JSON 参数的字段值
	ExtJSONParamValue() string
}

func marshal(obj interface{}) string {
	var bytes, err = json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}
