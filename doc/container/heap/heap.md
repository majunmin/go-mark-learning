# heap.Interface

## heap.Interface

```go
type Interface interface {
    sort.Interface
    Push(x interface{}) // add x as element Len()
    Pop() interface{}   // remove and return element Len() - 1.
}

// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
```

如果我们要自定义实现一个 堆的话,只要实现以上五个方法就可以了



### 成员函数

heap包中提供了几个最基本的堆操作函数,包括`Init` `Fix` `Push` `Pop` `Remove` (其中up, down函数为非导出函数)
这些函数都通过调用前面实现接口里的方法,对堆进行操作

```go

// 初始化 heap
func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}

// Push和Pop是一对标准堆操作,
// Push向堆添加一个新元素,Pop弹出并返回堆顶元素,而在push和pop操作不会破坏堆的结构
func Push(h Interface, x interface{}) {
    h.Push(x)
    up(h, h.Len()-1)
}

func Pop(h Interface) interface{} {
    n := h.Len() - 1
    h.Swap(0, n)
    down(h, 0, n)
    return h.Pop()
}

// 删除堆中的第i个元素,并保持堆的约束性
func Remove(h Interface, i int) interface{} {
    n := h.Len() - 1
    if n != i {
        h.Swap(i, n)
        if !down(h, i, n) {
            up(h, i)
        }
    }
    return h.Pop()
}

// 在修改第i个元素后,调用本函数修复堆,比删除第i个元素后插入新元素更有效率。
func Fix(h Interface, i int) {
    if !down(h, i, h.Len()) {
        up(h, i)
    }
}

func up(h Interface, j int) {
    for {
        i := (j - 1) / 2 // parent
        if i == j || !h.Less(j, i) {
            break
        }
        h.Swap(i, j)
        j = i
    }
}

func down(h Interface, i0, n int) bool {
    i := i0
    for {
        j1 := 2*i + 1
        if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
            break
        }
        j := j1 // left child
        if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
            j = j2 // = 2*i + 2  // right child
        }
        if !h.Less(j, i) {
            break
        }
        h.Swap(i, j)
        i = j
    }
    return i > i0
}

```


## 二叉堆的实现

堆的实现有很很多种，二叉堆,斐波那契堆,23堆
二叉堆实现起来相对容容易,性能相对较好
http://note.youdao.com/s/1iOp8Ha4


### 堆 插入元素 (Push)(大顶堆)

将元素插入到数组尾部,和其父节点比较大小,如果小于其父节点和其父节点交换(heapifyUp,up),直到其小于其父节点或者到达堆顶 [自底向上]




### 堆 删除最大值 (Pop)

arr[0]即为最大值
将数组arr[len-1]尾部元素放到数组头部arr[0],比较该值和其两个子节点的最大值，如果小于其子节点的值，则和 两个子节点的最大值交换 (heapifyDown, down),
直到不满足条件或者到达树的叶子节点 [自顶向下],最后截取数组长度 `arr[0:len-1]`


### 删除其中一个值(Remove)







