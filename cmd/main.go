package main

import (
	"flag"
	"fmt"
	v2gen2 "iochen.com/v2gen"
	"iochen.com/v2gen/infra/miniv2ray"
)

var (
	Version = "V1.3.* dev"
)

func main() {
	flag.Parse()

	if *v2gen2.FlagVersion {
		fmt.Println("v2gen version:", Version)
		fmt.Println("V2Ray Core version:", miniv2ray.CoreVersion())
		return
	}

	if *v2gen2.FlagInit {
		if err := v2gen2.InitV2GenConf(*v2gen2.FlagUserConf); err != nil {
			panic(err)
		}
		fmt.Println("v2gen config initialized")
		return
	}

	var err error

	if len(*v2gen2.FlagURIs) != 0 {
		err = v2gen2.GenFromSecRawData(*v2gen2.FlagURIs)
	} else if *v2gen2.FlagURL != "" {
		err = v2gen2.GenFromURL(*v2gen2.FlagURL)
	} else {
		panic("nothing to do")
	}

	if err != nil {
		panic(err)
	}
}
