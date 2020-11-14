/**
  @Author: majm@ushareit.com
  @date: 2020/11/13
  @note:
**/
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("%v \n", a)

	slice := a[2:5]
	for i := range slice { //i的值为  0, 1, 2
		slice[i]++
	}
	slice[0]++ // 相当于 a[2]++

	// 切片的长度是切片中的元素数. 切片的容量是从创建切片索引开始的底层数组中元素数。
	fmt.Printf("len(slice) is %v \n", len(slice))
	fmt.Printf("cap(slice) is %v \n", cap(slice))
	fmt.Printf("%v \n", a)

	// append 使其 capacity变为原来的2倍， len 即为 切片中元素的数量
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) // capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) // capacity of cars is doubled to 6

	// 切片的类型零值是 nil
	// 一个 nil 切片的长度和容量为 0。可以使用 append 函数将值追加到 nil 切片
	var ss []string
	fmt.Println(len(ss))
	fmt.Println(cap(ss))

	ss = append(ss, "nil", "Sam")
	fmt.Println(len(ss))
	fmt.Println(cap(ss))

	// 切片的函数传递
	var slia = []int{0, 14, 18}
	assembleSlice(slia)       // 切片传递的是地址， 数组传递至是 值的一份拷贝
	fmt.Printf("%v \n", slia) // [16 17 18]

}

func assembleSlice(slice []int) {
	slice[0] = 16
	slice[1] = 17
}
