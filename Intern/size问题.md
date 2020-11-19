# 关于CanalEventBatchEncoder的Size函数的实现

## 原有实现

```go
// Size implements the EventBatchEncoder interface
func (d *CanalEventBatchEncoder) Size() int {
	// TODO: avoid marshaling the messages every time for calculating the size of the packet
	err := d.refreshPacketBody()
	if err != nil {
		panic(err)
	}
	return proto.Size(d.packet)
}

// refreshPacketBody() marshals the messages to the packet body
func (d *CanalEventBatchEncoder) refreshPacketBody() error {
	oldSize := len(d.packet.Body)
	newSize := proto.Size(d.messages)
	if newSize > oldSize {
		// resize packet body slice
		d.packet.Body = append(d.packet.Body, make([]byte, newSize-oldSize)...)
	}
	_, err := d.messages.MarshalToSizedBuffer(d.packet.Body[:newSize])
	return err
}

```

也就是说，调用Size的时候，每次都会Mashal一下messages，然后调用protobuf的size方法去得到d.packet被Mashal后的大小。其中packet被Marshal后的大小取决于两个方面：

1. messages的大小，这部分取决于我们添加了多少DML，也就是RowChangeEvent，DML的数量和大小都会影响messages的大小。
2. canal.Packet结构体中的其他部分，这个部分在encoder初始化的时候就已经决定。

以上两个部分被Mashal后的buffer大小就是Size()应该返回的值。Size() = Len(Marshal(canal.Packet{others, Marshal(messages)}))

## 现有实现

现在，我们把RowChangeEvent暂存在TxnCache中，只有调用Build的时候，我们才会实际的去Mashal这些Event。每个事务一个canal.Packet，也就是说，对于一次Build()，产生出来的canal.Packet数目是不确定的，所以Size()是无法返回我们最终Build出来的MQmessage.Value的大小的。

我们可以假装Encoder只包含一个事务，并且Build的时候encode所有的event。基于这个假设，Size()可以这样实现：

1. 使用一个messageEncoder（这个encoder就是原有不支持事物的encoder的实现），在appendRowChangeEvent的时候同时添加到txnCache和messageEncoder，现有实现就直接调用用messageEncoder的Size()来获得结果。当外部调用Build()的时候，删掉原来的messageEncoder，创建一个新的messageEncoder，将剩下那些大于resolveTS的event重新加入messageEncoder。

   - 缺点
     - 实现会增加一倍的空间复杂度
     - 对于Build()，最坏会增加一倍的时间复杂度（所有的Event都大于resolvedTs）。
     - 对于Size()，时间 = 序列化Packet的时间
     - 和实际Build出来的value size会有差异
   - 优点
     - 实现简单

2. 做一个类似txnCache类似的canalMessageCache，同样提供txnCache类似的Split，Resolved功能，但是存储的不是txn，而是canal.packet，每次append event的时候判断StartTs和CommitTs，找到对应的canalSingleTableTxnMessageEncoder，然后调用实际的canalMessageEncoder去将event转化为canal.RowChangeEvent，调用Size()时，遍历unresolvedTxnMessage中的所有canalSingleTableTxnMessageEncoder，调用对应的canalMessageEncoder的Size()方法，累计相加即可。

   ```go
   // CanalEventBatchEncoder encodes the events into the byte of a batch into.
   type CanalEventBatchEncoder struct {
   	forceHkPk  bool
   	size       int
   	resolvedTs uint64
   	messageCache   *canalMessageCache
   }
   type canalMessageCache struct {
     unresolvedTxnMessage map[model.TableID][]*canalSingleTableTxnMessageEncoder
   }
   type canalSingleTableTxnMessageEncoder struct {
     Table     *TableName
   	StartTs   uint64
   	CommitTs  uint64
   	messageEncoder  *canalMessageEncoder
   	ReplicaID uint64
   }
   type canalMessageEncoder struct {
   	messages     *canal.Messages
   	packet       *canal.Packet
   	entryBuilder *canalEntryBuilder
   }
   func (c *canalMessageCache) Append(event) {
    	1. 找到event归属的canalSingleTableTxnMessageEncoder
     2. 调用canalSingleTableTxnMessageEncoder的Append方法，这个方法类似于model.SingleTableTxns的Append方法，会做一些检查
     3. 检查之后，会调用canalMessageEncoder的Append方法，这个方法做event => canal.RowChangeEvent的工作
   }
   func (c *CanalEventBatchEncoder) Size() {
    	1. 获得unresolvedTxnMessage
     2. 遍历所有canalSingleTableTxnMessageEncoder
     3. 调用对应的canalMessageEncoder.Size()
     4. 累加
     5. 返回
   }
   ```

   - 优点
     - 空间复杂度没有增加
     - 时间复杂度和原始不支持事务的CanalEventBatchEncoder的Size()复杂度类似
   - 缺点
     - 和txnCache实现类似，如果txnCache实现逻辑上有改动，那么可能同时需要修改canalMessageCache

canalMessageCache和txnCache功能类似，都是将所属不同事务的event封装起来，但是由于protobuf message需要Mashal之后才能知道size，为了更好的空间和时间复杂度，才需要去实现一个和txnCache功能类似的canalMessageCache，两者的区间仅限于canalSingleTableTxnMessageEncoder和SingleTableTxn的区别，前者具有“事务表示 + Encoding + Size”的功能，而后者只有"事务表示"的功能。问题的根源在于：

根据单一职责原则，encoder应该只有encoding这个功能，而txnCache应该只有txnGenerator的功能，现在encoder同时具备了encoding和txnGenerator的功能，违背了原则。而在现在的实现中，txnGenerator和encoding这两个职责比较独立，一旦调用一些需要encoding和txnGenerator同时工作的函数，就需要耗费很长时间，Size()和Build()都需要encoding和txnGenerator的参与，而Build()是一个一次性的函数，它的调用会减少下一次的复杂度，而Size()是一个会被频繁调用的函数，并且它的调用不会减小下一次调用的复杂度。

现在的支持事务的encoder，它的功能之一txnGenerator应该对标到mysqlSink，如果mysqlSink没有根据Size来判断是不是应该flush，那么canalEncoder也不应该根据Size来判断是不是应该flush；同时它的功能之一encoding应该对标到mqSink，mqSink要求根据Sink判断是否应该flush，那么canalEncoder需要根据Size判断flush。否则要么用更高的复杂度去支持Size()，要么用破坏DRY原则的代价去支持Size()。当然可能会有更好的设计，我没想出来。

