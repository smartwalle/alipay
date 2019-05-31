package alipay

import (
	"crypto"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/smartwalle/alipay/encoding"
)

var (
	kSignNotFound = errors.New("alipay: sign content not found")
)

type Client struct {
	appId              string
	apiDomain          string
	notifyVerifyDomain string
	appPrivateKey      *rsa.PrivateKey // 应用私钥
	aliPublicKey       *rsa.PublicKey  // 支付宝公钥
	Client             *http.Client
	SignType           string
}

// New 初始化支付宝客户端
// appId - 支付宝应用 id
// aliPublicKey - 支付宝公钥，创建支付宝应用之后，从支付宝后台获取
// privateKey - 应用私钥，开发者自己生成
// isProduction - 是否为生产环境，传 false 的时候为沙箱环境，用于开发测试，正式上线的时候需要改为 true
func New(appId, aliPublicKey, privateKey string, isProduction bool) (client *Client, err error) {
	pri, err := encoding.ParsePKCS1PrivateKey(encoding.FormatPrivateKey(privateKey))
	if err != nil {
		return nil, err
	}

	var pub *rsa.PublicKey
	if len(aliPublicKey) > 0 {
		pub, err = encoding.ParsePKCS1PublicKey(encoding.FormatPublicKey(aliPublicKey))
		if err != nil {
			return nil, err
		}
	}

	client = &Client{}
	client.appId = appId
	client.appPrivateKey = pri
	client.aliPublicKey = pub

	client.Client = http.DefaultClient
	if isProduction {
		client.apiDomain = kProductionURL
		client.notifyVerifyDomain = kProductionMAPIURL
	} else {
		client.apiDomain = kSandboxURL
		client.notifyVerifyDomain = kSandboxURL
	}
	client.SignType = K_SIGN_TYPE_RSA2
	return client, nil
}

func (this *Client) URLValues(param Param) (value url.Values, err error) {
	var p = url.Values{}
	p.Add("app_id", this.appId)
	p.Add("method", param.APIName())
	p.Add("format", kFormat)
	p.Add("charset", kCharset)
	p.Add("sign_type", this.SignType)
	p.Add("timestamp", time.Now().Format(kTimeFormat))
	p.Add("version", kVersion)

	if len(param.ExtJSONParamName()) > 0 {
		p.Add(param.ExtJSONParamName(), param.ExtJSONParamValue())
	}

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			p.Add(key, value)
		}
	}

	var hash crypto.Hash
	if this.SignType == K_SIGN_TYPE_RSA {
		hash = crypto.SHA1
	} else {
		hash = crypto.SHA256
	}
	sign, err := signWithPKCS1v15(p, this.appPrivateKey, hash)
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

	if this.aliPublicKey != nil {
		var dataStr = string(data)

		var rootNodeName = strings.Replace(param.APIName(), ".", "_", -1) + kResponseSuffix

		var rootIndex = strings.LastIndex(dataStr, rootNodeName)
		var errorIndex = strings.LastIndex(dataStr, kErrorResponse)

		var content string
		var sign string

		if rootIndex > 0 {
			content, sign = parserJSONSource(dataStr, rootNodeName, rootIndex)
		} else if errorIndex > 0 {
			content, sign = parserJSONSource(dataStr, kErrorResponse, errorIndex)
		} else {
			return kSignNotFound
		}

		if sign == "" {
			return kSignNotFound
		}

		if ok, err := verifyData([]byte(content), this.SignType, sign, this.aliPublicKey); ok == false {
			return err
		}
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
	return verifySign(data, this.aliPublicKey)
}

func parserJSONSource(rawData string, nodeName string, nodeIndex int) (content string, sign string) {
	var dataStartIndex = nodeIndex + len(nodeName) + 2
	var signIndex = strings.LastIndex(rawData, "\""+kSignNodeName+"\"")
	var dataEndIndex = signIndex - 1

	var indexLen = dataEndIndex - dataStartIndex
	if indexLen < 0 {
		return "", ""
	}
	content = rawData[dataStartIndex:dataEndIndex]

	var signStartIndex = signIndex + len(kSignNodeName) + 4
	sign = rawData[signStartIndex:]
	var signEndIndex = strings.LastIndex(sign, "\"}")
	sign = sign[:signEndIndex]

	return content, sign
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
	sig, err := encoding.SignPKCS1v15WithKey([]byte(src), privateKey, hash)
	if err != nil {
		return "", err
	}
	s = base64.StdEncoding.EncodeToString(sig)
	return s, nil
}

func VerifySign(data url.Values, key []byte) (ok bool, err error) {
	pub, err := encoding.ParsePKCS1PublicKey(encoding.FormatPublicKey(string(key)))
	if err != nil {
		return false, err
	}
	return verifySign(data, pub)
}

func verifySign(data url.Values, key *rsa.PublicKey) (ok bool, err error) {
	sign := data.Get("sign")
	signType := data.Get("sign_type")

	var keys = make([]string, 0, 0)
	for key, value := range data {
		if key == "sign" || key == "sign_type" {
			continue
		}
		if len(value) > 0 {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	var pList = make([]string, 0, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(data.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	var s = strings.Join(pList, "&")

	return verifyData([]byte(s), signType, sign, key)
}

func verifyData(data []byte, signType, sign string, key *rsa.PublicKey) (ok bool, err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	if signType == K_SIGN_TYPE_RSA {
		err = encoding.VerifyPKCS1v15WithKey(data, signBytes, key, crypto.SHA1)
	} else {
		err = encoding.VerifyPKCS1v15WithKey(data, signBytes, key, crypto.SHA256)
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
