# 关于shellscript传字符串的小坑.md
## 缘起
自己在`.bashrc`里面写了一个函数`gacp comment`来简化`git add . &&  git commit -m "comment" && git push`的过程，函数是这样的：
```shell
function gacp() {
    git add . &&  git commit -m $1 && git push
}
```
就是将参数作为commit的描述然后执行一套而已，然而在实际执行时总是发生一下错误：
```shell
qinggniq@qinggniq-G519:~/Git/Letters$ gacp "send 19-09-09.md"
error: pathspec '19-09-09.md' did not match any file(s) known to git.
```
## 原因
实验了

```shell
git commit -m send 19-09-09.md
```
得到和上述一样的错误后，终于回忆起来shell的一个坑了，就是`$1`在shell脚步里面会去掉引号，也就是说在你执行`gacp "comment"`时你预期shell脚步会执行`git add . &&  git commit -m "comment" && git push`，然而它执行的是`git add . &&  git commit -m  comment && git push`，如果`comment`中间有空格，那么就会被**git**识别为多个参数，从而报错。
## 解决
```shell
function gacp() {
    git add . &&  git commit -m "$1 " && git push
}
```
给参数手动加一个引号即可。