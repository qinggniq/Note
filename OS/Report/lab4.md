# lab4 内核线程管理
## 练习1：分配并初始化一个进程控制块（需要编码）
### 内容
`alloc_proc`函数（位于`kern/process/proc.c`中）负责分配并返回一个新的`struct proc_struct`结构，用于存储新建立的内核线程的管理信息。ucore需要对这个结构进行最基本的初始化，你需要完成这个初始化过程。
<!--more-->
### 答案
```c
struct proc_struct {
  enum proc_state state;  // Process state
  int pid;                // Process ID
  int runs;               // the running times of Proces
  uintptr_t kstack;       // Process kernel stack
  volatile bool
  need_resched;  // bool value: need to be rescheduled to release CPU?
  struct proc_struct *parent;  // the parent process
  struct mm_struct *mm;        // Process's memory management field
  struct context context;      // Switch here to run process
  struct trapframe *tf;        // Trap frame for current interrupt
  uintptr_t cr3;   // CR3 register: the base addr of Page Directroy Table(PDT)
  uint32_t flags;  // Process flag
  char name[PROC_NAME_LEN + 1];  // Process name
  list_entry_t list_link;        // Process link list
  list_entry_t hash_link;        // Process hash list
};
```
就是需要你去初始化**PCB**里面的内容，虽然注释上是让初始化`struct proc_struct`结构里面所有的字段，但是根据`proc_init`的内容我们可以发现：
- `pid`，`kstack`， `context`，`tf`，`mm`，`list_link`，`hash_link`在`do_fork()`时确定；
- `name`在`proc_init()`中确定。
所以只需要简单设置一下其他结构的初始值即可。
```c
static struct proc_struct *alloc_proc(void) {
  struct proc_struct *proc = kmalloc(sizeof(struct proc_struct));
  if (proc != NULL) {
    // LAB4:EXERCISE1 YOUR CODE
    /*
     * below fields in proc_struct need to be initialized
     *       enum proc_state state;                      // Process state
     *       int pid;                                    // Process ID
     *       int runs;                                   // the running times of
     * Proces uintptr_t kstack;                           // Process kernel
     * stack volatile bool need_resched;                 // bool value: need to
     * be rescheduled to release CPU?
     * struct proc_struct *parent; // the parent
     * process struct mm_struct *mm;                       // Process's memory
     * management field struct context context;                     // Switch
     * here to run process struct trapframe *tf;                       // Trap
     * frame for current interrupt uintptr_t cr3; // CR3 register: the base addr
     * of Page Directroy Table(PDT) uint32_t flags; // Process flag char
     * name[PROC_NAME_LEN + 1];               // Process name
     */
    memset(proc, 0, sizeof(struct proc_struct));
    proc->state = PROC_UNINIT;
    proc->pid = -1;
    proc->runs = 0;
    proc->need_resched = 0;
    proc->parent = current;
    proc->mm = NULL;

    proc->cr3 = boot_cr3;
    proc->flags = 0;
  }
  return proc;
}
```
### 问答
> 请说明`proc_struct`中`struct context context`和`struct trapframe *tf`成员变量含义和在本实验中的作用是啥？

- `context`用于上下文切换，线程切换的时候`context`用来存储线程的执行状态也就是`eip esp ebx ecx edx esi edi ebp`这一堆的寄存器的值。
- `tf`是实验一里面的用来存储中断栈的结构，当线程执行的时候发生了中断，实验一中断发生的时候中断信息（如**ERROR_CODE, EIP, CS, 一些寄存器信息**）是直接保存在栈上的，但是现在是保存在`proc->tf`结构体里面。

## 练习2：为新创建的内核线程分配资源（需要编码）
### 内容
创建一个内核线程需要分配和设置好很多资源。kernel_thread函数通过调用do_fork函数完成具体内核线程的创建工作
### 答案
照着注释和实验书上的要点写就好了，比较坑的是有关全局数据的处理需要禁止中断。
```c
/* do_fork -     parent process for a new child process
 * @clone_flags: used to guide how to clone the child process
 * @stack:       the parent's user stack pointer. if stack==0, It means to fork
 * a kernel thread.
 * @tf:          the trapframe info, which will be copied to child process's
 * proc->tf
 */
int do_fork(uint32_t clone_flags, uintptr_t stack, struct trapframe *tf) {
  int ret = -E_NO_FREE_PROC;
  struct proc_struct *proc;
  if (nr_process >= MAX_PROCESS) {
    goto fork_out;
  }
  ret = -E_NO_MEM;
  // LAB4:EXERCISE2 YOUR CODE
  /*
   * Some Useful MACROs, Functions and DEFINEs, you can use them in below
   * implementation. MACROs or Functions: alloc_proc:   create a proc struct and
   * init fields (lab4:exercise1) setup_kstack: alloc pages with size KSTACKPAGE
   * as process kernel stack copy_mm:      process "proc" duplicate OR share
   * process "current"'s mm according clone_flags if clone_flags & CLONE_VM,
   * then "share" ; else "duplicate" copy_thread:  setup the trapframe on the
   * process's kernel stack top and setup the kernel entry point and stack of
   * process hash_proc:    add proc into proc hash_list get_pid:      alloc a
   * unique pid for process wakeup_proc:  set proc->state = PROC_RUNNABLE
   * VARIABLES:
   *   proc_list:    the process set's list
   *   nr_process:   the number of process set
   */

  //    1. call alloc_proc to allocate a proc_struct
  //    2. call setup_kstack to allocate a kernel stack for child process
  //    3. call copy_mm to dup OR share mm according clone_flag
  //    4. call copy_thread to setup tf & context in proc_struct
  //    5. insert proc_struct into hash_list && proc_list
  //    6. call wakeup_proc to make the new child process RUNNABLE
  //    7. set ret vaule using child proc's pid
  if ((proc = alloc_proc()) == NULL) goto fork_out;
  if (setup_kstack(proc) != 0) goto bad_fork_cleanup_proc;
  if (copy_mm(clone_flags, proc) != 0) goto bad_fork_cleanup_kstack;
  copy_thread(proc, stack, tf);
  cprintf("up ok !\n");
  bool intr_flag;
  local_intr_save(intr_flag);
  {
    list_add(&proc_list, &proc->list_link);
    proc->pid = get_pid();
    hash_proc(proc);
    nr_process++;
  }
  local_intr_restore(intr_flag);
  cprintf("list add ok\n");
  wakeup_proc(proc);

  ret = proc->pid;
fork_out:
  return ret;

bad_fork_cleanup_kstack:
  put_kstack(proc);
bad_fork_cleanup_proc:
  kfree(proc);
  goto fork_out;
}
```

### 问答
> 请说明ucore是否做到给每个新fork的线程一个唯一的id？请说明你的分析和理由。

可以，根据`if (proc->pid == last_pid)`这条判断语句和之后的`++last_pid`可知，一旦循环时出现`proc->pid == last_pid`的情况就会立刻自增，不会构造出重复的id。


## 练习3：阅读代码，理解 proc_run 函数和它调用的函数如何完成进程切换的。
### 问答
> 在本实验的执行过程中，创建且运行了几个内核线程？

两个，一个`idleproc`，一个`initproc`。
> 语句local_intr_save(intr_flag);....local_intr_restore(intr_flag);在这里有何作用?

禁止中断，至于为什么，中断会破坏当前的执行吗？