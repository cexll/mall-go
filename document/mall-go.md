---
title: mall-go v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.11"

---

# mall-go

> v1.0.0

# Default

## GET Hello

GET /hello

> 返回示例

> 成功

```json
{
  "data": null,
  "message": "hello, world!",
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 用户

## GET 获取用户信息

GET /user/get

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|query|string| 否 |none|

> 返回示例

> 成功

```json
{
  "data": {
    "id": 1,
    "nickname": "cexll",
    "avatar_url": "1",
    "password": "$2a$10$Fv9FHT6/3CWuixswVHmknejzif1lPu57u2.sYWtosMlHHbTE85ay2",
    "mobile": "17772399000",
    "signature": "",
    "status": 0,
    "is_delete": 0,
    "created_at": "2022-07-02T12:27:10+08:00",
    "updated_at": "2022-07-02T12:27:10+08:00"
  },
  "message": "message",
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 增加用户

POST /user/create

> 返回示例

> 成功

```json
{
  "data": 9,
  "message": "ok",
  "status": 200
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

