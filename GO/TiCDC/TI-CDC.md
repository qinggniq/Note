#TICDC

## 作用

用于增量同步TIDB的数据变化，然后下发到MySQL，Mq，下游的服务。

## 模块

### entry模块

把tikv的数据变更转化为model模块中的格式格式。

### model模块

是一个中间层，tikv -> model -> canal/opencdc

### sink模块

把拉到的数据下发到消费端，比如MySQL和Kafka。

### puller模块

拉取tikv的数据变更到processer，然后交给entry模块去解析成model格式。然后放到sorter那里去sort，然后把这个行变化通过sink模块发送到下游。

### processer模块

连接所有模块的一个模块



## 问题

### 分区问题

一个表的rowdatachange是不是应该放到一个partition？  

- mysql的是一个表一个list，list里面按时间顺序存了row data change，然后根据resolved ts来做事务的划分
- mq是根据row的某些属性（表名、rowid、时间戳）来弄的

然后Kafka consumer需要的是decode mqsink发过来的事务，然后再sink到mysql，然后就可以验证正确性了。

- Kafka consumer目前消费的是mq messge type类型的消息，根据mq decoder来重新变成model里面的数据类型，然后再交给mysql sink去验证。



Emitcheckpoint 就是要求所有之前的timeStamp的rowdata change都被sink到下游。

