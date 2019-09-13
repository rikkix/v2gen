package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GenSettings(node Vmess, path string) map[string]string {
	Settings := make(map[string]string)

	// set user settings
	Settings, err := GetUserConf(path)
	checkErr(err)

	// set node settings
	Settings["address"] = node.Add
	Settings["serverPort"] = node.Port
	Settings["uuid"] = node.Id
	Settings["aid"] = node.Aid
	Settings["streamSecurity"] = node.Tls
	Settings["network"] = node.Net
	Settings["tls"] = node.Tls
	Settings["type"] = node.Type
	Settings["host"] = node.Host
	Settings["type"] = node.Type
	return Settings
}

func RawToVmessList(rawData string) ([]Vmess, error) {
	rawData, err := Base64Dec(rawData) // first time decode
	if err != nil {
		return nil, err
	}

	// get base64 List
	vmessURIList := strings.FieldsFunc(rawData, func(r rune) bool {
		if r == '\n' || r == ' ' {
			return true
		}
		return false
	})

	// get vmess struct
	vmessList := make([]Vmess, len(vmessURIList))
	for i := 0; i < len(vmessURIList); i++ {
		vmessList[i], err = VmessURItoVmess(vmessURIList[i])
		if err != nil {
			return nil, err
		}
	}
	return vmessList, err
}

func VmessURItoVmess(URI string) (Vmess, error) {
	vmess := Vmess{}
	j, err := RemoveProAndDec(URI)
	if err != nil {
		return vmess, err
	}
	err = json.Unmarshal([]byte(j), &vmess)
	return vmess, err
}

func prettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "\t")
	return out.Bytes(), err
}

func GetContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	return string(b), err
}

func Base64Dec(str string) (string, error) {
	de, err := base64.StdEncoding.DecodeString(str)
	return string(de), err
}

// remove protocol && second time decode
func RemoveProAndDec(vmessURI string) (string, error) {
	return Base64Dec(vmessURI[8:])
}

func SelectNode(vmessList []Vmess) (int, error) {
	var i int
	for i := 0; i < len(vmessList); i++ {
		fmt.Printf("[%d]%s\n", i, vmessList[i].Ps)
	}
	fmt.Print("=====================\nPlease Select: ")
	_, err := fmt.Scanf("%d", &i)
	return i, err
}
