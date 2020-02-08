package vmess

import (
	"encoding/json"
	"github.com/iochen/v2gen/common/encoding/base64"
)

func GenerateFromSecData(secData string) (Link, error) {
	node := Link{}

	j, err := base64.Decode(secData[8:])
	if err != nil {
		return node, err
	}

	err = json.Unmarshal([]byte(j), &node)
	return node, err
}
