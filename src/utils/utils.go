package utils

import (
	"encoding/base64"

	"golang.org/x/text/encoding/unicode"
)



var Backspace = []byte{8}

var NewLine = []byte{10}

var Nullbyte = []byte{0}

func SliceStringContains(s []string, e string) bool {
	for _, a :range s {
		if a == e {
			return true
		}
	}
	return false 
}

func SliceByteContains(s []byte, e byte) bool {
	for _, a :range s {
		if a == e {
			return true 
		}
	}
	return false
}

func Min(x int64, y int64) int64 {
	if x < y {
		return x
	}

	return y
}


type Iface struct {
	Name string
	IP   string
}

func Utf16leBase64(s string) (string, error) {
	var stringB64 = ""
	utfEncoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	ut16LeEncodedMessage, err := utfEncoder.String(s)
	if err == nil {
		stringB64 = base64.StdEncoding.EncodeToString([]byte(ut16LeEncodedMessage))
	}
	return stringB64, err
}