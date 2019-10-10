# uCore lab6 调度框架和调度算法设计与实现
## 练习0：填写已有实验
这真的是个坑爹的练习，Lab1-5的内容并不是全部拷贝到一起就可以的，里面还有要求你更新的内容，所以就会一不小心把要更新的注释给合并没了。

## 练习1: 使用 Round Robin 调度算法
完成练习0后，建议大家比较一下（可用kdiff3等文件比较软件）个人完成的lab5和练习0完成后的刚修改的lab6之间的区别，分析了解lab6采用RR调度算法后的执行过程。
### 问答
>请理解并分析sched_class中各个函数指针的用法，并结合Round Robin 调度算法描ucore的调度执行过程

正常几个调度点`trap(), do_wait(), do_yeild(), do_exit()`，然后`schedule`里面先入队当前的process，让这些就绪进程一起公平竞争，`pick_one`然后再出队即可。RoundRobin算法就是简单的先进先出队列，不涉及到任何排序。

> 请在实验报告中简要说明如何设计实现”多级反馈队列调度算法“，给出概要设计，鼓励给出详细设计

给process根据优先级先放到对应队列里面，设置好时间片，然后当前进程时间片用完的话出队，优先级降级，再根据优先级入队，设置时间片。选取下一个进程就是找到一个不是空的队列的队头进程运行即可。

## 练习2: 实现 Stride Scheduling 调度算法
### 内容
首先需要换掉RR调度器的实现，即用default_sched_stride_c覆盖default_sched.c。然后根据此文件和后续文档对Stride度器的相关描述，完成Stride调度算法的实现。
### 答案
Stride Scheduling算法是典型的最小堆的使用场景，所以uCore提供了一个节点结构的最小堆，然后根据`stride`的大小出队入队即可。完成这个练习首先需要确定**BIG_STRIDE**的值，只有**max_strice - min_stride <= BIG_STRIDE**的时候，即使`stride`的值溢出32位无符号整形数了依然可以判断大小。其次就是简单的照注释写代码了。。。
#### 确定**MAX_STRIDE**
这个比较玄学，首先这个值不可能是个奇奇怪怪的值，我们可以从**INT_MAX,INT_MIN, INT_MAX/2, INT_MIN/2**（32位）这几个比较特殊的值中选取，由于**MAX_STRIDE**这个值大于零，那么看一看**INT_MAX+1, INT_MAX/2+1**还可不可以，emm，最后定为**INT_MAX=2e31-1**。
```c
#define BIG_STRIDE 2147483647
```
#### 看注释写代码
```c
static void stride_init(struct run_queue *rq) {
  rq->lab6_run_pool = NULL;
  rq->proc_num = 0;
}
```
```c
static void stride_enqueue(struct run_queue *rq, struct proc_struct *proc) {
  rq->lab6_run_pool = skew_heap_insert(
      rq->lab6_run_pool, &(proc->lab6_run_pool), proc_stride_comp_f);
  if (proc->time_slice == 0 || proc->time_slice > rq->max_time_slice) {
    proc->time_slice = rq->max_time_slice;
  }
  proc->rq = rq;
  rq->proc_num++;
}
```
```c
static void stride_dequeue(struct run_queue *rq, struct proc_struct *proc) {
  rq->lab6_run_pool = skew_heap_remove(
      rq->lab6_run_pool, &(proc->lab6_run_pool), proc_stride_comp_f);
  rq->proc_num--;
}
```
```c
static struct proc_struct *stride_pick_next(struct run_queue *rq) {
  if (rq->lab6_run_pool == NULL) return NULL;
  struct proc_struct *p = le2proc(rq->lab6_run_pool, lab6_run_pool);
  if (p->lab6_priority == 0)
    p->lab6_stride += BIG_STRIDE;
  else
    p->lab6_stride += BIG_STRIDE / p->lab6_priority;
  return p;
}
```
```c
static void stride_proc_tick(struct run_queue *rq, struct proc_struct *proc) {
  if (proc->time_slice > 0) {
    proc->time_slice--;
  }
  if (proc->time_slice == 0) {
    proc->need_resched = 1;
  }
}
```