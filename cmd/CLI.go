package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iochen/v2gen/vmess"
)

func SelectNode(vmessList *[]vmess.Link) (int, error) {
	var i int

	if *flagNoPing {
		for i := 0; i < len(*vmessList); i++ {
			fmt.Printf("[%d] \t%s\n", i, (*vmessList)[i].Ps)
		}
	} else {
		for i := 0; i < len(*vmessList); i++ {
			fmt.Printf("[%d] \t%-25s\t[%s]\n", i, (*vmessList)[i].Ps, (*vmessList)[i].Ping())
		}
	}

	if *flagTest {
		return 0, nil
	}

	fmt.Print("=====================\nPlease Select: ")
	_, err := fmt.Scanf("%d", &i)
	return i, err
}

func AskIfPreview(config string) {
	var ifPreview string
	fmt.Print("=====================\nDo you want to preview the config?(y)es/(N)o: ")
	_, err := fmt.Scanf("%s", &ifPreview)
	checkErr(err)
	fmt.Println("=====================")
	if ifPreview == "y" || ifPreview == "Y" {
		fmt.Println(config)
		fmt.Println("=====================")
	}
}

func PrettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "\t")
	return out.Bytes(), err
}
