package alipay

import (
	"net/url"
)

// 支付宝个人协议页面签约接口 https://gw.alipayobjects.com/os/skylark-tools/public/files/4f3f6885c46963e01eb1025e0c8c724d.pdf
func (this *Client) AgreementPageSign(param Param) (result *url.URL, err error) {
	p, err := this.URLValues(param)
	if err != nil {
		return nil, err
	}

	result, err = url.Parse(this.apiDomain + "?" + p.Encode())
	if err != nil {
		return nil, err
	}
	return result, err
}

// 支付宝个人代扣协议查询接口 https://gw.alipayobjects.com/os/skylark-tools/public/files/acda03bc9adeb36c1e1ca590e02bc464.pdf
func (this *Client) AgreementQuery(param Param) (result *AgreementQueryRsp, err error) {
	err = this.DoRequest("POST", param, &result)
	return result, err
}

// 支付宝个人代扣协议解约接口 https://gw.alipayobjects.com/os/skylark-tools/public/files/b5a44abc569b201923cb1e2a19d687be.pdf
func (this *Client) AgreementUnsign(param Param) (result *AgreementUnsignRsp, err error) {
	err = this.DoRequest("POST", param, &result)
	return result, err
}
