package alipay

import (
	"context"
	"net/url"
	"strings"
)

// PublicAppAuthorize 用户信息授权接口(网站支付宝登录快速接入) https://docs.open.alipay.com/289/105656#s3 (https://docs.open.alipay.com/263/105809)
func (c *Client) PublicAppAuthorize(scopes []string, redirectURI, state string) (result *url.URL, err error) {
	var domain = kSandboxPublicAppAuthorize
	if c.isProduction {
		domain = kProductionPublicAppAuthorize
	}

	var p = url.Values{}
	p.Set("app_id", c.appId)
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
func (c *Client) SystemOauthToken(ctx context.Context, param SystemOauthToken) (result *SystemOauthTokenRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// UserInfoShare 支付宝会员授权信息查询接口 https://docs.open.alipay.com/api_2/alipay.user.info.share
func (c *Client) UserInfoShare(ctx context.Context, param UserInfoShare) (result *UserInfoShareRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// AppToAppAuth 第三方应用授权接口 https://docs.open.alipay.com/20160728150111277227/intro
func (c *Client) AppToAppAuth(redirectURI, state string) (result *url.URL, err error) {
	var domain = kSandboxAppToAppAuth
	if c.isProduction {
		domain = kProductionAppToAppAuth
	}

	var p = url.Values{}
	p.Set("app_id", c.appId)
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

// OpenAuthTokenApp 换取应用授权令牌接口 https://docs.open.alipay.com/api_9/alipay.open.auth.token.app
func (c *Client) OpenAuthTokenApp(ctx context.Context, param OpenAuthTokenApp) (result *OpenAuthTokenAppRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// OpenAuthTokenAppQuery 查询某个应用授权AppAuthToken的授权信息 https://opendocs.alipay.com/isv/04hgcp?pathHash=7ea21afe
func (c *Client) OpenAuthTokenAppQuery(ctx context.Context, param OpenAuthTokenAppQuery) (result *OpenAuthTokenAppQueryRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// AccountAuth 支付宝登录时, 帮客户端做参数签名, 返回授权请求信息字串接口 https://docs.open.alipay.com/218/105327
func (c *Client) AccountAuth(param AccountAuth) (result string, err error) {
	var values = url.Values{}
	values.Add("app_id", c.appId)
	values.Add("method", param.APIName())

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			values.Add(key, value)
		}
	}

	values.Add("sign_type", kSignTypeRSA2)

	signature, err := c.sign(values)
	if err != nil {
		return "", err
	}
	values.Add("sign", signature)

	return values.Encode(), err
}

// OpenAuthAppAuthInviteCreate ISV向商户发起应用授权邀约 https://opendocs.alipay.com/isv/06evao?pathHash=f46ecafa
// TODO OpenAuthAppAuthInviteCreate 接口未经测试
func (c *Client) OpenAuthAppAuthInviteCreate(param OpenAuthAppAuthInviteCreate) (result *url.URL, err error) {
	return c.BuildURL(param)
}
