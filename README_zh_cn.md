# v2gen

v2gen 是一个强大的 V2Ray 订阅客户端，使用 vmessping 代替 ICMP ping

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?)](https://pkg.go.dev/iochen.com/v2gen)
![GitHub top language](https://img.shields.io/github/languages/top/iochen/v2gen)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/iochen/v2gen) 
![Go](https://github.com/iochen/v2gen/workflows/Test/badge.svg) 

[English](README.md)

## 预览
```
[ 0] Server A      [451ms  (0 errors)]
[ 1] Server B      [452ms  (0 errors)]
[ 3] Server C      [251ms  (0 errors)]
...
[25] Server Z      [652ms  (2 errors)]
=====================
Please Select:
```



## 编译或下载

```sh
git clone https://github.com/iochen/v2gen/ && cd v2gen
env GOPRIVATE=github.com/v2ray/v2ray-core go build ./cmd/v2gen
```
或在 GitHub Release 上下载



## 快速开始

```sh
v2gen -u {{Your subscription link}} -o {{Your V2Ray config path}}
```



## 参数

```Param
Usage of v2gen:
  -best
        根据 ping 的结果选择最优节点
  -c int
        每个节点 ping 的次数 (default 3)
  -config string
        v2gen 配置文件路径 (default "/etc/v2ray/v2gen.ini")
  -dst string
        测试目标地址 (vmess ping only) (default "https://cloudflare.com/cdn-cgi/trace")
  -init
        初始化 v2gen 配置 (specify certain path with -config)
  -log string
        日志输出文件 (default "-")
  -loglevel string
        日志等级 (default "warn")
  -o string
        输出路径 (default "/etc/v2ray/config.json")
  -ping
        ping 节点 (default true)
  -pipe
        自动从 pipe 中读取 (default true)
  -random
        随机节点
  -template string
        V2Ray 模板路径
  -thread int
        ping 线程数 (default 3)
  -u string
        订阅链接 (URL)
  -v    展示版本
```



## v2gen 用户配置

你可以使用 `v2gen --init` 来生成初始化配置文件

```yaml
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

```

以下配置可能不对每个节点都有效

```yaml
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
```



## 模板文件

### 渲染流程

1. 将所有配置（节点信息，用户配置，默认配置 （排名分先后））转化为 `key-value` 对（哈系表）
2. 读取用户模板或使用默认模板，并加载
3. 将加载好的模板文件中 `{{foobar}}` 部分替换为 1 中哈系表中 `foobar` 所对的键值（[关键词会进行特殊处理](#关键词)）
4. 输出

### 关键词

#### china

如果值为 `true`，则将模板文件中 `{{china_ip}}` 与 `{{china_sites}}` 替换为

```
"geoip:cn",
```

和

```
{
	"type": "field",
	"outboundTag": "direct",
    "domain": ["geosite:cn"] 
},
```

否则，则将配置文件中 `{{china_ip}}` 与 `{{china_sites}}` 均替换为空白

#### tls

如果值为 `tls`，则将模板文件中 `{{tls}}` 替换为

```
{
 		 "serverName": "{{address}}",
 		 "allowInsecure": {{allowInsecure}},
 		 "alpn": ["http/1.1"]
}
```

否则则将模板文件中 `{{tls}}` 替换为 `null`，

#### network

值为 `kcp`：将 `{{kcp}}` 替换为

```
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
}
```

值为 `ws`：将 `{{ws}}` 替换为

```
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
}
```

值为 `http`：将`{{http}}` 替换为

```
{
		"host": [{{host}}],
		"path": "{{path}}"
}
```

并将哈系表中 `host` 所对值由 `foo,bar,foobar` 修改为 `"foo","bar","foobar"`

值为 `quic`：将`{{quic}}` 替换为

```
{
		  "security": "{{host}}",
		  "key": "{{path}}",
		  "header": {
		    "type": "{{type}}"
		  }
}
```

否则将不做修改

### 其他

#### 计划

使用 `text/template` 代替现有方案

#### 默认模板文件

```json
{
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
  "outbounds": [ 
	{
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
      "quicSettings": {{quic}},
	  "mux": {
  		"enabled": {{mux}},
      	"concurrency": {{concurrency}}
      }
    }
  	},
    {
      "protocol": "freedom",
      "settings": {},
      "tag": "direct"
    }
],
  "dns": {
    "servers": [
      "{{dns1}}",
      "{{dns2}}",
      "localhost"
    ]
  },
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
}
```

## LICENSE

MIT LICENSE
