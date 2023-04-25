package alipay

import (
	"bytes"
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/smartwalle/ncrypto"
)

var (
	ErrBadResponse          = errors.New("alipay: bad response")
	ErrSignNotFound         = errors.New("alipay: sign content not found")
	ErrAliPublicKeyNotFound = errors.New("alipay: alipay public key not found")
)

const (
	kAliPayPublicKeySN = "alipay-public-key"
	kAppAuthToken      = "app_auth_token"
)

type Client struct {
	mu                 sync.Mutex
	isProduction       bool
	appId              string
	apiGateway         string
	notifyVerifyDomain string
	Client             *http.Client
	location           *time.Location

	// 内容加密
	encryptNeed    bool
	encryptIV      []byte
	encryptType    string
	encryptKey     []byte
	encryptPadding ncrypto.Padding

	appPrivateKey    *rsa.PrivateKey // 应用私钥
	appPublicCertSN  string
	aliRootCertSN    string
	aliPublicCertSN  string
	aliPublicKeyList map[string]*rsa.PublicKey
}

type OptionFunc func(c *Client)

func WithTimeLocation(location *time.Location) OptionFunc {
	return func(c *Client) {
		c.location = location
	}
}

func WithHTTPClient(client *http.Client) OptionFunc {
	return func(c *Client) {
		c.Client = client
	}
}

func WithSandboxGateway(gateway string) OptionFunc {
	return func(c *Client) {
		if gateway == "" {
			gateway = kSandboxGateway
		}
		if !c.isProduction {
			c.apiGateway = gateway
		}
	}
}

func WithProductionGateway(gateway string) OptionFunc {
	return func(c *Client) {
		if gateway == "" {
			gateway = kProductionGateway
		}
		if c.isProduction {
			c.apiGateway = gateway
		}
	}
}

// New 初始化支付宝客户端
//
// appId - 支付宝应用 id
//
// privateKey - 应用私钥，开发者自己生成
//
// isProduction - 是否为生产环境，传 false 的时候为沙箱环境，用于开发测试，正式上线的时候需要改为 true
func New(appId, privateKey string, isProduction bool, opts ...OptionFunc) (client *Client, err error) {
	priKey, err := ncrypto.ParsePKCS1PrivateKey(ncrypto.FormatPKCS1PrivateKey(privateKey))
	if err != nil {
		priKey, err = ncrypto.ParsePKCS8PrivateKey(ncrypto.FormatPKCS8PrivateKey(privateKey))
		if err != nil {
			return nil, err
		}
	}
	client = &Client{}
	client.isProduction = isProduction
	client.appId = appId

	if client.isProduction {
		client.apiGateway = kProductionGateway
		client.notifyVerifyDomain = kProductionMAPIGateway
	} else {
		client.apiGateway = kSandboxGateway
		client.notifyVerifyDomain = kSandboxGateway
	}
	client.Client = http.DefaultClient
	client.location = time.Local

	client.appPrivateKey = priKey
	client.aliPublicKeyList = make(map[string]*rsa.PublicKey)

	for _, opt := range opts {
		if opt != nil {
			opt(client)
		}
	}

	return client, nil
}

func (this *Client) IsProduction() bool {
	return this.isProduction
}

// SetEncryptKey 接口内容加密密钥 https://opendocs.alipay.com/common/02mse3
func (this *Client) SetEncryptKey(key string) error {
	if key == "" {
		this.encryptNeed = false
		return nil
	}

	var data, err = base64.StdEncoding.DecodeString(key)
	if err != nil {
		return err
	}
	this.encryptNeed = true
	this.encryptIV = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	this.encryptType = "AES"
	this.encryptKey = data
	this.encryptPadding = ncrypto.NewPKCS5Padding()
	return nil
}

