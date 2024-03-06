package utils

import (
	"encoding/base64"
	"net/url"
)

// Encode encodes into base64, then escapes special chars with URI encoding
func Encode(s string) string {
	base64String := EncodeBase64String(s)
	result := EncodeURI(base64String)

	return result
}

// Decode decodes base64 => URI encoded string
// See the Encode function
func Decode(s string) string {
	var result string
	var resultBytes []byte
	var base64String string
	var err error

	base64String, err = url.QueryUnescape(s)

	if nil == err {
		resultBytes, err = base64.StdEncoding.DecodeString(base64String)

		if nil == err {
			result = string(resultBytes)
		}
	}

	return result
}

func EncodeBase64String(s string) string {
	byteString := []byte(s)
	result := EncodeBase64Bytes(byteString)

	return result
}

func EncodeBase64Bytes(b []byte) string {
	result := base64.StdEncoding.EncodeToString(b)

	return result
}

func EncodeURI(s string) string {
	result := url.QueryEscape(s)

	return result
}
