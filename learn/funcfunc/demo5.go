/**
  @Author: majm@ushareit.com
  @date: 2020/11/19
  @note:
**/
package main

import "fmt"

/**
defer 调用会在当前函数执行结束前才被执行，这些调用被称为延迟调用 。defer 中使用匿名函数依然是一个闭包。
*/
func main() {

	// 为什么在defer中的x是1而不是101呢？
	// 原因: 在defer定义时 已经将x的拷贝 1 复制给了defer, defer执行时使用的是当时defer定义时x的拷贝，而不是当前环境中x的值。
	x, y := 1, 2
	defer func(a int) {
		fmt.Printf("x:%d,  y:%d \n", a, y)
	}(x)

	x += 100
	y += 100
	fmt.Println(x, y)
}
