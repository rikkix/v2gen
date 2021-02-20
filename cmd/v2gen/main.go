package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"unicode/utf8"

	"github.com/remeh/sizedwaitgroup"

	"github.com/sirupsen/logrus"

	"iochen.com/v2gen/v2"
	"iochen.com/v2gen/v2/common/base64"
	"iochen.com/v2gen/v2/common/mean"
	"iochen.com/v2gen/v2/infra"
	"iochen.com/v2gen/v2/ping"
	"iochen.com/v2gen/v2/vmess"
)

var (
	Version = "v2.0.0-dev"

	FlagLoglevel = flag.String("loglevel", "warn", "log level")
	FlagLog      = flag.String("log", "-", "log output file")
	FlagAddr     = flag.String("u", "", "subscription address(URL)")
	FlagOut      = flag.String("o", "/etc/v2ray/config.json", "output path")
	FlagConf     = flag.String("config", "/etc/v2ray/v2gen.ini", "v2gen config path")
	FlagTPL      = flag.String("template", "", "V2Ray template path")
	FlagInit     = flag.Bool("init", false, "init v2gen config (specify certain path with -config)")
	FlagRandom   = flag.Bool("random", false, "random node index")
	FlagPing     = flag.Bool("ping", true, "ping nodes")
	FlagDest     = flag.String("dst", "https://cloudflare.com/cdn-cgi/trace", "test destination url (vmess ping only)")
	FlagCount    = flag.Int("c", 3, "ping count for each node")
	// FlagMedian   = flag.Bool("med", false, "use median instead of ArithmeticMean")
	FlagThreads = flag.Int("thread", 3, "threads used when pinging")
	FlagBest    = flag.Bool("best", false, "use best node judged by ping result")
	FlagPipe    = flag.Bool("pipe", true, "read from pipe")
	FlagVersion = flag.Bool("v", false, "show version")
)

/*
function main may be too long, here is a simple step list:
#################################################################################################
#   STEP 1 (READ):                                                                              #
#   1. read links from subscription(net) and pipe.                                              #
#                                                                                               #
#   STEP 2 (PROCESS):                                                                           #
#   	TYPE 1 (PING):                                                                          #
#   		SUBTYPE 1.1 (BEST):                                                                 #
#   			1. ping.                                                                        #
#   			2. choose the best node.                                                        #
#                                                                                               #
#   		SUBTYPE 1.2 (RANDOM):                                                               #
#   			1. ping.                                                                        #
#   			2. filter out available node list A.                                            #
#   				NOTE: if exist nodes that no error, then A would be them(it),               #
#   					  else A would be all of them.                                          #
#   			3. randomly choose one from A.                                                  #
#                                                                                               #
#   		SUBTYPE 1.3 (DEFAULT):                                                              #
#   			1. ping.                                                                        #
#   			2. print nodes and ping info.                                                   #
#   			3. wait for user's choosing.                                                    #
#                                                                                               #
#   	TYPE 2 (NOT PING):                                                                      #
#   		SUBTYPE 1.2 (RANDOM):                                                               #
#   			1. randomly choose one from nodes.                                              #
#                                                                                               #
#   		SUBTYPE 1.3 (DEFAULT):                                                              #
#   			1. print nodes and ping info.                                                   #
#   			2. wait for user's choosing.                                                    #
#                                                                                               #
#   STEP 3 (RENDER AND WRITE):                                                                  #
#   1. render and write.                                                                        #
#################################################################################################
*/

type PingInfo struct {
	Status   *ping.Status
	Duration ping.Duration
	Link     v2gen.Link
	Err      error
}

type PingInfoList []*PingInfo

func (pf *PingInfoList) Len() int {
	return len(*pf)
}

func (pf *PingInfoList) Less(i, j int) bool {
	if (*pf)[i].Err != nil {
		return false
	} else if (*pf)[j].Err != nil {
		return true
	}

	if len((*pf)[i].Status.Errors) != len((*pf)[j].Status.Errors) {
		return len((*pf)[i].Status.Errors) < len((*pf)[j].Status.Errors)
	}

	return (*pf)[i].Duration < (*pf)[j].Duration
}

