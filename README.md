# onebot-sdk
onebot-sdk
# 注意事项
- `*11.0.15*`之后（不包含）出现不兼容变动，变更仍在继续，预计持续1-2个月
## 配置变更
- 原配置
	```
	http:
	  schema: "http"
	  host: 127.0.0.1
	  port: 5700
	websocket:
	  host: 127.0.0.1
	  port: 6700
	SUPERUSERS: # 超级管理员账号
	 - 123456789
	```
- 新配置
	```
	api:
		type: websocket
		endpoint: "ws://127.0.0.1:6700/api"
		accessToken: ""
		secret: "" # 暂不支持
	event:
		type: websocket
		addr: "0.0.0.0:8080" # go-cqhttp配置为 ws-reverse api: ws://${程序启动ip}:8080/api
		accessToken: ""
		secret: "" # 支持不完善
	```
## 需要进行实例化
- API调用: 详情查看 [examples/api/main.go](./examples/api/main.go)
- EVENT回调: 详情查看 [examples/event/main.go](./examples/event/main.go)

# Wiki
介绍实现/命名规则等
* 地址：[Wiki](./wiki/)
# 待处理事项
- [x] http支持
- [x] 反向http支持
- [x] websocket支持
- [x] 反向websocket支持
- [x] access-token支持
- [x] 调用超时设置支持
	- [x] api调用可以配置超时时间，单位ms,默认值为10000
- [ ] 消息格式支持
	- [x] array支持
	- [ ] CQ码支持
- [ ] secret支持
- [ ] 支持[onebot-v12](https://12.onebot.dev/)
# 快速开始
## 引用
```
go get github.com/dezhishen/onebot-sdk@${tag}
```
`${tag}`请参考[releases](https://github.com/dezhishen/onebot-sdk/releases)中的版本号和对应的onebot协议
**11.x支持onebotv11协议，其他版本类似**
## 发送API
- API调用: 详情查看 [examples/api/main.go](./examples/api/main.go)
## 响应事件
WebSocket事件分为注册和启用两步
- EVENT回调: 详情查看 [examples/event/main.go](./examples/event/main.go)
# 功能点清单
## API
- <a href="https://docs.go-cqhttp.org/api/" target="_blank">go-cqhttp#api</a>
## EVENT
- <a href="https://docs.go-cqhttp.org/event/" target="_blank">go-cqhttp#event</a>
