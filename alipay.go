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
	"github.com/smartwalle/crypto4go"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

var (
	kSignNotFound         = errors.New("alipay: sign content not found")
	kAliPublicKeyNotFound = errors.New("alipay: alipay public key not found")
)

type Client struct {
	isProduction       bool
	appId              string
	apiDomain          string
	notifyVerifyDomain string
	appPrivateKey      *rsa.PrivateKey // 应用私钥
	aliPublicKey       *rsa.PublicKey  // 支付宝公钥
	Client             *http.Client

	appCertSN        string
	rootCertSN       string
	aliPublicCertSN  string
	aliPublicKeyList map[string]*rsa.PublicKey
}

// New 初始化支付宝客户端
// appId - 支付宝应用 id
// aliPublicKey - 支付宝公钥，创建支付宝应用之后，从支付宝后台获取
// privateKey - 应用私钥，开发者自己生成
// isProduction - 是否为生产环境，传 false 的时候为沙箱环境，用于开发测试，正式上线的时候需要改为 true
func New(appId, aliPublicKey, privateKey string, isProduction bool) (client *Client, err error) {
	pri, err := crypto4go.ParsePKCS1PrivateKey(crypto4go.FormatPKCS1PrivateKey(privateKey))
	if err != nil {
		pri, err = crypto4go.ParsePKCS8PrivateKey(crypto4go.FormatPKCS8PrivateKey(privateKey))
		if err != nil {
			return nil, err
		}
	}

	var pub *rsa.PublicKey
	if len(aliPublicKey) > 0 && isProduction == false {
		pub, err = crypto4go.ParsePublicKey(crypto4go.FormatPublicKey(aliPublicKey))
		if err != nil {
			return nil, err
		}
	}

	client = &Client{}
	client.isProduction = isProduction
	client.appId = appId
	client.appPrivateKey = pri
	client.aliPublicKey = pub

	client.Client = http.DefaultClient
	if client.isProduction {
		client.apiDomain = kProductionURL
		client.notifyVerifyDomain = kProductionMAPIURL
	} else {
		client.apiDomain = kSandboxURL
		client.notifyVerifyDomain = kSandboxURL
	}
	client.aliPublicKeyList = make(map[string]*rsa.PublicKey)
	return client, nil
}

func (this *Client) IsProduction() bool {
	return this.isProduction
}

func (this *Client) LoadAppPublicCert(s string) error {
	cert, err := crypto4go.ParseCertificate([]byte(s))
	if err != nil {
		return err
	}
	this.appCertSN = getCertSN(cert)
	return nil
}

func (this *Client) LoadAppPublicCertFromFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return this.LoadAppPublicCert(string(b))
}

func (this *Client) LoadAliPayPublicCert(s string) error {
	cert, err := crypto4go.ParseCertificate([]byte(s))
	if err != nil {
		return err
	}

	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if ok == false {
		return nil
	}

	this.aliPublicCertSN = getCertSN(cert)
	this.aliPublicKeyList[this.aliPublicCertSN] = key

	if this.aliPublicKey == nil {
		this.aliPublicKey = key
	}
	return nil
}

func (this *Client) LoadAliPayPublicCertFromFile(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return this.LoadAliPayPublicCert(string(b))
}

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

	this.rootCertSN = strings.Join(certSNList, "_")
	return nil
}

func (this *Client) LoadAliPayRootCertFromFile(filename string) error {
	b, err := ioutil.ReadFile(filename)

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
	p.Add("timestamp", time.Now().Format(kTimeFormat))
	p.Add("version", kVersion)
	if this.appCertSN != "" {
		p.Add("app_cert_sn", this.appCertSN)
	}
	if this.rootCertSN != "" {
		p.Add("alipay_root_cert_sn", this.rootCertSN)
	}

	bytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	p.Add("biz_content", string(bytes))

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
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

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dataStr = string(data)

	var rootNodeName = strings.Replace(param.APIName(), ".", "_", -1) + kResponseSuffix

	var rootIndex = strings.LastIndex(dataStr, rootNodeName)
	var errorIndex = strings.LastIndex(dataStr, kErrorResponse)

	var content string
	var certSN string
	var sign string

	if rootIndex > 0 {
		content, certSN, sign = parseJSONSource(dataStr, rootNodeName, rootIndex)
		if sign == "" {
			var errRsp *ErrorRsp
			if err = json.Unmarshal([]byte(content), &errRsp); err != nil {
				return err
			}
			if errRsp != nil {
				return errRsp
			}
			return kSignNotFound
		}
	} else if errorIndex > 0 {
		content, certSN, sign = parseJSONSource(dataStr, kErrorResponse, errorIndex)
		if sign == "" {
			var errRsp *ErrorRsp
			if err = json.Unmarshal([]byte(content), &errRsp); err != nil {
				return err
			}
			return errRsp
		}
	} else {
		return kSignNotFound
	}

	var publicKey *rsa.PublicKey

	if this.isProduction {
		publicKey = this.aliPublicKeyList[certSN]
	} else {
		publicKey = this.aliPublicKey
	}

	if publicKey == nil {
		if this.isProduction {
			// TODO https://docs.open.alipay.com/api_9/alipay.open.app.alipaycert.download 下载新的证书
		}
		return kAliPublicKeyNotFound
	}

	if ok, err := verifyData([]byte(content), sign, publicKey); ok == false {
		return err
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	return err
}

func (this *Client) DoRequest(method string, param Param, result interface{}) (err error) {
	return this.doRequest(method, param, result)
}

func (this *Client) VerifySign(data url.Values) (ok bool, err error) {
	var publicKey *rsa.PublicKey

	if this.isProduction {
		publicKey = this.aliPublicKeyList[this.aliPublicCertSN]
	} else {
		publicKey = this.aliPublicKey
	}

	//if publicKey == nil {
	//	if this.isProduction {
	//		// TODO https://docs.open.alipay.com/api_9/alipay.open.app.alipaycert.download 下载新的证书
	//	}
	//	return kAliPublicKeyNotFound
	//}

	return verifySign(data, publicKey)
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
	sign := data.Get("sign")

	var keys = make([]string, 0, 0)
	for key := range data {
		if key == "sign" || key == "sign_type" {
			continue
		}
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var pList = make([]string, 0, 0)
	for _, key := range keys {
		pList = append(pList, key+"="+data.Get(key))
	}
	var s = strings.Join(pList, "&")

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
