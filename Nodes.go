package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func GenSettings(node Vmess, path string) map[string]string {
	Settings := make(map[string]string)

	// set user settings
	Settings, err := GetUserConf(path)
	checkErr(err)

	// set node settings
	Settings["address"] = node.Add
	Settings["serverPort"] = fmt.Sprintf("%v", node.Port)
	Settings["uuid"] = node.Id
	Settings["aid"] = fmt.Sprintf("%v", node.Aid)
	Settings["streamSecurity"] = node.Tls
	Settings["network"] = node.Net
	Settings["tls"] = node.Tls
	Settings["type"] = node.Type
	Settings["host"] = node.Host
	Settings["type"] = node.Type
	Settings["path"] = node.Path
	return Settings
}

func rawToVmessList(rawData string) ([]Vmess, error) {
	vmessURIs, err := Base64Dec(rawData) // first time decode
	if err != nil {
		return nil, err
	}
	return VmessURIsToList(vmessURIs)
}

func VmessURIsToList(vmessURIs string) ([]Vmess, error) {
	// get base64 List
	vmessURIList := strings.FieldsFunc(vmessURIs, func(r rune) bool {
		if r == '\n' || r == ' ' || r == ';' {
			return true
		}
		return false
	})

	var err error
	// get vmess struct
	vmessList := make([]Vmess, len(vmessURIList))
	for k, v := range vmessURIList {
		vmessList[k], err = VmessURItoVmess(v)
		if err != nil {
			return nil, err
		}
	}

	return vmessList, nil
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

func Base64Dec(str string) (string, error) {
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

// remove protocol && second time decode
func RemoveProAndDec(vmessURI string) (string, error) {
	return Base64Dec(vmessURI[8:])
}

func SelectNode(vmessList []Vmess) (int, error) {
	var i int

	if *noPing {
		for i := 0; i < len(vmessList); i++ {
			fmt.Printf("[%d] \t%s\n", i, vmessList[i].Ps)
		}
	} else {
		for i := 0; i < len(vmessList); i++ {
			fmt.Printf("[%d] \t%-25s\t[%s]\n", i, vmessList[i].Ps, Ping(vmessList[i].Add))
		}
	}

	if *test {
		return 0, nil
	}

	fmt.Print("=====================\nPlease Select: ")
	_, err := fmt.Scanf("%d", &i)
	return i, err
}
