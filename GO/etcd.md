# etcd raft

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

