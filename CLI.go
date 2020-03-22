package v2gen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"iochen.com/v2gen/infra/vmess"
	"math/rand"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

func SelectNode(vmessList *[]vmess.Link) (int, error) {
	var n int
	if len(*vmessList) > 1 {
		if *FlagIndex != -1 {
			n = *FlagIndex
		} else if *FlagRandom {
			rand.Seed(time.Now().UnixNano())
			n = rand.Intn(len(*vmessList))
		} else if len(*vmessList) == 1 {
			n = 0
		} else {
			var npList *NodeStatusList

			if *FlagNoPing {
				for k, v := range *vmessList {
					fmt.Printf("[%2d] %s\n", k, v.Ps)
				}
			} else {
				npList = PingNodes(vmessList, !*FlagBest && !*FlagSort)

				if *FlagBest || *FlagSort {
					sort.Sort(npList)
					if *FlagBest {
						return (*npList)[0].Index, nil
					}
					for i := 0; i < len(*npList); i++ {
						PrintNode(i, vmessList, &(*npList)[i])
					}
				}

			}

		SELECT:
			var in int
			fmt.Print("=====================\nPlease Select: ")
			_, err := fmt.Scanf("%d", &in)
			if err != nil {
				goto SELECT
			}

			if *FlagNoPing {
				return in, nil
			} else {
				n = (*npList)[in].Index
			}
		}
	}
	return n, nil
}

func PrintNode(i int, vmessList *[]vmess.Link, np *NodePingStatus) {
	ps := (*vmessList)[np.Index].Ps

	if np.Err != nil {
		fmt.Printf("[%2d] %s%s[%s]\n", i, ps, spaceCount(35, ps), np.Err.Error())
		return
	}

	fmt.Printf("[%2d] %s%s[%-7s(%d errors)]\n", i, ps, spaceCount(30, ps), np.Result.Precision(1e6), np.PingStat.ErrCounter)
}

func spaceCount(i int, str string) string {
	rl := utf8.RuneCountInString(str)
	c := i - (len(str)+rl)/2
	if c < 0 {
		c = 0
	}
	return strings.Repeat(" ", c)
}

func PrettyPrint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "\t")
	return out.Bytes(), err
}
