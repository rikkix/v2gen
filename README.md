# v2gen

A powerful V2Ray config generator

You can use use vmess ping instead of ICMP ping

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?)](https://pkg.go.dev/iochen.com/v2gen)
![GitHub top language](https://img.shields.io/github/languages/top/iochen/v2gen)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/iochen/v2gen) 
![Go](https://github.com/iochen/v2gen/workflows/Test/badge.svg) 


[简体中文](README_zh_cn.md)

## Preview
```
[ 0] Server A      [451ms  (0 errors)]
[ 1] Server B      [452ms  (0 errors)]
[ 3] Server C      [251ms  (0 errors)]
...
[25] Server Z      [652ms  (2 errors)]
=====================
Please Select:
```

## Build or Download
```sh
git clone https://github.com/iochen/v2gen/ && cd v2gen
env GOPRIVATE=github.com/v2ray/v2ray-core go build ./cmd/v2gen
```
or Download in GitHub Releases  

## Quick start
```sh
v2gen -u {{Your subscription link}} -o {{Your V2Ray config path}}
```

## Param

```Param
Usage of ./v2gen:
  -best
        use best node judged by ping result
  -c int
        ping count for each node (default 3)
  -config string
        v2gen config path (default "/etc/v2ray/v2gen.ini")
  -dst string
        test destination url (vmess ping only) (default "https://cloudflare.com/cdn-cgi/trace")
  -init
        init v2gen config (specify certain path with -config)
  -log string
        log output file (default "-")
  -loglevel string
        log level (default "warn")
  -o string
        output path (default "/etc/v2ray/config.json")
  -ping
        ping nodes (default true)
  -pipe
        read from pipe (default true)
  -random
        random node index
  -template string
        V2Ray template path
  -thread int
        threads used when pinging (default 3)
  -u string
        subscription address(URL)
  -v    show version
```

## V2Gen user config

You can use `v2gen --init` to generate one

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

The following config may NOT work on every node

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

## LINCENSE

MIT LICENSE
