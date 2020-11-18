/**
  @Author: majm@ushareit.com
  @date: 2020/11/19
  @note:
**/
package main

import "fmt"

/*
对 golang 闭包的理解
1. 内函数对外函数的变量的修改，是对变量的引用. 共享一个在堆上的变量.  变量被引用后,它所在的函数结束,这变量也不会马上被销毁.相当于变相延长了函数的生命周期.
*/
func main() {
	func1 := OneFunc(18) // 这里定义了 n = 18, 并且执行了 n++, 返回函数return func() { fmt.Println(n) }
	func1()              // 19
	func1()              // 19
	func1()              // 19

	func2 := OtherFunc(18) // 这里定义了 n = 18, 返回函数return func() { n++; fmt.Println(n) }
	func2()                // 19
	func2()                // 20
	func2()                // 21
}

func OneFunc(n int) func() {
	n++
	return func() {
		fmt.Println(n)
	}
}

func OtherFunc(n int) func() {
	return func() {
		n++
		fmt.Println(n)
	}
}
