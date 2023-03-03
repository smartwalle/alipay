package alipay

import (
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

	"github.com/smartwalle/crypto4go"
)

var (
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
	apiDomain          string
	notifyVerifyDomain string
	Client             *http.Client
	location           *time.Location

	// 内容加密
	encryptNeed bool
	encryptIV   []byte
	encryptType string
	encryptKey  []byte

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

// New 初始化支付宝客户端
//
// appId - 支付宝应用 id
//
// privateKey - 应用私钥，开发者自己生成
//
// isProduction - 是否为生产环境，传 false 的时候为沙箱环境，用于开发测试，正式上线的时候需要改为 true
func New(appId, privateKey string, isProduction bool, opts ...OptionFunc) (client *Client, err error) {
	priKey, err := crypto4go.ParsePKCS1PrivateKey(crypto4go.FormatPKCS1PrivateKey(privateKey))
	if err != nil {
		priKey, err = crypto4go.ParsePKCS8PrivateKey(crypto4go.FormatPKCS8PrivateKey(privateKey))
		if err != nil {
			return nil, err
		}
	}
	client = &Client{}
	client.isProduction = isProduction
	client.appId = appId

	if client.isProduction {
		client.apiDomain = kProductionURL
		client.notifyVerifyDomain = kProductionMAPIURL
	} else {
		client.apiDomain = kSandboxURL
		client.notifyVerifyDomain = kSandboxURL
	}
	client.Client = http.DefaultClient
	client.location = time.Local

	client.appPrivateKey = priKey
	client.aliPublicKeyList = make(map[string]*rsa.PublicKey)

	for _, opt := range opts {
		opt(client)
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
	return nil
}

// LoadAliPayPublicKey 加载支付宝公钥
func (this *Client) LoadAliPayPublicKey(aliPublicKey string) error {
	var pub *rsa.PublicKey
	var err error
	if len(aliPublicKey) < 0 {
		return ErrAliPublicKeyNotFound
	}
	pub, err = crypto4go.ParsePublicKey(crypto4go.FormatPublicKey(aliPublicKey))
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
	cert, err := crypto4go.ParseCertificate([]byte(s))
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
	cert, err := crypto4go.ParseCertificate([]byte(s))
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

		var cert, _ = crypto4go.ParseCertificate([]byte(certStr))
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
	var p = url.Values{}
	p.Add("app_id", this.appId)
	p.Add("method", param.APIName())
	p.Add("format", kFormat)
	p.Add("charset", kCharset)
	p.Add("sign_type", kSignTypeRSA2)
	p.Add("timestamp", time.Now().In(this.location).Format(kTimeFormat))
	p.Add("version", kVersion)
	if this.appPublicCertSN != "" {
		p.Add("app_cert_sn", this.appPublicCertSN)
	}
	if this.aliRootCertSN != "" {
		p.Add("alipay_root_cert_sn", this.aliRootCertSN)
	}

	jsonBytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var content = string(jsonBytes)
	if this.encryptNeed {
		jsonBytes, err = crypto4go.AESCBCEncrypt(jsonBytes, this.encryptKey, this.encryptIV)
		if err != nil {
			return nil, err
		}
		content = base64.StdEncoding.EncodeToString(jsonBytes)
		p.Add("encrypt_type", this.encryptType)
	}
	p.Add("biz_content", content)

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			if key == kAppAuthToken && value == "" {
				continue
			}
			p.Add(key, value)
		}
	}

	sign, err := signWithPKCS1v15(p, this.appPrivateKey, crypto.SHA256)
	if err != nil {
		return nil, err
	}
	p.Add("sign", sign)
	return p, nil
}

