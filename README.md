## Summary
在看了好几天kratos文档以及代码示例example后，并且看了看官方的示例beer-shop 项目的代码后，基本了解了kratos项目该怎么写，于是我也终于开始施工我的第一个称得上项目的微服务项目，我暂定其名为学校管理系统，并暂定其由用户（学生，教师，管理员）管理微服务，选课成绩管理微服务，机构设置微服务组成，后续其他微服务等我想到了再网上加.
**项目地址为：<https://github.com/kasoushu/education-micro>**


## 暂定要使用的技术
|技术|说明| 链接 |
|---|---|----|
|kratos|微服务框架| <https://go-kratos.dev/docs/> |
|protobuf|生成grpc| <https://developers.google.com/protocol-buffers/docs/downloads> 
|jwt|用户验证| <https://github.com/dgrijalva/jwt-go> |
|grpc|远程调用| <https://grpc.io/docs/languages/go/quickstart/> |
|jaeger|链路追踪| <https://pkg.go.dev/go.opentelemetry.io/otel/exporters/jaeger> |
|consul|服务注册，服务发现| <https://github.com/golang-jwt/jwt>|
|ent|orm框架 | <https://github.com/ent/ent> |
|gorm|orm框架 | <https://github.com/go-gorm/gorm> |

## 目前进度
|step|status|
|---|--|
|用户管理微服务|已实现|
|选课成绩管理微服务|已实现|
|机构管理微服务|未施工|