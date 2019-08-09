package alipay_test

import (
	"fmt"
	"github.com/smartwalle/alipay"
	"testing"
)

func TestClient_PublicAppAuthorize(t *testing.T) {
	t.Log("========== PublicAppAuthorize ==========")
	var result, err = client.PublicAppAuthorize([]string{"auth_user"}, "http://127.0.0.1", "hhh")
	t.Log(result, err)
}

func TestClient_SystemOauthToken(t *testing.T) {
	t.Log("========== SystemOauthToken ==========")
	var p = alipay.SystemOauthToken{}
	p.GrantType = "authorization_code"
	p.Code = "0a59d92a24564661afd00221b0dbOC31"
	rsp, err := client.SystemOauthToken(p)
	fmt.Println(rsp, err)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.UserId, rsp.Content.AccessToken)
}
