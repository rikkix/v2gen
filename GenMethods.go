package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"

	"time"
)

func genFromURL() {
	// get raw data from url
	rawData, err := GetContent(*url)
	checkErr(err)

	vmessList, err := rawToVmessList(rawData)
	checkErr(err)
	genFromList(vmessList)
}

func genFromVmessURIs() {
	vmessList, err := VmessURIsToList(*vmessURIs)
	checkErr(err)
	genFromList(vmessList)
}

func genFromList(vmessList []Vmess) {
	var node Vmess
	if len(vmessList) > 1 {
		// Select Node
		var n int
		var err error
		if *numFlag != -1 {
			n = *numFlag
			checkErr(err)
		} else if *randChoose {
			rand.Seed(time.Now().UnixNano())
			n = rand.Intn(len(vmessList))
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

	if (*numFlag != -1) || *chooseYes || *randChoose || *test {
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
