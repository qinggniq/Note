# golang runtime

![image-20191129154138762](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191129154138762.png)

1.1 CPM协程模型

1.5 三色标记法

1.8 内存屏障





id，状态，栈，调度上下文信息（栈支持指针位置，运行到的程序位置），程序开始地址



调度

全局队列

![image-20191129155258215](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191129155258215.png)

GPM模型

P是一个资源：（内存、队列）

网络

![image-20191129160237662](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191129160237662.png)

使用epoll，把携程作为private数据，在fd ready的时候把携程放到可运行队列中，然后调度。

sysmon会定时拿出来net ready的携程

gc的时候也会拿出来

调度的时候也会

![image-20191129161327778](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191129161327778.png)

先cache、mcentral、mheap、sysAlloc(mmap)，锁的粒度不断增大，从无锁到大锁，增加性能

sysmon会5分钟检查mcache里面的内存是不是没有被使用，那么把mcache归还给mheap。



go的虚拟内存地址只会增加，但是物理内存会被归还

**tcmalloc**



## 三色标记法gc

![image-20191129162115808](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191129162115808.png)

白色没有标记

灰色 => 黑色 扫描的

最后把白色的（不可达）的对象gc。

go和java不一样，没有元数据。

标记和并发的关系。

### 并发问题



![image-20191129163037124](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191129163037124.png)



![image-20191129163506903](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191129163506903.png)

![image-20191129164019202](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191129164019202.png)

![image-20191129164747844](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191129164747844.png)

并发情况下多用channel（CSP模型）



携程need schedule的情况，运行长，阻塞（获得锁，IO）

线程控制携程，自己是不能让出cpu的。

![image-20191129172000511](/Users/qinggniq/Library/Application Support/typora-user-images/image-20191129172000511.png)

重启的时候会通过负载均衡把所有的流量清零，然后再重启。

通过网关维护长链接。

