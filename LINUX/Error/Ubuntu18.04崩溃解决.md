# Ubuntu18.04 桌面卡死解决
## 背景
发现**Ubuntu18.04LTS**的桌面经常被卡死，操作毫无反应，此时电脑的内存已经饱满，交换区也已经饱满，等待电脑反应过来，能给你一次关闭进程刷新页面的机会几乎是不存在的，而且此时想进入`tty`终端往往也是失败的，这时候的解决方法也许就是万能的关机重启了。或许能够进入到tty终端，但是登录时候也是一直卡壳，没有反应。或许成功了，成功杀死进程`Xorg`，重新登录到系统，但是会发现很快就又会变得卡壳了，原因是只是关闭了Ubuntu桌面程序，内存并没有释放掉。

所以最终原因：桌面测程序并不是导致卡壳的更远，而就是内存饱满，交换区饱满导致的，因此新方法是释放内存，释放交换区Swp，将电脑恢复接近到开机的状态。

## 定期清理内存

该操作可能导致部分浏览器页面内容丢失

转：https://blog.csdn.net/qq_21398167/article/details/51657977
```shell
# vim /root/satools/freemem.sh
#!/bin/bash
used=`free -m | awk 'NR==2' | awk '{print $3}'`
free=`free -m | awk 'NR==2' | awk '{print $4}'`
 
echo "===========================" >> /var/log/mem.log
date >> /var/log/mem.log
echo "Memory usage | [Use：${used}MB][Free：${free}MB]" >> /var/log/mem.log
 
if [ $free -le 100 ] ; then
                sync && echo 1 > /proc/sys/vm/drop_caches
                sync && echo 2 > /proc/sys/vm/drop_caches
                sync && echo 3 > /proc/sys/vm/drop_caches
                echo "OK" >> /var/log/mem.log
else
                echo "Not required" >> /var/log/mem.log
fi
```
将脚本添加到crond任务，定时执行。
```
# echo "*/1 * * * * root /root/satools/freemem.sh" >> /etc/crontab
或
crontab -e
添加
*/1 * * * * root /root/satools/freemem.sh
```
(切换到root用户下将上面那句话加入到crontab里，注意格式*之间的空格 )

加入自动释放内存脚本以后，再也没有死机过。
