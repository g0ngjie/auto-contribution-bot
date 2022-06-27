# 一个自动更新贡献度的机器人

:alembic: 此项目纯属自娱自乐，目前还处于实验阶段

通过定时任务方式，可用在 Nas、软路由、服务器或者个人 PC 上

[文件下载](https://github.com/g0ngjie/auto-contribution-bot/releases/tag/v1.0.0)

## conf.toml

与执行文件放到同级目录下

```toml
# 是否新增文件，默认为false
# 非新增，则在执行文件中添加新的行
new_file = false
# 文件类型，默认为 md （值需要设置目标文件的后缀名）
file_type = "md"
# 内容来源
# 可选值：0：默认读取new_content 1：一言
content_from = 1
# 新增内容，默认为 YYYY-MM-DD
new_content = ""
# git仓库地址
# 注* https 方式
git_url = "https://gitee.com/g0ngjie/test-auto-contribution-bot.git"
# git仓库登录用户名
git_user = "<username>"
# git仓库登录密码
git_pass = "<password>"
# git邮箱，如果登录用户名为邮箱登录，则一致
# 邮箱如果不填写，贡献度将无法更新
git_email = "<email>"
```
