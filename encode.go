package alipay

import (
	"bytes"
	"github.com/smartwalle/nsign"
	"net/url"
	"sort"
	"strings"
)

type Encoder struct {
}

// EncodeValues
// 1、将参数名及其对应的值进行升序排序
// 2、将排序后的参数名及参数名使用等号进行连接，例如：a=10
// 3、将组合之后的参数使用&号进行连接，例如：a=10&b=20&c=30&c=31
func (encoder *Encoder) EncodeValues(buffer *bytes.Buffer, values url.Values, opts *nsign.SignOptions) ([]byte, error) {
	if values == nil {
		return nil, nil
	}

	if opts.Prefix != "" {
		buffer.WriteString(opts.Prefix)
	}

	var pairs = make([]string, 0, len(values))
	for key := range values {
		if _, ok := opts.Ignores[key]; ok {
			continue
		}

		var nValues = values[key]
		for _, value := range nValues {
			pairs = append(pairs, key+"="+value)
		}
	}
	sort.Strings(pairs)

	buffer.WriteString(strings.Join(pairs, "&"))

	if opts.Suffix != "" {
		buffer.WriteString(opts.Suffix)
	}

	return buffer.Bytes(), nil
}

func (encoder *Encoder) EncodeBytes(buffer *bytes.Buffer, data []byte, opts *nsign.SignOptions) ([]byte, error) {
	if data == nil {
		return nil, nil
	}

	if opts.Prefix != "" {
		buffer.WriteString(opts.Prefix)
	}

	buffer.Write(data)

	if opts.Suffix != "" {
		buffer.WriteString(opts.Suffix)
	}

	return buffer.Bytes(), nil
}
