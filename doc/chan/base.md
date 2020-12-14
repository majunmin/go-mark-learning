# Chan

## 对通道的发送和接收操作都有哪些基本的特性？

它们的基本特性如下:
1. 对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的
2. 发送操作和接收操作中对元素值的处理都是不可分割的
3. 发送操作在完全完成之前会被阻塞,接收操作也是如此

## 1. 通道的接收 和发送 是互斥的

对于通道中的同一个元素值来说,发送操作和接收操作之间也是互斥的.
例如，虽然会出现，正在被复制进通道但还未复制完成的元素值,但是这时它绝不会被想接收它的一方看到和取走.
1. **这里要注意的一个细节是,元素值从外界进入通道时会被复制.更具体地说，进入通道的并不是在接收操作符右边的那个元素值,而是它的副本.**
2. **元素值从通道进入外界时会被移动.这个移动操作实际上包含了两步:1.是生成正在通道中的这个元素值的副本,并准备给到接收方 2.是删除在通道中的这个元素值.**



## 2. 发送操作和接收操作中对元素值的处理都是不可分割的

这里的`不可分割`的意思是,它们处理元素值时都是原子操作,绝不会被打断.
这既是为了保证通道中元素值的完整性,也是为了保证通道操作的唯一性.

## 3. 发送操作在完全完成之前会被阻塞,接收操作也是如此

如此阻塞代码其实就是为了实现操作的互斥和元素值的完整。

等待的,所有接收操作所在的 goroutine,都会按照先后顺序被放入通道内部的接收等待队列.


## 发送操作和接收操作在什么时候会引发 panic？

对于一个已初始化,但并未关闭的通道来说,收发操作一定不会引发 panic.
但是通道一旦关闭,再对它进行发送操作,就会引发 panic.

> 对于已关闭的通道执行发送操作，会引发 panic
> 对于已关闭的通道执行发送接收，不会引发panic,一般会收到通道中值类型的零值


## 单向通道

单向通道最主要的用途就是约束其他代码的行为,为了安全性

## select + chan

select语句只能与通道联用,它一般由若干个分支组成. 每次执行这种语句的时候,一般只有一个分支中的代码会被运行.

```golang
// 准备好几个通道。
intChannels := [3]chan int{ 
	make(chan int, 1),
	make(chan int, 1),
	make(chan int, 1),
}
// 随机选择一个通道，并向它发送元素值。
index := rand.Intn(3)fmt.Printf("The index: %d\n", index)intChannels[index] <- index // 哪一个通道中有可取的元素值，哪个对应的分支就会被执行。
select {
case <-intChannels[0]: fmt.Println("The first candidate case is selected.")
case <-intChannels[1]: fmt.Println("The second candidate case is selected.")
case elem := <-intChannels[2]: fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
default: fmt.Println("No candidate case is selected!")
}
```