# heap.Interface

二叉堆的性质:
父节点序号N: 左子节点索引 2N+1  右子节点索引: 2N+2



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


### 删除其中一个值(Remove(i))

与 Pop() 方法如出一辙


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

heap包中提供了几个最基本的堆操作函数,包括`Init` `Fix` `Push` `Pop` `Remove` (其中**up**, **down**函数为非导出函数)
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


### down(h Interface, i0, n int) bool  -- 自顶向下

return： bool : 表示 h[i0]是否交换(下沉)过

将有序数组h 的 节点 h[i0] 和其两个子节点交换 (一直交换到叶子节点)
最大堆: 找到两个子节点中的最大值 且 大于h[i0] 的值,则交换 h[i0] 与  子节点最大值
最小堆: 找到两个子节点中的最小值 且 小于h[i0] 的值,则交换 h[i0] 与  子节点最小值

最大堆而言:
当遍历到该节点比起两个子节点都大时,即可退跳出循环,



### heap.Init()

Init(): 将一个有序数组结构化为一个 heap (最大堆|最小堆)
down():
n/2-1 : 获取 最底层 最右侧 节点的父节点, 可以自己画图看看 (二叉堆的性质)


### up(h Interface, j int)  -- 自底向上

对于最大堆而言，

j = (i-1)/2 : 找到 i 节点的 父节点 j
如果 节点 h[i] > 节点h[j] 进行交换 上浮  (一直交换到根节点) 
     一旦 i == j || h[i] <= 节点h[j] 就跳出循环

### heap.Push()
1. 调用自定义 h.Push()
2. up(): 



### heap.Pop()

交换 h 头尾元素，并将 h 头元素下沉 (down 函数)到合适位置

最后将 调用 自定义的 h.Pop()


### heap.Fix(h Interface, i int)

当修改堆中的元素的值的时候,通过调用此函数来进行对堆的修复
h[i]能下沉就下沉, 否则上浮

```golang
func Fix(h Interface, i int) {
    if !down(h, i, h.Len()) {
        up(h, i)
    }
}
```


### heap.Remove(h Interface, i int) interface{}

删除 堆中 第i个元素, 并返回被删除的元素

如果 i== n(删除 h 中最后一个元素,则直接调用 自定义 h.Pop())

1. 交换h 中 第 i 与 第 n 个元素
2. 并调用 Fix() 修复堆
3. 调用自定义 h.Pop()







