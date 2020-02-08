/*
Intro :
	Generate V2Ray Json format from "vmess://" format

Author Info :
	Author	: Richard Chen
	Twitter : @realRichardChen
	GitHub	: @iochen
	Website	: https://iochen.com/

Software Info :
	Version			: V1.0.1
	Support format	: v2rayN/v2rayN/v2rayN/Mode/VmessQRCode.cs (Maybe, not tested all config types now)
	License			: MIT LICENSE
*/

package main

import (
	"flag"
	"fmt"
)

const ver = "V1.0.1"

var (
	flagURL      = flag.String("u", "", "The URL to get nodes info from")
	flagPath     = flag.String("p", "/etc/v2ray/config.json", "V2Ray json config output path")
	flagUserConf = flag.String("c", "/etc/v2ray/v2gen.ini", "V2Gen config path")
	flagTPL      = flag.String("tpl", "", "v2ray json tpl file path")
	flagURIs     = flag.String("vmess", "", "vmess://foo or vmess://foo;vmess://bar")
	flagInit     = flag.Bool("init", false, "if initialize V2Gen config")
	flagNum      = flag.Int("n", -1, "Choose node (auto add -y param)")
	flagYes      = flag.Bool("y", false, "select \"yes\" when asking if preview config")
	flagRandom   = flag.Bool("r", false, "select nodes at random")
	flagTest     = flag.Bool("test", false, "only for test")
	flagNoPing   = flag.Bool("noPing", false, "disable ping function")
	flagVersion  = flag.Bool("v", false, "version")
)

func main() {
	flag.Parse()

	if *flagVersion {
		fmt.Println("Version:", ver)
	}

	if *flagInit {
		if !checkErr(InitV2GenConf(*flagPath)) {
			fmt.Println("V2Gen config initialized")
		}
		return
	}

	if len(*flagURIs) != 0 {
		GenFromSecRawData(*flagURIs)
	} else if *flagURL != "" {
		GenFromURL(*flagURL)
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
