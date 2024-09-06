package alipay

import "context"

// SecurityRiskComplaintInfoBatchQuery 查询消费者投诉列表请求参数 https://opendocs.alipay.com/open/8ad1ac86_alipay.security.risk.complaint.info.batchquery?pathHash=e92f2f9f
func (c *Client) SecurityRiskComplaintInfoBatchQuery(ctx context.Context, param SecurityRiskComplaintInfoBatchQueryReq) (result *SecurityRiskComplaintInfoBatchQueryRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// SecurityRiskComplaintInfoQuery 查询消费者投诉详情 https://opendocs.alipay.com/open/271499b9_alipay.security.risk.complaint.info.query?pathHash=44398ac2
func (c *Client) SecurityRiskComplaintInfoQuery(ctx context.Context, param SecurityRiskComplaintInfoQueryReq) (result *SecurityRiskComplaintInfoQueryRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}

// SecurityRiskComplaintProcessFinish 处理消费者投诉 https://opendocs.alipay.com/open/da75e1ec_alipay.security.risk.complaint.process.finish?pathHash=b45c30c5
func (c *Client) SecurityRiskComplaintProcessFinish(ctx context.Context, param SecurityRiskComplaintProcessFinishReq) (result *SecurityRiskComplaintProcessFinishRsp, err error) {
	err = c.doRequest(ctx, "POST", param, &result)
	return result, err
}
