# etcd raft

### etcd特性

1. 线性一致读
2. 配置更改
3. 预选举
4. 快照、日志压缩
5. 

```go
type raftLog struct {
  logs []Entry
  commit int
}

func (l *raftLog) applyTo(idx int) {} 

```

```go
type raft struct {
  msg []Msg	//这个用来收集要发的msg，待到合适的时机去发
}
```

节点状态

```go
//State of raft object
const (
	StateFollower StateType = iota
	StateCandidate
	StateLeader
	StatePreCandidate
)
```



在每个raft内部都有不同的消息，有的消息是内部用的，有的消息是外部发过来的。内部的消息不需要放到消息队列里面。

### 所有的消息类型

```go
//MsgType of different msg
const (
	MsgApp MsgType = iota			//append请求  leader => follower
  MsgHup[local]							//hup请求			follower => follower
	MsgAppResp								//append响应	follower => leader
	MsgBeat[local]						//广播请求		 leader => leader
	MsgHeartBeat							//心跳消息		 leader	=> follower
	MsgHeartBeatResp					//心跳回复		 follower => leader
	MsgProp										//提议消息		 client => leader [follower => leader]
	MsgPropResp								//提议回复		 leader => client
	MsgVote										//投票消息		 candidate => follower
	MsgVoteResp								//投票回复		 follower => candidate
	MsgPreVote								//预投票消息		preCandidate => follower
	MsgPreVoteResp						//预投票回复		follower => preCandidate
)
```



### leader需要处理的Msg

- MsgBeat[local]
- MsgProp
- MsgAppResp 			//根据这个更新follower的matchIndex和nextIndex状态
- MsgHeartBeatResp   //根据这个决定一个follower是不是应该

### leader需要发送的Msg

- MsgHeartBeat
- MsgApp

### pre/candidate需要处理的Msg

- MsgHeartBeat
- MsgApp
- MsgVoteResp
- MsgPreVoteResp

### pre/candidate需要发送的Msg

- MsgVote
- MsgPreVote

### follower需要处理的Msg

- MsgApp
- MsgHeartBeat
- MsgVote
- MsgPreVote

### follower需要发送的Msg

- MsgAppResp
- MsgHeartBeatResp
- MsgVoteResp
- MsgPreVoteResp
- Msg

