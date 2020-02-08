package main

import (
	base642 "github.com/iochen/v2gen/common/encoding/base64"
	"github.com/iochen/v2gen/vmess"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func URL2Pri(URL string) string {
	resp, err := http.Get(URL)
	if err != nil {
		log.Println(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return string(b)
}

func Pri2Sec(priData string) []string {
	secRawData, err := base642.Decode(priData)
	if err != nil {
		log.Println(err)
	}

	return SecRaw2Sec(secRawData)
}

func SecRaw2Sec(secRawData string) []string {
	sep := map[rune]bool{
		' ':  true,
		'\n': true,
		',':  true,
		';':  true,
		'\t': true,
		'\f': true,
		'\v': true,
		'r':  true,
	}

	return strings.FieldsFunc(secRawData, func(r rune) bool {
		return sep[r]
	})
}

func GetVmessList(secDataList []string) *[]vmess.Link {
	var nodeList []vmess.Link

	for _, secData := range secDataList {
		node, err := vmess.GenerateFromSecData(secData)
		if err != nil {
			log.Println(err)
		}

		nodeList = append(nodeList, node)
	}

	return &nodeList
}
