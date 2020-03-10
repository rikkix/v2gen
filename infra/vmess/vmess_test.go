package vmess_test

import (
	"iochen.com/v2gen/infra/vmess"
	"testing"
)

const secData = `vmess://ewogICJhZGQiOiAid3d3LnYycmF5LmNvbSIsCiAgImhvc3QiOiAid3d3LnYycmF5LmNvbSIsCiAg
ImlkIjogIjI0MThkMDg3LTY0OGQtNDk5MC04NmU4LTE5ZGNhMWQwMDZkMyIsCiAgIm5ldCI6ICJ3
cyIsCiAgInBhdGgiOiAiL3JheSIsCiAgInBvcnQiOiA0NDMsCiAgInBzIjogInRlc3Qgbm9kZSIs
CiAgInRscyI6ICJ0bHMiLAogICJ2IjogMiwKICAiYWlkIjogMCwKICAidHlwZSI6ICJhdXRvIgp9
Cg==`

var node = vmess.Link{
	Ps:   "test node",
	Add:  "www.v2ray.com",
	Port: 443,
	Id:   "2418d087-648d-4990-86e8-19dca1d006d3",
	Aid:  0,
	Net:  "ws",
	Type: "auto",
	Host: "www.v2ray.com",
	Path: "/ray",
	TLS:  "tls",
}

func TestLink_Parse(t *testing.T) {
	t.Log(node.Parse())
}

func TestGenerateFromSecData(t *testing.T) {
	link, err := vmess.GenerateFromSecData(secData)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(link)
}
