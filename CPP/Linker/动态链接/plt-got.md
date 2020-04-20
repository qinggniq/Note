# 动态链接过程

## got表

`got`表存动态链接函数的地址，一开始的时候地址为000000。

动态库需要重定位的函数在`.got.plt`段。

动态库里面需要重定位的函数在.got.plt这个段里面，我们看下：[![动态链接库中函数的地址确定---PLT和GOT](http://blog.chinaunix.net/attachment/201209/16/24774106_1347796206Ly7h.png)](http://blog.chinaunix.net/attachment/201209/16/24774106_1347796206Ly7h.png)



 .got.plt这个段的起始地址是0x8049ff4。 .got.plt这个section大小为0x24 = 36,可是我们只有6个需要解析地址的function，4*6=24个字节，只需要24个字节就能存放这6个函数指针。多出来的12个字节是dynamic段地址，ModuleID 和 _dl_runtime_resolve的地址，如下图所示

[![动态链接库中函数的地址确定---PLT和GOT](http://blog.chinaunix.net/attachment/201209/16/24774106_1347798530g9dm.png)](http://blog.chinaunix.net/attachment/201209/16/24774106_1347798530g9dm.png)



