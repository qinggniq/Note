# ACID
## 原子性（Atomicity）
事务开始后所有操作，要么全部做完，要么全部不做，不可能停滞在中间环节。事务执行过程中出错，会回滚到事务开始前的状态，所有的操作就像没有发生一样。例如，如果一个事务需要新增 100 条记录，但是在新增了 10 条记录之后就失败了，那么数据库将回滚对这 10 条新增的记录。也就是说事务是一个不可分割的整体，就像化学中学过的原子，是物质构成的基本单位。

注：MySQL使用事务来实现原子性

## 一致性（Consistency）
指事务将数据库从一种状态转变为另一种一致的的状态。事务开始前和结束后，数据库的完整性约束没有被破坏。例如工号带有唯一属性，如果经过一个修改工号的事务后，工号变的非唯一了，则表明一致性遭到了破坏。

注：MySQL的一致性由程序保证
## 隔离性（Isolation）
要求每个读写事务的对象对其他事务的操作对象能互相分离，即该事务提交前对其他事务不可见。 也可以理解为多个事务并发访问时，事务之间是隔离的，一个事务不应该影响其它事务运行效果。这指的是在并发环境中，当不同的事务同时操纵相同的数据时，每个事务都有各自的完整数据空间。由并发事务所做的修改必须与任何其他并发事务所做的修改隔离。例如一个用户在更新自己的个人信息的同时，是不能看到系统管理员也在更新该用户的个人信息（此时更新事务还未提交）。

注：MySQL 通过锁机制来保证事务的隔离性。

## 持久性（Durability）
事务一旦提交，则其结果就是永久性的。即使发生宕机的故障，数据库也能将数据恢复，也就是说事务完成后，事务对数据库的所有更新将被保存到数据库，不能回滚。这只是从事务本身的角度来保证，排除 RDBMS（关系型数据库管理系统，例如 Oracle、MySQL 等）本身发生的故障。

注：MySQL 使用 redo log 来保证事务的持久性。