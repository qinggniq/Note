# 容错虚拟机的设计
## introduction
容错系统的基本方法是`主从复制`，backup server需要和primary server保持一个状态。那么复制状态的方法可以是将机器的所有状态，比如CPU、内存、IO设备的状态变化全部传输到backup server，然而由于带宽不够，特别是内存状态的变化是十分巨大的。

还有一种方法是**状态机方法**，把servers模拟为一个状态机，在初始时确定它们的初始状态一致，那么如果接受同样的输入，它们就会产生同样的输出。但是但有个问题是有些操作比如定时器、中断之类的指令是非确定性的，所以需要额外的协同来**同步**，因为需要额外的同步远远小于所有的状态，所以这个方法还是work的。

主机上这么做很困难，但是虚拟机这么做还是比较容易的，因为可以从外侧观测到虚拟的状态，所以那些额外需要同步的信息虚拟机管理器（hypervisor）是可以得知的。

## 基本设计
[图]
通过 logging channel进行通信。
### 确定性演绎的实现
确定性输入，非确定性输入（时钟脉冲），难点：
1. 正确的捕获所有的输出和必要的非确定性行为，保证backup server的执行。
2. 正确的应用输入和非确定性行为到backup server上。
3. 在保证1,2点的同时不降级性能。

非确定性行为就把实际发生的指令放到log里面。

### FT协议
为了保证容错性，必须遵循严格的FT协议，
**输出要求**
如果backup VM在primary VM挂起后接管了，那么backup VM会持续执行，保证它的执行和primary VM执行输出一致

输出要求可以通过延迟外部输出”只有backup VM接收到所有的信息“才输出出去。
> 中间一部分没看懂

**输出规则**
直到backup VM接受并知道（发出ACK）有关**产生输出**的日志之前，primary VM不会向外输出。
注：产生输出也是个操作。

TCP，操作系统保证在primary VM收到并且没有发给backup VM之前崩溃的话会重发而不会使消息丢失。


### 检测、响应 失败
如果backup VM failure，那么primary VM 继续存活，不再把操作发给logging channel了。如果primary VM failure，那么backup VM也继续存活，但是情况比较复杂，backup VM需要把收到的log演绎完了，然后才正常执行。然后VMware就会自动路由请求到这个backup VM，backup VM于是升级到primary VM。

VMware FT使用UDP心跳来检测failure。UDP跑在logging channel的信道里面（？？？存疑）。
但是可能发生脑裂的情况，所以使用虚拟磁盘的共享存储来解决脑裂问题，如果任何一个想go live，也就是上面发生failure的情况，继续存活，那么就会执行一下TAS(test-and-set)操作到共享存储上，如果成功，那么继续存活，如果失败，那么必然另外一个存活。（？？？这有啥？？？）

## FT的工程实现
可用、健壮、自动系统
### 开启和重启FT VMs
VMware VMotion允许迁移一个正在运行的虚拟机，同时也创建了logging channel，后面就是VMware FT可以根据一些原则在集群中选择合适的虚拟机去拷贝。
### 管理logging channel
channel为空的时候backup server选择挂起，channel满的时候primary server选择挂起，所以需要减小挂起的可能性。

主要是因为backup server的消耗速度过慢导致此类事件发生。

通过CPU控制primary server的执行速度不要过快，比backup server快太多就分配少一点CPU。

### FT VMs的操作
关机也需要同步

### 有关磁盘IO上面实现上的问题
对于同一块磁盘上的并行磁盘操作会导致不确定性行为，所以需要检测这些哦操作，然后让它们串行。

因为磁盘读是直接用的DMA的方法，所以实现方法是给页面加锁，如果VM读的那个页被锁保护了，那么就等待。

*bounce buffers*一个缓冲区。

primary VM挂了之后即使那个时候的磁盘操作成功了也标志为失败。由于之前把磁盘操作串行化了，所以可以复现失败的磁盘操作。

### 有关网络IO上的实现问题

优化
1. 聚簇优化，降低VM的陷入和中断。就是接到的数据不是立刻拷贝到内核缓冲区，而是等到一定量之后再拷贝。
2. 发log entirs 和 ack的时候不会发生任何线程上下文切换，直接在TCP栈上面跑，类似于tasklet。


## 设计上的选择
### 共享 VS 非共享磁盘
共享磁盘不用同步磁盘，非共享磁盘不用挂起写操作并且脑裂的哪个解决方式就不能用了。

### backup VM执行磁盘读
可以用磁盘读写来代替Logging  channel，但是在失效的时候会增加复杂度。
