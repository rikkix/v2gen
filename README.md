# v2gen

A powerful V2Ray config generator

You can use use vmess ping instead of ICMP ping

![GitHub top language](https://img.shields.io/github/languages/top/iochen/v2gen) ![Go](https://github.com/iochen/v2gen/workflows/Go/badge.svg) ![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/iochen/v2gen) 

[简体中文](README_zh_cn.md)

## Preview
```
[ 0] 中继香港C1 Media (HK)(1)      [518ms  (0 errors)]
[ 1] 中继香港C3 Media (HK)(1)      [527ms  (0 errors)]
[ 2] 中继香港C2 Media (HK)(1)      [536ms  (0 errors)]
[ 3] 中继香港C5 Media (HK)(1)      [451ms  (0 errors)]
[ 4] 中继香港C6 Media (HK)(1)      [452ms  (0 errors)]
[ 5] 中继香港G2 Media (HK)(1)      [904ms  (0 errors)]
[ 6] BGP中继香港 2 Media (HK)(1)   [468ms  (0 errors)]
[ 7] BGP中继香港 3 Media (HK)(1)   [778ms  (0 errors)]
[ 8] BGP中继香港 1 Media (HK)(1)   [881ms  (0 errors)]
[ 9] 中继香港G1 Media (HK)(1)      [1.35s  (1 errors)]
...
[50] 日本中继 3 Media (JP)(1)      [641ms  (0 errors)]
=====================
Please Select:
```

## How to use

Build it first
```sh
go get -u iochen.com/v2gen/cmd
```
or Download in GitHub Releases  
  
Then run

```sh
v2gen -u {{Your subscription link}} -o {{Your V2Ray config path}}
```

## Param

```Param
Usage of v2gen:
  -c string
        v2gen config path (default "/etc/v2ray/v2gen.ini")
  -ct int
        ping count for each node (default 3)
  -dest string
        test destination url (vmess ping only) (default "https://cloudflare.com/cdn-cgi/trace")
  -eto int
        timeout seconds for each request (default 8)
  -init
        initialize v2gen config
  -med
        use median instead of ArithmeticMean
  -n int
        node index (default -1)
  -np
        do not ping
  -o string
        output path (default "/etc/v2ray/config.json")
  -r    
        random node index
  -t    
        use ICMP ping instead of vmess ping
  -tpl string
        V2Ray tpl path
  -tto int
        timeout seconds for each node (default 25)
  -u string
        subscription URL
  -v    
        show version
  -vmess string
        vmess link(s)
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

## Version

<!--version-->
*V1.3.5*
<!--version-end-->
In case of any discrepancy or inconsistency between Version Number here and Tag, Tag shall prevail.

## LINCENSE

MIT LICENSE

## Notice

no support for version 1 format
