# Nagle算法
## 机制
- 一个 TCP 连接上最多只能有一个未被确认的未完成的小分组[^1]，在它到达目的地前，不能发送其它分组。
-  在上一个小分组未到达目的地前，即还未收到它的 ACK 前，TCP 会收集后来的小分组。当上一个小分组的 ack 收到后，TCP 就将收集的小分组合并成一个大分组发送出去。  
[^1]:小于MSS的分组

## 关闭
设置**TCP_NODELAY**就行。