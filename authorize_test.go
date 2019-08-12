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
	p.Code = "aac6328ccb5048f5b5896ca5b659OX31"
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
	p.AuthToken = "authusrB235b1ccbd56346d39c24a3280b2acX31"
	rsp, err := client.UserInfoShare(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.UserId)
}

func TestClient_AppToAppAuth(t *testing.T) {
	t.Log("========== AppToAppAuth ==========")
	var result, err = client.AppToAppAuth("http://127.0.0.1")
	t.Log(result, err)
}

func TestClient_OpenAuthTokenApp(t *testing.T) {
	t.Log("========== OpenAuthTokenApp ==========")
	var p = alipay.OpenAuthTokenApp{}
	p.GrantType = "authorization_code"
	p.Code = "5a14fd7482254120a351109daedbdX31"
	rsp, err := client.OpenAuthTokenApp(p)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.Content.Code != alipay.K_SUCCESS_CODE {
		t.Fatal(rsp.Content.Msg, rsp.Content.SubMsg)
	}
	t.Log(rsp.Content.AppAuthToken, rsp.Content.UserId)
}
