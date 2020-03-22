package vmessping

import (
	"iochen.com/v2gen/app/ping"
	mv2ray "iochen.com/v2gen/infra/miniv2ray"
	"iochen.com/v2gen/infra/vmess"
	"log"
	"time"
)

func VmessPing(lk *vmess.Link, count int, dest string, totalTimeout, eachTimeout time.Duration, verbose bool) (*ping.Status, error) {
	server, err := mv2ray.StartV2Ray(lk, verbose)
	if err != nil {
		log.Println(err)
	}

	if err := server.Start(); err != nil {
		log.Println(err)
	}

	defer func() {
		if err := server.Close(); err != nil {
			log.Println(err)
		}
	}()

	ps := &ping.Status{
		Durations: &ping.DurationList{},
	}

	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(totalTimeout * time.Second)
		timeout <- true
	}()

L:
	for round := 0; round < count; round++ {
		chDelay := make(chan time.Duration)
		go func() {
			delay, err := mv2ray.MeasureDelay(server, time.Second*time.Duration(eachTimeout), dest)
			if err != nil {
				ps.ErrCounter++
			}
			chDelay <- delay
		}()

		select {
		case delay := <-chDelay:
			if delay > 0 {
				*ps.Durations = append(*ps.Durations, ping.Duration(delay))
			}
		case <-timeout:
			break L
		}
	}

	return ps, nil
}
