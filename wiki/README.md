# WIKI
项目WIKI
## 命名规则
* struct与func 尽可能采用 驼峰命名
* json 均采用 `_`+**小写**的命名
### Message
* MessageSegment代表消息段，使用数组
* 不同的消息类型 使用 MessageElement+**消息类型**
    * 如：MessageElementAt MessageElementText
* 内部属性均使用对应的`json`转**首字母驼峰**
### API
API跟随oneBot11定义的地址，转为方法时，方法名为首字母驼峰
### Event
### Protobuf/GRPC
* 使用 **原来的名称**+`GRPC` 作为grpc对象 如 MessageElementAtGRPC
* 增加转换方法
```
    func (m *MessageElementAtGRPC) ToStruct()* MessageElementAt
    ...
    func (m *MessageElementAt) ToGRPC()*MessageElementAtGRPC
```
## MessageElement json序列化说明
### 目标
* 使用无额外成本
* 对修改关闭，对拓展开放
### 实现
#### MessageElement
* 所有的MessageElement需要注册`messageElementUnmarshalJSON`
    * 如 `MessageElementAt`
    ```
    func init() {
        messageElementUnmarshalJSONMap["at"] = func(data []byte) (MessageElement, error) {
            var result MessageElementAt
            err := json.Unmarshal(data, &result)
            return &result, err
        }
    }
    ```
* 实现接口`MessageElement`的`Type()string`方法
```
    func (msg MessageElementAt) Type() string {
        return "at"
    }
    ```
    * 重写`MessageSegment`的``方法：
    ```

    var messageElementUnmarshalJSONMap = make(map[string]func(data []byte) (MessageElement, error))

    func (msgSeg *MessageSegment) UnmarshalJSON(data []byte) error {
        if len(data) == 0 || data[0] == 'n' { // copied from the Q, can be improved
            return nil
        }
        messageType := gjson.GetBytes(data, "type").Str
        decoder, ok := messageElementUnmarshalJSONMap[messageType]
        if !ok {
            return fmt.Errorf("未找到指定的消息类型,%v", messageType)
        }
        element, err := decoder([]byte(gjson.GetBytes(data, "data").Raw))
        if !ok {
            return fmt.Errorf("未找到指定的消息类型,%v", messageType)
        }
        msgSeg.Type = messageType
        msgSeg.Data = element
        return err
    }

    type MessageElement interface {
        Type() string
    }

```

### 优缺点
* 优：新增消息类型时,只需要按照要求，实现`func Type()string`和注册`messageElementUnmarshalJSON`
* 缺：性能上有一定损耗
