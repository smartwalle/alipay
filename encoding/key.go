package encoding

import (
	"strings"
	"bytes"
)

func ParsePublicKey(raw string) (result []byte) {
	return parseKey(raw, "-----BEGIN PUBLIC KEY-----", "-----END PUBLIC KEY-----")
}

func ParsePrivateKey(raw string) (result []byte) {
	return parseKey(raw, "-----BEGIN RSA PRIVATE KEY-----", "-----END RSA PRIVATE KEY-----")
}

func parseKey(raw, prefix, suffix string) (result []byte) {
	if strings.HasPrefix(raw, prefix) {
		raw = strings.Replace(raw, prefix, "", 1)
	}
	if strings.HasSuffix(raw, suffix) {
		raw = strings.Replace(raw, suffix, "", 1)
	}

	raw = strings.Replace(raw, " ", "", -1)
	raw = strings.Replace(raw, "\n", "", -1)
	raw = strings.Replace(raw, "\r", "", -1)

	var ll = 64
	var sl = len(raw)
	var c = sl / ll
	if sl % ll > 0 {
		c = c + 1
	}

	var buf bytes.Buffer
	buf.WriteString(prefix+"\n")
	for i:=0; i<c; i++ {
		var b = i * ll
		var e = b + ll
		if e > sl {
			buf.WriteString(raw[b:])
		} else {
			buf.WriteString(raw[b:e])
		}
		buf.WriteString("\n")
	}
	buf.WriteString(suffix)
	return buf.Bytes()
}