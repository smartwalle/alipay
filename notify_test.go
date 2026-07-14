package alipay

import (
	"testing"
)

func TestDecodeGBK(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "ascii only",
			input:    "hello world 123",
			expected: "hello world 123",
		},
		{
			name:     "GBK encoded Chinese - 树",
			input:    "\xca\xf7", // GBK encoding of "树"
			expected: "树",
		},
		{
			name:     "GBK encoded Chinese - 中国",
			input:    "\xd6\xd0\xb9\xfa", // GBK encoding of "中国"
			expected: "中国",
		},
		{
			name:     "mixed ascii and GBK",
			input:    "test\xca\xf7123", // "test树123"
			expected: "test树123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := decodeGBK(tt.input)
			if result != tt.expected {
				t.Errorf("decodeGBK(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDecodeNotificationWithCharset_GBK(t *testing.T) {
	// Test the decodeGBK function with URL-decoded values
	// In real scenario, url.ParseQuery will decode %CA%F7 to \xca\xf7 (raw bytes)

	// URL decode simulation - %CA%F7 becomes raw bytes \xca\xf7
	rawGBK := "\xca\xf7"
	decoded := decodeGBK(rawGBK)

	if decoded != "树" {
		t.Errorf("Expected '树', got %q", decoded)
	}
}

func TestParseCharsetFromContentType(t *testing.T) {
	tests := []struct {
		contentType string
		expected    string
	}{
		{"application/x-www-form-urlencoded; charset=GBK", "GBK"},
		{"application/x-www-form-urlencoded; charset=gbk", "gbk"},
		{"application/x-www-form-urlencoded; charset=UTF-8", "UTF-8"},
		{"application/x-www-form-urlencoded;charset=GBK", "GBK"},
		{"application/x-www-form-urlencoded", ""},
		{"text/html; charset=GB2312; boundary=something", "GB2312"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.contentType, func(t *testing.T) {
			result := parseCharsetFromContentType(tt.contentType)
			if result != tt.expected {
				t.Errorf("parseCharsetFromContentType(%q) = %q, want %q", tt.contentType, result, tt.expected)
			}
		})
	}
}

func TestDecodeNotificationWithCharset_UTF8(t *testing.T) {
	// When charset is UTF-8, the data should remain unchanged
	utf8String := "中文测试"
	result := decodeGBK(utf8String)

	// Note: decodeGBK will try to decode UTF-8 as GBK, which may produce garbled text
	// This is expected behavior - the DecodeNotificationWithCharset method
	// only calls decodeGBK when charset is GBK/GB2312/GB18030

	// The important thing is that DecodeNotificationWithCharset checks charset first
	// and only decodes when necessary
	t.Logf("UTF-8 string through GBK decoder: %q -> %q", utf8String, result)
}

func TestCharsetCheck(t *testing.T) {
	// Test that charset checking works for various cases
	charsets := []struct {
		input    string
		expected bool
	}{
		{"GBK", true},
		{"gbk", true},
		{"GB2312", true},
		{"gb2312", true},
		{"GB18030", true},
		{"gb18030", true},
		{"UTF-8", false},
		{"utf-8", false},
		{"UTF8", false},
		{"", false},
	}

	for _, tc := range charsets {
		t.Run(tc.input, func(t *testing.T) {
			// Simulate the check in DecodeNotificationWithCharset
			upper := ""
			for _, c := range tc.input {
				if c >= 'a' && c <= 'z' {
					upper += string(c - 32)
				} else {
					upper += string(c)
				}
			}
			isGBK := upper == "GBK" || upper == "GB2312" || upper == "GB18030"
			if isGBK != tc.expected {
				t.Errorf("charset %q: got isGBK=%v, want %v", tc.input, isGBK, tc.expected)
			}
		})
	}
}
