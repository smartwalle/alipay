package alipay

import (
	"context"
	"crypto"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/smartwalle/ncrypto"
	"github.com/smartwalle/ngx"
	"github.com/smartwalle/nsign"
)

var (
	ErrBadResponse          = errors.New("alipay: bad response")
	ErrSignNotFound         = errors.New("alipay: sign content not found")
	ErrAliPublicKeyNotFound = errors.New("alipay: alipay public key not found")
)

const (
	kAliPayPublicKey = "alipay-public-key"
	kEmptyBizContent = "{}"
)

type Client struct {
	mu         sync.Mutex
	production bool

	host             string
	notifyVerifyHost string

	Client              *http.Client
	location            *time.Location
	receivedDataHandler ReceivedDataHandler

	// 内容加密
	needEncrypt    bool
	encryptIV      []byte
	encryptType    string
	encryptKey     []byte
	encryptPadding ncrypto.Padding

	// 应用相关信息
	appId         string
	appCertSN     string
	aliRootCertSN string
	aliCertSN     string

	// 签名和验签
	encoder   nsign.Encoder
	signer    Signer
	verifiers map[string]Verifier
}

type Signer interface {
	SignValues(values url.Values, opts ...nsign.SignOption) ([]byte, error)

	SignBytes(data []byte, opts ...nsign.SignOption) ([]byte, error)
}

type Verifier interface {
	VerifyValues(values url.Values, signature []byte, opts ...nsign.SignOption) error

	VerifyBytes(data []byte, signature []byte, opts ...nsign.SignOption) error
}

type OptionFunc func(c *Client)

func WithTimeLocation(location *time.Location) OptionFunc {
	return func(c *Client) {
		c.location = location
	}
}

func WithHTTPClient(client *http.Client) OptionFunc {
	return func(c *Client) {
		if client != nil {
			c.Client = client
		}
	}
}

// WithVerifier 创建一个自定义的验签器
// 适用于支付宝公钥由第三方托管，无法直接获取的场景
func WithVerifier(aliCertSN string, verifier Verifier) OptionFunc {
	return func(c *Client) {
		if aliCertSN != "" && verifier != nil {
			c.aliCertSN = aliCertSN
			c.verifiers[aliCertSN] = verifier
		}
	}
}

func WithSandboxGateway(gateway string) OptionFunc {
	return func(c *Client) {
		if gateway == "" {
			gateway = kNewSandboxGateway
		}
		if !c.production {
			c.host = gateway
		}
	}
}

func WithProductionGateway(gateway string) OptionFunc {
	return func(c *Client) {
		if gateway == "" {
			gateway = kProductionGateway
		}
		if c.production {
			c.host = gateway
		}
	}
}

func WithNewSandboxGateway() OptionFunc {
	return WithSandboxGateway(kNewSandboxGateway)
}

func WithPastSandboxGateway() OptionFunc {
	return WithSandboxGateway(kPastSandboxGateway)
}

// New 初始化支付宝客户端
//
//	appId - 支付宝应用 id
//	privateKey - 应用私钥，开发者自己生成
//	production - 是否为生产环境，传 false 的时候为沙箱环境，用于开发测试，正式上线的时候需要改为 true
func New(appId, privateKey string, production bool, opts ...OptionFunc) (client *Client, err error) {
	// 解析应用私钥
	priKey, err := ncrypto.DecodePrivateKey([]byte(privateKey)).PKCS1().RSAPrivateKey()
	if err != nil {
		priKey, err = ncrypto.DecodePrivateKey([]byte(privateKey)).PKCS8().RSAPrivateKey()
		if err != nil {
			return nil, err
		}
	}
	client = &Client{}
	client.production = production
	client.appId = appId

	if client.production {
		client.host = kProductionGateway
		client.notifyVerifyHost = kProductionMAPIGateway
	} else {
		client.host = kNewSandboxGateway
		client.notifyVerifyHost = kNewSandboxGateway
	}
	client.Client = http.DefaultClient
	client.location = time.Local

	client.encoder = &Encoder{}
	client.signer = nsign.New(nsign.WithMethod(nsign.NewRSAMethod(crypto.SHA256, priKey, nil)), nsign.WithEncoder(client.encoder))
	client.verifiers = make(map[string]Verifier)

	for _, opt := range opts {
		if opt != nil {
			opt(client)
		}
	}

	return client, nil
}