// LoadAliPayPublicKey 加载支付宝公钥
func (this *Client) LoadAliPayPublicKey(aliPublicKey string) error {
	var pub *rsa.PublicKey
	var err error
	if len(aliPublicKey) < 0 {
		return ErrAliPublicKeyNotFound
	}
	pub, err = ncrypto.ParsePublicKey(ncrypto.FormatPublicKey(aliPublicKey))
	if err != nil {
		return err
	}
	this.mu.Lock()
	this.aliPublicCertSN = kAliPayPublicKeySN
	this.aliPublicKeyList[this.aliPublicCertSN] = pub
	this.mu.Unlock()
	return nil
}

// LoadAppPublicCert 加载应用公钥证书
func (this *Client) LoadAppPublicCert(s string) error {
	cert, err := ncrypto.ParseCertificate([]byte(s))
	if err != nil {
		return err
	}
	this.appPublicCertSN = getCertSN(cert)
	return nil
}

// LoadAppPublicCertFromFile 加载应用公钥证书
func (this *Client) LoadAppPublicCertFromFile(filename string) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return this.LoadAppPublicCert(string(b))
}

// LoadAliPayPublicCert 加载支付宝公钥证书
func (this *Client) LoadAliPayPublicCert(s string) error {
	cert, err := ncrypto.ParseCertificate([]byte(s))
	if err != nil {
		return err
	}

	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if ok == false {
		return nil
	}

	this.mu.Lock()
	this.aliPublicCertSN = getCertSN(cert)
	this.aliPublicKeyList[this.aliPublicCertSN] = key
	this.mu.Unlock()

	return nil
}

// LoadAliPayPublicCertFromFile 加载支付宝公钥证书
func (this *Client) LoadAliPayPublicCertFromFile(filename string) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return this.LoadAliPayPublicCert(string(b))
}

// LoadAliPayRootCert 加载支付宝根证书
func (this *Client) LoadAliPayRootCert(s string) error {
	var certStrList = strings.Split(s, kCertificateEnd)

	var certSNList = make([]string, 0, len(certStrList))

	for _, certStr := range certStrList {
		certStr = certStr + kCertificateEnd

		var cert, _ = ncrypto.ParseCertificate([]byte(certStr))
		if cert != nil && (cert.SignatureAlgorithm == x509.SHA256WithRSA || cert.SignatureAlgorithm == x509.SHA1WithRSA) {
			certSNList = append(certSNList, getCertSN(cert))
		}
	}

	this.aliRootCertSN = strings.Join(certSNList, "_")
	return nil
}

// LoadAliPayRootCertFromFile 加载支付宝根证书
func (this *Client) LoadAliPayRootCertFromFile(filename string) error {
	b, err := os.ReadFile(filename)

	if err != nil {
		return err
	}

	return this.LoadAliPayRootCert(string(b))
}

func (this *Client) URLValues(param Param) (value url.Values, err error) {
	var values = url.Values{}
	values.Add("app_id", this.appId)
	values.Add("method", param.APIName())
	values.Add("format", kFormat)
	values.Add("charset", kCharset)
	values.Add("sign_type", kSignTypeRSA2)
	values.Add("timestamp", time.Now().In(this.location).Format(kTimeFormat))
	values.Add("version", kVersion)
	if this.appPublicCertSN != "" {
		values.Add("app_cert_sn", this.appPublicCertSN)
	}
	if this.aliRootCertSN != "" {
		values.Add("alipay_root_cert_sn", this.aliRootCertSN)
	}

	jsonBytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var content = string(jsonBytes)
	if this.encryptNeed && param.APIName() != kCertDownloadAPI {
		jsonBytes, err = ncrypto.AESCBCEncrypt(jsonBytes, this.encryptKey, this.encryptIV, this.encryptPadding)
		if err != nil {
			return nil, err
		}
		content = base64.StdEncoding.EncodeToString(jsonBytes)
		values.Add("encrypt_type", this.encryptType)
	}
	values.Add("biz_content", content)

	var params = param.Params()
	if params != nil {
		for key, value := range params {
			if key == kAppAuthToken && value == "" {
				continue
			}
			values.Add(key, value)
		}
	}

	sign, err := signWithPKCS1v15(values, this.appPrivateKey, crypto.SHA256)
	if err != nil {
		return nil, err
	}
	values.Add("sign", sign)
	return values, nil
}

