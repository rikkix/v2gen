package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func genFromURL() {
	// get raw data from url
	rawData, err := GetContent(*url)
	checkErr(err)

	vmessList, err := rawToVmessList(rawData)
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

	if *outPath != "" {
		// write V2Ray json config
		err = ioutil.WriteFile(*outPath, config, 0644)
		checkErr(err)
		fmt.Println("The config file has been written to", *outPath)
	}

	if *silent {
		fmt.Println(string(config))
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
}

func genFromVmessURIs() {
	vmessList, err := VmessURIsToList(*vmessURIs)
	checkErr(err)

	var node Vmess
	if len(vmessList) > 1 {
		// Select Node
		var n int
		if *silent {
			n, err = strconv.Atoi(os.Getenv("NODE_NUM"))
			checkErr(err)
		} else {
			n, err = SelectNode(vmessList)
			checkErr(err)
		}
		node = vmessList[n]
	} else if len(vmessList) == 1 {
		node = vmessList[0]
	} else {
		panic("no available vmess found")
	}

	Settings := GenSettings(node, *userConfPath)

	// Gen V2Ray json config
	config, err := GenConf(Settings)
	checkErr(err)

	if *outPath != "" {
		// write V2Ray json config
		err = ioutil.WriteFile(*outPath, config, 0644)
		checkErr(err)
		fmt.Println("The config file has been written to", *outPath)
	}

	if *silent {
		fmt.Println(string(config))
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

}
