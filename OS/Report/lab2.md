# uCore lab2 物理内存管理
## 练习0：填写已有实验
### 内容
本实验依赖实验1。请把你做的实验1的代码填入本实验中代码中有“LAB1”的注释相应部分。提示：可采用diff和patch工具进行半自动的合并（merge），也可用一些图形化的比较/merge工具来手动合并，比如meld，eclipse中的diff/merge工具，understand中的diff/merge工具等。
### 答案
`meld ./lab1 ./lab2`将`lab1`里面的做的练习的代码移到`lab2`里面去就行。
## 练习1：实现 first-fit 连续物理内存分配算法（需要编程）
### 实验
在实现first fit 内存分配算法的回收函数时，要考虑地址连续的空闲块之间的合并操作。提示:在建立空闲页块链表时，需要按照空闲页块起始地址来排序，形成一个有序的链表。可能会修改default_pmm.c中的default_init，default_init_memmap，default_alloc_pages， default_free_pages等相关函数。请仔细查看和理解default_pmm.c中的注释。
请在实验报告中简要说明你的设计实现过程。
### 答案
做练习一的主要内容就是理解页表项`Page`的结构
```c
struct Page {
    int ref;        // page frame's reference counter
    uint32_t flags; // array of flags that describe the status of the page frame
    unsigned int property;// the num of free block, used in first fit pm manager
    list_entry_t page_link;// free list link
};
```
其中
- `ref`代表页表的**引用计数**，在引入进程后每个进程都会有一个页表，所以和智能指针一样，决定是不是释放`Page`的页是需要`ref`来标识还有没有页表引用这个物理页的。
- `flags`这个标志主要标识两个信息，分别在第0位和第1位：
    1. **是不是保留页**，就是说是不是不能用的页，比如内核代码所在的页就是**保留页**，不能被用。
    2. **是不是连续物理块的页头**，在**首次适配物理分配算法**中，空闲的连续块被空闲链表组织，空闲链表里面是按地址排好序的连续物理块，连续物理块是由一个或多个物理页组成的，所以队头那个就被标志为物理块第一页了。
- `property`，如果当前页是页头，那么`property`代表这页所处块有多少个页。
- `page_link`，用于链接到空闲链表。

#### default_init
不用管，默认实现。
#### default_init_memmap
用于初始化最初的空闲物理页。要注意的点就是给的实现是`list_add(&free_list, &(base->page_link));`它会把新的空闲物理块插到`freelist`的队头去，如果初始化顺序是从小到大初始化的话，那么`freelist`里面的物理块的开始地址的顺序就是从大到小了，所以改为`   list_add_before(&free_list, &(base->page_link));`每次插到队尾即可。
```c
static void
default_init_memmap(struct Page *base, size_t n) {
    assert(n > 0);
    struct Page *p = base;
    for (; p != base + n; p ++) {
        assert(PageReserved(p));
        p->flags = p->property = 0;
        set_page_ref(p, 0);
    }
    SetPageProperty(base);
    base->property = n;

    nr_free += n;
    list_add_before(&free_list, &(base->page_link));
}
```
#### default_alloc_pages
向外提供页分配的接口，流程就是遍历`freelist`，找到第一个足够大就保存那个块的头页指针，然后看看能不能分裂，能的话就分裂成两个，前面的物理块是实际分配出去的物理块，从`freelist`里面删除，后面的设置好`flag`位，让第一个页变成队头，再插入空闲链表里面就行。

```c
static struct Page *
default_alloc_pages(size_t n) {
    assert(n > 0);
    if (n > nr_free) {
        return NULL;
    }
    struct Page *page = NULL;
    list_entry_t *le = &free_list;
    while ((le = list_next(le)) != &free_list) {
        struct Page *p = le2page(le, page_link);
        if (p->property >= n) {
            page = p;
            break;
        }
    }
	
    if (page != NULL) {
        if (page->property > n) {
            struct Page *p = page + n;
            SetPageProperty(p);
            p->property = page->property - n;
            list_add(&(page->page_link), &(p->page_link));
		}
		list_del(&(page->page_link));
        nr_free -= n;
        ClearPageProperty(page);
    }
    return page;
}
```

