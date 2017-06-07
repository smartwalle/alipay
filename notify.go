package alipay

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func NewRequest(method, url string, params url.Values) (*http.Request, error) {
	var m = strings.ToUpper(method)
	var body io.Reader
	if m == "GET" || m == "HEAD" {
		if len(params) > 0 {
			if strings.Contains(url, "?") {
				url = url + "&" + params.Encode()
			} else {
				url = url + "?" + params.Encode()
			}
		}
	} else {
		body = strings.NewReader(params.Encode())
	}
	return http.NewRequest(m, url, body)
}

func (this *AliPay) NotifyVerify(notifyId string) bool {
	var param = url.Values{}
	param.Add("service", "notify_verify")
	param.Add("partner", this.partnerId)
	param.Add("notify_id", notifyId)
	req, err := NewRequest("GET", this.apiDomain, param)
	if err != nil {
		return false
	}

	rep, err := this.client.Do(req)
	if err != nil {
		return false
	}
	defer rep.Body.Close()

	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return false
	}
	if string(data) == "true" {
		return true
	}
	return false
}
