package alipay

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"net/url"
	"testing"

	"github.com/smartwalle/nsign"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type recordingVerifier struct {
	t            *testing.T
	expectedData []byte
	expectedSign []byte
	called       bool
}

func (v *recordingVerifier) VerifyValues(_ url.Values, _ []byte, _ ...nsign.SignOption) error {
	v.t.Fatal("VerifyValues should not be called")
	return nil
}

func (v *recordingVerifier) VerifyBytes(data, signature []byte, _ ...nsign.SignOption) error {
	v.called = true
	if !bytes.Equal(data, v.expectedData) {
		v.t.Errorf("verified data changed before signature verification\ngot:  %x\nwant: %x", data, v.expectedData)
	}
	if !bytes.Equal(signature, v.expectedSign) {
		v.t.Errorf("signature = %x, want %x", signature, v.expectedSign)
	}
	return nil
}

func TestClientDecodeGBKResponseAfterVerifyingOriginalBytes(t *testing.T) {
	biz := encodeGBK(t, `{"code":"20000","msg":"Service Currently Unavailable","sub_code":"aop.unknow-error","sub_msg":"系统繁忙"}`)
	signature := []byte("test-signature")
	envelope := append([]byte(`{"test_response":`), biz...)
	envelope = append(envelope, []byte(`,"sign":"`+base64.StdEncoding.EncodeToString(signature)+`"}`)...)

	verifier := &recordingVerifier{
		t:            t,
		expectedData: biz,
		expectedSign: signature,
	}
	client := &Client{
		aliCertSN: "test-cert",
		verifiers: map[string]Verifier{"test-cert": verifier},
	}
	var received []byte
	client.OnReceivedData(func(_ context.Context, method string, data []byte) {
		if method != "test_response" {
			t.Errorf("received method = %q, want test_response", method)
		}
		received = append([]byte(nil), data...)
	})

	var response BillDownloadURLQueryRsp
	if err := client.decode(context.Background(), envelope, "test_response", true, &response); err != nil {
		t.Fatalf("decode GBK response: %v", err)
	}
	if !verifier.called {
		t.Fatal("response signature was not verified")
	}
	if !bytes.Equal(received, biz) {
		t.Fatalf("received data changed\ngot:  %x\nwant: %x", received, biz)
	}
	if got, want := response.SubMsg, "系统繁忙"; got != want {
		t.Fatalf("SubMsg = %q, want %q", got, want)
	}
}

func TestClientDecodeGBKUnsignedBusinessError(t *testing.T) {
	biz := encodeGBK(t, `{"code":"20000","msg":"Service Currently Unavailable","sub_code":"aop.unknow-error","sub_msg":"系统繁忙"}`)
	envelope := append([]byte(`{"test_response":`), biz...)
	envelope = append(envelope, '}')

	client := &Client{}
	var response BillDownloadURLQueryRsp
	err := client.decode(context.Background(), envelope, "test_response", true, &response)
	var responseErr *Error
	if !errors.As(err, &responseErr) {
		t.Fatalf("decode error = %v, want *Error", err)
	}
	if got, want := responseErr.SubMsg, "系统繁忙"; got != want {
		t.Fatalf("SubMsg = %q, want %q", got, want)
	}
}

func TestClientDecodeGBKTopLevelError(t *testing.T) {
	errorResponse := encodeGBK(t, `{"code":"20000","msg":"Service Currently Unavailable","sub_code":"aop.unknow-error","sub_msg":"系统繁忙"}`)
	envelope := append([]byte(`{"error_response":`), errorResponse...)
	envelope = append(envelope, '}')

	client := &Client{}
	var response BillDownloadURLQueryRsp
	err := client.decode(context.Background(), envelope, "test_response", true, &response)
	var responseErr *Error
	if !errors.As(err, &responseErr) {
		t.Fatalf("decode error = %v, want *Error", err)
	}
	if got, want := responseErr.SubMsg, "系统繁忙"; got != want {
		t.Fatalf("SubMsg = %q, want %q", got, want)
	}
}

func TestUnmarshalResponseJSONLeavesUTF8Unchanged(t *testing.T) {
	var response Error
	if err := unmarshalResponseJSON([]byte(`{"sub_msg":"系统繁忙"}`), &response); err != nil {
		t.Fatalf("unmarshal UTF-8 response: %v", err)
	}
	if got, want := response.SubMsg, "系统繁忙"; got != want {
		t.Fatalf("SubMsg = %q, want %q", got, want)
	}
}

func encodeGBK(t *testing.T, value string) []byte {
	t.Helper()
	encoded, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(value))
	if err != nil {
		t.Fatalf("encode GBK fixture: %v", err)
	}
	return encoded
}
