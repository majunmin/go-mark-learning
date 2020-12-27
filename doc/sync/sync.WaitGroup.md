# sync.WaitGroup

`sync.WaitGroup`类型是开箱即用的，也是并发安全的 同时,与我们前面讨论的几个同步工具一样,它一旦被真正使用就不能被复制了

WaitGroup类型拥有三个指针方法:
Add()
Done()
Wait()

> 你可以想象该类型中有一个计数器,它的默认值是0
> 通过调用该类型值的 Add() 来增加,或者调用 Done() 减少这个计数器的值
> 一般情况下，我们用这个Add() 来记录 并行执行 groutine 的数量
> Wait() 的作用阻塞当前 gorouttine， 直到计数器 归0
> 


`sync.WaitGroup` 使用方式:
**我们最好用"先统一Add,再并发Done,最后Wait"这种标准方式,来使用WaitGroup值.**
尤其不要在调用 Wait() 的同时,并发地通过调用 Add() 去增加其计数器的值,因为这也有可能引发 panic





## sync.Once 类型的值怎么保证 仅执行一次

```golang
type Once struct {
    done uint32
    m    Mutex
}

func (o *Once) Do(f func ()) {

    if atomic.LoadUint32(&o.done) == 0 {
        // Outlined slow-path to allow inlining of the fast-path.
        o.doSlow(f)
    }
}

func (o *Once) doSlow(f func ()) {
    o.m.Lock()
    defer o.m.Unlock()
    if o.done == 0 {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}
```

`sync.Once`类型也属于结构体类型，同样也是开箱即用和并发安全的 由于这个类型中包含了一个`sync.Mutex`类型的字段 所以,复制该类型的值也会导致功能的失效

`done`: 记录其所属的 Do() 被调用的次数,该字段的值仅为 0|1,一旦Do()首次调用完成,他的值就会从 0 -> 1


## Do() 功能上的两个特点

1. 由于 Do() 只会在参数函数执行结束之后把done字段的值变为1.
   因此,如果参数函数的执行需要很长时间或者根本就不会结束(比如执行一些守护任务),那么就有可能会导致相关 goroutine 的同时阻塞
   
2. Do方法在参数函数执行结束后,对done字段的赋值用的是原子操作. 并且,这一操作是被挂在defer语句中的.
   因此,不论参数函数的执行会以怎样的方式结束(panic...),done字段的值都会变为1.
