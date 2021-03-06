# 延迟确认
## 机制
服务器在收到客户端数据后并不立刻返回**ACK**，而是等待一段时间（40ms**弄清楚具体时间**），再发**ACK**。
## 目的 
- **累积确认**：服务器极有可能在这 40ms 里又收到了客户端发送过来的多个 TCP 报文段，40ms 后就可以对这些报文进行累积确认，也就是只返回一个 ACK 报文就行。这样就能减少网络中 ACK 报文的数量。
- **捎带确认**：40ms 里，服务器也可能会返回数据给客户端，如果服务器有数据返回给客户端，那不如把这个 ACK 连同数据一起返回给客户端吧，等一下下是值得的。
## 启动
`man 7 tcp`，查看TCP选项，其中有一个**QUICKACK**选项，关掉它就可以开启延迟确认了。
```c
int opt = 0;
setsockopt(sockfd, IPPROTO_TCP, TCP_QUICKACK, &opt, sizeof(int));
```
每次`recv`后都需要重新设置。

