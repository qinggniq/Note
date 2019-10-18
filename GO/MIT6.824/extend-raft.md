# 分布式共识算法（扩展）
## 摘要
**Raft**适用于管理复制日志的分布式共识算法，相比于**Paxos**算法，它更易理解，提供了更好的基础来实现系统。为了让人们更好地理解它，**Raft**分离了**共识问题**中的关键元素——**选主**、**日志复制**、**安全性**，并且它有更低的一致性要求，从而降低了**共识问题**中状态的数量。
## 介绍
**共识算法**允许许多机器在一些机器失败的情况下依然能协同工作，所以被广泛用于大型软件系统。最近的十年里**Paxos**[^15][^16]共识算法一直是学术界和工业界的标准算法。然而**Paxos**过于复杂，很难理解和编码。所以作者发明了一种新的共识算法——**Raft**算法用于教学和系统构建，这个算法的首要目标就是**易于理解**。
**Raft**算法通过分解共识问题中的概念，将其分为**选主**、**日志复制**、**安全性**三个部分和降低一致性要求来降低状态数量的方法来让其更好理解。**Raft**算法在很多方面和其他的共识算法很像，但是有三个新特性：
- **强领导：**日志文件只能从Leader发给follower，这样可以简化复制日志的管理。
- **选主：**使用随机定时器来选主。
- **角色转换：**使用新的**联合共识**机制来改变集合中机器的状态，从而使得集群能在配置更改的时候继续工作。

## 复制状态机

# Draft
## 选主 
- 请求投票
## 正常操作
- `requestVote()`
- `duplicateLogEntry()`
- 心跳，不带`command`的`duplicateLogEntry()`操作
- 绝大多数**Follower**执行并提交之后才发`response`。=> 确保了在**Leader**挂了之后的选主阶段选出来的新**Leader**包含了所有的`committed logs`
- 

## 安全性和一致性
### leader挂了
- `checkStatus()`归纳法证明正确性。
- `(index, term)`确定选**Leader**的时候选哪个。 => 用于保证**Leader**里面有所有`committed logs`。
### follower的一致性
- `next_index, (index, term), checkStatus()`来检查直到第一个状态一致的log。
### follower挂了
- **Leader**一直retry就行。
## 废弃的Leader
- 旧的**Leader**可能是网卡了一下，然后新的**Leader**被选出来了，所以旧**Leader**会重新`duplicateLogEntry()`
- `(index, term)`用于确定`follower()`是不是再认回来原来的**Leader**。
- $term_{old\_leader} > term_{follower\_or\_new\_leader}$，那么现在的**Leader**退位给旧**Leader**，否则旧**Leader**变成**Follower**。
## 用户规则
- 设置`timeout`，超时重发即可。
- 保证`command`**正好一次**执行。
- 大于等于一次：执行完成才提交。
- 不超过一次：通过给已提交的`command`添加`unique id`确保不会重复执行。
## 配置更改
- 2phase提交。
- *joint consensus*阶段用于确保在$C_{old,new}$的时候不会产生新的**Leader**。

## 问题
- `committed log`是**follower**自己确定的还是**leader**确定的 => **Leader**确定的。