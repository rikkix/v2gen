package v2gen

import (
	"fmt"

	"iochen.com/v2gen/v2/ping"
)

type Link interface {
	fmt.Stringer
	Ping(round int, dst string) (ping.Status, error)
	DestAddr() string
	Description() string
	Safe() string
	Config
}
