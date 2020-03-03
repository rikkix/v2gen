package vmess

import (
	"errors"
	mv2ray "github.com/iochen/v2gen/infra/miniv2ray"
	"log"
	"sort"
	"time"
)

type Method uint8

const (
	ArithmeticMean Method = iota
	Median
)

type PingStat struct {
	Duration   time.Duration
	durations  []time.Duration
	ErrCounter uint
}

func VmessPing(lk *Link, mtd Method, count int, dest string, totalTimeout , v2timeout time.Duration, verbose bool) (*PingStat, error) {
	link := &mv2ray.Link{
		Add: lk.Add,
		Port : lk.Port,
		Id:lk.Id,
		Aid:lk.Aid,
		Net:lk.Net,
		Type: lk.Type,
		Host:lk.Host,
		Path:lk.Path,
		TLS:lk.TLS,
	}

	server, err := mv2ray.StartV2Ray(link, verbose)
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

	ps := &PingStat{}

	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(totalTimeout*time.Second)
		timeout <- true
	}()

L:
	for round := 0; round < count; round++ {
		chDelay := make(chan time.Duration)
		go func() {
			delay, err := mv2ray.MeasureDelay(server, time.Second*time.Duration(v2timeout), dest)
			if err != nil {
				ps.ErrCounter++
			}
			chDelay <- delay
		}()

		select {
		case delay := <-chDelay:
			if delay > 0 {
				ps.durations = append(ps.durations, delay)
			}
		case <-timeout:
			break L
		}
	}

	switch mtd {
	case ArithmeticMean:
		return ps.aMean(), nil
	case Median:
		return ps.median(), nil
	default:
		return ps, errors.New("wrong Method")
	}

}

func (ps *PingStat) aMean() *PingStat {
	var sum time.Duration
	durations := ps.durations
	l := len(durations)
	if l == 0 {
		return nil
	}
	for i := 0; i < l; i++ {
		sum += durations[i]
	}
	(*ps).Duration = sum / time.Duration(l)
	return ps
}

func (ps *PingStat) median() *PingStat {
	durations := ps.durations
	l := len(durations)
	if l == 0 {
		return nil
	}
	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < durations[j]
	})

	if l%2 == 0 {
		(*ps).Duration = (durations[l/2] + durations[l/2-1]) / 2
	} else {
		(*ps).Duration = durations[(l-1)/2]
	}
	return ps
}
