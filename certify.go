package alipay

import "net/url"

// UserCertifyOpenInitialize 身份认证初始化服务接口 https://docs.open.alipay.com/api_2/alipay.user.certify.open.initialize
func (c *Client) UserCertifyOpenInitialize(param UserCertifyOpenInitialize) (result *UserCertifyOpenInitializeRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// UserCertifyOpenCertify 身份认证开始认证接口 https://docs.open.alipay.com/api_2/alipay.user.certify.open.certify
func (c *Client) UserCertifyOpenCertify(param UserCertifyOpenCertify) (result *url.URL, err error) {
	return c.BuildURL(param)
}

// UserCertifyOpenQuery 身份认证记录查询接口 https://docs.open.alipay.com/api_2/alipay.user.certify.open.query/
func (c *Client) UserCertifyOpenQuery(param UserCertifyOpenQuery) (result *UserCertifyOpenQueryRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// UserCertDocCertVerifyPreConsult 实名证件信息比对验证预咨询 https://opendocs.alipay.com/apis/api_2/alipay.user.certdoc.certverify.preconsult
func (c *Client) UserCertDocCertVerifyPreConsult(param UserCertDocCertVerifyPreConsult) (result *UserCertDocCertVerifyPreConsultRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}

// UserCertDocCertVerifyConsult 实名证件信息比对验证咨询 https://opendocs.alipay.com/apis/api_2/alipay.user.certdoc.certverify.consult
func (c *Client) UserCertDocCertVerifyConsult(param UserCertDocCertVerifyConsult) (result *UserCertDocCertVerifyConsultRsp, err error) {
	err = c.doRequest("POST", param, &result)
	return result, err
}
