# fopen打开模式
| fopen() mode | open() flags                  |
| :--: | :--: |
|     r       | O_RDONLY                      | 
|     w       | O_WRONLY | O_CREAT | O_TRUNC  |  |     a       | O_WRONLY | O_CREAT | O_APPEND |
|     r+      | O_RDWR                        |
|     w+      | O_RDWR | O_CREAT | O_TRUNC    |
|     a+      | O_RDWR | O_CREAT | O_APPEND   |
              
