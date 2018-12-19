## stealtoken 暴力破解数字币

#### 安装
go get github.com/learnergo/stealtoken

#### 运行
```
./setup up  启动
./setup down  停止
./setup restart  重启
```

#### 日志

```
stealtoken.err 错误日志
stealtoken.out debug 日志（如果开启debug。不建议开启，会爆）
stealtoken.pid 后台运行id
stealtoken.suc 成功日志（非常渺茫）
```

#### TODO
配置公私钥发现有token立即转移

#### 注意
程序实现每一个币种对应一个协程，是因为如果频繁访问第三方网址会被禁止访问。


