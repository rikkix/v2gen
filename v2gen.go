package v2gen

import "flag"

var (
	FlagURL      = flag.String("u", "", "subscription URL")
	FlagPath     = flag.String("o", "/etc/v2ray/config.json", "output path")
	FlagUserConf = flag.String("c", "/etc/v2ray/v2gen.ini", "v2gen config path")
	FlagTPL      = flag.String("tpl", "", "V2Ray tpl path")
	FlagURIs     = flag.String("vmess", "", "vmess link(s)")
	FlagInit     = flag.Bool("init", false, "initialize v2gen config")
	FlagIndex    = flag.Int("n", -1, "node index")
	FlagRandom   = flag.Bool("r", false, "random node index")
	FlagNoPing   = flag.Bool("np", false, "do not ping")
	FlagDest     = flag.String("dest", "https://cloudflare.com/cdn-cgi/trace", "test destination url (vmess ping only)")
	FlagCount    = flag.Int("ct", 3, "ping count for each node")
	FlagETO      = flag.Int("eto", 8, "timeout seconds for each request")
	FlagTTO      = flag.Int("tto", 25, "timeout seconds for each node")
	FlagICMP     = flag.Bool("ic", false, "use ICMP ping instead of vmess ping")
	FlagMedian   = flag.Bool("med", false, "use median instead of ArithmeticMean")
	FlagThreads  = flag.Int("t", 5, "threads used when pinging")
	FlagVersion  = flag.Bool("v", false, "show version")
)
