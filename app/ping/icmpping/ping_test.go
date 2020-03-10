package icmpping_test

import (
	"iochen.com/v2gen/app/ping/icmpping"
	"iochen.com/v2gen/infra/vmess"
	"testing"
)

var lk = &vmess.Link{
	Ps:   "test node",
	Add:  "8.8.8.8",
	Port: 443,
	Id:   "2418d087-648d-4990-86e8-19dca1d006d3",
	Aid:  0,
	Net:  "ws",
	Type: "auto",
	Host: "www.v2ray.com",
	Path: "/ray",
	TLS:  "tls",
}

func TestPing(t *testing.T) {
	d, err := icmpping.Ping(lk, 10, 30, 2)
	if err != nil {
		t.Error(err)
	}
	t.Log(d.Durations, d.ErrCounter)
}
