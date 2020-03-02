# v2gen

V2Ray config generator

[简体中文](README_zh_cn.md)

## Preview
```
[0]     中继香港G1 Media (HK)(1)        [80.01ms (cu2.example.com)]
...
[12]    BGP中继香港 5 Media (HK)(1)     [43.22ms (hk5.example.com)]
[13]    BGP中继香港 6 Media (HK)(1)     [22.62ms (hk6.example.com)]
[14]    BGP中继香港 7 Media (HK)(1)     [20.01ms (hk7.example.com)]
[15]    BGP中继香港 8 Media (HK)(1)     [19.71ms (hk8.example.com)]
[16]    BGP中继香港 9 Media (HK)(1)     [19.73ms (hk9.example.com)]
[17]    BGP中继香港 10 Media (HK)(1)    [37.28ms (hk10.example.com)]
[18]    BGP中继香港 11 Media (HK)(1)    [43.71ms (hk11.example.com)]
...
[47]    美国GIA 3 Media (US)(0.7)       [192.90ms (us3.example.com)]
[48]    香港 8 (HK)(1)                  [64.42ms (hk9.example.com)]
[49]    香港 9 (HK)(1)                  [72.44ms (hk10.example.com)]
[50]    香港负载均衡 1 Test (HK)(1)      [18.62ms (93.184.216.34)]
=====================
Please Select:
```

## How to use

Build it first

```sh
go get -u github.com/iochen/v2gen/cmd
```
  
Then run

```sh
v2gen -u {{Your subscription link}} -o {{Your V2Ray config path}}
```

## Param

```Param
Usage of v2gen:
  -c string
        v2gen config path (default "/etc/v2ray/v2gen.ini")
  -init
        initialize v2gen config
  -n int
        node index (default -1)
  -np
        do not ping
  -o string
        output path (default "/etc/v2ray/config.json")
  -r    random node index
  -tpl string
        V2Ray tpl path
  -u string
        subscription URL
  -v    show version
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
dns1: 9.9.9.9
dns2: 1.1.1.1

# If China sites and ips directly connect
# ( true | false )
china: true

```

The following config may NOT work on every node

```yaml
# If allow insecure connection ( true | false )
allowInsecure false

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

*V1.1.2*

## LINCENSE

MIT LICENSE

## Notice

no support for version 1 format
