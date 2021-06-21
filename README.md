#整体框架
```
使用gin框架实现一下功能的服务
（1）管理后台调用 - 创建礼品码
（2）管理后台调用 - 查询礼品码信息
（3）客户端调用 - 验证礼品码
```

#目录结构
```
├── README.md                   #介绍
├── config                      #配置redis
│   └── redis.go
├── ctrl                        #处理router的请求
│   └── routerCtrl.go
├── module                      
│   └── gift.go           #保存gift的数据结构
│   └── redisSer.go       #操作redis
├── gift_test.go                #单元测试
├── go.mod
├── handle                      #处理具体业务逻辑
│   └── handle.go
├── main.go
├── router                      #路由转发       
│   └── router.go
└── util                        #工具方法 生成随机礼品码
    └── getRandCode.go
```
#代码逻辑分层  gift-Redis
| 层     | 文件夹|主要职责 |调用关系|
| :----: | :----|:---- | :-----|
|router  | /router|路由转发 |调用ctrl|
|ctrl    | /ctrl  | 请求参数验证，处理请求后构造回复消息|调用handle|
|handle  | /handle|处理具体的业务逻辑 |调用module util|
|module  | /gift |数据模型 操作数据库 |被handle调用|
|util    | /util | 通用工具 | 被handle调用|


#存储设计

##gift
| 内容 | 类型 |  说明 |
| :----:| :---- | :----|
|GiftCode |string |礼包码
|Description |string|礼包描述
|GiftNum |string	|可领取数
|ValidPeriod |string|有效期
|GiftContent| string | 礼包内容
|CreatePer |string  |创建用户
|CreatTime| string  |创建时间
|AvailableNum| string |已领取次数
|ClaimList |string   |领取列表

#接口设计
##1。创建礼品码

```
接口地址 
/giftCode/adminCreatGiftcode 
```
### 请求方式
GET
###请求示例
```
http://127.0.0.1:8080/giftCode/adminCreatGiftcode?des=description&GN=3&VP=5&GC=G:100:D:20&CP=llq
```
### 参数  说明

``` 
des 类型string Description描述礼包的字段
```
``` 
GN 类型string GiftNum礼包的可领取数
```
``` 
VP 类型string ValidPeriod礼包的有效时间
```
``` 
GC 类型string GiftContent 礼包的内容
```
``` 
CP 类型string GiftContent CreatePer礼包的创建者
```

```
成功示例 
{
    "GiftCode": "r1czr5u2",
    "condition": "success"
}

错误示列 
{
    "GiftCode": {},
    "condition": "error"
}
```

##2。查询礼品码
```
接口地址 
/giftCode/admininquireGiftCode 
```
### 请求方式
GET
### 请求示例
```
http://127.0.0.1:8080/giftCode/admininquireGiftCode?giftCode=r1czr5u2
```

### 参数  说明

``` 
giftCode 类型string 此字段为需要查询的礼包码
```


```
成功示例 
{
    "condition": "success",
    "data": {
        "AvailableNum": "0",
        "ClaimList": "",
        "CreatTime": "2021-06-09 18:47:07",
        "CreatePer": "llq",
        "Description": "description",
        "GiftContent": "10k",
        "GiftNum": "3",
        "ValidPeriod": "5",
        "giftCode": "r1czr5u2"
    }
}
错误示列 
{
    "condition": "error",
    "giftCode": "GiftCode is error"
}
```

##3。验证礼品码

```
接口地址 
/giftCode/client 
```
### 请求方式
GET
### 请求示例
```
http://127.0.0.1:8080/giftCode/client?giftCode=r1czr5u2&usr=nna
```

### 参数  说明

``` 
giftCode 类型string 此字段为需要验证的礼包码
```
``` 
usr 类型string 此字段用来输入用户名称
```

```
成功示例 
{
    "GiftContent": "10k",
    "condition": "success"
}
错误示列 
{
    "GiftCode": "input usr",
    "condition": "error"
}
```

#第三方库
## redis
```
用于  操作redis 
代码 https://github.com/go-redis/redis
```

## gin
```
用于  开发服务器框架
代码  https://github.com/gin-gonic/gin 
```
