package alipay

import (
	"encoding/json"
)

type AliPayParam interface {
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
