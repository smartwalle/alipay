package alipay

import (
	"bytes"
	"context"
	"encoding/base64"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/smartwalle/nsign"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type notificationVerifier struct {
	t               *testing.T
	expectedSubject string
	expectedSign    []byte
	called          bool
}

func (v *notificationVerifier) VerifyValues(values url.Values, signature []byte, _ ...nsign.SignOption) error {
	v.called = true
	if got := values.Get("subject"); got != v.expectedSubject {
		v.t.Errorf("subject changed before signature verification\ngot:  %x\nwant: %x", []byte(got), []byte(v.expectedSubject))
	}
	if !bytes.Equal(signature, v.expectedSign) {
		v.t.Errorf("signature = %x, want %x", signature, v.expectedSign)
	}
	return nil
}

func (v *notificationVerifier) VerifyBytes(_ []byte, _ []byte, _ ...nsign.SignOption) error {
	v.t.Fatal("VerifyBytes should not be called")
	return nil
}

func TestDecodeNotificationDecodesAllGBKValuesAfterVerification(t *testing.T) {
	subject := string(encodeGBK(t, "测试商品"))
	signature := []byte("notification-signature")
	values := url.Values{
		"charset":             {"GBK"},
		"subject":             {subject},
		"body":                {string(encodeGBK(t, "商品描述"))},
		"buyer_logon_id":      {string(encodeGBK(t, "买家账号"))},
		"refund_reason":       {string(encodeGBK(t, "退款原因"))},
		"passback_params":     {string(encodeGBK(t, "回传参数"))},
		"voucher_detail_list": {string(encodeGBK(t, `[{"name":"优惠券"}]`))},
		"sign_type":           {"RSA2"},
		"sign":                {base64.StdEncoding.EncodeToString(signature)},
	}
	verifier := &notificationVerifier{
		t:               t,
		expectedSubject: subject,
		expectedSign:    signature,
	}
	client := notificationTestClient(verifier)

	notification, err := client.DecodeNotification(context.Background(), values)
	if err != nil {
		t.Fatalf("DecodeNotification: %v", err)
	}
	if !verifier.called {
		t.Fatal("notification signature was not verified")
	}
	if values.Get("subject") != subject {
		t.Fatal("DecodeNotification modified the signed input values")
	}
	assertNotificationText(t, notification)
}

func TestDecodeNotificationWithCharsetSupportsGB18030(t *testing.T) {
	subjectBytes, err := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte("𠀀商品"))
	if err != nil {
		t.Fatalf("encode GB18030 fixture: %v", err)
	}
	signature := []byte("notification-signature")
	values := url.Values{
		"subject":   {string(subjectBytes)},
		"sign_type": {"RSA2"},
		"sign":      {base64.StdEncoding.EncodeToString(signature)},
	}
	verifier := &notificationVerifier{
		t:               t,
		expectedSubject: string(subjectBytes),
		expectedSign:    signature,
	}
	client := notificationTestClient(verifier)

	notification, err := client.DecodeNotificationWithCharset(context.Background(), values, "GB18030")
	if err != nil {
		t.Fatalf("DecodeNotificationWithCharset: %v", err)
	}
	if got, want := notification.Subject, "𠀀商品"; got != want {
		t.Fatalf("Subject = %q, want %q", got, want)
	}
}

func TestDecodeNotificationPrefersSignedCharset(t *testing.T) {
	subject := string(encodeGBK(t, "测试商品"))
	signature := []byte("notification-signature")
	values := url.Values{
		"charset":   {"GBK"},
		"subject":   {subject},
		"sign_type": {"RSA2"},
		"sign":      {base64.StdEncoding.EncodeToString(signature)},
	}
	client := notificationTestClient(&notificationVerifier{
		t:               t,
		expectedSubject: subject,
		expectedSign:    signature,
	})

	notification, err := client.DecodeNotificationWithCharset(context.Background(), values, "UTF-8")
	if err != nil {
		t.Fatalf("DecodeNotificationWithCharset: %v", err)
	}
	if got, want := notification.Subject, "测试商品"; got != want {
		t.Fatalf("Subject = %q, want %q", got, want)
	}
}

func TestGetTradeNotificationUsesContentTypeCharsetFallback(t *testing.T) {
	subject := string(encodeGBK(t, "测试商品"))
	signature := []byte("notification-signature")
	values := url.Values{
		"subject":   {subject},
		"sign_type": {"RSA2"},
		"sign":      {base64.StdEncoding.EncodeToString(signature)},
	}
	request := httptest.NewRequest("POST", "/notify", bytes.NewBufferString(values.Encode()))
	request.Header.Set("Content-Type", `application/x-www-form-urlencoded; charset="GBK"`)
	client := notificationTestClient(&notificationVerifier{
		t:               t,
		expectedSubject: subject,
		expectedSign:    signature,
	})

	notification, err := client.GetTradeNotification(request)
	if err != nil {
		t.Fatalf("GetTradeNotification: %v", err)
	}
	if got, want := notification.Subject, "测试商品"; got != want {
		t.Fatalf("Subject = %q, want %q", got, want)
	}
}

func TestDecodeNotificationLeavesUTF8ValuesUnchanged(t *testing.T) {
	signature := []byte("notification-signature")
	values := url.Values{
		"charset":   {"UTF-8"},
		"subject":   {"测试商品"},
		"sign_type": {"RSA2"},
		"sign":      {base64.StdEncoding.EncodeToString(signature)},
	}
	client := notificationTestClient(&notificationVerifier{
		t:               t,
		expectedSubject: "测试商品",
		expectedSign:    signature,
	})

	notification, err := client.DecodeNotification(context.Background(), values)
	if err != nil {
		t.Fatalf("DecodeNotification: %v", err)
	}
	if got, want := notification.Subject, "测试商品"; got != want {
		t.Fatalf("Subject = %q, want %q", got, want)
	}
}

func notificationTestClient(verifier Verifier) *Client {
	return &Client{
		aliCertSN: "test-cert",
		verifiers: map[string]Verifier{"test-cert": verifier},
	}
}

func assertNotificationText(t *testing.T, notification *Notification) {
	t.Helper()
	checks := map[string]struct {
		got  string
		want string
	}{
		"Subject":           {notification.Subject, "测试商品"},
		"Body":              {notification.Body, "商品描述"},
		"BuyerLogonId":      {notification.BuyerLogonId, "买家账号"},
		"RefundReason":      {notification.RefundReason, "退款原因"},
		"PassbackParams":    {notification.PassbackParams, "回传参数"},
		"VoucherDetailList": {notification.VoucherDetailList, `[{"name":"优惠券"}]`},
	}
	for name, check := range checks {
		if check.got != check.want {
			t.Errorf("%s = %q, want %q", name, check.got, check.want)
		}
	}
}
