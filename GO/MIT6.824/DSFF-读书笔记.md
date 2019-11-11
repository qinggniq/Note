分布式程序解决两个问题：

1. 信息以光速传输（速度）
2. 独立的事独立地失败（多种事情）

distance time and consistency models

DRDTs CALM理论

## Basic (terms and concepts)

**goals**
scalability, availability, performance, latency, fault tolerance (?same with availability)

## Up and down the level of abstraction

CAP theorem, FLP impossibility result, number of consistency models

## Time and Order

time and order and clocks

## Replication: preventing divergence

replication method for mantaining single-copy consistency (2PC to Paxos)

## Replication: accepting divergence

weak consistency guarantees, CRDTs and CALM

## 0x01 High level of distributed systems

- finite resource
- unsuitable for hardware upgrade
- 