# V2Gen
Generate V2Ray Json format from "vmess://{{base 64 encoded}}" URI format.

## How to use
Build it first
```
git clone https://github.com/iochen/V2Gen.git
cd ./V2Gen
go build -o ./v2gen
``` 
Then run 
```
./v2gen -u {{Your subscription link}} -p {{Your V2Ray config path}}
```

## Param
```
Usage of ./v2gen:
  -c string
        user config path (default "/etc/v2ray/v2gen.ini")
  -p string
        config output path (default "/etc/v2ray/config.json")
  -u string (necessary)
        The URL to get configs from
```

## Version
*V0.1.3*

## LINCENSE
MIT LICENSE
