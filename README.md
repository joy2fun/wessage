
# wessage

  通过命令行发送微信公众号模板消息给指定微信用户。

  [下载地址](https://github.com/joy2fun/wessage/releases)

## 准备测试账号

 [开通测试账号](https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login) ，登录测试账号管理后台，扫描测试公众号二维码并关注，获得微信号 open id。

 新增测试模板，模板标题随意，模板内容填写：`{{msg.DATA}}`

## 基本配置

```sh
# 配置公众号 appid & secret
wessage config set wx_appid {your appid}
wessage config set wx_secret {your secret}

# 可选配置：默认模板 id 和接收者微信号
wessage config set wx_template_id {模板 id}
wessage config set wx_receiver {微信号 open id}
```

## 发送消息

```sh
wessage send {content}

# 默认会使用配置的模板 id 和接收者微信号 open id
# 也可以使用命令行选项覆盖默认配置
wessage send {content} -t {模板 id} -r {微信号 open id}
```

更多命令和选项：

```sh
# 查看公众号已有模板列表
wessage template
# 查看关注用户列表
wessage user
# 查看配置信息
wessage config list
# 默认配置文件为 $HOME/.wessage.json，可以通过选项 --config 来自指定配置文件，如：
wessage config list --config /path/to/config.json
# 更多帮助信息
wessage --help
```