// NewWithCustomSigner 创建一个使用自定义签名器的支付宝客户端
// 适用于私钥由第三方托管，无法直接获取的场景
// 参数:
//   - appId: 支付宝分配的应用ID
//   - isProduction: true 为生产环境，false 为沙箱环境
//   - signer: 自定义签名器，不能为 nil
//   - opts: 可选配置函数
func NewWithCustomSigner(appId string, isProduction bool, signer nsign.Signer, opts ...OptionFunc) (nClient *Client, err error) {
	if signer == nil {
		return nil, errors.New("signer cannot be nil")
	}

	nClient = &Client{}
	nClient.isProduction = isProduction
	nClient.appId = appId

	if nClient.isProduction {
		nClient.host = kProductionGateway
		nClient.notifyVerifyHost = kProductionMAPIGateway
	} else {
		nClient.host = kNewSandboxGateway
		nClient.notifyVerifyHost = kNewSandboxGateway
	}
	nClient.Client = http.DefaultClient
	nClient.location = time.Local

	nClient.encoder = &Encoder{}
	nClient.signer = signer
	nClient.verifiers = make(map[string]Verifier)

	for _, opt := range opts {
		if opt != nil {
			opt(nClient)
		}
	}

	return nClient, nil
}

func (c *Client) Production() bool {
	return c.production
}

// SetEncryptKey 接口内容加密密钥 https://opendocs.alipay.com/common/02mse3
func (c *Client) SetEncryptKey(key string) error {
	if key == "" {
		c.needEncrypt = false
		return nil
	}

	var data, err = base64.StdEncoding.DecodeString(key)
	if err != nil {
		return err
	}
	c.needEncrypt = true
	c.encryptIV = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	c.encryptType = "AES"
	c.encryptKey = data
	c.encryptPadding = ncrypto.PKCS7Padding{}
	return nil
}

func (c *Client) loadVerifier(sn string, pub *rsa.PublicKey) Verifier {
	c.aliCertSN = sn
	var verifier = nsign.New(nsign.WithMethod(nsign.NewRSAMethod(crypto.SHA256, nil, pub)), nsign.WithEncoder(c.encoder))
	c.verifiers[c.aliCertSN] = verifier
	return verifier
}

