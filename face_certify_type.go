package alipay

type FaceCertifyInitialize struct {
	AuxParam

	// 可选
	AppAuthToken string `json:"-"`

	// 必选 商户请求的唯一标识，商户要保证其唯一性，值为32位长度的字母数字组合。
	// 建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列。
	// 示例：ZGYD201809132323000001234
	OuterOrderNo string `json:"outer_order_no"`

	// 必选 H5人脸核身场景码。入参支持的场景码。
	// 示例：FUTURE_TECH_BIZ_FACE_SDK
	BizCode CertifyBizCode `json:"biz_code"`

	// 必选 需要验证的身份信息
	IdentityParam IdentityParam `json:"identity_param"`

	// 必选 商户个性化配置信息
	MerchantConfig MerchantConfig `json:"merchant_config"`
}

func (u FaceCertifyInitialize) APIName() string {
	return "datadigital.fincloud.generalsaas.face.certify.initialize"
}

func (u FaceCertifyInitialize) Params() map[string]string {
	m := make(map[string]string)
	m["app_auth_token"] = u.AppAuthToken
	return m
}

type FaceCertifyInitializeRsp struct {
	Error

	// 必选 本次申请操作的唯一标识，商户需要记录，后续的操作都需要用到
	// 示例：2109b5e671aa3ff2eb4851816c65828f
	CertifyId string `json:"certify_id"`
}

type FaceCertifyVerify struct {
	AuxParam

	// 可选
	AppAuthToken string `json:"-"`

	// 必选 本次申请操作的唯一标识，由H5人脸核身初始化接口调用后生成，后续的操作都需要用到
	// 示例：OC201809253000000393900404029253
	CertifyId string `json:"certify_id"`
}

func (u FaceCertifyVerify) APIName() string {
	return "datadigital.fincloud.generalsaas.face.certify.verify"
}

func (u FaceCertifyVerify) Params() map[string]string {
	m := make(map[string]string)
	m["app_auth_token"] = u.AppAuthToken
	return m
}

type FaceCertifyVerifyRsp struct {
	Error

	// 必选 返回用于唤起刷脸页面的url
	// 示例：https://openapi.alipay.com/gateway.do?alipay_sdk=alipay-sdk-java-dynamicVersionNo&app_id=2015111100758155&biz_content=%7B%22certify_id%22%3A%22ZM201611253000000121200404215172%22%7D&charset=GBK&format=json&method=datadigital.fincloud.generalsaas.face.certify.verify&sign=MhtfosO8AKbwctDgfGitzLvhbcvi%2FMv3iBES7fRnIXn%2BHcdwq9UWltTs6mEvjk2UoHdLoFrvcSJipiE3sL8kdJMd51t87vcwPCfk7BA5KPwa4%2B1IYzYaK6WwbqOoQB%2FqiJVfni602HiE%2BZAomW7WA3Tjhjy3D%2B9xrLFCipiroDQ%3D&sign_type=RSA2×tamp=2016-11-25+15%3A00%3A59&version=1.0&sign=MhtfosO8AKbwctDgfGitzLvhbcvi%2FMv3iBES7fRnIXn%2BHcdwq9UWltTs6mEvjk2UoHdLoFrvcSJipiE3sL8kdJMd51t87vcwPCfk7BA5KPwa4%2B1IYzYaK6WwbqOoQB%2FqiJVfni602HiE%2BZAomW7WA3Tjhjy3D%2B9xrLFCipiroDQ%3D
	CertifyUrl string `json:"certify_url"`
}

type FaceCertifyQuery struct {
	AuxParam

	// 可选
	AppAuthToken string `json:"-"`

	// 必选 本次申请操作的唯一标识，通过datadigital.fincloud.generalsaas.face.certify.initialize 接口同步响应获取。
	// 示例：03cdsfsss20048373
	CertifyId string `json:"certify_id"`
}

func (u FaceCertifyQuery) APIName() string {
	return "datadigital.fincloud.generalsaas.face.certify.query"
}

func (u FaceCertifyQuery) Params() map[string]string {
	m := make(map[string]string)
	m["app_auth_token"] = u.AppAuthToken
	return m
}

type FaceCertifyQueryRsp struct {
	Error

	// 必选 是否通过，通过为T，不通过为F。
	// 示例：T
	Passed string `json:"passed"`
}

// FaceVerificationInitialize 人脸核身初始化请求参数 https://opendocs.alipay.com/open/07260073_datadigital.fincloud.generalsaas.face.verification.initialize?scene=common&pathHash=0572cc86
type FaceVerificationInitialize struct {
	AuxParam
	AppAuthToken string `json:"-"` // 可选

	OuterOrderNo string `json:"outer_order_no"` // 必选 商户请求的唯一标识，商户要保证其唯一性，值为64位长度的字母数字组合。建议：前面几位字符是商户自定义的简称，中间可以使用一段时间，后段可以使用一个随机或递增序列
	BizCode      string `json:"biz_code"`       // 必选 人脸核身具体类型目前仅支持：DATA_DIGITAL_BIZ_CODE_FACE_VERIFICATION
	IdentityType string `json:"identity_type"`  // 必选 认证类型，固定值为：CERT_INFO
	CertType     string `json:"cert_type"`      // 必选 证件类型，当前枚举支持： IDENTITY_CARD：身份证 RESIDENCE_HK_MC：港澳居民居住证 RESIDENCE_TAIWAN：台湾居民居住证
	CertName     string `json:"cert_name"`      // 必选 真实姓名
	CertNo       string `json:"cert_no"`        // 必选 证件号
}

func (t FaceVerificationInitialize) APIName() string {
	return "datadigital.fincloud.generalsaas.face.verification.initialize"
}

func (t FaceVerificationInitialize) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = t.AppAuthToken
	return m
}

// VerificationInitializeRsp 人脸核身初始化响应参数
type VerificationInitializeRsp struct {
	Error
	CertifyId string `json:"certify_id"` // 认证单据号，请保留以便排查问题。
	WebURL    string `json:"web_url"`    // 人脸核身url
}

// FaceVerificationQuery 人脸核身结果查询请求参数 https://opendocs.alipay.com/open/9438eff0_datadigital.fincloud.generalsaas.face.verification.query?scene=common&pathHash=1608a398
type FaceVerificationQuery struct {
	AuxParam
	AppAuthToken string `json:"-"` // 可选

	CertifyId string `json:"certify_id"` // 必选 填入人脸核身初始化阶段获取到的certify_id
}

func (t FaceVerificationQuery) APIName() string {
	return "datadigital.fincloud.generalsaas.face.verification.query"
}

func (t FaceVerificationQuery) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = t.AppAuthToken
	return m
}

// FaceVerificationQueryRsp 人脸核身结果查询响应参数
type FaceVerificationQueryRsp struct {
	Error
	CertifyState string   `json:"certify_state"` // 人脸认证状态。PROCESSING：初始化；SUCCESS：认证通过；FAIL：认证不通过。
	MetaInfo     MetaInfo `json:"meta_info"`     // 人脸认证元数据信息
	Score        string   `json:"score"`         // double值，活体检测结果分数
	Quality      string   `json:"quality"`       // double值，人脸图片质量分
	AlivePhoto   string   `json:"alive_photo"`   // base64过后的图片
	AttackFlag   string   `json:"attack_flag"`   // 本次认证是否存在安全风险，true：检测到安全风险；false：未检测到安全风险。
}

type MetaInfo struct {
	DeviceType string `json:"device_type"` // 设备操作系统类型 鸿蒙系统: harmony iOS系统: ios 安卓系统: android H5页面: h5
}
