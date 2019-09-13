/*
Intro :
	Generate V2Ray Json format from "vmess://{{base 64 encoded}}" format

Author Info :
	Author	: Richard Chen
	Twitter : @realRichardChen
	GitHub	: @iochen
	Website	: https://iochen.com/

Software Info :
	Version			: V0.2.7
	Support format	: v2rayN/v2rayN/v2rayN/Mode/VmessQRCode.cs (Maybe, not tested all config types now)
	License			: MIT LICENSE
*/

package main

import (
	"flag"
	"fmt"
)

type Vmess struct {
	Ps, Add, Id, Aid, Net, Type, Host, Path, Tls string

	// in some vmess URIs is int type, some are string type
	Port interface{}
}

var (
	url          = flag.String("u", "", "The URL to get nodes info from")
	outPath      = flag.String("p", "/etc/v2ray/config.json", "V2Ray json config output path")
	userConfPath = flag.String("c", "/etc/v2ray/v2gen.ini", "V2Gen config path")
	initUserConf = flag.Bool("init", false, "if initialize V2Gen config")
	silent       = flag.Bool("silent", false,
		"if you want to keep it silent (Select node by reading env NODE_NUM)")
	vmessURIs = flag.String("vmess", "", "vmess://foo or vmess://foo;vmess://bar")
)

func main() {
	flag.Parse()

	if *initUserConf {
		if !checkErr(InitV2GenConf(*userConfPath)) {
			fmt.Println("V2Gen config initialized")
		}
		return
	}

	if len(*vmessURIs) != 0 {
		genFromVmessURIs()
	} else if *url != "" {
		genFromURL()
	} else {
		panic("nothing to do")
	}

	fmt.Println("All is done!", "Please restart your V2Ray Service.")
}

func checkErr(err error) bool {
	if err != nil {
		panic(err)
		return true
	}
	return false
}
