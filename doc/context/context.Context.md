# context.Context

```golang

type Context interface {

	// 返回当前 context 截止时间， 如果不存在截止时间, ok = false
	Deadline() (deadline time.Time, ok bool)

	// 返回一个只读信道
	Done() <-chan struct{}
	
	// 返回 context cancel 原因,其值 = context.Canceled or context.DeadlineExceeded
	Err() error
	
	// 获取 当前context 的 key 对应的Value
	Value(key interface{}) interface{}
}

```

`context`包中还包含了四个用于繁衍Context值的函数
WithCancel()  : 会返回一个可撤销的parent子值
WithDeadline(): 会产生一个定时撤销的子值
WithTimeout() : 会产生一个定时撤销的子值
WithValue()   : 会产生一个会携带额外数据的parent的子值



## 撤销信号是如何在上下文中传播的

撤销操作:
对应的Context值会先关闭它内部的接收通道,也就是它的 Done() 会返回的那个通道. 然后,它会向它的所有子值(或者说子节点)传达撤销信号.
这些子值会如法炮制,把撤销信号继续传播下去.  最后，这个Context值会断开它与其父值之间的关联.

![在上下文树中传播撤销信号](https://raw.githubusercontent.com/majunmin/image/master/img%E5%9C%A8%E4%B8%8A%E4%B8%8B%E6%96%87%E6%A0%91%E4%B8%AD%E4%BC%A0%E6%92%AD%E6%92%A4%E9%94%80%E4%BF%A1%E5%8F%B7.png)

> 通过调用 context.WithValue() 得到的Context值是不可撤销的.
> 撤销信号在被传播时,若遇到它们则会直接跨过,并试图将信号直接传给它们的子值



## context.Context 如何携带数据，数据有事如何传递的

WithValue() 在产生新的Context值(以下简称含数据的Context值)的时候需要三个参数
`parentContext`、`key`和`val`  
与"字典对于键的约束"类似,这里键的类型必须是可判等的。

> 1. Context 查找 键值的行为:
> 
> Context类型的Value方法就是被用来获取数据的.在我们调用含数据的Context#Value() 时，它会先判断给定的键，是否与当前值中存储的键相等
> 如果相等就把该值中存储的值直接返回，否则就到其父值中继续查找。
> 如果其父值中仍然未存储相等的键,那么该方法就会沿着上下文根节点的方向一路查找下去.
> 
> 2. 除了含数据的Context值以外,其他几种Context值都是无法携带数据的
>    因此,Context值的Value方法在沿路查找的时候,会直接跨过那几种值





```golang

func contextWithVal() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	valCtx := context.WithValue(ctx, "trace_id", "45454545454")

	go watch(valCtx)

	time.Sleep(time.Second * 10)
	cancelFunc()

	// for protect
	time.Sleep(2 * time.Second)
}

func watch(ctx context.Context) {

	for true {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value("trace_id"), "停止了...", ctx.Err())
			return
		default:
			//取出值
			fmt.Println(ctx.Value("trace_id"), "goroutine process ...")
			time.Sleep(2 * time.Second)
		}
	}
}

```

> Context接口并没有提供改变数据的方法



## "可撤销的"在context包中代表着什么? "撤销"一个Context值又意味着什么?


Done() 会返回一个元素类型为 `struct{}` 的接收通道. 
不过,这个接收通道的用途并不是传递元素值,而是让调用方去感知"撤销"当前Context值的那个信号
