package v2gen

import (
	"iochen.com/v2gen/app/ping"
	"iochen.com/v2gen/app/ping/icmpping"
	"iochen.com/v2gen/app/ping/vmessping"
	"iochen.com/v2gen/infra/vmess"
	"time"
)

type NodePing struct {
	Index    int
	PingStat *ping.PingStat
	Err      error
}

func PingNodes(lks *[]vmess.Link, print bool) *[]NodePing {
	var npList []NodePing
	npCh := make(chan NodePing)
	max := *FlagThreads

	go func() {
		var doneCh = make(chan bool)
		var waitList int
		for k, v := range *lks {
			for {
				if waitList >= max {
					<-doneCh
				} else {
					break
				}
			}
			waitList++
			go MakeNode(k, v, &waitList, npCh, doneCh)
		}

	}()

	for i := 0; i < len(*lks); i++ {
		select {
		case np := <-npCh:
			npList = append(npList, np)
			if print {
				PrintNode(i, lks, &np)
			}
		}
	}

	return &npList
}

func MakeNode(k int, v vmess.Link, waitList *int, npCh chan NodePing, doneCh chan bool) {
	ps, err := Ping(&v)
	np := NodePing{
		Index:    k,
		PingStat: ps,
		Err:      err,
	}
	npCh <- np
	*waitList--
	doneCh <- true
}

func Ping(lk *vmess.Link) (*ping.PingStat, error) {
	var ps *ping.PingStat
	var err error
	if *FlagICMP {
		ps, err = icmpping.Ping(lk, *FlagCount, time.Duration(*FlagTTO), time.Duration(*FlagETO))
	} else {
		ps, err = vmessping.VmessPing(lk, *FlagCount, *FlagDest, time.Duration(*FlagTTO), time.Duration(*FlagETO), false)
	}

	if err != nil {
		return nil, err
	}

	return ps, nil
}
