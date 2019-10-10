# uCore Lab3虚拟内存管理
## 前言
相比于**Lab1**、**Lab2**而言，**Lab3**确实简单许多，只是在**Lab2**的基础上增加了`mm_area, vma_area`数据结构和页替换机制，然而硬件方面的并不需要我们去理解，几个练习看注释也能写出来。
## 练习0 填写已有实验
## 内容
本实验依赖实验1/2。请把你做的实验1/2的代码填入本实验中代码中有“LAB1”,“LAB2”的注释相应部分。
## 答案
`meld lab2/ lab3/`转移即可

## 练习1 给未被映射的地址映射上物理页（需要编程）
### 内容
完成do_pgfault（mm/vmm.c）函数，给未被映射的地址映射上物理页。
### 答案
要完成实验需要理解`struct Page`新增的两个熟悉：
```c
struct Page {
      int ref;                        // page frame's reference counter
      uint32_t flags;                 // array of flags that describe the status of the page frame
      unsigned int property;          // the num of free block, used in first fit pm manager
      list_entry_t page_link;         // free list link
 +   list_entry_t pra_page_link;     // used for pra (page replace algorithm)
 +   uintptr_t pra_vaddr;            // used for pra (page replace algorithm)
};
```
`pra_page_link`是页替换算法用于组织物理页的链表结构，`pra_vaddr`表明这个物理页对应哪个虚拟地址。
**页异常**主要由三个原因导致：
1. **缺页**：即虚拟地址对应的物理地址不存在。
2. **读写权限异常**：如写只读页。
3. **访问权限异常**：如用户访问内核区域。

下面两个情况简单返回错误即可，第一个情况根据虚拟地址是否映射到物理地址分为两种情况，还没有映射到物理地址的话就分配一个`Page`然后插入即可了；已经分配了页的虚拟地址然而物理页的内容被别的虚拟地址替换下来了，说明之前有可能通过这个虚拟地址写过这个页，然后换入到了磁盘中，所以需要再分配一个物理页然后从磁盘读入之前的数据到这个物理页。要注意的是要使用`do_pgfault`这个函数，这个函数直接完成了从获得虚拟页表项，到物理页到虚拟页的映射的过程，并且还执行了FIFO算法中入队的过程。
```c
 ptep = get_pte(mm->pgdir, addr, 1);
  if (*ptep == 0) {
    //(2) if the phy addr isn't exist, then alloc a page & map the phy addr with
    // logical addr
    // struct Page *page = alloc_page();
    pgdir_alloc_page(mm->pgdir, addr, perm);
  } else {
    /*LAB3 EXERCISE 2: YOUR CODE
     * Now we think this pte is a  swap entry, we should load data from disk to
     * a page with phy addr, and map the phy addr with logical addr, trigger
     * swap manager to record the access situation of this page.
     *
     *  Some Useful MACROs and DEFINEs, you can use them in below
     * implementation. MACROs or Functions: swap_in(mm, addr, &page) : alloc a
     * memory page, then according to the swap entry in PTE for addr, find the
     * addr of disk page, read the content of disk page into this memroy page
     *    page_insert ： build the map of phy addr of an Page with the linear
     * addr la swap_map_swappable ： set the page swappable
     */
    if (swap_init_ok) {
      struct Page *page = NULL;
      //(1）According to the mm AND addr, try to load the content of right disk
      // page
      //    into the memory which page managed.
      swap_in(mm, addr, &page);
      //(2) According to the mm, addr AND page, setup the map of phy addr <--->
      page_insert(mm->pgdir, page, addr, perm);
      // logical addr (3) make the page swappable.
      swap_map_swappable(mm, addr, page, 0);
    } else {
      cprintf("no swap_init_ok but ptep is %x, failed\n", *ptep);
      goto failed;
    }
  }
```
注意要使用`pgdir_alloc_page`这个函数，这个函数在
### 问答
> 请描述页目录项（Page Directory Entry）和页表项（Page Table Entry）中组成部分对ucore实现页替换算法的潜在用处。

可以使用页目录项（Page Directory Entry）和页表项（Page Table Entry）中的**保留位**来表示页替换算法所需要的信息，如时钟置换算法就需要保留位来表示该物理页是否被修改-访问。

> 如果ucore的缺页服务例程在执行过程中访问内存，出现了页访问异常，请问硬件要做哪些事情？

所以这个问题和Lab2的问题就是换了个条件？一个是在执行过程中，一个是在中断服务例程中，那还是没有区别，把**Cr2**里的内容压栈，然后把当前又缺的线性地址放入**Cr2**中，然后压栈**error_code, CS, EIP**乱七八糟的。

