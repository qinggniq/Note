# golang中http.Get.Body.ReadString做法

- 调用的fd.中的read，一次最多读1G。
- 使用ReadString -> ReadSlice -> fd.Read
  - 维护一个4k的buffer，每次读的时候就读满，然后ReadString的时候就会从已经读的地方找`delim`。

