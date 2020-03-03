package vmess

import (
	"errors"
	"fmt"
	"github.com/iochen/v2gen/utils/icmpping"
	"time"
)

type PingMethod uint8

const (
	ICMP PingMethod = iota
	VMESS
)

func (lk *Link) Ping(pm PingMethod, mtd Method, dest string, count, eto, tto int, verbose bool) (string, error) {
	defer func() {
		recover()
	}()

	switch pm {
	case ICMP:
		return icmpping.Ping(lk.Add)
	case VMESS:
		ps, err := VmessPing(lk, mtd, count, dest, time.Duration(tto), time.Duration(eto), verbose)
		if err != nil {
			return err.Error(), nil
		}

		return fmt.Sprintf("%s (%d errors)", ps.Duration.String(), ps.ErrCounter), nil
	}
	return "", errors.New("unknown method")
}
