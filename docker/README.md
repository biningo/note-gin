## 生产环境Docker部署教程

### 1、使用七牛图片存储

注意，要使用七牛图片存储的话需要在 `config/AppConfig.yaml` 加上下面两个字段 `QiniuAccessKey`  `QiniuSecretKey` ，当然为了不泄露自己的密钥，我这里就取消这两个配置，**不影响运行**

```yaml
PageSize: 13
MakeMigration: false
QiniuAccessKey: <you-key> #这两个
QiniuSecretKey: <you-key> #这个
LogFilePath: pkg/logging/log.log
JwtSecretKey: "note-gin"
```

<br>

## 2、关于Docker注意的点

axios在请求后端资源的时候，发出的http请求其实是在宿主机层面的，不是在容器层面的，所以请求需要发送到后端docker容器在宿主机射影的地址上

