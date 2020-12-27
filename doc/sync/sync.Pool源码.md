
```golang
type Pool struct {
	// go1.7 引入的一个静态检查机制
	// 因为 Pool 不希望被复制，所以结构体里有一个 noCopy 的字段，使用 `go vet` 工具可以检测到用户代码是否复制了 Pool
	noCopy noCopy
	
	// 每个P的本地队列，实际类型为 [P]poolLocal 
	local     unsafe.Pointer // local fixed-size per-P pool, actual type is [P]poolLocal 
	// [P]poolLocal的size
	localSize uintptr        // size of the local array

	victim     unsafe.Pointer // local from previous cycle
	victimSize uintptr        // size of victims array

	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
    // 自定义的对象创建回调函数, 当 pool 中无可用对象时会调用此函数
	New func() interface{}
}


// 本地池列表
type poolLocalInternal struct {
    private interface{} // Can be used only by the respective P. 私有临时对象
    shared  poolChain   // Local P can pushHead/popHead; any P can popTail.  共享临时对象列表
                        // poolChain是一个双端队列 poolDequeue 的实现
}
```


Pool结构体
![Pool](https://raw.githubusercontent.com/majunmin/image/master/imgPool.png)

![sync.Pool.Put流程](https://raw.githubusercontent.com/majunmin/image/master/imgsync.Pool.Put.png)

![sync.Pool.Get流程](https://raw.githubusercontent.com/majunmin/image/master/imgsync.Pool.Get.png)









## noCopy

```golang
// noCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
type noCopy struct{}

// Lock is a no-op used by -copylocks checker from `go vet`.(静态检查)
func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

```

## pool.pin()

> pin 的作用就是将当前 groutine 和 P 绑定在一起，禁止抢占. 并且返回对应的 poolLocal 以及 P 的 id
> 
> 

```golang
// 调用方必须在完成取值后调用 runtime_procUnpin() 来取消抢占。
func (p *Pool) pin() (*poolLocal, int) {
	pid := runtime_procPin()
	// In pinSlow we store to local and then to localSize, here we load in opposite order.
	// Since we've disabled preemption, GC cannot happen in between.
	// Thus here we must observe local at least as large localSize.
	// We can observe a newer/larger local, it is fine (we must observe its zero-initialized-ness).
	s := atomic.LoadUintptr(&p.localSize) // load-acquire
	l := p.local                          // load-consume
	// 因为可能存在动态的 P(运行时调整 P 的个数)
	if uintptr(pid) < s {
		return indexLocal(l, pid), pid
	}
	return p.pinSlow()
}

```


## getSlow()

> 如果在 shared 里没有获取到缓存对象，则继续调用 Pool.getSlow()，尝试从其他 P 的 poolLocal 偷取：



## pack/unpack

> 它们实际上是一组绑定、解绑 head 和 tail 指针的两个函数。




## GC

对于 Pool 而言，并不能无限扩展，否则对象占用内存太多了，会引起内存溢出。
GO在GC时会自动清理未使用的对象

```golang
// pool.go
func init() {
	runtime_registerPoolCleanup(poolCleanup)
}
```
向运行时注册 池清理函数，会在gc发生时执行




-----

[深度解密 Go 语言之 sync.Pool](https://www.cnblogs.com/qcrao-2018/p/12736031.html)

