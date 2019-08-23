/*
Intro :
	Generate V2Ray Json format from "vmess://{{base 64 encoded}}" format

Author Info :
	Author	: Richard Chen
	Twitter : @realRichardChen
	GitHub	: @iochen
	Website	: https://iochen.com/

Software Info :
	Version			: V0.1.2
	Support format	: v2rayN/v2rayN/v2rayN/Mode/VmessQRCode.cs (Maybe, not tested all config type now)
	License			: MIT LICENSE
*/

package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Vmess struct {
	Ps, Add, Port, Id, Aid, Net, Type, Host, Path, Tls string
}

func main() {
	url := flag.String("u", "", "The URL to get configs from")
	outPath := flag.String("p", "/etc/v2ray/config.json", "config output path")
	userConfPath := flag.String("c", "/etc/v2ray/v2gen.ini", "user config path")
	flag.Parse()

	if *url == "" {
		panic("Please use -h to get help info.")
	}

	rawData, err := GetContent(*url) // get raw data from url
	checkErr(err)

	t := time.Now()

	rawData, err = Base64Dec(rawData) // first time decode
	checkErr(err)

	// get json List
	jsonList := strings.FieldsFunc(rawData, func(r rune) bool {
		if r == '\n' || r == ' ' {
			return true
		}
		return false
	})

	// remove "vmess://"
	jsonList, err = RemovePro(jsonList)
	checkErr(err)

	// get vmess struct
	vmessList := make([]Vmess, len(jsonList))
	for i := 0; i < len(jsonList); i++ {
		err := json.Unmarshal([]byte(jsonList[i]), &vmessList[i])
		checkErr(err)
	}

	del := time.Since(t)

	// choose node - CLI
	n, err := NodeSelect(vmessList)
	checkErr(err)
	node := vmessList[n]

	t = time.Now()

	Settings := make(map[string]string)

	// set user settings
	Settings, err = GetUserConf(*userConfPath)
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

	// Gen V2Ray json config
	config := GenConf(Settings)

	// make it prettier
	prettyConfig, err := prettyPrint([]byte(config))
	checkErr(err)

	// write V2Ray json config
	err = ioutil.WriteFile(*outPath, prettyConfig, 0644)
	checkErr(err)
	fmt.Println("The config file has been written to",*outPath)

	del = del + time.Since(t)

	ifYes := map[string]bool{
		"y":   true,
		"Y":   true,
		"yes": true,
		"Yes": true,
		"YES": true,
	}

	// If preview config - CLI
	var ifPreview string
	fmt.Print("=====================\nDo you want to preview the config?(y)es/(N)o: ")
	_, err = fmt.Scanf("%s", &ifPreview)
	checkErr(err)
	fmt.Println("=====================")
	if ifYes[ifPreview] {
		fmt.Println(string(prettyConfig))
		fmt.Println("=====================")
	}

	fmt.Println("All is done!", "Please restart your V2Ray Service.")
	fmt.Println("The generation process took a total of", del)
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

// remove protocol
func RemovePro(dataList []string) ([]string, error) {
	var err error
	for i := 0; i < len(dataList); i++ {
		dataList[i], err = Base64Dec(dataList[i][8:])
		if err != nil {
			return nil, err
		}
	}
	return dataList, err
}

func NodeSelect(vmessList []Vmess) (int, error) {
	var i int
	for i := 0; i < len(vmessList); i++ {
		fmt.Printf("[%d]%s\n", i, vmessList[i].Ps)
	}
	fmt.Print("=====================\nPlease Select: ")
	_, err := fmt.Scanf("%d", &i)
	return i, err
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