func (this *Client) doRequest(method string, param Param, result interface{}) (err error) {
	var body io.Reader
	if param != nil {
		var values url.Values
		values, err = this.URLValues(param)
		if err != nil {
			return err
		}
		body = strings.NewReader(values.Encode())
	}

	req, err := http.NewRequest(method, this.apiGateway, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", kContentType)

	rsp, err := this.Client.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	bodyBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	var data = string(bodyBytes)

	var rootNodeName = strings.Replace(param.APIName(), ".", "_", -1) + kResponseSuffix

	var rootIndex = strings.LastIndex(data, rootNodeName)
	var errorIndex = strings.LastIndex(data, kErrorResponse)

	var content string
	var certSN string
	var sign string
	var contentBytes []byte
	var jsonBytes []byte

	if rootIndex > 0 {
		content, certSN, sign = parseJSONSource(data, rootNodeName, rootIndex)
		contentBytes = []byte(content)
	} else if errorIndex > 0 {
		content, certSN, sign = parseJSONSource(data, kErrorResponse, errorIndex)
		contentBytes = []byte(content)
	} else {
		return ErrBadResponse
	}

	// 没有签名数据直接返回
	if sign == "" && param.APIName() != kCertDownloadAPI {
		var rErr *ErrorRsp
		if err = json.Unmarshal(contentBytes, &rErr); err != nil {
			return err
		}
		if rErr.Code != CodeSuccess {
			return rErr
		}
	}

	// 解密并重组 JSON 数据
	jsonBytes, err = this.decrypt(data, content)
	if err != nil {
		return err
	}

	// 验签
	if param.APIName() != kCertDownloadAPI {
		publicKey, err := this.getAliPayPublicKey(certSN)
		if err != nil {
			return err
		}
		if ok, err := verifyBytes(contentBytes, sign, publicKey); ok == false {
			return err
		}
	}

	err = json.Unmarshal(jsonBytes, result)
	if err != nil {
		return err
	}

	return err
}

func (this *Client) decrypt(data, content string) ([]byte, error) {
	if len(content) > 1 && content[0] == '"' {
		ciphertext, err := base64.StdEncoding.DecodeString(content[1 : len(content)-1])
		if err != nil {
			return nil, err
		}
		plaintext, err := ncrypto.AESCBCDecrypt(ciphertext, this.encryptKey, this.encryptIV, this.encryptPadding)
		if err != nil {
			return nil, err
		}
		data = strings.Replace(data, content, string(plaintext), 1)
	}
	return []byte(data), nil
}

func (this *Client) DoRequest(method string, param Param, result interface{}) (err error) {
	return this.doRequest(method, param, result)
}

func (this *Client) VerifySign(values url.Values) (ok bool, err error) {
	var certSN = values.Get(kCertSNNodeName)
	publicKey, err := this.getAliPayPublicKey(certSN)
	if err != nil {
		return false, err
	}

	return verifyValues(values, publicKey)
}

func (this *Client) getAliPayPublicKey(certSN string) (key *rsa.PublicKey, err error) {
	this.mu.Lock()
	defer this.mu.Unlock()

	if certSN == "" {
		certSN = this.aliPublicCertSN
	}

	key = this.aliPublicKeyList[certSN]

	if key == nil {
		if this.isProduction {
			cert, err := this.downloadAliPayCert(certSN)
			if err != nil {
				return nil, err
			}

			var ok bool
			key, ok = cert.PublicKey.(*rsa.PublicKey)
			if ok == false {
				return nil, ErrAliPublicKeyNotFound
			}
		} else {
			return nil, ErrAliPublicKeyNotFound
		}
	}
	return key, nil
}

func (this *Client) CertDownload(param CertDownload) (result *CertDownloadRsp, err error) {
	err = this.doRequest("POST", param, &result)
	return result, err
}

func (this *Client) downloadAliPayCert(certSN string) (cert *x509.Certificate, err error) {
	var param = CertDownload{}
	param.AliPayCertSN = certSN
	rsp, err := this.CertDownload(param)
	if err != nil {
		return nil, err
	}
	certBytes, err := base64.StdEncoding.DecodeString(rsp.Content.AliPayCertContent)
	if err != nil {
		return nil, err
	}

	cert, err = ncrypto.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}

	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if ok == false {
		return nil, nil
	}

	this.aliPublicCertSN = getCertSN(cert)
	this.aliPublicKeyList[this.aliPublicCertSN] = key

	return cert, nil
}

