# unsafe pointer

> Go 的指针是不支持指针运算和转换
> Go 是一门静态语言，所有的变量都必须为标量类型。不同的类型不能够进行赋值、计算等跨类型的操作。
> 那么指针也对应着相对的类型，也在 Compile 的静态类型检查的范围内。
> 同时静态语言，也称为强类型。也就是一旦定义了，就不能再改变它



- 任何类型的指针值都可以转换为 Pointer
- Pointer 可以转换为任何类型的指针值
- uintptr 可以转换为 Pointer
- Pointer 可以转换为 uintptr



## unsafe.Offsetof()

```go
type Num struct{
    i string
    j int64
}

func main(){
    n := Num{i: "EDDYCJY", j: 1}
    nPointer := unsafe.Pointer(&n)

    niPointer := (*string)(unsafe.Pointer(nPointer))
    *niPointer = "煎鱼"

    njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
    *njPointer = 2

    fmt.Printf("n.i: %s, n.j: %d", n.i, n.j)
}

```

在剖析这段代码做了什么事之前，我们需要了解结构体的一些基本概念：

1. 结构体的成员变量在内存存储上是一段连续的内存
2. **结构体的初始地址就是第一个成员变量的内存地址**
3. 基于结构体的成员地址去计算偏移量. 就能够得出其他成员变量的内存地址
4. 再回来看看上述代码，得出执行流程：
  * 修改 `n.i` 值：i 为第一个成员变量。因此不需要进行偏移量计算，直接取出指针后转换为 Pointer，再强制转换为字符串类型的指针值即可
  * 修改 `n.j` 值：j 为第二个成员变量。需要进行偏移量计算，才可以对其内存地址进行修改。在进行了偏移运算后，当前地址已经指向第二个成员变量。接着重复转换赋值即可
  
1. uintptr：uintptr 是 Go 的内置类型。返回无符号整数，可存储一个完整的地址。后续常用于指针运算
```
type uintptr uintptr
```

2. unsafe.Offsetof：返回成员变量 x 在结构体当中的偏移量。更具体的讲，就是返回结构体初始位置到 x 之间的字节数。需要注意的是入参 `ArbitraryType` 表示任意类型，并非定义的 int。它实际作用是一个占位符
```
func Offsetof(x ArbitraryType) uintptr
```



