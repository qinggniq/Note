# MergeSort分析

## 朴素的MergeSort

单线程，$O(n)$额外空间分配的归并算法在时间上已经战胜了go内置的`sort`，显然归并排序可以通过多核的优势并行排序原始数组的各个部分，使用多协程来优化`MergeSort`的时间性能。

![image-20200212204331668](/Users/qinggniq/Library/Application Support/typora-user-images/image-20200212204331668.png)

## 多线程优化时间

由于排序过程是一个**CPU密集型**的任务，所以我们每个核分配数组的一个部分就行了。假如是2核CPU，数组有100个元素，那么可以给前50个数分配一个goroutine，后50个分配一个goroutine。找到分配的边界的任务可以递归完成。

![image-20200212210327261](/Users/qinggniq/Library/Application Support/typora-user-images/image-20200212210327261.png)

物理核为4核，逻辑核为8核的情况下，性能比单线程快了大约4倍，符合预期。

优化完时间复杂度之后，优化空间复杂度。

## 原地归并优化空间

C++中有原地归并算法，网上也有人移植到各种语言，Go语言内置的归并排序也是使用的这种合并算法。算法来源于这篇论文[Jyrki Katajainen, Tomi Pasanen, Jukka Teuhola. ``Practical in-place mergesort''. Nordic Journal of Computing, 1996](http://akira.ruc.dk/~keld/teaching/algoritmedesign_f04/Artikler/04/Huang88.pdf)，大致思想就是通过交换来实现原地合并。

把网上http://thomas.baudel.name/Visualisation/VisuTri/inplacestablesort.html的实现移植到了golang上面。

![image-20200212233312429](/Users/qinggniq/Library/Application Support/typora-user-images/image-20200212233312429.png)

可以看见时间上来说和单线程的sort相差不大（因为`merge`过程的时间复杂度从$O(n)$变成了$O(nlog(n))$，所以总的时间复杂度也变成了$O(nlognlogn)$），但是空间分配相比于原来的归并算法少了很多。

