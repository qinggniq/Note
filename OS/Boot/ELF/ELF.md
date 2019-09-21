# ELF文件格式
## 作用
用于描述**目标文件**的格式。
## 类别
1. 可重定位文件（*relocatable file*）包含适合链接到其他文件的代码和数据。
2. 可执行文件（*executable file*）。
3. 共享目标文件（*shared object file*）包含适合两种链接的代码和数据。
## 文件格式
![elf](https://github.com/qinggniq/Note/Images/ELF_view.png)
## ELF Header
```c
#define EI_NIDENT 16
typedef struct {
    unsigned char e_ident[EI_NIDENT];       //机器相关数据
    Elf32_Half e_type;                                          //文件类型
    Elf32_Half e_machine;                                  //体系结构
    Elf32_Word e_version;                                   //版本
    Elf32_Addr e_entry;                                         //转移控制权的地址
    Elf32_Off e_phoff;                                             //程序头偏移量
    Elf32_Off e_shoff;                                              //区头偏移量
    Elf32_Word e_flags;             
    Elf32_Half e_ehsize;
    Elf32_Half e_phentsize;
    Elf32_Half e_phnum;
    Elf32_Half e_shentsize;
    Elf32_Half e_shnum;
    Elf32_Half e_shstrndx;
 } Elf32_Ehdr;
```