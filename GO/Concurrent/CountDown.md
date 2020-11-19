# CountDownLatch

## 概念

在 Java 中**CountDownLatch** 类是一个同步组件，允许一个线程在其他一个或多个线程执行完成之后再执行。

在GoLang中对应的是**sync.WaitGroup**

## 使用场景

一个线程需要其他线程执行完成之后才能执行。比如一个多线程计数程序，从线程计数，主线程打印计数结果，这个时候就需要主线程等待从线程执行完计数之后才能执行自己的打印，需要用到**CountDownLatch**.

```go
func main() {
  var mutex sync.mutex
  var count int
  var wg sync.WaitGroup
  wg.Add(2)
  countFunc := func(up int) {
    defer wg.Down()
    for i := 0; i < up; i++ {
      mutex.lock()
      count++
      mutex.unlock()
    }
  }
  go countFunc(10)
  go countFunc(100)
  wg.Wait()
  println(count)
}
```

## 实现[Go]

在Golang中，`sync.waitGroup`提供三个接口：

- `Add(int)`：相当于`sync.WaitGroup`在内部维护了一个计数器，当需要等待一个线程执行的时候，`wg.Add(1)`，如果想等待多个线程，那么就对应`wg.Add(N)`相应的线程数。
- `Done()`：一个线程执行结束之后需要调用`wg.Done()`来使计数器减一。
- `Wait()`：当计数器为0之后，`wg.Wait()`返回，线程可以继续执行了。

一个朴素的做法是在`sync.WaitGroup`里面维护一个原子整数，然后对应的三个API实现如下：

```go
type WaitGroup struct{
  count int32
}

func (wg *WaitGroup) Add(delta int) {
  atomic.AddInt32(&wg.count, delta)
}

func (wg *WaitGroup) Done() {
  wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
  for remain := atomic.LoadInt32(*wg.count); reamin > 0{}
}
```

在这个实现下，`Done()`个数大于`Add()`所添加的个数依然可以执行，但是这和语意不符。所以可以给`Add()`添加一个判断，当count为负数时`panic`。

```go
func (wg *WaitGroup) Add(delta int) {
  atomic.AddInt32(&wg.count, delta)
  if remain := atomic.Load(&wg.count); remain < 0 {
    panic("negative WaitGroup Counter")
  }
}
```

虽然增加了错误处理，但是这个`Wait()`十分低效，它是**busy-waiting**，而不会被调度到其他线程去执行，在`c/c++`中我们可以使用`mutex + condition variable`的组合去实现**线程阻塞和唤醒**操作，那么在GoLang中如何实现。

如下是[golang](https://golang.org/src/sync/waitgroup.go)中的实现，去掉了检测数据竞争的代码。

```go
package sync

import (
	"sync/atomic"
	"unsafe"
)

type WaitGroup struct {
  // WaitGroup不能被拷贝
	noCopy noCopy
  // stat1分为counter/64部分和sem/32部分
  // 其中
  // - counter/64中高32位统计Add添加的计数
  // - counter/64中低32位统计调用Wait()的线程数
  // - sem/32用于实现阻塞和唤醒操作
  // 而counter部分是在state1前面还是后面取决于机器是32位还是64位
	state1 [3]uint32
}


func (wg *WaitGroup) state() (statep *uint64, semap *uint32) {
  // 检测是不是64位机器，地址数被8整除说明是64位机器
	if uintptr(unsafe.Pointer(&wg.state1))%8 == 0 {
    // 64位机器counter部分在前面，sem在后面
		return (*uint64)(unsafe.Pointer(&wg.state1)), &wg.state1[2]
	} else {
    // 32位机器counter部分在后面，sem在前面
		return (*uint64)(unsafe.Pointer(&wg.state1[1])), &wg.state1[0]
	}
}

func (wg *WaitGroup) Add(delta int) {
	statep, semap := wg.state()
  // 计数器 +delta
	state := atomic.AddUint64(statep, uint64(delta)<<32)
  // 获得计数
	v := int32(state >> 32)
  // 获得等待线程数
	w := uint32(state)
	if v < 0 {
		panic("sync: negative WaitGroup counter")
	}
  // 在调用Wait之后才有Add操作，可能是Add函数和Wait并发，但是正确的用法是Add操作和Wait操作需要在同一个线程。
	if w != 0 && delta > 0 && v == int32(delta) {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	if v > 0 || w == 0 {
		return
	}
  // 在Add的时候发生了并发调用了Add、Done、Wait操作导致状态变化
	if *statep != state {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
  // 此时计数器为0并且有等待计数器为0的线程
	*statep = 0
  // 唤醒
	for ; w != 0; w-- {
		runtime_Semrelease(semap, false, 0)
	}
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	statep, semap := wg.state()
	for {
		state := atomic.LoadUint64(statep)
		v := int32(state >> 32)
		w := uint32(state)
		if v == 0 {
			return
		}
		// 如果是第一次循环，增加counter中等待线程计数器
		if atomic.CompareAndSwapUint64(statep, state, state+1) {
      // 阻塞，直到被唤醒
			runtime_Semacquire(semap)
			if *statep != 0 {
				panic("sync: WaitGroup is reused before previous Wait has returned")
			}
			return
		}
	}
}

```

