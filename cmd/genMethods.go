package main

import (
	"fmt"
	"github.com/iochen/v2gen/vmess"
	"io/ioutil"
)

func GenFromURL(URL string) {
	genFromVmessList(GetVmessList(Pri2Sec(URL2Pri(URL))))
}

func GenFromSecRawData(secRawData string) {
	genFromVmessList(GetVmessList(SecRaw2Sec(secRawData)))
}

func genFromVmessList(vmessList *[]vmess.Link) {
	n, err := SelectNode(vmessList)

	Settings := GenSettings((*vmessList)[n], *flagUserConf)

	// Generate V2Ray json config
	config, err := GenConf(Settings)
	checkErr(err)

	if *flagPath != "" {
		// write V2Ray json config
		err = ioutil.WriteFile(*flagPath, config, 0644)
		checkErr(err)
		fmt.Println("The config file has been written to", *flagPath)
	}

	if (*flagNum != -1) || *flagYes || *flagRandom || *flagTest {
		fmt.Println(string(config))
		return
	}

	AskIfPreview(string(config))
}
