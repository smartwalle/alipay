package alipay

import (
	"crypto"
	"encoding/base64"
	"github.com/smartwalle/going/encoding"
	"net/url"
	"strings"
	"sort"
	"time"
	"net/http"
	"io/ioutil"
	"io"
)

const (
	K_ALI_PAY_SANDBOX_API_URL    = "https://openapi.alipaydev.com/gateway.do"
	K_ALI_PAY_PRODUCTION_API_URL = "https://openapi.alipay.com/gateway.do"
)

type AliPay struct {
	appId string
	apiDomain string
	publickKey []byte
	privateKey []byte
}

func New(appId string, publicKey, privateKey []byte, isProduction bool) (client *AliPay) {
	client = &AliPay{}
	client.appId = appId
	client.privateKey = privateKey
	client.publickKey = publicKey
	if isProduction {
		client.apiDomain = K_ALI_PAY_PRODUCTION_API_URL
	} else {
		client.apiDomain = K_ALI_PAY_SANDBOX_API_URL
	}
	return client
}

func (this *AliPay)URLValues(param AliPayParam) url.Values {
	var p = url.Values{}
	p.Add("app_id", this.appId)
	p.Add("method", param.APIName())
	p.Add("format", "JSON")
	p.Add("charset", "utf-8")
	p.Add("sign_type", "RSA2")
	p.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	p.Add("version", "1.0")

	if len(param.ExtJSONParamName()) > 0 {
		p.Add(param.ExtJSONParamName(), param.ExtJSONParamValue())
	}

	var ps = param.Params()
	if ps != nil {
		for key, value := range ps {
			p.Add(key, value)
		}
	}

	var keys = make([]string, 0, 0)
	for key, _ := range p {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	p.Add("sign", sign(keys, p, this.privateKey))

	return p
	//return http.JSONRequest("POST", K_ALI_PAY_SANDBOX_API_URL, p)
}

func (this *AliPay) CreateWapPayment(param AliPayTradeWapPay) (html string, err error) {
	var body io.Reader
	body = strings.NewReader(this.URLValues(param).Encode())

	var request *http.Request
	request, err = http.NewRequest("POST", this.apiDomain, body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return "", err
	}

	rep, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer rep.Body.Close()

	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return "", err
	}
	html = string(data)
	return html, err
}

//var request = http.NewRequest("POST", this.)


func sign(keys []string, param url.Values, privateKey []byte) (s string) {
	var pList = make([]string, 0, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	var src = strings.Join(pList, "&")
	var sig, err = encoding.SignPKCS1v15([]byte(src), privateKey, crypto.SHA256)
	if err != nil {
		return ""
	}
	s = base64.StdEncoding.EncodeToString(sig)
	return s
}
