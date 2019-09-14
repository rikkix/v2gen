package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func GetUserConf(path string) (map[string]string, error) {
	Settings := make(map[string]string)

	//default settings
	Settings["loglevel"] = "warning"
	Settings["socksPort"] = "1080"
	Settings["udp"] = "true"
	Settings["httpPort"] = "1081"
	Settings["security"] = "aes-256-gcm"
	Settings["mux"] = "true"
	Settings["concurrency"] = "8"
	Settings["dns1"] = "9.9.9.9"
	Settings["dns2"] = "1.1.1.1"
	Settings["china"] = "true"
	Settings["tls"] = "null"
	Settings["kcp"] = "null"
	Settings["ws"] = "null"
	Settings["quic"] = "null"
	Settings["http"] = "null"
	Settings["allowInsecure"] = "false"
	Settings["mtu"] = "1350"
	Settings["tti"] = "20"
	Settings["up"] = "5"
	Settings["down"] = "20"
	Settings["congestion"] = "false"
	Settings["readBufferSize"] = "1"
	Settings["writeBufferSize"] = "1"

	// If user config not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return Settings, nil
	}

	// read user config file
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	for k, v := range ParseV2GenConf(b) {
		Settings[k] = v
	}

	return Settings, nil
}

func GenConf(Settings map[string]string) ([]byte, error) {
	conf := ConfigTpl

	if *tpl != "" {
		b, err := ioutil.ReadFile(*tpl)
		if err != nil {
			return nil, err
		}
		conf = string(b)
	}

	// set china setting
	if Settings["china"] == "true" {
		conf = strings.ReplaceAll(conf, "{{china_ip}}", "\n"+`"geoip:cn",`)
		conf = strings.ReplaceAll(conf, "{{china_sites}}", chinaSites)
	} else {
		conf = strings.ReplaceAll(conf, "{{china_ip}}", "")
		conf = strings.ReplaceAll(conf, "{{china_sites}}", "")
	}

	// set stream
	if Settings["tls"] == "tls" {
		conf = strings.ReplaceAll(conf, "{{tls}}", TLSObject)
	} else {
		conf = strings.ReplaceAll(conf, "{{tls}}", "null")
	}

	switch Settings["network"] {
	case "kcp":
		conf = strings.ReplaceAll(conf, "{{kcp}}", KcpObject)
	case "ws":
		conf = strings.ReplaceAll(conf, "{{ws}}", wsObject)
	case "http":
		conf = strings.ReplaceAll(conf, "{{http}}", HttpObject)
		Settings["host"] = ParseHost(Settings["host"])
	case "quic":
		conf = strings.ReplaceAll(conf, "{{quic}}", QuicObject)
	}

	// set other settings
	for k, v := range Settings {
		conf = strings.ReplaceAll(conf, "{{"+k+"}}", v)
	}

	return prettyPrint([]byte(conf))
}

// from "aaa.ltd,bbb.ltd" to ""aaa.ltd","bbb.ltd""
func ParseHost(s string) string {
	arr := strings.Split(s, ",")
	for i := range arr {
		arr[i] = "\"" + arr[i] + "\""
	}
	s = strings.Join(arr, ",")
	return s
}

const ConfigTpl = `{
  "log": {
    "loglevel": "{{loglevel}}"
  },
  "inbounds": [
    {
      "port": {{socksPort}},
      "protocol": "socks",
      "settings": {
		"udp": {{udp}}
      }
    },
    {
      "port": {{httpPort}},
      "protocol": "http",
      "settings": {
		"udp": {{udp}}
      }
    }
  ],
  "outbound": {
    "protocol": "vmess",
    "settings": {
      "vnext": [
        {
          "address": "{{address}}",
          "port": {{serverPort}},
          "users": [
            {
              "id": "{{uuid}}",
              "alterId": {{aid}},
              "security": "{{security}}"
            }
          ]
        }
      ]
    },
    "streamSettings": {
      "network": "{{network}}",
      "security": "{{streamSecurity}}",
      "tlsSettings": {{tls}},
      "kcpSettings": {{kcp}},
      "wsSettings": {{ws}},
      "httpSettings": {{http}},
      "quicSettings": {{quic}}
    },
    "mux": {
      "enabled": {{mux}},
      "concurrency": {{concurrency}}
    }
  },
  "dns": {
    "servers": [
      "{{dns1}}",
      "{{dns2}}",
      "localhost"
    ]
  },
  "outboundDetour": [
    {
      "protocol": "freedom",
      "settings": {},
      "tag": "direct"
    }
  ],
	"routing": {
		"strategy": "rules",
			"settings": {
			"domainStrategy": "IPIfNonMatch",
				"rules": [{{china_sites}}
					{
    			    "type": "field",
    			    "outboundTag": "direct",
     			    "ip": [{{china_ip}}
       				    "geoip:private"
					]
				}
			]
		}
	}
}`

const chinaSites = `
{
	"type": "field",
	"outboundTag": "direct",
    "domain": ["geosite:cn"] 
},`

const TLSObject = `{
 		 "serverName": "{{address}}",
 		 "allowInsecure": {{allowInsecure}},
 		 "alpn": ["http/1.1"]
		}`

const wsObject = `{
 		 "path": "{{path}}",
 		 "headers": {
  		  "Host": "{{host}}"
 		 }
		}`

const KcpObject = `
{
		"mtu": {{mtu}},
		"tti": {{tti}},
		"uplinkCapacity": {{up}},
		"downlinkCapacity": {{down}},
		"congestion": {{congestion}},
		"readBufferSize": {{readBufferSize}},
		"writeBufferSize": {{writeBufferSize}},
		"header": {
		"type": "{{type}}"
		}
		}`
const (
	HttpObject = `{
		"host": [{{host}}],
		"path": "{{path}}"
		}`
	QuicObject = `{
		  "security": "{{host}}",
		  "key": "{{path}}",
		  "header": {
		    "type": "{{type}}"
		  }
		}`
)
