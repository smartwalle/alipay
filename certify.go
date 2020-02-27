package alipay

import "net/url"

// UserCertifyOpenInitialize 身份认证初始化服务接口 https://docs.open.alipay.com/api_2/alipay.user.certify.open.initialize
func (this *Client) UserCertifyOpenInitialize(param UserCertifyOpenInitialize) (result *UserCertifyOpenInitializeRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// UserCertifyOpenCertify 身份认证开始认证接口 https://docs.open.alipay.com/api_2/alipay.user.certify.open.certify
func (this *Client) UserCertifyOpenCertify(param UserCertifyOpenCertify) (result *url.URL, err error) {
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

// UserCertifyOpenQuery 身份认证记录查询接口 https://docs.open.alipay.com/api_2/alipay.user.certify.open.query/
func (this *Client) UserCertifyOpenQuery(param UserCertifyOpenQuery) (result *UserCertifyOpenQueryRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}
