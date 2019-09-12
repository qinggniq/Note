# send
|   返回值 n    |                          返回值含义                          |
| :-----------: | :----------------------------------------------------------: |
|    大于 0     |                      成功发送 n 个字节                       |
|       0       |                         对端关闭连接                         |
| 小于 0（ -1） | 出错或者被信号中断或者对端 TCP 窗口太小数据发不出去（send）或者当前网卡缓冲区已无数据可收（recv） |

## 发送0个字节的数据
`send`实际上不会发送0字节的数据，从而保证了返回值为0时的行为的一致性。