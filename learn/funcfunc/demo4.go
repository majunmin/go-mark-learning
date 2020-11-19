/**
  @Author: majm@ushareit.com
  @date: 2020/11/19
  @note:
**/
package main

import (
	"fmt"
	"unsafe"
)

/*
当closure所在函数重新调用时: 其closure是新的，其context引用的变量也是重新在heap定义过的。
*/
func main() {
	nextInt := intSeq()
	fmt.Println(unsafe.Pointer(&nextInt))
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	nextInt2 := intSeq()
	fmt.Println(unsafe.Pointer(&nextInt2))
	fmt.Println(nextInt2()) // 1
	fmt.Println(nextInt2()) // 2
	fmt.Println(nextInt2()) // 3
}

func intSeq() func() int {
	i := 0
	fmt.Println("closure init ....")
	fmt.Println(unsafe.Pointer(&i))

	return func() int {
		i += 1
		fmt.Println("in closure")
		fmt.Println(unsafe.Pointer(&i))
		return i
	}
}
