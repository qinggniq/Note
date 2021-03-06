# 写一个RPC框架需要什么

## RPC框架是什么

- 是一套工具集
- 能让你的代码远程调用另一个另一个进程的代码片段

## 需要考虑什么

- 高性能？
- 支持的语言多样性？
- 协议的多样性？
- API简洁性？

## 需要写什么

- 数据传输
- 数据序列化/反序列化
- 代码生成（根据proto文件生成相应的服务代码）
- 应用层
  - 客户端
  - 服务器
  - 线程池
- 协议层
  - 兼容？
  - 二进制格式
  - 压缩格式
  - Json格式
- 网络传输层
  - Http
  - Buffered

![Apache Thrift Layered Architecture](%E5%86%99%E4%B8%80%E4%B8%AARPC%E6%A1%86%E6%9E%B6%E9%9C%80%E8%A6%81%E4%BB%80%E4%B9%88/thrift-layers.png)

thrift的主要内容分两部分，一个是compiler，通过thrift.pd文件生成代码，第二个部分是各个语言的编码/解码、通信、序列化/反序列化的一些代码。

thrift和dubbo的区别，一个是rpc框架，一个是微服务治理的框架，可以负责服务的熔断、重试，负载均衡，服务注册。