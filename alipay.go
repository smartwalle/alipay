package alipay

import (
	"crypto"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/smartwalle/alipay/encoding"
)

type AliPay struct {
	appId              string
	apiDomain          string
	notifyVerifyDomain string
	//partnerId          string
	privateKey      []byte
	AliPayPublicKey []byte
	Client          *http.Client
	SignType        string
}

func New(appId, aliPublicKey, privateKey string, isProduction bool) (client *AliPay) {
	client = &AliPay{}
	client.appId = appId
	//client.partnerId = partnerId
	client.privateKey = encoding.FormatPrivateKey(privateKey)
	client.AliPayPublicKey = encoding.FormatPublicKey(aliPublicKey)
	client.Client = http.DefaultClient
	if isProduction {
		client.apiDomain = kProductionURL
		client.notifyVerifyDomain = kProductionMAPIURL
	} else {
		client.apiDomain = kSandboxURL
		client.notifyVerifyDomain = kSandboxURL
	}
	client.SignType = K_SIGN_TYPE_RSA2
	return client
}

func (this *AliPay) URLValues(param AliPayParam) (value url.Values, err error) {
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
	sign, err := signWithPKCS1v15(p, this.privateKey, hash)
	if err != nil {
		return nil, err
	}
	p.Add("sign", sign)
	return p, nil
}

func (this *AliPay) doRequest(method string, param AliPayParam, results interface{}) (err error) {
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

	if len(this.AliPayPublicKey) > 0 {
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
			return nil
		}

		if sign != "" {
			if ok, err := verifyData([]byte(content), this.SignType, sign, this.AliPayPublicKey); ok == false {
				return err
			}
		}
	}

	err = json.Unmarshal(data, results)
	if err != nil {
		return err
	}

	return err
}

func (this *AliPay) DoRequest(method string, param AliPayParam, results interface{}) (err error) {
	return this.doRequest(method, param, results)
}

func (this *AliPay) VerifySign(data url.Values) (ok bool, err error) {
	return verifySign(data, this.AliPayPublicKey)
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

func signWithPKCS1v15(param url.Values, privateKey []byte, hash crypto.Hash) (s string, err error) {
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
	sig, err := encoding.SignPKCS1v15([]byte(src), privateKey, hash)
	if err != nil {
		return "", err
	}
	s = base64.StdEncoding.EncodeToString(sig)
	return s, nil
}

func VerifySign(data url.Values, key []byte) (ok bool, err error) {
	return verifySign(data, key)
}

func verifySign(data url.Values, key []byte) (ok bool, err error) {
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

func verifyData(data []byte, signType, sign string, key []byte) (ok bool, err error) {
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false, err
	}

	if signType == K_SIGN_TYPE_RSA {
		err = encoding.VerifyPKCS1v15(data, signBytes, key, crypto.SHA1)
	} else {
		err = encoding.VerifyPKCS1v15(data, signBytes, key, crypto.SHA256)
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
