# onebot-sdk
onebot-sdk
# Wiki
介绍实现/命名规则等
* 地址：[Wiki](./wiki/)
# 待处理事项
- [x] model 增加protobuf
- [ ] 增加普通对象与GRPC对象转换方法
# 快速开始
## 发送API
直接调用[./pkg/api/](./pkg/api/)下面的方法,例如
```
package test

import (
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/dezhishen/onebot-sdk/pkg/api"
)

func do(){
	msg := model.PrivateMsg{
		UserID: userId,
		Message: []*model.MessageSegment{
			{
				Type: "text",
				Data: &model.MessageElementText{
					Text: "测试消息",
				},
			},
		},
	}
	got, err := api.SendPrivateMsg(&msg)
	if err != nil {
		t.Errorf("SendPrivate() error = %v", err)
		return
	}
	print(got)
}
```

## 响应事件
WebSocket事件分为注册和启用两步,
### 注册
直接调用[./pkg/event/](./pkg/event/)下面的`ListenXxxXxx`方法
```
import (

	log "github.com/sirupsen/logrus"

	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/dezhishen/onebot-sdk/pkg/event"
)

func TestListenMessagePrivate() {
    //注册第一个处理器
	event.ListenMessagePrivate(func(data model.EventMessagePrivate) error {
		log.Printf("1:%v", data.Message)
		return nil
	})
    //注册第二个
	event.ListenMessagePrivate(func(data model.EventMessagePrivate) error {
		log.Printf("2:%v", data.Message)
		return nil
	})
    //会依次执行注册的处理器
    ...	
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	go func() {
		time.Sleep(time.Second * 5)
		//模拟退出
		log.Info("模拟退出...")
		cancel()
	}()
	defer cancel()
	event.StartWsWithContext(ctx)
}

```
### 开启

直接调用[./pkg/event/](./pkg/event/)下面的`StartWsWithContext(ctx context.Context)`方法

```
import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/dezhishen/onebot-sdk/pkg/model"
	"github.com/dezhishen/onebot-sdk/pkg/event"
)

func TestListenMessagePrivate() {
	...
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	go func() {
		time.Sleep(time.Second * 5)
		//模拟退出
		log.Info("模拟退出...")
		cancel()
	}()
	defer cancel()
	event.StartWsWithContext(ctx)
}

```

# 功能点清单
## API
### 消息
- [x] `send_private_msg` 发送私聊消息
- [x] `send_group_msg` 发送群消息
- [x] `send_msg` 发送消息
- [x] `delete_msg` 撤回消息
- [x] `get_msg` 获取消息
- [x] `get_forward_msg` 获取合并转发消息

### 群组
- [x] `set_group_kick` 群组踢人
- [x] `set_group_ban` 群组单人禁言
- [x] `set_group_anonymous_ban` 群组匿名用户禁言
- [x] `set_group_whole_ban` 群组全员禁言
- [x] `set_group_admin` 群组设置管理员
- [x] `set_group_anonymous` 群组匿名
- [x] `set_group_card` 设置群名片（群备注）
- [x] `set_group_name` 设置群名
- [x] `set_group_leave` 退出群组
- [x] `set_group_special_title` 设置群组专属头衔
- [x] `set_group_add_request` 处理加群请求／邀请
- [x] `get_group_info` 获取群信息
- [x] `get_group_list` 获取群列表
- [x] `get_group_member_info` 获取群成员信息
- [x] `get_group_member_list` 获取群成员列表
- [x] `get_group_honor_info` 获取群荣誉信息

### 好友
- [x] `send_like` 发送好友赞
- [x] `set_friend_add_request` 处理加好友请求
- [x] `get_friend_list` 获取好友列表

### 账号
- [x] `get_login_info` 获取登录号信息
- [x] `get_stranger_info` 获取陌生人信息
- [x] `get_cookies` 获取 Cookies
- [x] `get_csrf_token` 获取 CSRF Token 
- [x] `get_credentials` 获取 QQ 相关接口凭证
- [x] `get_record` 获取语音
- [x] `get_image` 获取图片

### 系统
- [x] `get_status` 获取运行状态
- [x] `get_version_info` 获取版本信息
- [x] `set_restart` 重启 OneBot 实现
- [x] `clean_cache` 清理缓存
- [x] `can_send_image` 检查是否可以发送图片
- [x] `can_send_record` 检查是否可以发送语音

## 事件
### 监听命令
- [x] 开始监听
- [x] 消息事件适配器
- [x] 通知事件适配器
- [x] 请求事件适配器
- [x] 元事件适配器

### 消息事件注册
- [x] 群聊消息
- [x] 私聊消息
### 通知事件注册
- [x] 群文件上传
- [x] 群管理员变动
- [x] 群成员减少
- [x] 群成员增加
- [x] 群禁言
- [x] 好友添加
- [x] 群消息撤回
- [x] 好友消息撤回
- [x] 群内戳一戳
- [x] 群红包运气王
- [x] 群成员荣誉变更
### 请求事件注册
- [x] 加好友请求
- [x] 加群请求／邀请
### 元事件注册
- [x] 生命周期
- [x] 心跳
- [x] 相关配置
