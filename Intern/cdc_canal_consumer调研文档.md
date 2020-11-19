### cdc canal consumer调研文档

#### 背景

cdc中canal协议目前没有完备的支持和测试，需要编写canal consumer从而进行集成测试用于验证canal协议的正确性。目前有两套方案：

#### 方案一

使用go编写canal consumer。仿照kafka consumer编写canal consumer，然后使用mysqlSink同步上下游TiDB服务器，以此验证正确性。

##### 优点

- 对原有的集成测试修改不大
- 编码和现有的kafka consumer类似

##### 缺点

- 需要对现有的kafka consumer有较深的理解
- 开发时间未知，可能需要写很多单元测试

#### 方案二

使用alibaba canal server中提供的client.adapter实现同步。配置canal client adpater消费kafka的消息并同步到指定mysql数据库。

##### 优点

- 不需要自己写canal consumer，consumer本身的正确性可以得到保证

##### 缺点

- 可能会对现有的集成测试脚本有较大修改
- 不确定现有的mqSink和alibaba canal client通信是否兼容

#### 选择

选择方案二。

#### 工期

3-4天。

#### 实现思路

- 根据[网上教程](https://www.jianshu.com/p/5bcf97335e71)，下载canal.adpater，修改配置文件中mq的地址为mqSink的kafka地址
- 启动两个TIDB，开启cdc server，和canal.adpater，输入sql语句测试是否兼容。
- 编写shell脚本进行更完备的测试。

