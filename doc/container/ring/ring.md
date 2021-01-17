# ring

```go
type Ring struct {
	next, prev *Ring
	Value      interface{} // for use by client; untouched by this library
}


```

```go
type Ring
    func New(n int) *Ring               // 用于创建一个新的 Ring, 接收一个整形参数，用于初始化 Ring 的长度  
    func (r *Ring) Len() int            // 环长度
    
    func (r *Ring) Next() *Ring         // 返回当前元素的下个元素
    func (r *Ring) Prev() *Ring         // 返回当前元素的上个元素
    func (r *Ring) Move(n int) *Ring    // 指针从当前元素开始向后移动或者向前(n 可以为负数)

    // Link & Unlink 组合起来可以对多个链表进行管理
    func (r *Ring) Link(s *Ring) *Ring  // 将两个 ring 连接到一起 (r 不能为空)
    func (r *Ring) Unlink(n int) *Ring  // 从当前元素开始，删除 n 个元素

    func (r *Ring) Do(f func(interface{}))  // Do 会依次将每个节点的 Value 当作参数调用这个函数 f, 实际上这是策略方法的引用，通过传递不同的函数以在同一个 ring 上实现多种不同的操作。
```