# uCore lab7报告 同步互斥机制的设计与实现
## 练习0：填写已有实验
### 内容
本实验依赖实验1/2/3/4/5/6。请把你做的实验1/2/3/4/5/6的代码填入本实验中代码中有“LAB1”/“LAB2”/“LAB3”/“LAB4”/“LAB5”/“LAB6”的注释相应部分。并确保编译通过。注意：为了能够正确执行lab7的测试应用程序，可能需对已完成的实验1/2/3/4/5/6的代码进行进一步改进。
### 答案
- 注意把lab6写的`stride_sched.c`也拷过去，然后把默认的**RoundRobin**实现给删了；
- `trap.c`中将时钟中断的处理换成`run_timer_list()`；
```c
case IRQ_OFFSET + IRQ_TIMER:
+    run_timer_list();
-     sched_class_proc_tick(current);
      break;
```
## 练习1: 理解内核级信号量的实现和基于内核级信号量的哲学家就餐问题
### 内容
完成练习0后，建议大家比较一下（可用`meld`等文件`diff`比较软件）个人完成的lab6和练习0完成后的刚修改的lab7之间的区别，分析了解lab7采用信号量的执行过程。执行`make grade`，大部分测试用例应该通过。

### 答案
> 请在实验报告中给出内核级信号量的设计描述，并说明其大致执行流程。

- 主要数据结构
```c
//信号量结构体
typedef struct {
    int value;
    wait_queue_t wait_queue;
} semaphore_t;

//等待队列
typedef struct {
    list_entry_t wait_head;
} wait_queue_t;

//控制信息
typedef struct {
    struct proc_struct *proc;
    uint32_t wakeup_flags;
    wait_queue_t *wait_queue;
    list_entry_t wait_link;
} wait_t;
```
通过`wait_t`将等待队列与进程关联。
- 主要执行流程
    - `down`判断信号量值是否大于0，大于零更新信号量之后直接返回，否则将其加入信号量的等待队列，设置当前进程状态为**SLEEPING**，然后重新调度，当重新执行的时候只可能是信号量可用的时候被唤醒的情况，然后出队返回。
    - `up`判断信号量是否大于零，大于零的话说明等待队列里面没有要被唤醒的进程，否则就唤醒队头的那个进程。
> 请在实验报告中给出给用户态进程/线程提供信号量机制的设计方案，并比较说明给内核级提供信号量机制的异同。

- `struct semaphore`信号量结构
- `semaphore_init(struct semaphore* sem, size_t val)`初始化信号量，直接调用内核的`void sem_init(semaphore_t *sem, int value);`即可；
- `P(struct semaphore* sem)`尝试获取资源，调用`down`即可；
- `V(struct semaphore* sem)`释放资源，调用呢`up`。
所有的操作都得做成系统调用。

## 练习2: 完成内核级条件变量和基于内核级条件变量的哲学家就餐问题
### 内容
首先掌握管程机制，然后基于信号量实现完成条件变量实现，然后用管程机制实现哲学家就餐问题的解决方案（基于条件变量）。
### 答案
看伪代码写代码。。。
```c
// Unlock one of threads waiting on the condition variable.
void cond_signal(condvar_t *cvp) {
  // LAB7 EXERCISE1: YOUR CODE
  cprintf(
      "cond_signal begin: cvp %x, cvp->count %d, cvp->owner->next_count %d\n",
      cvp, cvp->count, cvp->owner->next_count);
  /*
   *      cond_signal(cv) {
   *          if(cv.count>0) {
   *             mt.next_count ++;
   *             signal(cv.sem);
   *             wait(mt.next);
   *             mt.next_count--;
   *          }
   *       }
   */
  if (cvp->count > 0) {
    cvp->owner->next_count++;
    up(&cvp->sem);
    down(&cvp->owner->next);
    cvp->owner->next_count--;
  }
  cprintf("cond_signal end: cvp %x, cvp->count %d, cvp->owner->next_count %d\n",
          cvp, cvp->count, cvp->owner->next_count);
}
```

```c
void cond_wait(condvar_t *cvp) {
  // LAB7 EXERCISE1: YOUR CODE
  cprintf(
      "cond_wait begin:  cvp %x, cvp->count %d, cvp->owner->next_count %d\n",
      cvp, cvp->count, cvp->owner->next_count);
  /*
   *         cv.count ++;
   *         if(mt.next_count>0)
   *            signal(mt.next)
   *         else
   *            signal(mt.mutex);
   *         wait(cv.sem);
   *         cv.count --;
   */
  cvp->count++;
  if (cvp->owner->next_count)
    up(&cvp->owner->next);
  else
    up(&cvp->owner->mutex);
  down(&cvp->sem);
  cvp->count--;
  cprintf("cond_wait end:  cvp %x, cvp->count %d, cvp->owner->next_count %d\n",
          cvp, cvp->count, cvp->owner->next_count);
}
```
