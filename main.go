/*
Intro :
	Generate V2Ray Json format from "vmess://{{base 64 encoded}}" format

Author Info :
	Author	: Richard Chen
	Twitter : @realRichardChen
	GitHub	: @iochen
	Website	: https://iochen.com/

Software Info :
	Version			: V0.2.10
	Support format	: v2rayN/v2rayN/v2rayN/Mode/VmessQRCode.cs (Maybe, not tested all config types now)
	License			: MIT LICENSE
*/

package main

import (
	"flag"
	"fmt"
)

const ver = "V0.2.10"

type Vmess struct {
	Ps, Add, Id, Net, Type, Host, Path, Tls string

	// in some vmess URIs is int type, some are string type
	Port, Aid interface{}
}

var (
	url          = flag.String("u", "", "The URL to get nodes info from")
	outPath      = flag.String("p", "/etc/v2ray/config.json", "V2Ray json config output path")
	userConfPath = flag.String("c", "/etc/v2ray/v2gen.ini", "V2Gen config path")
	tpl          = flag.String("tpl", "", "v2ray json tpl file path")

	vmessURIs = flag.String("vmess", "", "vmess://foo or vmess://foo;vmess://bar")

	initUserConf = flag.Bool("init", false, "if initialize V2Gen config")
	numFlag      = flag.Int("n", -1, "Choose node (auto add -y param)")
	chooseYes    = flag.Bool("y", false, "select \"yes\" when asking if preview config")
	randChoose   = flag.Bool("r", false, "select nodes at random")
	test         = flag.Bool("test", false, "only for test")
	noPing       = flag.Bool("noPing", false, "disable ping function")
	version      = flag.Bool("v", false, "version")
)

func main() {
	flag.Parse()

	if *version {
		fmt.Println("Version:", ver)
	}

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

}

func checkErr(err error) bool {
	if err != nil {
		panic(err)
		return true
	}
	return false
}
