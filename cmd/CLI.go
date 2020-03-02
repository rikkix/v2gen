package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iochen/v2gen/vmess"
	"log"
	"math/rand"
	"time"
)

func SelectNode(vmessList *[]vmess.Link) (int, error) {
	var n int

	if len(*vmessList) > 1 {
		if *flagIndex != -1 {
			n = *flagIndex
		} else if *flagRandom {
			rand.Seed(time.Now().UnixNano())
			n = rand.Intn(len(*vmessList))
		} else {

			if *flagNoPing {
				for i := 0; i < len(*vmessList); i++ {
					fmt.Printf("[%d] \t%s\n", i, (*vmessList)[i].Ps)
				}
			} else {
				for i := 0; i < len(*vmessList); i++ {
					fmt.Printf("[%d] \t%-25s\t[%s]\n", i, (*vmessList)[i].Ps, (*vmessList)[i].Ping())
				}
			}

			fmt.Print("=====================\nPlease Select: ")
			_, err := fmt.Scanf("%d", &n)
			if err != nil {
				log.Printf("%v\nSelect again!\n\n", err)
				return SelectNode(vmessList)
			}

		}
	} else if len(*vmessList) == 1 {
		n = 0
	} else {
		panic("no available vmess found")
	}

	return n, nil
}

func PrettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "\t")
	return out.Bytes(), err
}
