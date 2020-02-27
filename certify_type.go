package alipay

type CertifyBizCode string

const (
	CertifyBizCodeFace          CertifyBizCode = "FACE"            // 多因子人脸认证
	CertifyBizCodeCertPhoto     CertifyBizCode = "CERT_PHOTO"      // 多因子证照认证
	CertifyBizCodeCertPhotoFace CertifyBizCode = "CERT_PHOTO_FACE" // 多因子证照和人脸认证
	CertifyBizCodeSmartFace     CertifyBizCode = "SMART_FACE"      // 多因子快捷认证
)

// UserCertifyOpenInitialize 身份认证初始化服务接口请求参数 https://docs.open.alipay.com/api_2/alipay.user.certify.open.initialize
type UserCertifyOpenInitialize struct {
	AppAuthToken        string         `json:"-"`                               // 可选
	OuterOrderNo        string         `json:"outer_order_no"`                  // 必选  商户请求的唯一标识，商户要保证其唯一性，值为32位长度的字母数字组合。建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列
	BizCode             CertifyBizCode `json:"biz_code"`                        // 必选 认证场景码。入参支持的认证场景码和商户签约的认证场景相关，取值如下: FACE：多因子人脸认证 CERT_PHOTO：多因子证照认证 CERT_PHOTO_FACE ：多因子证照和人脸认证 SMART_FACE：多因子快捷认证
	IdentityParam       IdentityParam  `json:"identity_param"`                  // 必选
	MerchantConfig      MerchantConfig `json:"merchant_config"`                 // 必选 商户个性化配置，格式为json，详细支持的字段说明为： return_url：需要回跳的目标地址，必填，一般指定为商户业务页面
	FaceContrastPicture string         `json:"face_contrast_picture,omitempty"` // 可选 自定义人脸比对图片的base64编码格式的string字符串
}

type IdentityParam struct {
	IdentityType string `json:"identity_type"` // 身份信息参数类型，必填，必须传入CERT_INFO
	CertType     string `json:"cert_type"`     // 证件类型，必填，当前支持身份证，必须传入IDENTITY_CARD
	CertName     string `json:"cert_name"`     // 真实姓名，必填，填写需要验证的真实姓名
	CertNo       string `json:"cert_no"`       // 证件号码，必填，填写需要验证的证件号码
}

type MerchantConfig struct {
	ReturnURL string `json:"return_url"`
}

func (this UserCertifyOpenInitialize) APIName() string {
	return "alipay.user.certify.open.initialize"
}

func (this UserCertifyOpenInitialize) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

// UserCertifyOpenInitializeRsp 身份认证初始化服务接口响应参数
type UserCertifyOpenInitializeRsp struct {
	Content struct {
		Code      Code   `json:"code"`
		Msg       string `json:"msg"`
		SubCode   string `json:"sub_code"`
		SubMsg    string `json:"sub_msg"`
		CertifyId string `json:"certify_id"`
	} `json:"alipay_user_certify_open_initialize_response"`
	Sign string `json:"sign"`
}

// UserCertifyOpenCertify 身份认证开始认证接口请求参数 https://docs.open.alipay.com/api_2/alipay.user.certify.open.certify
type UserCertifyOpenCertify struct {
	AppAuthToken string `json:"-"`          // 可选
	CertifyId    string `json:"certify_id"` // 必选 本次申请操作的唯一标识，由开放认证初始化接口调用后生成，后续的操作都需要用到
}

func (this UserCertifyOpenCertify) APIName() string {
	return "alipay.user.certify.open.certify"
}

func (this UserCertifyOpenCertify) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

// UserCertifyOpenCertifyRsp 身份认证开始认证接口响应参数
type UserCertifyOpenCertifyRsp struct {
	Content struct {
		Code    Code   `json:"code"`
		Msg     string `json:"msg"`
		SubCode string `json:"sub_code"`
		SubMsg  string `json:"sub_msg"`
	} `json:"alipay_user_certify_open_certify_response"`
	Sign string `json:"sign"`
}

// UserCertifyOpenQuery 身份认证记录查询接口请求参数 https://docs.open.alipay.com/api_2/alipay.user.certify.open.query/
type UserCertifyOpenQuery struct {
	AppAuthToken string `json:"-"`          // 可选
	CertifyId    string `json:"certify_id"` // 必选 本次申请操作的唯一标识，由开放认证初始化接口调用后生成，后续的操作都需要用到
}

func (this UserCertifyOpenQuery) APIName() string {
	return "alipay.user.certify.open.query"
}

func (this UserCertifyOpenQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

// UserCertifyOpenQueryRsp 身份认证记录查询接口响应参数
type UserCertifyOpenQueryRsp struct {
	Content struct {
		Code         Code   `json:"code"`
		Msg          string `json:"msg"`
		SubCode      string `json:"sub_code"`
		SubMsg       string `json:"sub_msg"`
		Passed       string `json:"passed"`
		IdentityInfo string `json:"identity_info"`
		MaterialInfo string `json:"material_info"`
	} `json:"alipay_user_certify_open_query_response"`
	Sign string `json:"sign"`
}