func (this *Client) doRequest(method string, param Param, result interface{}) (err error) {
	var buf io.Reader
	if param != nil {
		p, err := this.URLValues(param)
		if err != nil {
			return err
		}
		buf = strings.NewReader(p.Encode())
	}

	req, err := http.NewRequest(method, this.apiDomain, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", kContentType)

	resp, err := this.Client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var body = string(bodyBytes)

	var rootNodeName = strings.Replace(param.APIName(), ".", "_", -1) + kResponseSuffix

	var rootIndex = strings.LastIndex(body, rootNodeName)
	var errorIndex = strings.LastIndex(body, kErrorResponse)

	var content string
	var certSN string
	var sign string
	var contentBytes []byte
	var jsonBytes []byte

	if rootIndex > 0 {
		content, certSN, sign = parseJSONSource(body, rootNodeName, rootIndex)
		contentBytes = []byte(content)
	} else if errorIndex > 0 {
		content, certSN, sign = parseJSONSource(body, kErrorResponse, errorIndex)
		contentBytes = []byte(content)
	} else {
		return ErrSignNotFound
	}

	// 没有签名数据直接返回
	if sign == "" {
		var errRsp *ErrorRsp
		if err = json.Unmarshal(contentBytes, &errRsp); err != nil {
			return err
		}
		if errRsp.Code != CodeSuccess {
			return errRsp
		}
	}

	// 解密并重组 JSON 数据
	jsonBytes, err = this.decrypt(body, content)
	if err != nil {
		return err
	}

	// 验签
	publicKey, err := this.getAliPayPublicKey(certSN)
	if err != nil {
		return err
	}
	if ok, err := verifyData(contentBytes, sign, publicKey); ok == false {
		return err
	}

	err = json.Unmarshal(jsonBytes, result)
	if err != nil {
		return err
	}

	return err
}

func (this *Client) decrypt(body, content string) ([]byte, error) {
	if len(content) > 1 && content[0] == '"' {
		ciphertext, err := base64.StdEncoding.DecodeString(content[1 : len(content)-1])
		if err != nil {
			return nil, err
		}
		plaintext, err := crypto4go.AESCBCDecrypt(ciphertext, this.encryptKey, this.encryptIV)
		if err != nil {
			return nil, err
		}
		body = strings.Replace(body, content, string(plaintext), 1)
	}
	return []byte(body), nil
}

func (this *Client) DoRequest(method string, param Param, result interface{}) (err error) {
	return this.doRequest(method, param, result)
}

func (this *Client) VerifySign(data url.Values) (ok bool, err error) {
	var certSN = data.Get(kCertSNNodeName)
	publicKey, err := this.getAliPayPublicKey(certSN)
	if err != nil {
		return false, err
	}

	return verifySign(data, publicKey)
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
	var cp = CertDownload{}
	cp.AliPayCertSN = certSN
	rsp, err := this.CertDownload(cp)
	if err != nil {
		return nil, err
	}
	certBytes, err := base64.StdEncoding.DecodeString(rsp.Content.AliPayCertContent)
	if err != nil {
		return nil, err
	}

	cert, err = crypto4go.ParseCertificate(certBytes)
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

func parseJSONSource(rawData string, nodeName string, nodeIndex int) (content, certSN, sign string) {
	var dataStartIndex = nodeIndex + len(nodeName) + 2
	var signIndex = strings.LastIndex(rawData, "\""+kSignNodeName+"\"")
	var certIndex = strings.LastIndex(rawData, "\""+kCertSNNodeName+"\"")
	var dataEndIndex int

	if signIndex > 0 && certIndex > 0 {
		dataEndIndex = int(math.Min(float64(signIndex), float64(certIndex))) - 1
	} else if certIndex > 0 {
		dataEndIndex = certIndex - 1
	} else if signIndex > 0 {
		dataEndIndex = signIndex - 1
	} else {
		dataEndIndex = len(rawData) - 1
	}

	var indexLen = dataEndIndex - dataStartIndex
	if indexLen < 0 {
		return "", "", ""
	}
	content = rawData[dataStartIndex:dataEndIndex]

	if certIndex > 0 {
		var certStartIndex = certIndex + len(kCertSNNodeName) + 4
		certSN = rawData[certStartIndex:]
		var certEndIndex = strings.Index(certSN, "\"")
		certSN = certSN[:certEndIndex]
	}

	if signIndex > 0 {
		var signStartIndex = signIndex + len(kSignNodeName) + 4
		sign = rawData[signStartIndex:]
		var signEndIndex = strings.LastIndex(sign, "\"")
		sign = sign[:signEndIndex]
	}

	return content, certSN, sign
}

func signWithPKCS1v15(param url.Values, privateKey *rsa.PrivateKey, hash crypto.Hash) (s string, err error) {
	if param == nil {
		param = make(url.Values, 0)
	}

	var pList = make([]string, 0, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList)
	var src = strings.Join(pList, "&")
	sig, err := crypto4go.RSASignWithKey([]byte(src), privateKey, hash)
	if err != nil {
		return "", err
	}
	s = base64.StdEncoding.EncodeToString(sig)
	return s, nil
}

func verifySign(data url.Values, key *rsa.PublicKey) (ok bool, err error) {
	sign := data.Get(kSignNodeName)

	var keys = make([]string, 0, 0)
	for k := range data {
		if k == kSignNodeName || k == kSignTypeNodeName || k == kCertSNNodeName {
			continue
		}
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var buf strings.Builder

	for _, k := range keys {
		vs := data[k]
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	s := buf.String()
	return verifyData([]byte(s), sign, key)
}

func verifyData(data []byte, sign string, key *rsa.PublicKey) (ok bool, err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	if err = crypto4go.RSAVerifyWithKey(data, signBytes, key, crypto.SHA256); err != nil {
		return false, err
	}
	return true, nil
}

func getCertSN(cert *x509.Certificate) string {
	var value = md5.Sum([]byte(cert.Issuer.String() + cert.SerialNumber.String()))
	return hex.EncodeToString(value[:])
}
