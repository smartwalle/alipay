package alipay_test

import (
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
	p.Code = "0bb679f49d444197bdb1e7e08cafWX31"
	rsp, err := client.SystemOauthToken(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.UserId, rsp.Content.AccessToken)
}

func TestClient_UserInfoShare(t *testing.T) {
	t.Log("========== UserInfoShare ==========")
	var p = alipay.UserInfoShare{}
	p.AuthToken = "authusrB82e8d260db754bbf8e02ed7beb195E31"
	rsp, err := client.UserInfoShare(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.UserId, rsp.Content.Avatar)
}
