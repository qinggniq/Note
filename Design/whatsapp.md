# WhatsApp的设计

1. Group Massage（群消息）
2. Sent + Delivered + Read（收发消息）
3. Online/Last time  Seen（在线/最近在线时间）
4. Image Sharing（图片）
5. Chats temporary/Perment（消息历史存储位置）
6. one to one 

## ONE TO ONE 单点通信

![image-20191211211556740](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191211211556740.png)

`session`里面存的用户到网关的映射。

> 为什么要用session而不是传统的数据库

- 因为用户到网关的映射会经常发生变化。

> 通信用什么协议？

- 最好不要用普通的http协议，因为http协议是一个CS架构，server不能向client发送请求。（但是可以用[long polling](https://www.jianshu.com/p/d3f66b1eb748)**[TODO]**技术），但是long polling不是实时的，有延迟。
- websocket([wss](https://www.zhihu.com/question/20215561/answer/40316953))

> 聊天记录在session发的时候失败了怎么办？

- 用一个`chat`数据库持久化消息，失败重试。

## 上线/下线时间

- 用一个表来存，user/timestamp。
- 注意区分用户发的消息和系统消息（不会更新last seen时间）的区别。

## 群消息

