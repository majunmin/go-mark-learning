/**
  @Author: majm@ushareit.com
  @date: 2020/11/19
  @note:
**/
package main

import "fmt"

func main() {
	A := closure()
	fmt.Println(A(1))
	fmt.Println(A(2))
	fmt.Println(A(1))
}

func closure() func(int) int {
	var x int
	return func(a int) int {
		x++
		return a + x
	}
}
