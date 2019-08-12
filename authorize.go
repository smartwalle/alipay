package alipay

import (
	"net/url"
	"strings"
)

// PublicAppAuthorize 用户信息授权(网站支付宝登录快速接入) https://docs.open.alipay.com/289/105656#s3 (https://docs.open.alipay.com/263/105809)
func (this *Client) PublicAppAuthorize(scopes []string, redirectURI, state string) (result *url.URL, err error) {
	var domain = kSandboxPublicAppAuthorize
	if this.isProduction {
		domain = kProductionPublicAppAuthorize
	}

	var p = url.Values{}
	p.Set("app_id", this.appId)
	p.Set("scope", strings.Join(scopes, ","))
	p.Set("redirect_uri", redirectURI)
	if state != "" {
		p.Set("state", state)
	}

	result, err = url.Parse(domain + "?" + p.Encode())
	if err != nil {
		return nil, err
	}
	return result, err
}

// https://docs.open.alipay.com/api_9/alipay.system.oauth.token
func (this *Client) SystemOauthToken(param SystemOauthToken) (result *SystemOauthTokenRsp, err error) {
	err = this.doRequest("POST", param, &result)
	if result != nil {
		if result.Error != nil {
			result.Content.Code = result.Error.Code
			result.Content.Msg = result.Error.Msg
			result.Content.SubCode = result.Error.SubCode
			result.Content.SubMsg = result.Error.SubMsg
		} else {
			result.Content.Code = K_SUCCESS_CODE
		}
	}
	return result, err
}

// https://docs.open.alipay.com/api_2/alipay.user.info.share
func (this *Client) UserInfoShare(param UserInfoShare) (result *UserInfoShareRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}
