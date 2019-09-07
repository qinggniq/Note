# 新的系统开始一些必要的步骤
## 动机
领了一台新机器，从装Ubuntu系统到系统真正能够提供使用，如基本的翻墙功能、Git免密登录到输入法的设置，每次装的时候都需要在网上各个步骤所需要的软件命令，十分麻烦，所以记录下来以供下次新装系统使用。
## 大致内容
1. 翻墙
2. 给GitHub添加公钥
3. 中文输入法的设置
## 翻墙
### Chrome的安装
过程略。
### shadowsocks的安装
1. 可以先装一个`python` ，然后使用`sudo pip install shadowsocks`安装，然而`pip`里面的**shadowsocks**是有问题的，所以可以使用第二种方法。
2. 使用`sudo apt install shadowsocks`安装ubuntu里面默认软件源的包。
### SwitchOmega插件安装
1. 用于在Chrome中设置代理的插件。
2. 在Chrome Store里面可以安装，然而没有翻墙前是访问不了Chrome Store的，然而翻墙又需要这个插件，所以陷入了**鸡生蛋，蛋生鸡**的困局，应该离线下载这个插件。
3. https://github.com/FelisCatus/SwitchyOmega/releases
4. 安装完成，打开Chrome的插件页，开启**开发者模式**，然后`Ctrl-R`，将下载的插件拖到页面上，Chrome会询问是否安装。
### SwitchOmega配置
1. 在**情景模式**proxy改成如下格式
| 代理协议| 代理服务器 | 代理端口 |
| :--:  |   :--:|   :--:  |   :--:|   
| SOCKS5| 127.0.0.1|1080|
2. 规则列表输入网址[](https://raw.githubusercontent.com/calfzhou/autoproxy-gfwlist/trunk/gfwlist.txt)。
### 本机启动shadowsocks
1. 命令行输入`sslocal  -d start -p {server_port} -k {password} - s {server_ip} -m {encrypt method} - l {local_port}`即可翻墙。
## 给GitHub添加公钥
### 配置Git
```shell
$ cd work_space
$ git config --global user.name "name"
$ git config --global user.email "email"
```
### 产生公钥
```shell
$ ssh-keygen
$ cat $HOME/.ssh/idsa.pub
```
### 打开GitHub的设置页
将公钥粘贴到GitHub的`SSH KEY`列表里面。

## 中文输入法
### 下载Sogou输入法
`dpkg`安装即可。
### 系统/语言支持
1. 将中文提到高亮处。
2. **键盘输入法系统**更改为**fcitx**。
3. 重启机器。
### fcitx配置
1. 在坐下角的**+** 号找到**Sogou输入法**。

