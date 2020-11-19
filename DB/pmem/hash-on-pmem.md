# **Dash: Scalable Hashing on Persistent Memory**

## 有意思的点

1. 论文说对于hash的场景来说，pmem上的read操作其实是比write操作更需要优化的点。因为对于Hash场景，不管是Get还是Set，都涉及到一系列的查找操作，其中涉及到很多cache miss，成为影响性能的关键。
2. 并发控制大部分是用的分区锁，然而分区锁的时候又要涉及到PMEM的读，所以作者是用cas来代替锁，读的时候无锁。
3. 通过在每个bucket加一个footfinger来过滤掉不必要的比较和读取，这样就减少的cache miss。
4. 通过精密设计的探测操作，让hash表的负载因子尽量的大，

