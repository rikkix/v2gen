package vmess

import (
	"encoding/json"
	"errors"
	"fmt"
	"unicode"

	"iochen.com/v2gen/v2/common/base64"
	"iochen.com/v2gen/v2/common/split"
)

type Link struct {
	Ps   string      `json:"ps"`
	Add  string      `json:"add"`
	Port interface{} `json:"port"`
	ID   string      `json:"id"`
	Aid  interface{} `json:"aid"`
	Net  string      `json:"net"`
	Type string      `json:"type"`
	Host string      `json:"host"`
	Path string      `json:"path"`
	TLS  string      `json:"tls"`
}

var (
	ErrWrongProtocol = errors.New("wrong protocol")
)

// ParseSingle parses single vmess URL into Link structure
func ParseSingle(vmessURL string) (*Link, error) {
	if len(vmessURL) < 8 {
		return &Link{}, errors.New(fmt.Sprint("wrong url:", vmessURL))
	}
	if vmessURL[:8] != "vmess://" {
		return &Link{}, ErrWrongProtocol
	}

	j, err := base64.Decode(vmessURL[8:])
	if err != nil {
		return &Link{}, err
	}

	lk := &Link{}
	err = json.Unmarshal([]byte(j), lk)
	return lk, err
}

func Parse(s string) ([]*Link, error) {
	var vl []*Link
	urlList := split.Split(s)
	for i := 0; i < len(urlList); i++ {
		lk, err := ParseSingle(urlList[i])
		if err != nil {
			if err == ErrWrongProtocol {
				continue
			} else {
				return nil, err
			}
		}
		vl = append(vl, lk)
	}
	return vl, nil
}

func (lk *Link) Config() map[string]string {
	var config = make(map[string]string)
	// set node settings
	config["address"] = lk.Add
	config["serverPort"] = fmt.Sprintf("%v", lk.Port)
	config["uuid"] = lk.ID
	config["aid"] = fmt.Sprintf("%v", lk.Aid)
	config["streamSecurity"] = lk.TLS
	config["network"] = lk.Net
	config["tls"] = lk.TLS
	config["type"] = lk.Type
	config["host"] = lk.Host
	config["type"] = lk.Type
	config["path"] = lk.Path
	return config
}

// String converts vmess link to vmess:// URL
func (lk *Link) String() string {
	b, _ := json.Marshal(lk)
	return "vmess://" + base64.Encode(string(b))
}

func redact(str string) string {
	var result []rune
	for _, v := range str {
		if unicode.IsDigit(v) {
			result = append(result, '0')
			continue
		}

		if unicode.IsUpper(v) {
			result = append(result, 'X')
			continue
		}

		if unicode.IsLower(v) {
			result = append(result, 'x')
			continue
		}

		result = append(result, v)
	}
	return string(result)
}

func (lk *Link) Safe() string {
	safeLink := &Link{
		Ps:   lk.Ps,
		Add:  redact(lk.Add),
		Port: lk.Port,
		ID:   redact(lk.ID),
		Aid:  lk.Aid,
		Net:  lk.Net,
		Type: lk.Type,
		Host: redact(lk.Host),
		Path: redact(lk.Path),
		TLS:  lk.TLS,
	}
	b, _ := json.Marshal(safeLink)
	return string(b)
}

func (lk *Link) DestAddr() string {
	return lk.Add
}

func (lk *Link) Description() string {
	return lk.Ps
}
