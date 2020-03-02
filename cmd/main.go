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

const ver = "V1.0.2"

var (
	flagURL      = flag.String("u", "", "subscription URL")
	flagPath     = flag.String("o", "/etc/v2ray/config.json", "output path")
	flagUserConf = flag.String("c", "/etc/v2ray/v2gen.ini", "v2gen config path")
	flagTPL      = flag.String("tpl", "", "V2Ray tpl path")
	flagURIs     = flag.String("vmess", "", "vmess link(s)")
	flagInit     = flag.Bool("init", false, "initialize v2gen config")
	flagIndex    = flag.Int("n", -1, "node index")
	flagRandom   = flag.Bool("r", false, "random node index")
	flagNoPing   = flag.Bool("np", false, "do not ping")
	flagVersion  = flag.Bool("v", false, "show version")
)

func main() {
	flag.Parse()

	if *flagVersion {
		fmt.Println("Version:", ver)
		return
	}

	if *flagInit {
		if !checkErr(InitV2GenConf(*flagUserConf)) {
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
