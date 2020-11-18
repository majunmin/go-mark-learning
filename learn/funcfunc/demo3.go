/**
  @Author: majm@ushareit.com
  @date: 2020/11/19
  @note:
**/
package main

import (
	"fmt"
	"time"
)

func main() {
	strs := []string{"a", "b", "c"}
	for _, str := range strs {
		go func() {
			fmt.Println(str)
		}() // 函数类型()
	}

	// 输出 c c c ?
	// 原因: 由于上面是异步执行的， 此时 打印的str的值，取决 该匿名函数执行的时机
	//       当循环执行完 再执行打印语句时 会是c c c
	// 也验证了  闭包函数内部持有外部变量的`引用`
	time.Sleep(time.Second * 1)

	// 2, 如何 使程序还打印 a b c ？
	// 只需要每次将变量 str 的拷贝传进函数即可，但此时就不是使用的上下文环境中的变量了。
	for _, str := range strs {
		go func(str string) {
			fmt.Println(str)
		}(str) // 函数类型()
	}
}
