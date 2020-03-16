# V2Gen

V2Ray配置文件生成器

*注：我们建议参阅英文资料*  
[English](README.md)

## 预览

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


## 如何使用

先编译它（请确保您的`$GOPATH/bin/`已添加至`$PATH`中）

```sh
go get -u iochen.com/v2gen/cmd
```
或到 GitHub Releases 中下载    
  
然后运行

```sh
v2gen -u {{你的订阅链接}} -o {{你V2Ray的配置文件路径}}
```

## 参数

```Usage
Usage of v2gen:
  -c string
        v2gen 配置文件路径 (default "/etc/v2ray/v2gen.ini")
  -ct int
        ping 次数 (default 3)
  -dest string
        测试链接 (vmess ping only) (default "https://cloudflare.com/cdn-cgi/trace")
  -eto int
        单个超时时间 (default 8)
  -init
        初始化 v2gen 配置
  -med
        使用中位数而不是算术平均 
  -n int
        节点引索 (default -1)
  -np
        别 ping
  -o string
        输出路径 (default "/etc/v2ray/config.json")
  -r    
        随机节点引索
  -t    
        使用 ICMP ping 而不是 vmess ping
  -tpl string
        V2Ray 模板路径
  -tto int
        单个节点超时时间 (default 25)
  -u string
        订阅链接
  -v    
        展示版本
  -vmess string
        vmess 链接（们）
```

## V2Gen 用户配置

你可以使用 `v2gen --init` 来生成一个新的

```yaml
# V2Ray 日志等级
# ( debug | info | warning | error | none )
loglevel: warning

# Socks 端口
socksPort: 1080

# Http 端口
httpPort: 1081

# 是否允许UDP流量
# ( true | false )
udp: true

# 安全
# ( aes-128-gcm | aes-256-gcm | chacha20-poly1305 | auto | none )
security: aes-256-gcm

# 是否开启 mux
# ( true | false )
mux: true

# Mux 并发数
concurrency: 8

# DNS 服务器
dns1: 9.9.9.9
dns2: 1.1.1.1

# 中国IP与网站是否直连
# ( true | false )
china: true

```

下面的配置可能不会在所有节点上生效

```yaml
# 是否允许不安全连接 ( true | false )
allowInsecure: false

# KCP mtu 值
mtu: 1350

# KCP tti 值
tti: 20

# KCP 最大上行速度
# 单位: MB/s
up: 5

# KCP 最大下行速度
# 单位: MB/s
down: 20

# 是否开启 UDP 拥堵控制 ( true | false )
congestion: false

# 读缓冲区大小
# 单位: MB
readBufferSize: 1

# 写缓冲区大小
# 单位: MB
writeBufferSize: 1
```

## 版本

*V1.3.1*

## 协议

MIT LICENSE

## 注意

不支持第一版本格式
