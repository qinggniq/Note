# uCore lab5 用户进程管理
## 练习0：填写已有实验
除了将Lab1-4中练习的内容拷贝到Lab5的代码中，之前实验练习的内容还需要更新。

- 首先是中断描述符表的初始化，实验一是全部设置为特权级0，然后**T_SWITCH_TOU**设置为特权级3，现在由于需要实现系统调用，所以需要把**T_SYSCALL**也设置为特权级3；
```c
void idt_init(void) {
    for (int i = 0; i <= 255; i++) {
        SETGATE(idt[i], 0, GD_KTEXT, __vectors[i], DPL_KERNEL);
    }
    SETGATE(idt[T_SWITCH_TOU], 1, GD_KTEXT, __vectors[T_SWITCH_TOU], DPL_USER);
+  SETGATE(idt[T_SYSCALL], 1, GD_KTEXT, __vectors[T_SYSCALL], DPL_USER);
     lidt(&idt_pd);
}
```
- 其次是Lab4中有关`alloc_proc`中对于进程描述符的初始化，因为增加了新的属性`wait_state, cpro, ypro, opro`用来处理子进程的回收工作。

|    属性     |        用途        |
| :---------: | :----------------: |
| waite_state |    进程等待状态    |
|    cpro     | 进程的第一个子进程 |
|    ypro     |  进程的前一个兄弟  |
|    opro     |  进程的后一个兄弟  |

```c
static struct proc_struct *alloc_proc(void) {
  struct proc_struct *proc = kmalloc(sizeof(struct proc_struct));
  if (proc != NULL) {
    memset(proc, 0, sizeof(struct proc_struct));
    proc->state = PROC_UNINIT;
    proc->pid = 0;
    proc->runs = 0;
    proc->need_resched = 0;
    proc->parent = current;
    proc->mm = NULL;
    proc->cr3 = boot_cr3;
    proc->flags = 0;

+  proc->wait_state = 0;
+  proc->cptr = NULL;
+  proc->yptr = NULL;
+  proc->optr = NULL;
  }
  return proc;
}
```
还有就是`trap.c/trap()`函数里面处理时钟中断的时候需要添加`current->need_resched = 1;`，十分坑爹的是如果你用`make grade`来验证你的程序是不是正确的话，那么得把Lab1中打印100个时钟脉冲的`cprintf()`语句去掉，否则是拿不了满分的。。

## 练习1: 加载应用程序并执行（需要编码）
### 内容
`do_execv`函数调用`load_icode`（位于`kern/process/proc.c`中）来加载并解析一个处于内存中的ELF执行文件格式的应用程序，建立相应的用户内存空间来放置应用程序的代码段、数据段等，且要设置好`proc_struct`结构中的成员变量`trapframe`中的内容，确保在执行此进程后，能够从应用程序设定的起始执行地址开始执行。需设置正确的trapframe内容。
### 答案
答案尽在注释里。。
```c
/* LAB5:EXERCISE1 YOUR CODE
   * should set tf_cs,tf_ds,tf_es,tf_ss,tf_esp,tf_eip,tf_eflags
   * NOTICE: If we set trapframe correctly, then the user level process can
   * return to USER MODE from kernel. So tf_cs should be USER_CS segment (see
   * memlayout.h) tf_ds=tf_es=tf_ss should be USER_DS segment tf_esp should be
   * the top addr of user stack (USTACKTOP) tf_eip should be the entry point of
   * this binary program (elf->e_entry) tf_eflags should be set to enable
   * computer to produce Interrupt
   */
  tf->tf_cs = USER_CS;
  tf->tf_ds = tf->tf_es = tf->tf_ss = USER_DS;
  tf->tf_eip = elf->e_entry;
  tf->tf_esp = USTACKTOP;
  tf->tf_eflags |= FL_IF;
```
### 问答
> 请在实验报告中描述当创建一个用户态进程并加载了应用程序后，CPU是如何让这个应用程序最终在用户态执行起来的。即这个用户态进程被ucore选择占用CPU执行（RUNNING态）到具体执行应用程序第一条指令的整个经过。

设置为**RUNNING**后，设置`current`为该用户进程，然后调用`swich_to`函数，记得当时`copy_thread`的时候设置的`eip`为`forkret`，恢复了上下文之后`ret`会返回到`forkret`，然后`forkret`又会跳到`__trapret`，然后会将`proc->trap`里面存储的关于段寄存器的值全部存储到相应的寄存器里面，通过`trapframe`我们可以完成特权级的切换，然后就会跳到`tf->eip`那里去执行，而在`copy_thread`函数中我们知道`tf->eip = fn`，于是跳到了用户态进程执行。
## 练习2: 父进程复制自己的内存空间给子进程
### 内容
创建子进程的函数do_fork在执行中将拷贝当前进程（即父进程）的用户内存地址空间中的合法内容到新进程中（子进程），完成内存资源的复制。具体是通过copy_range函数（位于kern/mm/pmm.c中）实现的，请补充copy_range的实现，确保能够正确执行。

### 答案
考察函数的使用。。。
```c
    uintptr_t src_kvaddr = page2kva(page);
    uintptr_t dst_kvaddr = page2kva(npage);

    memcpy(dst_kvaddr, src_kvaddr, PGSIZE);
    ret = page_insert(to, npage, start, perm);
    assert(ret == 0);
```
## 练习3: 阅读分析源代码，理解进程执行 fork/exec/wait/exit 的实现，以及系统调用的实现
### 问答
> 请分析fork/exec/wait/exit在实现中是如何影响进程的执行状态的？

| 函数  |                                                                                  影响                                                                                  |
| :---: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------: |
| fork  |                                                                             创建小新的PCB                                                                              |
| exec  |                                                                    清空`mm`，`load_icode`执行新代码                                                                    |
| wait  | 如果有子进程，看看有没有僵死的子进程，有的话释放资源（PCB和stack）返回状码，如果没有僵死的子进程，那么设置进程为**SLEEPING**状态，重新调度；没有子进程的话那么返回错误 |
| exit  |                               销毁`mm`，设置状态为僵死状态，如果父进程在等待的话，那么唤醒父进程，然后将改进程的子进程全部给init进程收养                               |
> 请给出ucore中一个用户态进程的执行状态生命周期图（包执行状态，执行状态之间的变换关系，以及产生变换的事件或函数调用）

[RUNNING] -{do_wait}-> [SLEEPING] 

[SLEEPING] -{wakeup_proc}-> [RUNNING]

[RUNNING] -{do_exit}-> [ZOMBI]

[ZOMBI] --{be_waited}-> [END]



