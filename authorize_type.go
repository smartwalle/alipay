package alipay

const (
	kProductionPublicAppAuthorize = "https://openauth.alipay.com/oauth2/publicAppAuthorize.htm"
	kSandboxPublicAppAuthorize    = "https://openauth.alipaydev.com/oauth2/publicAppAuthorize.htm"
)

type SystemOauthToken struct {
	AppAuthToken string `json:"-"`          // 可选
	GrantType    string `json:"grant_type"` // 值为 authorization_code 时，代表用code换取；值为refresh_token时，代表用refresh_token换取
	Code         string `json:"code"`
	RefreshToken string `json:"refresh_token"`
}

func (this SystemOauthToken) APIName() string {
	return "alipay.system.oauth.token"
}

func (this SystemOauthToken) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	m["grant_type"] = this.GrantType
	if this.Code != "" {
		m["code"] = this.Code
	}
	if this.RefreshToken != "" {
		m["refresh_token"] = this.RefreshToken
	}
	return m
}

func (this SystemOauthToken) ExtJSONParamName() string {
	return "biz_content"
}

func (this SystemOauthToken) ExtJSONParamValue() string {
	return marshal(this)
}

type SystemOauthTokenRsp struct {
	Content struct {
		Code         string `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		UserId       string `json:"user_id"`
		AccessToken  string `json:"access_token"`
		ExpiresIn    int64  `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		ReExpiresIn  int64  `json:"re_expires_in"`
	} `json:"alipay_system_oauth_token_response"`
	Error *struct {
		Code    string `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"error_response"` // 不要访问此结构体
	Sign string `json:"sign"`
}
