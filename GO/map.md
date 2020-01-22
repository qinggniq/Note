# Golang 中的 map

- 一般性能在于 `hash`函数和拉链出来的链表里面找。

### 一般的存储模型

- map -[1 : n]> bucket -[1:8]> cell
- 额外装不下的就用 

### key类型上面的区别

- 使用如下的类型编译器可以将其替换为更高效的实现

| key 类型 | 查找                                                         |
| :------- | :----------------------------------------------------------- |
| uint32   | mapaccess1_fast32(t *maptype, h *hmap, key uint32) unsafe.Pointer |
| uint32   | mapaccess2_fast32(t *maptype, h *hmap, key uint32) (unsafe.Pointer, bool) |
| uint64   | mapaccess1_fast64(t *maptype, h *hmap, key uint64) unsafe.Pointer |
| uint64   | mapaccess2_fast64(t *maptype, h *hmap, key uint64) (unsafe.Pointer, bool) |
| string   | mapaccess1_faststr(t *maptype, h *hmap, ky string) unsafe.Pointer |
| string   | mapaccess2_faststr(t *maptype, h *hmap, ky string) (unsafe.Pointer, bool) |

这些函数的参数类型直接是具体的 uint32、unt64、string，在函数内部由于提前知晓了 key 的类型，所以内存布局是很清楚的，因此能节省很多操作，提高效率。

上面这些函数都是在文件 `src/runtime/hashmap_fast.go` 里。



### map扩容

- 装载因子 $map\_count \div bucket\_count$

  1. 装载因子超过阈值，源码里定义的阈值是 6.5。
  2. overflow 的 bucket 数量过多：当 $B < 15$，也就是 bucket 总数 $2^B < 2^{15}$ 时，如果 overflow 的 bucket 数量超过 $2^B$；当 $B \ge 15$，也就是 bucket 总数 $2^B \ge 2^{15}$，如果 overflow 的 bucket 数量超过$ 2^{15}$。

  ```go
  // If we hit the max load factor or we have too many overflow buckets,
  	// and we're not already in the middle of growing, start growing.
  	if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
  		hashGrow(t, h)
  		goto again // Growing the table invalidates everything, so try again
  	}
  
  // overLoadFactor reports whether count items placed in 1<<B buckets is over loadFactor.
  func overLoadFactor(count int, B uint8) bool {
  	return count > bucketCnt && uintptr(count) > loadFactorNum*(bucketShift(B)/loadFactorDen)
  }
  // tooManyOverflowBuckets reports whether noverflow buckets is too many for a map with 1<<B buckets.
  // Note that most of these overflow buckets must be in sparse use;
  // if use was dense, then we'd have already triggered regular map growth.
  func tooManyOverflowBuckets(noverflow uint16, B uint8) bool {
  	// If the threshold is too low, we do extraneous work.
  	// If the threshold is too high, maps that grow and shrink can hold on to lots of unused memory.
  	// "too many" means (approximately) as many overflow buckets as regular buckets.
  	// See incrnoverflow for more details.
  	if B > 15 {
  		B = 15
  	}
  	// The compiler doesn't see here that B < 16; mask B to generate shorter shift code.
  	return noverflow >= uint16(1)<<(B&15)
  }
  ```

  

### 两种不同的查询效率低的情况

1. 第 1 点：我们知道，**每个 bucket 有 8 个空位**，在没有溢出，且所有的桶都装满了的情况下，装载因子算出来的结果是 8。因此当装载因子超过 6.5 时，表明很多 bucket 都快要装满了，查找效率和插入效率都变低了。在这个时候进行扩容是有必要的。

2. 第 2 点：是对第 1 点的补充。就是说在装载因子比较小的情况下，这时候 map 的查找和插入效率也很低，而第 1 点识别不出来这种情况。表面现象就是计算装载因子的分子比较小，**即 map 里元素总数少，但是 bucket 数量多（真实分配的 bucket 数量多，包括大量的 overflow bucket**）。

### 两种查询效率低情况下的解决办法

1. 对于条件 1，元素太多，而 bucket 数量太少，很简单：将 B 加 1，bucket 最大数量（2^B）直接变成原来 bucket 数量的 2 倍。于是，就有新老 bucket 了。注意，这时候元素都在老 bucket 里，还没迁移到新的 bucket 来。而且，新 **bucket 只是最大数量变为原来最大数量（2^B）的 2 倍（2^B * 2）**。

2. 对于条件 2，其实元素没那么多，但是 overflow bucket 数特别多，说明很多 bucket 都没装满。**解决办法就是开辟一个新 bucket 空间，将老 bucket 中的元素移动到新 bucket，使得同一个 bucket 中的 key 排列地更紧密**。这样，原来，在 overflow bucket 中的 key 可以移动到 bucket 中来。结果是节省空间，提高 bucket 利用率，map 的查找和插入效率自然就会提升。
   1. 但是如果是因为hash值都弄到一个bucket里面造成的话，即使新开辟一个也没有用。

```go
//实际的扩容代码
func hashGrow(t *maptype, h *hmap) {
	// If we've hit the load factor, get bigger.
	// Otherwise, there are too many overflow buckets,
	// so keep the same number of buckets and "grow" laterally.
  // 扩容一倍
	bigger := uint8(1)
  // 如果是情况二的话就不需要扩容了
	if !overLoadFactor(h.count+1, h.B) {
		bigger = 0
		h.flags |= sameSizeGrow
	}
	oldbuckets := h.buckets
	newbuckets, nextOverflow := makeBucketArray(t, h.B+bigger, nil)

	flags := h.flags &^ (iterator | oldIterator)
	if h.flags&iterator != 0 {
		flags |= oldIterator
	}
	// commit the grow (atomic wrt gc)
	h.B += bigger
	h.flags = flags
	h.oldbuckets = oldbuckets
	h.buckets = newbuckets
  //搬迁进度
	h.nevacuate = 0
  //overflow bucket的数量
	h.noverflow = 0

	if h.extra != nil && h.extra.overflow != nil {
		// Promote current overflow buckets to the old generation.
		if h.extra.oldoverflow != nil {
			throw("oldoverflow is not nil")
		}
		h.extra.oldoverflow = h.extra.overflow
		h.extra.overflow = nil
	}
	if nextOverflow != nil {
		if h.extra == nil {
			h.extra = new(mapextra)
		}
		h.extra.nextOverflow = nextOverflow
	}

```



### 扩容后旧数据的处理办法

- 渐进式hash（和redis一样），在插入、删除的时候对旧数据进行搬迁。

### 和redis的区别

- redis里面的bucket里面的正常cell不是固定的（8），overflow设计不一样
- 渐进式的搬运都是搬的bucket。