#### default_free_pages
向外提供释放物理块的接口，流程就是先把物理块里面的页的引用计数全部清零，然后再遍历`freelist`看看能不能合并，要点就是保存合并后的物理块在`freelist`里合适的插入位置。
```c
static void
default_free_pages(struct Page *base, size_t n) {
    assert(n > 0);
    struct Page *p = base;
    for (; p != base + n; p ++) {
        assert(!PageReserved(p) && !PageProperty(p));
        p->flags = 0;
        set_page_ref(p, 0);
    }
    base->property = n;
    SetPageProperty(base);
    list_entry_t *le = list_next(&free_list);
	list_entry_t *nxt = &free_list;
    while (le != &free_list) {
        p = le2page(le, page_link);
        le = list_next(le);
        cprintf("%08p\n", p);
        if (base + base->property == p) {
            base->property += p->property;
            p->property = 0;
            ClearPageProperty(p);
			nxt = (p->page_link).next;
            list_del(&(p->page_link));
        } else if (p + p->property == base) {
            p->property += base->property;
            base->property = 0;
            ClearPageProperty(base);
            base = p;
			nxt = (p->page_link).next;
            list_del(&(p->page_link));
        } else if (base + base->property < p && nxt == NULL) {
			nxt = le;
            break;
		}         
    }
    nr_free += n;
    list_add_before(nxt, &(base->page_link));
}
```
意思是如果合并了，那么就更新插入位置为新合并的那个位置，如果没有可以合并的，那么找到第一个合适的插入位置，由于`freelist`根据首地址排好了序，就找到第一个地址大于块尾地址的就行了。

