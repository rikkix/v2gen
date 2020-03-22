package v2gen

import (
	"fmt"
	"iochen.com/v2gen/app/ping"
	"iochen.com/v2gen/app/ping/icmpping"
	"iochen.com/v2gen/app/ping/vmessping"
	"iochen.com/v2gen/infra/mean"
	"iochen.com/v2gen/infra/vmess"
	"time"
)

type NodePingStatus struct {
	Index    int
	PingStat *ping.Status
	Result   ping.Duration
	Err      error
}

type NodeStatusList []NodePingStatus

func (sList *NodeStatusList) Len() int {
	return len(*sList)
}

func (sList *NodeStatusList) Less(i, j int) bool {
	if (*sList)[i].PingStat.ErrCounter != (*sList)[j].PingStat.ErrCounter {
		return (*sList)[i].PingStat.ErrCounter < (*sList)[j].PingStat.ErrCounter
	}

	return (*sList)[i].Result < (*sList)[j].Result
}

func (sList *NodeStatusList) Swap(i, j int) {
	(*sList)[i], (*sList)[j] = (*sList)[j], (*sList)[i]
}

func PingNodes(lks *[]vmess.Link, print bool) *NodeStatusList {
	var npList NodeStatusList
	npCh := make(chan NodePingStatus)
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
			var result mean.Value
			if *FlagMedian {
				result = mean.Median(np.PingStat.Durations)
			} else {
				result = mean.ArithmeticMean(np.PingStat.Durations)
			}

			if result != nil {
				np.Result = result.(ping.Duration)
			}

			npList = append(npList, np)
			if print {
				PrintNode(i, lks, &np)
			} else {
				fmt.Printf("\rPinging %d", i)
			}
		}
	}
	fmt.Print("\r")

	return &npList
}

func MakeNode(k int, v vmess.Link, waitList *int, npCh chan NodePingStatus, doneCh chan bool) {
	ps, err := Ping(&v)
	np := NodePingStatus{
		Index:    k,
		PingStat: ps,
		Err:      err,
	}
	npCh <- np
	*waitList--
	doneCh <- true
}

func Ping(lk *vmess.Link) (*ping.Status, error) {
	ps := &ping.Status{
		Durations: &ping.DurationList{},
	}
	var err error
	if *FlagICMP {
		ps, err = icmpping.Ping(lk, *FlagCount, time.Duration(*FlagTTO), time.Duration(*FlagETO))
	} else {
		ps, err = vmessping.VmessPing(lk, *FlagCount, *FlagDest,
			time.Duration(*FlagTTO), time.Duration(*FlagETO), false)
	}

	return ps, err
}
