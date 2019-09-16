# fopen
## fopen打开模式
| fopen() mode | open() flags                  |
| :--: | :--: |
|     r       | O_RDONLY                      | 
|     w       | O_WRONLY | O_CREAT | O_TRUNC  |  |     a       | O_WRONLY | O_CREAT | O_APPEND |
|     r+      | O_RDWR                        |
|     w+      | O_RDWR | O_CREAT | O_TRUNC    |
|     a+      | O_RDWR | O_CREAT | O_APPEND   |
              
## fopen结构体定义
```c++ 
struct _iobuf {
  char *ptr;
  int cnt;
  char *base;
  int flag;
  handle_t file;
  int charbuf;
  int bufsiz;
  int phndl;
};
```