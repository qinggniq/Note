# aurora

## background

AWS中的EC2给用户提供虚拟机，用户跑自己的应用：

- 对于无状态的WebServer扩容方便
- 但是DB不方便

在没有EBS之前：

- DB挂了，那么你的数据也没了（因为DB跑在EC2上面）

有了EBS：

- DB所在的EC2连上EBS，DB挂了，可以再开一个EC2跑DB连上之前的EBS

## aurora设计

- 网络负载是有容错数据库（主从复制）的主要瓶颈
- aurora使用redo log来减小同步复制
- database和storage分离，使用qrourm来提供弹性扩容和容错
- 监控storage的日志状态，以实现非qrourm读
- 异步恢复。