## 练习2：实现寻找虚拟地址对应的页表项（需要编程）
### 实验
请在实验报告中简要说明你的设计实现过程。
### 答案
`get_pte(pde_t *pgdir, uintptr_t la, bool create)`函数是让我们根据**页表目录起始地址**和**线性地址**（虚拟地址）来得到此虚拟地址的**页表项(page table entry)**，要点就是要理解好**页目录项和页表项**里面的结构。[page_entry](https://github.com/qinggniq/Note/OS/ELF/format_of_page_entry.png)
可以看到，高20位是页表地址/页框地址（注意里面是物理地址），低12位是标志位。我们通过页目录的起始地址可以知道改虚拟地址的二级页表所在的页目录项，通过`pde_t *pdep = &pgdir[PDX(la)]`。然后根据**PTE_P**标识位得知是否有对应的二级页表，如果没有，那么根据`create`标识是否需要新分配一个页来作为二级页表。新分配页表先给`page`设置一下页表引用计数，然后清理一下二级页表所在页的内容（因为后面的程序会根据页表项的**PTE_P**表示判断是不是有虚拟地址到物理地址的映射），最后设置一下页目录项的访问权限标识。最后根据二级页表的起始地址找到虚拟地址所在的页表项，返回即可。
```c
pte_t *
get_pte(pde_t *pgdir, uintptr_t la, bool create) {
	pde_t* pdep = &pgdir[PDX(la)];
	if (!(*pdep & PTE_P)) {
		if (create) {
			struct Page *page = alloc_page();
			if (page == NULL) {
				return NULL;
			}
			set_page_ref(page, 1);
			uintptr_t pa = page2pa(page);
			memset(KADDR(pa), 0, PGSIZE);
			*pdep = pa| PTE_USER;
		} else {
			return NULL;
		}
	}
	uintptr_t* pt_va = KADDR(PDE_ADDR(*pdep));
	//cprintf("here \n");
	pte_t* ptep = &pt_va[PTX(la)];		 
	return ptep;
}
```
## 问答
> 请描述页目录项（Page Directory Entry）和页表项（Page Table Entry）中每个组成部分的含义以及对ucore而言的潜在用处。

| Bit Position | Contents | Use for uCore |
| --| -- | -- |
|0(p) | 存在位，用于表示页表项是否有效|  减小实际的页表所占空间  |
| 1(R/W) |访问控制位，1可写，0可读|  可用于写时拷贝、防止写代码段之类的指令 |
|2(U/S) |权限位，0用户模式，1内核模式|  限制用户访问非法区域  |
|3(PWT)|缓存位，Write-through 数据总是直接写入磁盘，Write-back (or write-behind or Write caching) 数据不是直接被写入磁盘|  感觉有特殊用途 |
| 4(PCD) | 禁止页级缓冲 | 不知|
|5(A) | 访问位 | 不知 |
| 6  | 无用| 无用|
| 7(PS) | 页大小，必须为0|无用|
| 8 - 11| 无用| 无用|
| 12-31 | 地址| 不知|
>  如果ucore执行过程中访问内存，出现了页访问异常，请问硬件要做哪些事情？

- CR0 — Contains system control flags that control operating mode and states of the processor.
- CR1 — Reserved.
- CR2 — Contains the page-fault linear address (the linear address that caused a page fault).
- CR3 — Contains the physical address of the base of the paging-structure hierarchy and two flags (PCD and PWT). Only the most-significant bits (less the lower 12 bits) of the base address are specified; the lower 12 bits of the address are assumed to be 0. The first paging structure must thus be aligned to a page (4-KByte) boundary. The PCD and PWT flags control caching of that paging structure in the processor’s internal data caches (they do not control TLB caching of page-directory information). When using the physical address extension, the CR3 register contains the base address of the page-directory-pointer table. In IA-32e mode, the CR3 register contains the base address of the PML4 table.
- CR4 — Contains a group of flags that enable several architectural extensions, and indicate operating system or executive support for specific processor capabilities.

五个控制寄存器的含义，访问内存异常后，要压入当前的线性地址到`cr2`中，然后就是正常的执行中断服务例程的操作了，压入**EFLAGS**，压入**CS**，压入**EIP**，压入**ERROR_CODE**。

## 练习3：释放某虚地址所在的页并取消对应二级页表项的映射（需要编程）
### 实验
当释放一个包含某虚地址的物理内存页时，需要让对应此物理内存页的管理数据结构Page做相关的清除处理，使得此物理内存页成为空闲；另外还需把表示虚地址与物理地址对应关系的二级页表项清除。请仔细查看和理解page_remove_pte函数中的注释。为此，需要补全在 kern/mm/pmm.c中的page_remove_pte函数。
### 答案
释放比较简单，就是看看页表项是不是真的指向了一个页框，是的话取出指向的页，然后判断页表引用计数是不是为1，为1意味着就是最后一个指向该页的页表也要释放它了，然后就调用`pmm_manager`的`free_page`，free掉就行了，最后把页表项的内容清空即可。
``` c
static inline void
page_remove_pte(pde_t *pgdir, uintptr_t la, pte_t *ptep) {
	if ((*ptep & PTE_P)) {
		struct Page *page = pte2page(*ptep);
		if (page == NULL) return;
		page_ref_dec(page);
		if (page->ref == 0) {
			free_page(page);
			tlb_invalidate(pgdir, la);
		}	
	}
	((pte_t *)KADDR(PDE_ADDR(pgdir[PDX(la)])))[PTX(la)] = NULL;
}
```
### 问答
> 数据结构Page的全局变量（其实是一个数组）的每一项与页表中的页目录项和页表项有无对应关系？如果有，其对应关系是啥？ 

有对应关系，页表项或页目录项如果有**PTE_P**标志的话，那么其中存的物理地址左移12位就是对应的page结构。

> 如果希望虚拟地址与物理地址相等，则需要如何修改lab2，完成此事？ 鼓励通过编程来具体完成这个问题 

把那些物理地址转虚拟地址的宏用到的**KERNBASE**改成0即可。

## Challenge 完成Buddy算法
