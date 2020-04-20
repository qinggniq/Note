# etcd raft

## 角色

### PreCandidate

在Basic Raft的基础上新增加了**PreCandidate**角色，用来避免因为网络分区的时候无法选出`Leader`导致`Term`无限制的增长的一个优化手段。具体步骤就是：

- 当定时器超时的时候也没有收到**Leader的心跳**或者**其他PreCandidata/Candidate的预投票/投票请求**，那么从**Folloer**转换为**PreCandidate**状态。
- 成为**PreCandiate**之后先进行一次**预投票**，这个时候不增加`Term`，只是用来检测集群各个节点的网络状态，
  - 如果集群**达到选出Leader的条件**（也就是健康的节点大于所有节点数目的一半），那么就转换为**Candidate**状态，`Term++`，开始真正的选举
  - 如果不能**达到选出Leader的条件**，什么都不做，等待定时器超时。
  - 其他节点收到**预选举**的请求，（？？？直接返回？？？）

### Leaner

这个状态并不在core raft算法状态机转换里面，而是为了解决**新节点加入集群导致集群不稳定**的问题而引入的一个状态，在**Leaner**状态，这个节点只用来同步**Leader**发过来的同步消息，它不能投票和自己选举，所以在**Vote**和**Log Replicate**的时候**Leaner**节点不算**quorum**值，不会影响选主和提交日志的决策。关于**Leaner**状态更细节的解释可以看[etcd doc](https://etcd.io/docs/v3.3.12/learning/learner/)。

不过在实际处理的时候，leaner还是会进行投票。具体看[实际代码的处理流程](https://github.com/qinggniq/etcd/blob/reading/raft/raft.go#L941)

### leader transfer

有些情况下**Leader**需要将领导权转让给其他节点：

- **Leader**节点需要重启，所以需要将自己的领导权转让给其他节点，虽然可以不用转让，直接shutdown，然后让其他节点自动超时导致重新选举出新的节点。但是会造成更多不必要的步骤
- 有其他节点更适合当领导，比如一个广域网的集群，通过选择合适的**Leader**可以使节点通信延迟降到最小。

转让步骤

- 原先的**Leader**停止接受客户端请求
- 发送给目标节点**Leader**所有日志。
- 发送TimeOut给目标节点，目标节点收到之后开始选举（`term++`，然后发送**Vote**给其他节点），原来的Leader因为收到了`term`更大的请求，就会转换为**Follower**状态。

### Leader CheckQuorum

在**etcd**里面，**Leader**不仅有心跳超时器，还有选举超时器，选举超时器超时的时候就会发一条消息来确认当前集群的状态。并且如果这个时候还在**领导换届**的状态，说明换届没成功，此时取消换届。。

### Leader Election 

```go
func (r *raft) compaign(t CompaignType) {
  //如果不是Leaner并且没有被移除集群
  if !canElection {
    return
  }
	//根据当前选举类型进行状态转换
  if t == compaignPreElection {
    r.becomePreCandidate()
  }else{
    r.becomeCandidate()
  }
  //如果集群里面只有自己一个，那么就直接变成Leader。
  if 加上自己投票之后获得了绝大多数节点的同意 {
    if 类型是“预选举” {
      进行正式选举
    } else if 类型是“正式选举” {
      成为Leader
    }
    return
  }
  for 节点列表 {
    给其他节点发送“请求[预投票/投票]消息”
  }
}
```



### 消息的处理

Step里面公共的部分是安全性原则。



不同版本的状态转换becomeXXX，不同的处理消息的函数step，公共的Step函数用于所有角色统一的处理。

### 集群节点状态追踪

tracker追踪各个节点的状态。

- MatchIndex 和 NextIndex 论文上面的
- VoteMap 投票状态
- Inflight 消息缓存，表示正在发送的消息，类似TCP的滑动窗口
- Status 节点的状态，用于控制对发送速度



### Soft状态/Hard状态

soft状态用于日志和debug。

- Lead
- Status

Hard状态是需要持久化到磁盘上的数据。

### Ready结构体

是对消息的封装，

- 持久化到磁盘，
- 提交
- 发送给其他节点

### Node处理

一个节点需要处理不同的消息，并且发出一些消息。

- PropC 客户端发给Leader的提议
- recvC 其他节点发过来的消息
- confC 某个地方？发过来的配置更改请求
- tickC 驱动算法超时机制的运转
- readyc 节点收集自己产生的消息，包装成Ready发送给readyC让它发给其他节点并且持久化
- advancec ？？？？
- stopC 用于停止节点的运行

但是那处理消息的入口都是raft.Step

### RawNode

RawNode就是对Raft基本算法的一个封装，加上的Raft算法需要持久化或者用于log的一些状态。

### Storage

存储里面有三个结构，一个

- HardStatus，用于存储Raft节点需要保存到磁盘上的数据，比如Vote，Commit，。。。

- SnapShot，已经压缩后的日志，任何想要查询这个部分里面的日志的请求，都会直接发送快照，而不是再解压。。
- ents，还没有合并到快照里面的log

但是index不是按照ents里面的算的，而是snapshot压缩的长度加上ents的长度。

### raftLog

对应Raft算法里面的日志复制状态机。保存四个东西：

```go
type raftLog struct {
	//持久化过的日志
	storage Storage

	//将要被持久化的日志
	unstable unstable

	//提交index
	committed uint64
	//应用到机器上的最后一条日志的index
	applied uint64
}
```

在日志复制里面，分不同角色有对日志有不同的操作：

#### Leader

>  接受客户端的请求，然后把客户端的命令封装为Log添加到Raftlog里面去，然后更新自己的matchIndex和NextIndex，然后检查是不是可以更新提交日志了。

`raft.stepLeader -> raft.appendEntry -> raft.raftLog.append`  

#### Candidate

Candidate对日志的操作基本没有，因为如果它收到了日志复制的消息，要么是旧Leader发过来的，这个时候丢弃它；要么是新的Leader发过来，这个时候需要变成Follower状态。

### Follower

Follower有两个地方需要和RaftLog打交道：

- Candidate发起Vote的时候，Follower会根据LogTerm和LogIndex去检查Candidate的日志是不是至少比自己新
- Leader发Append消息的时候，Follower需要检查这些日志是不是和自己匹配，然后append日志

`raft.stepFollower -> raft.handleAppendEntries -> raft.Log.maybeAppend -> raft.Log.matchTerm `



