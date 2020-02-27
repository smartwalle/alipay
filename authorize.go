package alipay

import (
	"crypto"
	"net/url"
	"strings"
)

// PublicAppAuthorize 用户信息授权接口(网站支付宝登录快速接入) https://docs.open.alipay.com/289/105656#s3 (https://docs.open.alipay.com/263/105809)
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
	return result, nil
}

// SystemOauthToken 换取授权访问令牌接口 https://docs.open.alipay.com/api_9/alipay.system.oauth.token
func (this *Client) SystemOauthToken(param SystemOauthToken) (result *SystemOauthTokenRsp, err error) {
	err = this.doRequest("POST", param, &result)
	if result != nil {
		if result.Error != nil {
			result.Content.Code = Code(result.Error.Code)
			result.Content.Msg = result.Error.Msg
			result.Content.SubCode = result.Error.SubCode
			result.Content.SubMsg = result.Error.SubMsg
		} else {
			result.Content.Code = CodeSuccess
		}
	}
	return result, err
}

// UserInfoShare 支付宝会员授权信息查询接口 https://docs.open.alipay.com/api_2/alipay.user.info.share
func (this *Client) UserInfoShare(param UserInfoShare) (result *UserInfoShareRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// AppToAppAuth 第三方应用授权接口 https://docs.open.alipay.com/20160728150111277227/intro
func (this *Client) AppToAppAuth(redirectURI string) (result *url.URL, err error) {
	var domain = kSandboxAppToAppAuth
	if this.isProduction {
		domain = kProductionAppToAppAuth
	}

	var p = url.Values{}
	p.Set("app_id", this.appId)
	p.Set("redirect_uri", redirectURI)

	result, err = url.Parse(domain + "?" + p.Encode())
	if err != nil {
		return nil, err
	}
	return result, nil
}

// OpenAuthTokenApp 换取应用授权令牌接口 https://docs.open.alipay.com/api_9/alipay.open.auth.token.app
func (this *Client) OpenAuthTokenApp(param OpenAuthTokenApp) (result *OpenAuthTokenAppRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

// AccountAuth 支付宝登录时, 帮客户端做参数签名, 返回授权请求信息字串接口 https://docs.open.alipay.com/218/105327
func (this *Client) AccountAuth(param AccountAuth) (result string, err error) {
	var p = url.Values{}
	p.Add("app_id", this.appId)
	p.Add("method", param.APIName())

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			p.Add(key, value)
		}
	}

	p.Add("sign_type", kSignTypeRSA2)

	sign, err := signWithPKCS1v15(p, this.appPrivateKey, crypto.SHA256)
	if err != nil {
		return "", err
	}
	p.Add("sign", sign)

	return p.Encode(), err
}
