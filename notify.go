package alipay

import (
	"io/ioutil"
	"net/url"

	"github.com/smartwalle/going/request"
)

func (this *AliPay) NotifyVerify(notifyId string) bool {
	var param = url.Values{}
	param.Add("service", "notify_verify")
	param.Add("partner", this.partnerId)
	param.Add("notify_id", notifyId)
	req, err := request.NewRequest("GET", this.apiDomain, param)
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