func (pf *PingInfoList) Swap(i, j int) {
	(*pf)[i], (*pf)[j] = (*pf)[j], (*pf)[i]
}

func main() {
	flag.Parse()

	/*
	   LOG PART
	*/
	logger := logrus.New()
	if *FlagLog != "-" && *FlagLog != "" {
		file, err := os.Create(*FlagLog)
		if err != nil {
			logrus.Fatal(err)
		}
		defer file.Close()
		_, err = file.Write([]byte(version() + "\n"))
		if err != nil {
			panic("cannot write into log file")
		}
		logger.Out = file
	}

	// set log level
	level, err := logrus.ParseLevel(*FlagLoglevel)
	if err != nil {
		logger.Panic(err)
	}
	logger.SetLevel(level)

	/*
	   FLAG PART
	*/
	// if -v || trace, debug, info
	if *FlagVersion {
		fmt.Println(version())
		return
	} else if level > logrus.ErrorLevel {
		fmt.Println(version())
	}

	// if -init
	if *FlagInit {
		err := ioutil.WriteFile(*FlagConf, []byte(infra.DefaultV2GenConf), 0644)
		if err != nil {
			panic(err)
			return
		}
		logger.Info("v2gen config initialized")
		return
	}

	/*
	   LINK PART
	*/
	var linkList []v2gen.Link // combine links from different sources
	// read from subscribe address(net)
	if *FlagAddr != "" {
		logger.Infof("Reading from %s...", *FlagAddr)
		resp, err := http.Get(*FlagAddr)
		if err != nil {
			logger.Fatal(err)
		}
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.Fatal(err)
		}
		links, err := ParseLinks(bytes)
		if err != nil {
			logger.Fatal(err)
		}
		linkList = append(linkList, links...)
	}

	// check whether reading from pipe
	if fi, _ := os.Stdin.Stat(); (fi.Mode()&os.ModeCharDevice) == 0 && *FlagPipe {
		logger.Info("Reading from pipe...")
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		links, err := ParseLinks(bytes)
		if err != nil {
			logger.Fatal(err)
		}
		linkList = append(linkList, links...)

	}

	// if no Link, then exit
	if len(linkList) == 0 {
		logger.Warn("no available links, nothing to do")
		os.Exit(0)
	}

	var chosenLink v2gen.Link
	var spaceCount = func(i int, str string) string {
		rl := utf8.RuneCountInString(str)
		c := i - (len(str)+rl)/2
		if c < 0 {
			c = 0
		}
		return strings.Repeat(" ", c)
	}
	if *FlagPing { // if ping
		// make ping info list
		pingInfoList := make(PingInfoList, len(linkList))
		wg := sizedwaitgroup.New(*FlagThreads)
		for i := range linkList {
			wg.Add()
			go func(i int) {
				logger.Debugf("[%d/%d]Pinging %s\n", i, len(linkList)-1, linkList[i].Safe())
				if level > logrus.ErrorLevel {
					fmt.Printf("\rPinging %d/%d", i, len(linkList)-1)
				}
				defer func() {
					wg.Done()
				}()
				pingInfoList[i] = &PingInfo{
					Link: linkList[i],
				}
				status, err := linkList[i].Ping(*FlagCount, *FlagDest)
				if status.Durations == nil || len(*status.Durations) == 0 {
					pingInfoList[i].Err = errors.New("all error")
					status.Durations = &ping.DurationList{-1}
				}
				if err != nil {
					pingInfoList[i].Err = err
					pingInfoList[i].Status = &ping.Status{
						Durations: &ping.DurationList{},
					}
				} else {
					pingInfoList[i].Status = &status
				}
			}(i)
		}
		wg.Wait()
		fmt.Println()

		for i := range pingInfoList {
			var ok bool
			pingInfoList[i].Duration, ok = mean.ArithmeticMean(pingInfoList[i].Status.Durations).(ping.Duration)
			if !ok {
				pingInfoList[i].Duration = 0
			}
		}
		sort.Sort(&pingInfoList)
		if *FlagBest { // if ping && best
			chosenLink = pingInfoList[0].Link
		} else if *FlagRandom { // if ping && rand
			pingInfoList = AvailableLinks(pingInfoList)
			i, err := Random(len(pingInfoList))
			if err != nil {
				logger.Fatal(err)
			}
			chosenLink = pingInfoList[i].Link
		} else { // if ping && not rand && not best
			for i := range pingInfoList {
				fmt.Printf("[%2d] %s%s[%-7s(%d errors)]\n",
					i, pingInfoList[i].Link.Description(),
					spaceCount(30, pingInfoList[i].Link.Description()),
					pingInfoList[i].Duration.Precision(1e6), len(pingInfoList[i].Status.Errors))
			}
			i := Select(len(pingInfoList))
			chosenLink = pingInfoList[i].Link
		}
	} else { // if not ping
		if *FlagRandom { // if not ping && rand
			i, err := Random(len(linkList))
			if err != nil {
				logger.Fatal(err)
			}
			chosenLink = linkList[i]
		} else { // if not ping && not rand
			for i := range linkList {
				fmt.Printf("[%2d] %s%s\n",
					i, linkList[i].Description(),
					spaceCount(30, linkList[i].Description()))
			}
			i := Select(len(linkList))
			chosenLink = linkList[i]
		}
	}

	/*
		CONFIG PART
	*/

	var template []byte
	template = []byte(infra.ConfigTpl)
	if *FlagTPL != "" {
		tpl, err := ioutil.ReadFile(*FlagTPL)
		if err != nil {
			logrus.Error(err, "using default template...")
		} else {
			template = tpl
		}
	}

	v2genConf := infra.V2genConfig{}
	confFile, err := ioutil.ReadFile(*FlagConf)
	if err == nil {
		v2genConf = infra.ParseV2genConf(confFile)
	}
	conf := infra.DefaultConf()
	bytes, err := infra.GenV2RayConf(*conf.Append(v2genConf).Append(chosenLink.Config()), template)
	if err != nil {
		logrus.Fatal(err)
	}

	if *FlagOut == "-" || *FlagOut == "" {
		fmt.Println(string(bytes))
		return
	} else {
		err := ioutil.WriteFile(*FlagOut, bytes, 0644)
		if err != nil {
			logrus.Fatal(err)
		} else {
			if level > logrus.ErrorLevel {
				fmt.Printf("config has been written to %s\n", filepath.Clean(*FlagOut))
			}
		}
	}
}

