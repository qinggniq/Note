# lsof
## 全称
**list open files**
## 详细
**lsof**（list open files）是一个列出当前系统打开文件的工具。由于在Linux系统里面*万物皆为文件*，所以只要实现了`open(), read(), write(), close()`接口东西都可以被列出来，比如`socket`。
## 功能
1. 查看进程**打开的文件**。
2. 打开文件的进程。
3. 进程打开的端口。
4. 找回/恢复删除的文件。

打开的文件可以是
1. 普通文件
2. 目录
3. 网络文件系统的文件
4. 字符或设备文件
5. (函数)共享库
6. 管道，命名管道
7. 符号链接
8. 网络文件（例如：NFS file、网络socket，unix域名socket）
9. 还有其它类型的文件，等等
## 命令参数
```shell
-a          列出打开文件存在的进程
-c          <进程名> 列出指定进程所打开的文件
-g          列出GID号进程详情
-d          <文件号> 列出占用该文件号的进程
+d          <目录>  列出目录下被打开的文件
+D          <目录>  递归列出目录下被打开的文件
-n          <目录>  列出使用NFS的文件
-i          <条件>  列出符合条件的进程。（4、6、协议、:端口、 @ip ）
-p          <进程号> 列出指定进程号所打开的文件
-u          列出UID号进程详情
-h          显示帮助信息
-v          显示版本信息
```

|COMMAND| PID  | USER| FD | TYPE | DEVICE|SIZE|NODE|NAME|
|:---:|:---:| :---:|:---:| :---:|:---:| :---:|:---:| :---:|:---:| :---:|:---:| 
|   进程名 | 进程号  | 进程所有者| 文件描述符|文件类型|所在磁盘|文件大小|索引节点|文件名|

## 常用操作
### 打开某个文件的进程
```shell
$ sudo lsof /bin/bash
```

### 递归查看某个目录的文件信息
```shell
$ sudo lsof +D /root/dir
```

### 列出某个用户打开的文件信息
```shell
$ sudo lsof -u username
```

### 列出某个程序进程所打开的文件信息
```shell
$ sudo lsof -c mysql
```
### 列出某个用户以及某个进程所打开的文件信息
```shell
$ sudo lsof  -u test -c mysql 
```