package vmess

import (
	"log"
	"time"

	"iochen.com/v2gen/v2/ping"
)

func (lk *Link) Ping(round int, dst string) (ping.Status, error) {
	server, err := startV2Ray(lk, false, false)
	if err != nil {
		return ping.Status{}, err
	}

	defer func() {
		if err := server.Close(); err != nil {
			log.Println(err)
		}
	}()

	ps := ping.Status{
		Durations: &ping.DurationList{},
	}

	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(3 * time.Duration(round) * time.Second)
		timeout <- true
	}()

L:
	for count := 0; count < round; count++ {
		chDelay := make(chan time.Duration)
		go func() {
			delay, err := measureDelay(server, 3*time.Second, dst)
			if err != nil {
				ps.Errors = append(ps.Errors, &err)
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
