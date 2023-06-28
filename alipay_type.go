package alipay

import (
	"encoding/json"
	"fmt"
)

const (
	kPastSandboxGateway    = "https://openapi.alipaydev.com/gateway.do"
	kNewSandboxGateway     = "https://openapi-sandbox.dl.alipaydev.com/gateway.do"
	kProductionGateway     = "https://openapi.alipay.com/gateway.do"
	kProductionMAPIGateway = "https://mapi.alipay.com/gateway.do"

	kFormat       = "JSON"
	kCharset      = "utf-8"
	kVersion      = "1.0"
	kSignTypeRSA2 = "RSA2"
	kContentType  = "application/x-www-form-urlencoded;charset=utf-8"
	kTimeFormat   = "2006-01-02 15:04:05"
)

const (
	kResponseSuffix    = "_response"
	kErrorResponse     = "error_response"
	kSignFieldName     = "sign"
	kSignTypeFieldName = "sign_type"
	kCertSNFieldName   = "alipay_cert_sn"
	kCertificateBegin  = "-----BEGIN CERTIFICATE-----"
	kCertificateEnd    = "-----END CERTIFICATE-----"
)

// Code 支付宝接口响应错误码 https://doc.open.alipay.com/docs/doc.htm?treeId=291&articleId=105806&docType=1
type Code string

func (c Code) IsSuccess() bool {
	return c == CodeSuccess
}

func (c Code) IsFailure() bool {
	return c != CodeSuccess
}

// 公共错误码 https://opendocs.alipay.com/common/02km9f#API%20%E5%85%AC%E5%85%B1%E9%94%99%E8%AF%AF%E7%A0%81
const (
	CodeSuccess                Code = "10000" // 接口调用成功
	CodeUnknowError            Code = "20000" // 服务不可用
	CodeInvalidAuthToken       Code = "20001" // 授权权限不足
	CodeMissingParam           Code = "40001" // 缺少必选参数
	CodeInvalidParam           Code = "40002" // 非法的参数
	CodeInsufficientConditions Code = "40003" // 条件异常
	CodeBusinessFailed         Code = "40004" // 业务处理失败
	CodeCallLimited            Code = "40005" // 调用频次超限
	CodePermissionDenied       Code = "40006" // 权限不足
)

type Param interface {
	// APIName 用于提供访问的 method
	APIName() string

	// Params 返回公共请求参数
	Params() map[string]string

	// FileParams 返回文件参数
	FileParams() map[string]*FileItem

	// NeedEncrypt 返回该接口是否支持内容加密，有的接口不支持内容加密，比如文件上传接口：alipay.open.file.upload
	NeedEncrypt() bool
}

type AuxParam struct {
}

func (this AuxParam) FileParams() map[string]*FileItem {
	return nil
}

func (this AuxParam) NeedEncrypt() bool {
	return true
}

type FileItem struct {
	Name     string
	Filename string
	Filepath string
}

type Payload struct {
	method  string
	Encrypt bool
	param   map[string]string
	biz     map[string]interface{}
	files   map[string]*FileItem
}

func NewPayload(method string) *Payload {
	var nPayload = &Payload{}
	nPayload.method = method
	nPayload.Encrypt = true
	nPayload.param = make(map[string]string)
	nPayload.biz = make(map[string]interface{})
	return nPayload
}

func (this *Payload) APIName() string {
	return this.method
}

func (this *Payload) Params() map[string]string {
	return this.param
}

func (this *Payload) FileParams() map[string]*FileItem {
	return this.files
}

func (this *Payload) NeedEncrypt() bool {
	return this.Encrypt
}

// AddParam 添加公共请求参数
//
// 例如：https://opendocs.alipay.com/apis/api_1/alipay.trade.query/#%E5%85%AC%E5%85%B1%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0
func (this *Payload) AddParam(key, value string) *Payload {
	this.param[key] = value
	return this
}

// Set 添加请求参数(业务相关)，这里添加的参数会序列化成 JSON 字符串，然后通过 biz_content 参数传递
//
// 例如：https://opendocs.alipay.com/apis/api_1/alipay.trade.query/#%E8%AF%B7%E6%B1%82%E5%8F%82%E6%95%B0
func (this *Payload) Set(key string, value interface{}) *Payload {
	this.biz[key] = value
	return this
}

// AddFile 添加需要上传的文件
//
// name: 参数名称
// filename: 文件名称
// filepath: 本地文件完整路径
func (this *Payload) AddFile(name, filename, filepath string) {
	if this.files == nil {
		this.files = make(map[string]*FileItem)
	}

	if filename == "" {
		filename = name
	}
	this.files[name] = &FileItem{Name: name, Filename: filename, Filepath: filepath}
}

func (this *Payload) MarshalJSON() ([]byte, error) {
	return json.Marshal(this.biz)
}

type Error struct {
	Code    Code   `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

func (this Error) Error() string {
	return fmt.Sprintf("%s - %s", this.Code, this.SubMsg)
}

func (this Error) IsSuccess() bool {
	return this.Code.IsSuccess()
}

func (this Error) IsFailure() bool {
	return this.Code.IsFailure()
}

const (
	kCertDownloadAPI = "alipay.open.app.alipaycert.download"
)

// CertDownload 应用支付宝公钥证书下载 https://opendocs.alipay.com/common/06ue2z
type CertDownload struct {
	AuxParam
	AppAuthToken string `json:"-"`              // 可选
	AliPayCertSN string `json:"alipay_cert_sn"` // 支付宝公钥证书序列号
}

func (this CertDownload) APIName() string {
	return kCertDownloadAPI
}

func (this CertDownload) Params() map[string]string {
	var m = make(map[string]string)
	m["app_auth_token"] = this.AppAuthToken
	return m
}

func (this CertDownload) NeedEncrypt() bool {
	return false
}

type CertDownloadRsp struct {
	Error
	AliPayCertContent string `json:"alipay_cert_content"`
}
