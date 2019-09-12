# 非阻塞connect的行为
## 阻塞vs分阻塞
非阻塞模式的 `socket`，一般用于需要支持高并发多 **QPS** 的场景下（如服务器程序），但是正如前文所述，这种模式让程序执行流和控制逻辑变复杂；相反，阻塞模式逻辑简单，程序结构简单明了，常用于一些特殊的场景。
## Linux
在 Linux 系统上一个 `socket` 没有建立连接之前，用 `select` 函数检测其是否可写，你也会得到可写的结果。正确的做法是，connect 之后，不仅要用 `select `检测可写，还要检测此时 `socket` 是否出错，通过错误码来检测确定是否连接上，错误码为 0 表示连接上，反之为未连接上。

就是在`select`返回以后再`getsockopt`检查一下：
```c++
 if (select(clientfd + 1, NULL, &writeset, NULL, &tv) != 1) {
		close(clientfd);
		return -1;
	}
	int err;
    socklen_t len = static_cast<socklen_t>(sizeof err);
    if (::getsockopt(clientfd, SOL_SOCKET, SO_ERROR, &err, &len) < 0) {
        close(clientfd);
		return -1;
	}
```
