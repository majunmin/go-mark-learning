package main

import "fmt"

/*
这里需要理解 golang 是 值传递 or 引用传递
1. 函数参数的传递都是值传递，也就是传递原值的一个副本，对于`整形 字符串 bool 数组` 等非引用类型，  传递的是值的副本,
   还是 `数组切片(slice) 映射(map) 通道(channel) 方法 函数`等引用类型,                       传递的是引用的副本,
2.

summary:
 go中只存在 值传递 (值的副本 or 指针的副本)
 */
func main() {
    i := 10
    ip := &i
    // 1. 非引用类型传值
    fmt.Printf("main: i的值是 %v, i 的内存地址是: %v, i的指针ip的内存地址: %v \n", i, ip, &ip)
    modify(i)
    fmt.Printf("main: i的值是 %v, i 的内存地址是: %v, i的指针ip的内存地址: %v \n", i, ip, &ip)

    // 2. 非引用类型传地址
    fmt.Printf("main: i的值是 %v, i 的内存地址是: %v, i的指针ip的内存地址: %v \n", i, ip, &ip)
    modifyP(&i)
    fmt.Printf("main: i的值是 %v, i 的内存地址是: %v, i的指针ip的内存地址: %v \n", i, ip, &ip)

    // 2. 引用类型
    // slice 的内存地址 == slice总第一个元素的 内存地址
    arr := []int{1, 2, 3}
    fmt.Printf("arr : %p \n", arr)
    fmt.Printf("&arr[0] : %p \n", &arr[0])
    fmt.Printf("&arr : %p \n", &arr)
    // arr :     0xc000014140
    // &arr[0] : 0xc000014140
    // &arr :    0xc00000c060

}

// main:   i的值是 10, i 的内存地址是: 0xc0000b4008, i的指针ip的内存地址: 0xc0000ae018
// modify: i的值是 10, i 的内存地址是: 0xc0000b4018, i的指针ip的内存地址: 0xc0000ae028
// main:   i的值是 10, i 的内存地址是: 0xc0000b4008, i的指针ip的内存地址: 0xc0000ae018
func modify(i int) {
    ip := &i
    fmt.Printf("modify: i的值是 %v, i 的内存地址是: %v, i的指针ip的内存地址: %v \n", i, ip, &ip)
    i = 11
}

//这里传递的是 i的内存地址的副本,(所以 i的的内存地址是一样的， i的内存地址的内存地址是不一样的)
//
// main:    i的值是 10, i 的内存地址是: 0xc0000180b0, i的指针ip的内存地址: 0xc00000e028
// modifyP: i的值是 10, i 的内存地址是: 0xc0000180b0, i的指针ip的内存地址: 0xc00000e038
// main:    i的值是 11, i 的内存地址是: 0xc0000180b0, i的指针ip的内存地址: 0xc00000e028
func modifyP(i *int) {
    fmt.Printf("modify: i的值是 %v, i 的内存地址是: %v, i的指针ip的内存地址: %v \n", *i, i, &i)
    *i = 11
}
