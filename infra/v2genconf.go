package infra

import (
	"strings"
)

type V2genConfig map[string]string

func (config *V2genConfig) Config() map[string]string {
	return map[string]string(*config)
}

func (config *V2genConfig) Append(conf map[string]string) *V2genConfig {
	for k, v := range conf {
		(*config)[k] = v
	}
	return config
}

// ParseV2genConf parses the file into a [string]string map
func ParseV2genConf(b []byte) V2genConfig {
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
			if r == ' ' || r == '\t' {
				return true
			}
			return false
		})

		// check if is {k,v}
		if len(line) != 2 {
			continue
		}

		line[0] = strings.TrimRight(line[0], ":")

		V2GenSettings[line[0]] = line[1] // Set Settings[k]=v
	}

	return V2genConfig(V2GenSettings)
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
dns1: https://1.1.1.1/dns-query
dns2: https://dns.quad9.net/dns-query

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
