/**
  @Author: majm@ushareit.com
  @date: 2020/11/30
  @note:
**/
package unsafe_pointer

import (
	"fmt"
	"unsafe"
)

/*
// error 编译失败
	num := 5
	numPointer := &num

	flnum := (*float32)(numPointer)
	fmt.Println(*flnum)
*/
func Testcase1() {

	num := 5
	numPointer := &num

	flnum := (*float32)(unsafe.Pointer(numPointer))
	fmt.Printf("fnum: %f", *flnum)
}

type Num struct {
	i string
	j int64
}

// unsafe.Pointer 只能作用于 指针的转换
func Testcase2() {
	n := &Num{
		i: "Encrypt",
		j: 100,
	}

	nPointer := unsafe.Pointer(n)
	niPointer := (*string)(unsafe.Pointer(nPointer))
	*niPointer = "majm"

	njPointer := (*int64)(unsafe.Pointer(uintptr(nPointer) + unsafe.Offsetof(n.j)))
	*njPointer = 1000

	fmt.Printf("n.i : %s, n.j: %d", n.i, n.j)
}