## 练习2：补充完成基于FIFO的页面替换算法（需要编程）
### 内容
完成vmm.c中的do_pgfault函数，并且在实现FIFO算法的swap_fifo.c中完成map_swappable和swap_out_victim函数。
### 答案
`map_swappable`意思是来一个页面入队，`swap_out_victim`选一个页面出队。
```c
static int _fifo_map_swappable(struct mm_struct *mm, uintptr_t addr,
                               struct Page *page, int swap_in) {
  list_entry_t *head = (list_entry_t *)mm->sm_priv;
  list_entry_t *entry = &(page->pra_page_link);

  assert(entry != NULL && head != NULL);
  // record the page access situlation
  /*LAB3 EXERCISE 2: YOUR CODE*/
  //(1)link the most recent arrival page at the back of the pra_list_head
  // qeueue.
  page->pra_vaddr = addr;
  list_add_before(head, entry);
  cprintf("called swapple\n");
  return 0;
}
```
```c
static int _fifo_swap_out_victim(struct mm_struct *mm, struct Page **ptr_page,
                                 int in_tick) {
  list_entry_t *head = (list_entry_t *)mm->sm_priv;
  assert(head != NULL);
  assert(in_tick == 0);
  /* Select the victim */
  /*LAB3 EXERCISE 2: YOUR CODE*/
  //(1)  unlink the  earliest arrival page in front of pra_list_head qeueue
  //(2)  assign the value of *ptr_page to the addr of this page
  struct Page *res = le2page(head->next, pra_page_link);
  *ptr_page = res;
  list_del(head->next);
  return 0;
}
```
没有什么要点，唯一要注意的就是入队的时候把`page`结构里面的对应的虚拟地址填一下。
### 问答
> 如果要在ucore上实现"extended clock页替换算法"请给你的设计方案，现有的swap_manager框架是否足以支持在ucore中实现此算法？如果是，请给你的设计方案。如果不是，请给出你的新的扩展和基此扩展的设计方案。

支持，**extend clock页替换算法**需要的信息都可以通过`get_pte`获得页表项里面保留位的信息获得。

## Challenge 1：实现识别dirty bit的 extended clock页替换算法（需要编程）
原本以为在程序访问某个虚地址的时候需要我们手动设置页表项的`access`位，后来发现原来硬件自己就能设置，于是代码就好写了很多，只要替换一下`swap_victim`函数就行了，主要就是循环链表，然后检查每个对应页表项的`access`位和`dirty`位，如果同时为0那么就是返回的结果。还需要一个`list_entry_t *cur`指向当前的位置，由于标志在查找替换页时应该从哪个页面开始。
```c
#define IS_VISITED(pte) ((pte)&PTE_A)
#define IS_WRITED(pte) ((pte)&PTE_D)
```
```c
struct swap_manager swap_manager_fifo = {
    .name = "fifo swap manager",
    .init = &_fifo_init,
    .init_mm = &_extended_clock_init_mm,
    .tick_event = &_fifo_tick_event,
    .map_swappable = &_extended_clock_map_swapple,
    .set_unswappable = &_fifo_set_unswappable,
    .swap_out_victim = &_extended_clock_swap_out_victim,
    .check_swap = &_fifo_check_swap,
};
```
```c
list_entry_t *cur;
static int _extended_clock_init_mm(struct mm_struct *mm) {
  list_init(&pra_list_head);
  mm->sm_priv = &pra_list_head;
  cur = &pra_list_head;
  return 0;
}

static int _extended_clock_map_swapple(struct mm_struct *mm, uintptr_t addr,
                                       struct Page *page, int swap_in) {
  list_entry_t *head = (list_entry_t *)mm->sm_priv;
  list_entry_t *entry = &(page->pra_page_link);

  assert(entry != NULL && head != NULL);
  page->pra_vaddr = addr;
  list_add_before(head, entry);
  // only change when first in or loop to find the victim
  if (cur == head) {
    cur = head->next;
  }
  cprintf("called swapple\n");
  return 0;
}

static int _extended_clock_swap_out_victim(struct mm_struct *mm,
                                           struct Page **ptr_page,
                                           int in_tick) {
  list_entry_t *head = (list_entry_t *)mm->sm_priv;
  assert(head != NULL);
  assert(in_tick == 0);
  struct Page *res = NULL;
  list_entry_t *le = cur;
  assert(head != le);
  for (;; le = le->next) {
    if (le == head) continue;
    struct Page *page = le2page(le, pra_page_link);
    pte_t *pted = get_pte(mm->pgdir, page->pra_vaddr, 0);
    assert(pted != NULL);

    if (!(IS_VISITED(*pted)) && !(IS_WRITED(*pted))) {
      *ptr_page = page;
      if (le->next == head) {
        cur = le->next->next;
      } else {
        cur = le->next;
      }
      list_del(le);
      return 0;
    } else if (IS_WRITED(*pted)) {
      *pted &= (~PTE_D);
    } else if (IS_VISITED(*pted)) {
      *pted &= (~PTE_A);
    }
  }
}
```
然而使用全局变量而不是`mm_area`去存下一个起始页的不好的地方在于当引入用户进程后，`mm_area`结构不止一个，那么可能A进程的页指向的是B进程的页，但是uCore原本用于存储需要换出页的链表`pra_page_link`好像就是全局的，也就没所谓了。
### 问答
> 1. 需要被换出的页的特征是什么？

**访问位**和**修改位**为0的页。
> 2. 在ucore中如何判断具有这样特征的页？

检查一下页表项的**访问位**和**修改位**是不是1。

> 3. 何时进行换入和换出操作？

在`alloc_page`返回**NULL**时进行换出，在缺页并且内容在磁盘中时进行换入。
