package v2gen

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func InitV2GenConf(p string) error {
	// If user config  exist
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return ioutil.WriteFile(p, []byte(DefaultV2GenConf), 0644)
	}
	fmt.Println("=====================")
	fmt.Print("File already exists, if rewrite it?(y)es/(N)o: ")

	var ifContinue string
	_, err := fmt.Scanf("%s", &ifContinue)
	if err != nil {
		return err
	}
	fmt.Println("=====================")
	if ifContinue == "y" || ifContinue == "Y" {
		return ioutil.WriteFile(p, []byte(DefaultV2GenConf), 0644)
	}
	return errors.New("file already exists")
}

func ParseV2GenConf(b []byte) map[string]string {
	V2GenSettings := make(map[string]string)

	// file to lines
	lines := strings.Split(string(b), "\n")

	for _, s := range lines {
		s = strings.TrimLeft(s, " ")
		s = strings.TrimLeft(s, "\t")
		if s == "" {
			continue
		}
		if s[0] == '#' {
			continue
		}

		// Split "k v" to {k,v}
		line := strings.FieldsFunc(s, func(r rune) bool {
			if r == ' ' || r == '\t' || r == ':' {
				return true
			}
			return false
		})

		// check if is {k,v}
		if len(line) != 2 {
			continue
		}

		V2GenSettings[line[0]] = line[1] // Set Settings[k]=v
	}

	return V2GenSettings
}

const DefaultV2GenConf = `
#####################
# v2gen user config #
#####################

# V2Ray log level
# ( debug | info | warning | error | none )
loglevel: warning

# Socks port
socksPort: 1080

# Http port
httpPort: 1081

# If allow UDP traffic
# ( true | false )
udp: true

# Security
# ( aes-128-gcm | aes-256-gcm | chacha20-poly1305 | auto | none )
security: aes-256-gcm

# If enable mux
# ( true | false )
mux: true

# Mux concurrency num
concurrency: 8

# DNS server
dns1: 9.9.9.9
dns2: 1.1.1.1

# If China sites and ips directly connect
# ( true | false )
china: true


#@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
#@ NOTICE:                                            @
#@ The following settings may NOT work on every node  @
#@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

# If allow insecure connection ( true | false )
allowInsecure: false

# KCP mtu num
mtu: 1350

# KCP tti num
tti: 20

# KCP max upload speed
# Unit: MB/s
up: 5

# KCP max download speed
# Unit: MB/s
down: 20

# If enable UDP congestion control ( true | false )
congestion: false

# Read buffer size
# Unit: MB
readBufferSize: 1

# Write buffer size
# Unit: MB
writeBufferSize: 1

##############################
# Thank you for using v2gen! #
##############################
`
