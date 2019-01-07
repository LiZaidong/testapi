# Test API Server

| 版本号 | 说明 | 更新日期 | 更新者 |
| ----- | --- | ---------- | ---- |
|  V1.0 | 初版 | 2018.12.30 | 杨飞 |

[TOC]

## 项目介绍

为了更好了整合本人的技术栈，也为了能够帮助到大家学习go语言，尽快掌握go语言的工程实践。

我通过一个简单的业务demo，也就是用户信息的增删查改，来实现一个golang后端API应用。

主要涵盖了下面这些技术点：
* 路由
* 过滤器
* 日志
* 静态配置文件
* 控制器
* 模型
* redis缓存
* MySQL数据库
* 多国语言处理

以后可能会增加更多的技术内容进来，比如etcd, nsq等内容，也欢迎大家一起来完善本项目，造福更多的新人朋友。

## 项目操作

### HTTP头部设置要求

| key | value | 必填 | 说明 |
| --- | ----- | ---- | --- |
| x-access-token | test | Y | token，鉴别用户权限。因为是demo, 所以这里是写死的。 |
| accept-language | zh-cn | N | 服务端有多语言支持的时候，会采用这里的语言返回数据。 |

### 新增用户

- 方法和路由

```
PUT 127.0.0.1:6666/test/api/v1/user
```

- 请求参数

| 参数   | 类型   | 必填  | 说明    |
| ----  | ------ | ---- | ------- |
| name  | String | N    | 用户姓名 |
| age   | Int    | N    | 用户年龄 |
| email | String | N    | 电子邮件 |

- 请求参数demo

```json
{
	"name": "test",
	"age": 18,
	"email": "test@163.com"
}
```

- 响应

|   参数   | 类型   | 必填  | 说明    |
|   ----  | ------ | ---- | ------- |
| code    | String | Y    | 用户姓名 |
| message | Int    | Y    | 用户年龄 |

- 响应demo

```json
{
  "code": 0,
  "message": "用户添加成功!"
}
```

### 查询用户

- 方法和路由

```
GET 127.0.0.1:6666/test/api/v1/user?id=1
```

- 请求参数

| 参数 | 类型 | 必填  | 说明   |
| --- | ---- | ---- | -----  |
| id  |  int |  Y   | 数据库中生成的用户ID字段 |

- 响应

| 参数   | 类型   | 必填  | 说明    |
| ----  | ------ | ---- | ------  |
| id    | Int    | Y    | 用户ID  |
| name  | String | Y    | 用户姓名 |
| age   | Int    | Y    | 用户年龄 |
| email | String | Y    | 电子邮件 |

- 响应demo

```json
{
  "id": 1,
  "name": "test",
  "age": 18,
  "email": "test@163.com"
}
```

### 查询用户, 分页

- 方法和路由

```
GET 127.0.0.1:6666/test/api/v1/users?page_num=1&page_size=5
```

- 请求参数

| 参数 | 类型 | 必填  | 说明   |
| --- | ---- | ---- | -----  |
| page_num  |  int |  Y   | 第几页 |
| page_size  |  int |  Y   | 每页有多少条记录 |

- 响应

| 参数        | 类型   | 必填  | 说明    |
| ---------  | ------ | ---- | ------  |
| total_num  | Int    | Y    | 一共返回了多少条记录  |
| list       | []User | Y    | 返回用户列表，如果没有记录返回 null |

- 响应demo

```json
{
  "total_num": 0,
  "list": null
}
```

### 修改用户

- 方法和路由

```
POST 127.0.0.1:6666/test/api/v1/user
```

- 请求参数

| 参数   | 类型   | 必填  | 说明    |
| ----  | ------ | ---- | ------- |
| id    | Int    | Y    | 用户ID  |
| name  | String | N    | 用户姓名 |
| age   | Int    | N    | 用户年龄 |
| email | String | N    | 电子邮件 |

- 请求参数demo

```json
{
  "id": 1,
  "name": "fei",
  "age": 26
}
```

- 响应

|   参数   | 类型   | 必填  | 说明    |
|   ----  | ------ | ---- | ------- |
| code    | String | Y    | 用户姓名 |
| message | Int    | Y    | 用户年龄 |

- 响应demo

```json
{
  "code": 0,
  "message": "用户更新成功!"
}
```

### 删除用户

- 方法和路由

```
DELETE 127.0.0.1:6666/test/api/v1/user
```

- 请求参数

| 参数   | 类型   | 必填  | 说明   |
| ----  | ------ | ---- | ------ |
| id    | Int    | Y    | 用户ID |

- 请求参数demo

```json
{
  "id": 1
}
```

- 响应

|   参数   | 类型   | 必填  | 说明    |
|   ----  | ------ | ---- | ------- |
| code    | String | Y    | 用户姓名 |
| message | Int    | Y    | 用户年龄 |

- 响应demo

```json
{
  "code": 0,
  "message": "用户删除成功!"
}
```

## 其他

更多内容，陆续添加中，敬请期待。。。