// LoadAliPayPublicKey 加载支付宝公钥
func (c *Client) LoadAliPayPublicKey(s string) error {
	if len(s) < 0 {
		return ErrAliPublicKeyNotFound
	}

	var pub, err = ncrypto.DecodePublicKey([]byte(s)).PKIX().RSAPublicKey()
	if err != nil {
		return err
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.loadVerifier(kAliPayPublicKey, pub)
	return nil
}

// LoadAppPublicCert 加载应用公钥证书
//
// Deprecated: use LoadAppCertPublicKey instead.
func (c *Client) LoadAppPublicCert(s string) error {
	return c.LoadAppCertPublicKey(s)
}

// LoadAppPublicCertFromFile 加载应用公钥证书
//
// Deprecated: use LoadAppCertPublicKeyFromFile instead.
func (c *Client) LoadAppPublicCertFromFile(filename string) error {
	return c.LoadAppCertPublicKeyFromFile(filename)
}

func (c *Client) loadAppCertPublicKey(b []byte) error {
	var cert, err = ncrypto.DecodeCertificate(b)
	if err != nil {
		return err
	}
	c.appCertSN = getCertSN(cert)
	return nil
}

// LoadAppCertPublicKey 加载应用公钥证书
func (c *Client) LoadAppCertPublicKey(s string) error {
	return c.loadAppCertPublicKey([]byte(s))
}

// LoadAppCertPublicKeyFromFile 从文件加载应用公钥证书
func (c *Client) LoadAppCertPublicKeyFromFile(filename string) error {
	var certBytes, err = os.ReadFile(filename)
	if err != nil {
		return err
	}
	return c.loadAppCertPublicKey(certBytes)
}

// LoadAliPayPublicCert 加载支付宝公钥证书
//
// Deprecated: use LoadAlipayCertPublicKey instead.
func (c *Client) LoadAliPayPublicCert(s string) error {
	return c.LoadAlipayCertPublicKey(s)
}

// LoadAliPayPublicCertFromFile 加载支付宝公钥证书
//
// Deprecated: use LoadAlipayCertPublicKeyFromFile instead.
func (c *Client) LoadAliPayPublicCertFromFile(filename string) error {
	return c.LoadAlipayCertPublicKeyFromFile(filename)
}

// loadAlipayCertPublicKey 加载支付宝公钥证书
func (c *Client) loadAlipayCertPublicKey(b []byte) error {
	var cert, err = ncrypto.DecodeCertificate(b)
	if err != nil {
		return err
	}
	var pub, ok = cert.PublicKey.(*rsa.PublicKey)
	if ok == false {
		return nil
	}
	c.loadVerifier(getCertSN(cert), pub)
	return nil
}

// LoadAlipayCertPublicKey 支付宝公钥证书
func (c *Client) LoadAlipayCertPublicKey(s string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.loadAlipayCertPublicKey([]byte(s))
}

// LoadAlipayCertPublicKeyFromFile 从文件加载支付宝公钥证书
func (c *Client) LoadAlipayCertPublicKeyFromFile(filename string) error {
	var certBytes, err = os.ReadFile(filename)
	if err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.loadAlipayCertPublicKey(certBytes)
}

// LoadAliPayRootCert 加载支付宝根证书
func (c *Client) LoadAliPayRootCert(s string) error {
	var certStrList = strings.Split(s, kCertificateEnd)
	var certSNList = make([]string, 0, len(certStrList))
	for _, certStr := range certStrList {
		certStr = strings.Replace(certStr, kCertificateBegin, "", 1)
		var cert, _ = ncrypto.DecodeCertificate([]byte(certStr))
		if cert != nil && (cert.SignatureAlgorithm == x509.SHA256WithRSA || cert.SignatureAlgorithm == x509.SHA1WithRSA) {
			certSNList = append(certSNList, getCertSN(cert))
		}
	}
	c.aliRootCertSN = strings.Join(certSNList, "_")
	return nil
}

// LoadAliPayRootCertFromFile 加载支付宝根证书
func (c *Client) LoadAliPayRootCertFromFile(filename string) error {
	var certBytes, err = os.ReadFile(filename)
	if err != nil {
		return err
	}
	return c.LoadAliPayRootCert(string(certBytes))
}

func (c *Client) URLValues(param Param) (url.Values, error) {
	var values = url.Values{}
	values.Add(kFieldAppId, c.appId)
	values.Add(kFieldMethod, param.APIName())
	values.Add(kFieldFormat, kFormat)
	values.Add(kFieldCharset, kCharset)
	values.Add(kFieldSignType, kSignTypeRSA2)
	values.Add(kFieldTimestamp, time.Now().In(c.location).Format(kTimeFormat))
	values.Add(kFieldVersion, kVersion)
	if c.appCertSN != "" {
		values.Add(kFieldAppCertSN, c.appCertSN)
	}
	if c.aliRootCertSN != "" {
		values.Add(kFieldAliPayRootCertSN, c.aliRootCertSN)
	}

	bizBytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var bizContent = string(bizBytes)
	if bizContent != kEmptyBizContent {
		if c.needEncrypt && param.NeedEncrypt() {
			bizBytes, err = ncrypto.AESCBCEncrypt(bizBytes, c.encryptKey, c.encryptIV, c.encryptPadding)
			if err != nil {
				return nil, err
			}
			bizContent = base64.StdEncoding.EncodeToString(bizBytes)
			values.Add(kFieldEncryptType, c.encryptType)
		}
		values.Add(kFieldBizContent, bizContent)
	}

	var params = param.Params()
	for key, value := range params {
		if value == "" {
			continue
		}
		values.Add(key, value)
	}

	signature, err := c.sign(values)
	if err != nil {
		return nil, err
	}

	values.Add(kFieldSign, signature)
	return values, nil
}

func (c *Client) doRequest(ctx context.Context, method string, param Param, dest interface{}) error {
	var req = ngx.NewRequest(method, c.host, ngx.WithClient(c.Client))
	req.ContentType = kContentType
	if param != nil {
		values, err := c.URLValues(param)
		if err != nil {
			return err
		}
		req.Form = values
		req.FileForm = param.FileParams()
	}

	resp, err := req.Do(ctx)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apiName = param.APIName()
	var bizFieldName = strings.Replace(apiName, ".", "_", -1) + kResponseSuffix

	return c.decode(ctx, bodyBytes, bizFieldName, param.NeedVerify(), dest)
}

func (c *Client) decode(ctx context.Context, data []byte, bizFieldName string, needVerifySign bool, dest interface{}) (err error) {
	var raw = make(map[string]json.RawMessage)
	if err = json.Unmarshal(data, &raw); err != nil {
		return err
	}

	var signBytes = raw[kFieldSign]
	var certBytes = raw[kFieldAliPayCertSN]
	var bizBytes = raw[bizFieldName]
	var errBytes = raw[kErrorResponse]

	if len(certBytes) > 1 {
		certBytes = certBytes[1 : len(certBytes)-1]
	}
	if len(signBytes) > 1 {
		signBytes = signBytes[1 : len(signBytes)-1]
	}

	if len(bizBytes) == 0 {
		if len(errBytes) > 0 {
			var rErr *Error
			if err = json.Unmarshal(errBytes, &rErr); err != nil {
				return err
			}
			return rErr
		}
		return ErrBadResponse
	}

	// 数据解密
	var plaintext []byte
	if plaintext, err = c.decrypt(bizBytes); err != nil {
		return err
	}

	// 验证签名
	if needVerifySign {
		if c.receivedDataHandler != nil {
			c.receivedDataHandler(ctx, bizFieldName, plaintext)
		}

		if len(signBytes) == 0 {
			// 没有签名数据，返回的内容一般为错误信息
			var rErr *Error
			if err = json.Unmarshal(plaintext, &rErr); err != nil {
				return err
			}
			return rErr
		}

		// 验证签名
		if err = c.verify(ctx, string(certBytes), bizBytes, signBytes); err != nil {
			return err
		}
	}

	if err = json.Unmarshal(plaintext, dest); err != nil {
		return err
	}
	return nil
}

func (c *Client) decrypt(data []byte) ([]byte, error) {
	var plaintext = data
	if len(data) > 1 && data[0] == '"' {
		var ciphertext, err = base64decode(data[1 : len(data)-1])
		if err != nil {
			return nil, err
		}
		plaintext, err = ncrypto.AESCBCDecrypt(ciphertext, c.encryptKey, c.encryptIV, c.encryptPadding)
		if err != nil {
			return nil, err
		}
	}
	return plaintext, nil
}

func (c *Client) VerifySign(ctx context.Context, values url.Values) (err error) {
	var verifier Verifier
	if verifier, err = c.getVerifier(ctx, values.Get(kFieldAliPayCertSN)); err != nil {
		return err
	}

	var signBytes []byte
	if signBytes, err = base64.StdEncoding.DecodeString(values.Get(kFieldSign)); err != nil {
		return err
	}

	return verifier.VerifyValues(values, signBytes, nsign.WithIgnore(kFieldSign, kFieldSignType, kFieldAliPayCertSN))
}

func (c *Client) getVerifier(ctx context.Context, certSN string) (verifier Verifier, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if certSN == "" {
		certSN = c.aliCertSN
	}

	verifier = c.verifiers[certSN]

	if verifier == nil {
		if !c.production {
			return nil, ErrAliPublicKeyNotFound
		}

		cert, err := c.downloadAliPayCert(ctx, certSN)
		if err != nil {
			return nil, err
		}

		pub, ok := cert.PublicKey.(*rsa.PublicKey)
		if !ok {
			return nil, ErrAliPublicKeyNotFound
		}
		verifier = c.loadVerifier(getCertSN(cert), pub)
	}
	return verifier, nil
}

func (c *Client) CertDownload(ctx context.Context, param CertDownload) (result *CertDownloadRsp, err error) {
	err = c.doRequest(ctx, http.MethodPost, param, &result)
	return result, err
}

func (c *Client) downloadAliPayCert(ctx context.Context, certSN string) (cert *x509.Certificate, err error) {
	var param = CertDownload{}
	param.AliPayCertSN = certSN
	resp, err := c.CertDownload(ctx, param)
	if err != nil {
		return nil, err
	}
	certBytes, err := base64.StdEncoding.DecodeString(resp.AliPayCertContent)
	if err != nil {
		return nil, err
	}

	cert, err = ncrypto.DecodeCertificate(certBytes)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func (c *Client) sign(values url.Values) (signature string, err error) {
	signBytes, err := c.signer.SignValues(values)
	if err != nil {
		return "", err
	}
	signature = base64.StdEncoding.EncodeToString(signBytes)
	return signature, nil
}

func (c *Client) verify(ctx context.Context, certSN string, data, signature []byte) (err error) {
	var verifier Verifier
	if verifier, err = c.getVerifier(ctx, certSN); err != nil {
		return err
	}

	if signature, err = base64decode(signature); err != nil {
		return err
	}

	if err = verifier.VerifyBytes(data, signature); err != nil {
		return err
	}
	return nil
}

func (c *Client) Request(ctx context.Context, param Param, result interface{}) (err error) {
	return c.doRequest(ctx, http.MethodPost, param, result)
}

func (c *Client) BuildURL(param Param) (*url.URL, error) {
	value, err := c.URLValues(param)
	if err != nil {
		return nil, err
	}
	return url.Parse(c.host + "?" + value.Encode())
}

func (c *Client) EncodeParam(param Param) (string, error) {
	value, err := c.URLValues(param)
	if err != nil {
		return "", err
	}
	return value.Encode(), nil
}

func (c *Client) OnReceivedData(handler ReceivedDataHandler) {
	c.receivedDataHandler = handler
}

func (c *Client) SignValues(values url.Values, opts ...nsign.SignOption) ([]byte, error) {
	return c.signer.SignValues(values, opts...)
}

func (c *Client) SignBytes(data []byte, opts ...nsign.SignOption) ([]byte, error) {
	return c.signer.SignBytes(data, opts...)
}

func base64decode(data []byte) ([]byte, error) {
	var buffer = make([]byte, base64.StdEncoding.DecodedLen(len(data)))
	n, err := base64.StdEncoding.Decode(buffer, data)
	return buffer[:n], err
}

func getCertSN(cert *x509.Certificate) string {
	var value = md5.Sum([]byte(cert.Issuer.String() + cert.SerialNumber.String()))
	return hex.EncodeToString(value[:])
}