func version() string {
	return fmt.Sprintf("v2gen %s, V2Ray %s (%s %dcores %s/%s)", Version, vmess.CoreVersion(),
		runtime.Version(), runtime.NumCPU(), build.Default.GOOS, build.Default.GOARCH)
}

func ParseLinks(b []byte) ([]v2gen.Link, error) {
	s, err := base64.Decode(string(b))
	if err != nil {
		return nil, err
	}
	linkList, err := vmess.Parse(s)
	if err != nil {
		return nil, err
	}
	links := make([]v2gen.Link, len(linkList))
	for i := range linkList {
		links[i] = linkList[i]
	}
	return links, err
}

func AvailableLinks(pil PingInfoList) PingInfoList {
	var pingInfoList PingInfoList
	for i := range pil {
		if pil[i].Err != nil && len(pil[i].Status.Errors) == 0 {
			pingInfoList = append(pingInfoList, pil[i])
		}
	}

	if len(pingInfoList) != 0 {
		return pingInfoList
	} else {
		return pil
	}
}

// Select returns an int [0,max)
func Select(max int) int {
	var in int
	fmt.Print("=====================\nPlease Select: ")
	_, err := fmt.Scanf("%d", &in)
	if err != nil || in < 0 || in >= max {
		fmt.Println("wrong number, please reselect")
		return Select(max)
	}
	return in
}

func Random(max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}

	return int(n.Int64()), nil
}
