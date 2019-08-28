package encoding

import (
	"crypto/x509"
	"encoding/pem"
)

func LoadCertificate(b []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(b)
	if block == nil {
		return nil, nil
	}
	csr, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	return csr, nil
}
