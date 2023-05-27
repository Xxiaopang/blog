---
title: blog v1.0.0
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
generator: "@tarslib/widdershins v4.0.15"

---

# blog

> v1.0.0

# api

## GET 验证码

GET /api/captcha

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 主页

GET /api/home

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# api/user

## POST 用户注册

POST /api/register

> Body 请求参数

```json
{
  "username": "admin",
  "password": "123456",
  "mobile": "11122233345",
  "email": "123456@qq.com"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» username|body|string| 是 |none|
|» password|body|string| 是 |none|
|» mobile|body|string| 是 |none|
|» email|body|string| 是 |none|

> 返回示例

> 成功

```json
{
  "code": 1,
  "data": null,
  "msg": "注册成功"
}
```

```json
{
  "code": 1,
  "data": null,
  "msg": "用户已经存在"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|null|true|none||none|
|» msg|string|true|none||none|

## POST 用户登录

POST /api/login

> Body 请求参数

```json
{
  "password": "123456",
  "mobile": "12345678910",
  "captcha_id": "e1qHyUcNR4yCyYeoFoq6",
  "captcha": "z93n"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» mobile|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "data": "string",
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» data|string|true|none||none|
|» msg|string|true|none||none|

# api/articles

## GET 查询文章

GET /api/articles/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 1,
  "data": null,
  "msg": "查询失败"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE delete

DELETE /api/articles/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT update

PUT /api/articles/{id}

> Body 请求参数

```json
{
  "title": "第一本书",
  "content": "易经",
  "tags": [
    {
      "name": "山"
    },
    {
      "name": "命"
    }
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|string| 是 |none|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» content|body|string| 是 |none|
|» tags|body|[object]| 是 |none|
|»» name|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 文章列表

GET /api/articles/

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|limit|query|integer| 否 |none|
|page|query|integer| 否 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 成功

```json
{
  "code": 0,
  "data": null,
  "msg": "Invalid token!"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 添加文章

POST /api/articles

> Body 请求参数

```json
{
  "title": "第一本书",
  "content": "易经",
  "tags": [
    {
      "name": "山"
    },
    {
      "name": "命"
    }
  ]
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|
|» title|body|string| 是 |none|
|» content|body|string| 是 |none|
|» tags|body|[object]| 是 |none|
|»» name|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET tag

GET /api/article/

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|tag|query|string| 否 |none|
|page|query|integer| 否 |none|
|limit|query|integer| 否 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# api/tags

## GET list

GET /api/tags/

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST ADD

POST /api/tags/

> Body 请求参数

```
|-
{
    "name": "山"
}

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# api/file

## POST upload_file

POST /api/file/upload

> Body 请求参数

```yaml
file: file://D:\chrome_file\天下第一相书 云谷山人.pdf

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|
|body|body|object| 否 |none|
|» file|body|string(binary)| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET download_file

GET /api/file/download/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |none|
|Authorization|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET list

GET /api/file/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|page|query|integer| 否 |none|
|limit|query|integer| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

