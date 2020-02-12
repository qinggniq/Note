# Channels

## channel特性

- goroutine安全
- 能在goroutine之间传递消息
- 提供先入先出语意
- 能使gorountine阻塞和非阻塞

## channel实现

```go
type hchan struct{
  buf 
  sendIdx
  reciveIdx
 	mutex
}
```

- 在heap上分配
- chan是个指针

## send & Recive

- 在发送和接收数据的时候使用的是**copy**而不是共享内存，多个goroutine唯一共享的是hchan中的buf，然而buf的访问被mutex所保护，所以goroutine安全。（！！！！！Linux多线程服务端编程里面的写时拷贝技发？？？）
- goroutine的挂起是通过go runtime实现的，因为goroutine是用户线程。

## CPM调度模型

- M OS线程
- G goroutine
- P 调度上下文，维护一个阻塞的goroutine队列
- ![image-20191222224720365](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191222224720365.png)



## 优化

- 在send的时候直接把数据从g1的栈里面拷贝到g2的栈，这样就少了一次拷贝，也少了两次获取锁的操作。