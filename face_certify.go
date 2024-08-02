package alipay

import "context"

// FaceCertifyInitialize Web人脸核身初始化 https://opendocs.alipay.com/open/02zloa?scene=common&pathHash=b0b7fece
func (c *Client) FaceCertifyInitialize(ctx context.Context, param FaceCertifyInitialize) (result *FaceCertifyInitializeRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// FaceCertifyVerify Web人脸核身开始认证 https://opendocs.alipay.com/open/02zlob?scene=common&pathHash=20eba12a
func (c *Client) FaceCertifyVerify(ctx context.Context, param FaceCertifyVerify) (result *FaceCertifyVerifyRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// FaceCertifyQuery Web人脸核身记录查询 https://opendocs.alipay.com/open/02zloc?scene=common&pathHash=b1141506
func (c *Client) FaceCertifyQuery(ctx context.Context, param FaceCertifyQuery) (result *FaceCertifyQueryRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// FaceVerificationInitialize 人脸核身初始化 https://opendocs.alipay.com/open/07260073_datadigital.fincloud.generalsaas.face.verification.initialize?scene=common&pathHash=0572cc86
func (c *Client) FaceVerificationInitialize(ctx context.Context, param FaceVerificationInitialize) (result *VerificationInitializeRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// FaceVerificationQuery 人脸核身结果查询 https://opendocs.alipay.com/open/9438eff0_datadigital.fincloud.generalsaas.face.verification.query?scene=common&pathHash=1608a398
func (c *Client) FaceVerificationQuery(ctx context.Context, param FaceVerificationQuery) (result *FaceVerificationQueryRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}
