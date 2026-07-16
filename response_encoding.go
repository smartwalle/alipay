package alipay

import (
	"encoding/json"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// unmarshalResponseJSON 兼容支付宝偶发返回的 GBK 编码响应。
//
// SDK 请求声明的编码是 UTF-8，但部分业务错误响应的中文字段
// 仍可能以 GBK 编码返回。encoding/json 会将 JSON 字符串中的非法
// UTF-8 字节替换为 U+FFFD，从而丢失原始错误信息。
//
// 对于带签名的响应，调用方必须先使用原始字节完成验签，
// 再调用本函数，避免转码改变待验签内容。
func unmarshalResponseJSON(data []byte, dest interface{}) error {
	if !utf8.Valid(data) {
		decoded, err := simplifiedchinese.GBK.NewDecoder().Bytes(data)
		if err != nil {
			return err
		}
		data = decoded
	}
	return json.Unmarshal(data, dest)
}
