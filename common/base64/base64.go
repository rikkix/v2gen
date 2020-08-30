package base64

import (
	"encoding/base64"
	"errors"
)

func Decode(str string) (string, error) {
	de, err := base64.StdEncoding.DecodeString(str)
	if err == nil {
		return string(de), err
	}

	de, err = base64.RawStdEncoding.DecodeString(str)
	if err == nil {
		return string(de), err
	}

	de, err = base64.URLEncoding.DecodeString(str)
	if err == nil {
		return string(de), err
	}

	de, err = base64.RawURLEncoding.DecodeString(str)
	if err == nil {
		return string(de), err
	}

	return "", errors.New("no proper base64 decode method for: " + str)
}

func Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
