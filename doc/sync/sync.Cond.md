# 条件变量  与 互斥锁

条件变量是基于互斥锁的,必须有互斥锁的支撑才能发挥作用

> 条件变量并不是被用来保护临界区和共享资源的,它是用于协调想要访问共享资源的那些线程的
> 当共享资源的状态发生变化时,它可以被用来通知被互斥锁阻塞的线程
> 条件变量最大的优点是在效率上的提升



## 条件变量怎么与互斥锁合作使用

**条件变量的初始化离不开互斥锁,并且它的方法有的也是基于互斥锁的**

条件变量提供的方法有三个: `等待通知(wait)` `单发通知(signal) ` `广播通知(broadcast)`

```golang
	// 0: 表示有消息 1： 表示无消息
	var mailbox int8

	// lock变量的 Lock() 和 Unlock() 分别用于对其中写锁的锁定和解锁,它们与 sendCond变量 的含义是对应的
	var lock sync.RWMutex

	// sendCond是专门为放置情报而准备的条件变量，向信箱里放置情报,可以被视为`对共享资源的写操作`
	sendCond := sync.NewCond(&lock)
	// recvCond变量代表的是专门为获取情报而准备的条件变量,,可以看做是`对共享资源的读操作`
	recvCond := sync.NewCond(lock.RLocker())

	sign := make(chan struct{}, 1)

	max := 5

	// 发信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 1; i <= max; i++ {
			time.Sleep(1 * time.Second)
			lock.Lock()

			for mailbox == 1 { //信箱中有信 等待
				sendCond.Wait()
			}
			fmt.Printf("sender [%d]: the mailbox is empty. \n", i)
			mailbox = 1
			fmt.Printf("sender [%d]: the mailbox has message. \n", i)
			lock.Unlock()
			recvCond.Signal()
		}
	}(max)

	// 收信
	go func(max int) {
		defer func() {
			sign <- struct{}{}
		}()

		for i := 1; i <= max; i++ {
			time.Sleep(2* time.Second)

			lock.RLock()
			for mailbox == 0 {
				recvCond.Wait()
			}
			fmt.Printf("receiver [%d]: the mailbox is full. \n", i)
			mailbox = 0
			fmt.Printf("receiver [%d]: the letter has been received. \n", i)
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)

	<- sign
	<- sign
```

## 条件变量的Wait方法做了什么？

> 1. 为什么先要锁定条件变量基于的互斥锁, 才能调用它的Wait方法？
> 2. 为什么要用for语句来包裹调用其Wait方法的表达式, 用if语句不行吗？
 

**条件变量的Wait方法主要做了四件事**

1. 把`调用它的 goroutine(也就是当前的 goroutine)加入到当前条件变量的通知队列中`
2. `解锁`当前的条件变量基于的那个互斥锁
3. 让`当前的 goroutine 处于等待状态`,等到通知到来时再决定是否唤醒它.此时,这个 goroutine 就会阻塞在调用这个`Wait()` 的那行代码上
4. 如果通知到来并且决定唤醒这个 goroutine, 那么就在唤`醒它之后重新锁定当前条件变量基于的互斥锁`.自此之后,当前的 goroutine 就会继续执行后面的代码了


> Answer
> 1. 上述第二点
> 2. 为了确保条件改变,这主要是为了保险起见.
>    如果一个 goroutine 因收到通知而被唤醒,但却发现共享资源的状态,依然不符合它的要求,那么就应该再次调用条件变量的Wait(),并继续等待下次通知的到来




## sync.Cond 的 Signal() 和 Broadcast() 有哪些异同？

Signal(): 会唤醒一个 因此而等待的 goroutine
Broadcast(): 会唤醒所有 因此而等待的 goroutine

> 条件变量的 Wait() 总会把当前的 goroutine 添加到通知队列的队尾,而它的 Signal() 总会从通知队列的队首开始, 查找可被唤醒的 goroutine. 
> 所以，因Signal()的通知,而被唤醒的 goroutine 一般都是最早等待的那一个 


**注意**
与Wait()不同,条件变量的 Signal() 和 Broadcast() 并不需要在互斥锁的保护下执行.
恰恰相反,我们最好在解锁条件变量基于的那个互斥锁之后,再去调用它的这两个方法.这更有利于程序的运行效率



