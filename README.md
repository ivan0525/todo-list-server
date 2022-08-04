# todo-list-server

> 主要技术：`Gin`、 `GORM`、`MySql`、`Redis`

## 目录结构

```
todo-list-server
├── api             用于定义接口函数
├── cache           放置redis缓存
├── conf            用于存储配置文件
├── middleware      应用中间件
├── model           应用数据库模型
├── pkg
|    ├── e          封装错误码
|    └── util       工具函数
├── routes          路由逻辑处理   
├── serializer      将数据序列化为 json 的函数
└── server          接口函数的实现
```