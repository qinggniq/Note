# gcc相关选项
```
GCC
    -g  增加gdb的调试信息
    -Wall   显示告警信息
    -O2     优化处理 (有 0，1，2，3，0是不优化)
    -fno-builtin   只接受以"__"开头的内建函数
    -ggdb   让gcc为gdb生成比较丰富的调试信息
    -m32    编译32位程序
    -gstabs     此选项以stabs格式生成调试信息，但是不包括gdb调试信息
    -nostdinc   不在标准系统目录中搜索头文件，只在-l指定的目录中搜索
    -fstack-protector-all   启用堆栈保护，为所有函数插入保护代码
    -E  仅做预处理，不进行编译，汇编和链接
    -x c  指明使用的语言为C语言

LDD Flags
    -nostdlib   不连接系统标准启动文件和标准库文件，只把指定的文件传递给连接器
    -m elf\_i386    使用elf_i386模拟器
    -N      把text和data节设置为可读写，同时取消数据节的页对齐，取消对共享库的链接
    -e func     以符号func的位置作为程序开始运行的位置
    -Ttext addr  是连接时将初始地址重定向为addr （若不注明此，则程序的起始地址为0）
```