package alipay

import (
	"net/url"
)

// AgreementPageSign 支付宝个人协议页面签约接口 https://docs.open.alipay.com/api_2/alipay.user.agreement.page.sign
func (this *Client) AgreementPageSign(param AgreementPageSign) (result *url.URL, err error) {
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

// AgreementQuery 支付宝个人代扣协议查询接口 https://docs.open.alipay.com/api_2/alipay.user.agreement.query
func (this *Client) AgreementQuery(param AgreementQuery) (result *AgreementQueryRsp, err error) {
	err = this.DoRequest("POST", param, &result)
	return result, err
}

// AgreementUnsign 支付宝个人代扣协议解约接口 https://docs.open.alipay.com/api_2/alipay.user.agreement.unsign
func (this *Client) AgreementUnsign(param AgreementUnsign) (result *AgreementUnsignRsp, err error) {
	err = this.DoRequest("POST", param, &result)
	return result, err
}
