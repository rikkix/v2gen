/*
Intro :
	Generate V2Ray Json format from "vmess://{{base 64 encoded}}" format

Author Info :
	Author	: Richard Chen
	Twitter : @realRichardChen
	GitHub	: @iochen
	Website	: https://iochen.com/

Software Info :
	Version			: V0.2.0
	Support format	: v2rayN/v2rayN/v2rayN/Mode/VmessQRCode.cs (Maybe, not tested all config types now)
	License			: MIT LICENSE
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Vmess struct {
	Ps, Add, Port, Id, Aid, Net, Type, Host, Path, Tls string
}

func main() {
	url := flag.String("u", "", "The URL to get nodes info from")
	outPath := flag.String("p", "/etc/v2ray/config.json", "V2Ray json config output path")
	userConfPath := flag.String("c", "/etc/v2ray/v2gen.ini", "V2Gen config path")
	initUserConf := flag.Bool("init", false, "if initialize V2Gen config")
	silent := flag.Bool("silent", false,
		"if you want to keep it silent (Select node by reading env NODE_NUM)")
	flag.Parse()

	if *initUserConf {
		if !checkErr(InitV2GenConf(*userConfPath)) {
			fmt.Println("V2Gen config initialized")
		}
		return
	}

	if *url == "" {
		panic("no url found")
	}

	// get raw data from url
	rawData, err := GetContent(*url)
	checkErr(err)

	vmessList, err := RawToVmessList(rawData)
	checkErr(err)

	// Select Node
	var n int
	if *silent {
		n, err = strconv.Atoi(os.Getenv("NODE_NUM"))
		checkErr(err)
	} else {
		n, err = SelectNode(vmessList)
		checkErr(err)
	}
	node := vmessList[n]

	Settings := GenSettings(node, *userConfPath)

	// Gen V2Ray json config
	config, err := GenConf(Settings)
	checkErr(err)

	// write V2Ray json config
	err = ioutil.WriteFile(*outPath, config, 0644)
	checkErr(err)
	fmt.Println("The config file has been written to", *outPath)

	if *silent {
		fmt.Println("=====================")
		fmt.Println(string(config))
		fmt.Println("=====================")
		fmt.Println("All is done!", "Please restart your V2Ray Service.")
		return
	}

	// If preview config - CLI
	var ifPreview string
	fmt.Print("=====================\nDo you want to preview the config?(y)es/(N)o: ")
	_, err = fmt.Scanf("%s", &ifPreview)
	checkErr(err)
	fmt.Println("=====================")
	if ifPreview == "y" || ifPreview == "Y" {
		fmt.Println(string(config))
		fmt.Println("=====================")
	}

	fmt.Println("All is done!", "Please restart your V2Ray Service.")
}

func checkErr(err error) bool {
	if err != nil {
		log.Fatalln(err)
		return true
	}
	return false
}
