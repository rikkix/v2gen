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
	"github.com/iochen/v2gen/infra/miniv2ray"
)

var (
	V2GEN_VER = "V1.2.1"
)

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
	flagDest     = flag.String("dest", "https://cloudflare.com/cdn-cgi/trace", "test destination url (vmess ping only)")
	flagCount      = flag.Int("ct", 3, "ping count for each node (vmess ping only)")
	flagETO      = flag.Int("eto", 8, "timeout seconds for each request (vmess ping only)")
	flagTTO      = flag.Int("tto", 25, "timeout seconds for each node (vmess ping only)")
	flagICMP	= flag.Bool("t",false,"use ICMP ping instead of vmess ping")
	flagMedian = flag.Bool("med",false,"use median instead of ArithmeticMean (vmess ping only)")
	flagVersion  = flag.Bool("v", false, "show version")
)

func main() {
	flag.Parse()

	if *flagVersion {
		fmt.Println("v2gen version:", V2GEN_VER)
		fmt.Println("V2Ray Core version:", miniv2ray.CoreVersion())
		return
	}

	if *flagInit {
		if !checkErr(InitV2GenConf(*flagUserConf)) {
			fmt.Println("v2gen config initialized")
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
