# Linux中listen系统调用中backlog

## 问题
1. backlog指的什么？

## 回答
```c++
int listen(int sockfd, int backlog);
```
> The backlog argument defines the maximum length to which the queue of pending connections for sockfd may grow. 
 
> If a connection request arrives  when the queue is full,the client may receive an error with an indication of ECONNREFUSED or, if the underlying protocol supports retransmission, the request may be ignored so that a later reattempt at connection succeeds.

1. backlog是`the queue of pending connections`，可以指
    1. 已经完成三次握手的连接队列
    2. 已经完成三次握手的连接队列 + 没有完成三次握手的队列

对于现代Linux一般指第一种。

## 问题2
1. `the queue of pending connections`满了会如何处理客户端发过来的请求？
2. 实际上的两个队列（连接、未连接）的长度由什么决定？

## 回答
1. 当队列已满的时候，直接返回**reset**是否合适(ECONNREFUSED)这样处理有个坏处，就是客户端无法区别是未处理队列满了，还是访问了系统未监听的端口，所以还是直接忽略这个**SYN**比较好。
2. 由两个内核参数决定：
    1. `net.core.somaxconn` => **[ESTABLISHED]已连接队列**
    2. `tcp_max_syn_backlog` => **[SYN-RECEIVED]未完成连接队列**

## tips
当**ESTABLISHED队列**达到上限的时候，服务端会启动定时器重新发送SYN-ACK。
当**SYN-RECV队列**达到上限的时候，服务器直接丢弃了客户端发送的SYN包。