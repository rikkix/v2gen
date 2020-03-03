package vmess_test

import (
	"github.com/iochen/v2gen/vmess"
	"os"
	"testing"
	"time"
)

func TestVmessPing(t *testing.T) {
	secData := os.Getenv("VMESS_LINK")

	lk, err := vmess.GenerateFromSecData(secData)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", lk)

	ps, err := vmess.VmessPing(&lk, vmess.Median, 3, "https://cloudflare.com/cdn-cgi/trace", time.Duration(20), time.Duration(10),true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log( *ps)
}