func parseJSONSource(data string, nodeName string, nodeIndex int) (content, certSN, sign string) {
	var dataStartIndex = nodeIndex + len(nodeName) + 2
	var signIndex = strings.LastIndex(data, "\""+kSignNodeName+"\"")
	var certIndex = strings.LastIndex(data, "\""+kCertSNNodeName+"\"")
	var dataEndIndex int

	if signIndex > 0 && certIndex > 0 {
		dataEndIndex = int(math.Min(float64(signIndex), float64(certIndex))) - 1
	} else if certIndex > 0 {
		dataEndIndex = certIndex - 1
	} else if signIndex > 0 {
		dataEndIndex = signIndex - 1
	} else {
		dataEndIndex = len(data) - 1
	}

	var indexLen = dataEndIndex - dataStartIndex
	if indexLen < 0 {
		return "", "", ""
	}
	content = data[dataStartIndex:dataEndIndex]

	if certIndex > 0 {
		var certStartIndex = certIndex + len(kCertSNNodeName) + 4
		certSN = data[certStartIndex:]
		var certEndIndex = strings.Index(certSN, "\"")
		certSN = certSN[:certEndIndex]
	}

	if signIndex > 0 {
		var signStartIndex = signIndex + len(kSignNodeName) + 4
		sign = data[signStartIndex:]
		var signEndIndex = strings.LastIndex(sign, "\"")
		sign = sign[:signEndIndex]
	}

	return content, certSN, sign
}

func signWithPKCS1v15(values url.Values, privateKey *rsa.PrivateKey, hash crypto.Hash) (sign string, err error) {
	if values == nil {
		values = make(url.Values, 0)
	}

	var pList = make([]string, 0, 0)
	for key := range values {
		var value = strings.TrimSpace(values.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList)
	var src = strings.Join(pList, "&")
	signBytes, err := ncrypto.RSASignWithKey([]byte(src), privateKey, hash)
	if err != nil {
		return "", err
	}
	sign = base64.StdEncoding.EncodeToString(signBytes)
	return sign, nil
}

func verifyValues(values url.Values, publicKey *rsa.PublicKey) (ok bool, err error) {
	sign := values.Get(kSignNodeName)

	var keys = make([]string, 0, 0)
	for key := range values {
		if key == kSignNodeName || key == kSignTypeNodeName || key == kCertSNNodeName {
			continue
		}
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var buffer bytes.Buffer
	for _, key := range keys {
		vs := values[key]
		for _, v := range vs {
			if buffer.Len() > 0 {
				buffer.WriteByte('&')
			}
			buffer.WriteString(key)
			buffer.WriteByte('=')
			buffer.WriteString(v)
		}
	}

	return verifyBytes(buffer.Bytes(), sign, publicKey)
}

func verifyBytes(data []byte, sign string, key *rsa.PublicKey) (ok bool, err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	if err = ncrypto.RSAVerifyWithKey(data, signBytes, key, crypto.SHA256); err != nil {
		return false, err
	}
	return true, nil
}

func getCertSN(cert *x509.Certificate) string {
	var value = md5.Sum([]byte(cert.Issuer.String() + cert.SerialNumber.String()))
	return hex.EncodeToString(value[:])
}
