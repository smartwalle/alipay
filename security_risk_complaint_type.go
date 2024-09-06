package alipay

// SecurityRiskComplaintInfo 投诉详情信息
type SecurityRiskComplaintInfo struct {
	Id                     int64                            `json:"id"`                        //投诉主表的主键id，查询详情时使用本id进行查询
	OppositePid            string                           `json:"opposite_pid"`              //被投诉人pid
	OppositeName           string                           `json:"opposite_name"`             //投诉单中被投诉方的名称
	ComplainAmount         string                           `json:"complain_amount"`           //投诉单涉及交易总金额（单位：人民币元）
	Contact                string                           `json:"contact"`                   //联系方式
	GmtComplain            string                           `json:"gmt_complain"`              //投诉时间
	GmtProcess             string                           `json:"gmt_process"`               //处理时间
	GmtOverdue             string                           `json:"gmt_overdue"`               //过期时间
	ComplainContent        string                           `json:"complain_content"`          //用户投诉内容
	TradeNo                string                           `json:"trade_no"`                  //投诉交易单号
	Status                 string                           `json:"status"`                    //投诉状态
	StatusDescription      string                           `json:"status_description"`        //投诉单状态枚举值描述
	ProcessCode            string                           `json:"process_code"`              //商家处理结果码
	ProcessMessage         string                           `json:"process_message"`           //商家处理结果码对应描述
	ProcessRemark          string                           `json:"process_remark"`            //商家处理备注
	ProcessImgUrlList      []string                         `json:"process_img_url_list"`      //商家处理备注图片url列表
	GmtRiskFinishTime      string                           `json:"gmt_risk_finish_time"`      //推送时间
	ComplaintTradeInfoList []SecurityRiskComplaintTradeInfo `json:"complaint_trade_info_list"` //涉及的交易信息
	TaskId                 string                           `json:"task_id"`                   //投诉单号id
	ComplainUrl            string                           `json:"complain_url"`              //投诉网址
	CertifyInfo            []string                         `json:"certify_info"`              //投诉凭证图片信息 【注意事项】仅对定向商户开放
}

// SecurityRiskComplaintTradeInfo 投诉单涉及的交易信息
type SecurityRiskComplaintTradeInfo struct {
	Id                int64  `json:"id"`                  //交易信息表主键id
	ComplaintRecordId int64  `json:"complaint_record_id"` //投诉主表id
	TradeNo           string `json:"trade_no"`            //支付宝交易单号
	OutNo             string `json:"out_no"`              //商家订单号
	GmtTrade          string `json:"gmt_trade"`           //交易时间
	GmtRefund         string `json:"gmt_refund"`          //退款时间
	Status            string `json:"status"`              //交易投诉状态
	StatusDescription string `json:"status_description"`  //交易投诉状态描述
	Amount            string `json:"amount"`              //交易单金额（单位：人民币元）
}

// SecurityRiskComplaintInfoBatchQueryReq 查询消费者投诉列表请求参数 https://opendocs.alipay.com/open/8ad1ac86_alipay.security.risk.complaint.info.batchquery?pathHash=e92f2f9f
type SecurityRiskComplaintInfoBatchQueryReq struct {
	AuxParam
	AppAuthToken      string   `json:"-"`
	CurrentPageNum    int64    `json:"current_page_num,omitempty"`    // 分页查询页码，不传则默认为1
	PageSize          int64    `json:"page_size,omitempty"`           // 分页查询每次查询的数据量，不传则默认为10
	GmtComplaintStart string   `json:"gmt_complaint_start,omitempty"` // 按投诉时间范围进行查询：时间范围的下界【示例值】2020-01-01 00:00:00
	GmtComplaintEnd   string   `json:"gmt_complaint_end,omitempty"`   // 按投诉时间范围进行查询：时间范围的上界【示例值】2020-01-01 00:00:02
	TradeNo           string   `json:"trade_no,omitempty"`            // 查询条件：交易单号
	StatusList        []string `json:"status_list,omitempty"`         // 查询条件：投诉状态列表
	GmtProcessStart   string   `json:"gmt_process_start,omitempty"`   // 按处理时间范围进行查询：时间范围的下界
	GmtProcessEnd     string   `json:"gmt_process_end,omitempty"`     // 按处理时间范围进行查询：时间范围的上界
	TaskId            string   `json:"task_id,omitempty"`             // 投诉单号
	TaskIdList        []string `json:"task_id_list,omitempty"`        // 投诉单号列表
}

