package alipay_test

import (
	"context"
	"testing"

	"github.com/smartwalle/alipay/v3"
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
	p.Code = "647f16afe0b44c49a8eb1cb3c02aXX31"
	rsp, err := client.SystemOauthToken(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_UserInfoShare(t *testing.T) {
	t.Log("========== UserInfoShare ==========")
	var p = alipay.UserInfoShare{}
	p.AuthToken = "authusrB133e40c363934488a9c3e25e17fd9X31"
	rsp, err := client.UserInfoShare(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_AppToAppAuth(t *testing.T) {
	t.Log("========== AppToAppAuth ==========")
	var result, err = client.AppToAppAuth("http://127.0.0.1", "")
	t.Log(result, err)
}

func TestClient_OpenAuthTokenApp(t *testing.T) {
	t.Log("========== OpenAuthTokenApp ==========")
	var p = alipay.OpenAuthTokenApp{}
	p.GrantType = "authorization_code"
	p.Code = "5a14fd7482254120a351109daedbdX31"
	rsp, err := client.OpenAuthTokenApp(context.Background(), p)
	if err != nil {
		t.Fatal(err)
	}

	if rsp.IsFailure() {
		t.Fatal(rsp.Msg, rsp.SubMsg)
	}
	t.Logf("%v", rsp)
}

func TestClient_AccountAuth(t *testing.T) {
	t.Log("========== AccountAuth ==========")
	var p = alipay.AccountAuth{}
	p.Pid = "2088123456789012"
	p.TargetId = "kkkkk091125"
	p.AuthType = "AUTHACCOUNT"
	result, err := client.AccountAuth(p)
	if err != nil {
		t.Fatal(err)
	}
	if result == "" {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestClient_OpenAuthAppAuthInviteCreate(t *testing.T) {
	t.Log("========== OpenAuthAppAuthInviteCreate ==========")
	var p = alipay.OpenAuthAppAuthInviteCreate{}
	result, err := client.OpenAuthAppAuthInviteCreate(p)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
