package v2gen

import (
	"fmt"
	"io/ioutil"
	"iochen.com/v2gen/infra/vmess"
)

func GenFromURL(URL string) error {
	return genFromVmessList(GetVmessList(Pri2Sec(URL2Pri(URL))))
}

func GenFromSecRawData(secRawData string) error {
	return genFromVmessList(GetVmessList(SecRaw2Sec(secRawData)))
}

func genFromVmessList(vmessList *[]vmess.Link) error {
	n, err := SelectNode(vmessList)

	Settings := GenSettings((*vmessList)[n], *FlagUserConf)

	// Generate V2Ray json config
	config, err := GenConf(Settings)
	if err != nil {
		return err
	}

	op := *FlagPath

	if op == "" || op == "-" {
		fmt.Println(string(config))
	} else {
		// write V2Ray json config
		if err = ioutil.WriteFile(*FlagPath, config, 0644); err != nil {
			return err
		}
		fmt.Println("The config file has been written to", *FlagPath)
	}
	return nil
}