// SecurityRiskComplaintInfoBatchQueryRsp 查询消费者投诉列表请求结果
type SecurityRiskComplaintInfoBatchQueryRsp struct {
	Error
	TotalSize     int64                       `json:"total_size"`     // 满足条件的数据总条数
	CurrentPage   int64                       `json:"current_page"`   // 分页查询时的当前页码
	PageSize      int64                       `json:"page_size"`      // 分页查询时每页大小
	ComplaintList []SecurityRiskComplaintInfo `json:"complaint_list"` // 投诉详情信息列表
}

func (req SecurityRiskComplaintInfoBatchQueryReq) APIName() string {
	return "alipay.security.risk.complaint.info.batchquery"
}

func (req SecurityRiskComplaintInfoBatchQueryReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = req.AppAuthToken
	return m
}

// SecurityRiskComplaintInfoQueryReq 查询消费者投诉详情 https://opendocs.alipay.com/open/271499b9_alipay.security.risk.complaint.info.query?pathHash=44398ac2
type SecurityRiskComplaintInfoQueryReq struct {
	AuxParam
	AppAuthToken string `json:"-"`
	ComplainId   int64  `json:"complain_id"` //投诉主表主键id
}

// SecurityRiskComplaintInfoQueryRsp 查询消费者投诉详情结果
type SecurityRiskComplaintInfoQueryRsp struct {
	Error
	SecurityRiskComplaintInfo
}

func (req SecurityRiskComplaintInfoQueryReq) APIName() string {
	return "alipay.security.risk.complaint.info.query"
}

func (req SecurityRiskComplaintInfoQueryReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = req.AppAuthToken
	return m
}

// SecurityRiskComplaintProcessFinishReq 处理消费者投诉 https://opendocs.alipay.com/open/da75e1ec_alipay.security.risk.complaint.process.finish?pathHash=b45c30c5
type SecurityRiskComplaintProcessFinishReq struct {
	AuxParam
	AppAuthToken string                                `json:"-"`
	IdList       []int64                               `json:"id_list"`                 // 本次进行处理的投诉id列表(主表主键)
	ProcessCode  string                                `json:"process_code"`            //商家处理投诉结果码(查阅文档)
	Remark       string                                `json:"remark,omitempty"`        //本次投诉处理的备注信息
	ImgFileList  []SecurityRiskComplaintProcessImgFile `json:"img_file_list,omitempty"` //投诉处理附加上传的图片文件列表
}

// SecurityRiskComplaintProcessImgFile  投诉处理附加上传的图片
type SecurityRiskComplaintProcessImgFile struct {
	ImgUrl    string `json:"img_url"`     //调用投诉文件上传接口返回的文件url
	ImgUrlKey string `json:"img_url_key"` //调用投诉文件上传接口返回的文件key
}

// SecurityRiskComplaintProcessFinishRsp 处理消费者投诉结果
type SecurityRiskComplaintProcessFinishRsp struct {
	Error
	ComplaintProcessSuccess bool `json:"complaint_process_success"` //本次投诉处理是否成功，表示系统后台是否成功收到本次请求并完成处理流程
}

func (req SecurityRiskComplaintProcessFinishReq) APIName() string {
	return "alipay.security.risk.complaint.process.finish"
}

func (req SecurityRiskComplaintProcessFinishReq) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = req.AppAuthToken
	return m
}
