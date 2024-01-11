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
