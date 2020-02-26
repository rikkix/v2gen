# V2Gen

从 `vmess://` 格式中生成 V2Ray json 文件

*注：我们建议参阅英文资料*  
[English](README.md)

## 如何使用

先编译它

```sh
git clone https://github.com/iochen/V2Gen.git
cd ./V2Gen
go build -o ./v2gen
```
  
然后运行

```sh
./v2gen -u {{你的订阅链接}} -p {{你V2Ray的配置文件路径}}
```

## 参数

```Usage
Usage of v2set:
-c string
    V2Gen config path (default "/etc/v2ray/v2gen.ini")
-init
    if initialize V2Gen config
-n int
    Choose node (auto add -y param) (default -1)
-noPing
    disable ping function
-p string
    V2Ray json config output path (default "/etc/v2ray/config.json")
-r    select nodes at random
-test
    only for test
-tpl string
    v2ray json tpl file path
-u string
    The URL to get nodes info from
-v    version
-vmess string
    vmess://foo or vmess://foo;vmess://bar
-y    select "yes" when asking if preview config
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

*V1.0.1*

## 协议

MIT LICENSE

## 注意

不支持第一版本格式
