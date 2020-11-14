/**
  @Author: majm@ushareit.com
  @date: 2020/11/13
  @note:
**/
package main

import "fmt"

func main() {
	var a [3]int
	fmt.Printf("cap(a) is %v \n", cap(a))
	fmt.Printf("len(a) is %v \n", len(a))
	assembleArr(a)
	fmt.Printf("%v \n", a) // [0 0 0]

	b := a
	b[0] = 14
	fmt.Printf("%v \n", a) // [0 0 0]

	ca := [...]string{"hello", "world", "cc", "come on"}
	fmt.Printf("len(ca) is %v \n", len(ca))
	fmt.Printf("cap(ca) is %v \n", cap(ca))

	// 二维数组
	sarr := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
	}

	printArr(sarr)

}

// 二维数组打印
func printArr(sarr [3][2]string) {
	for _, v1 := range sarr {
		for _, v2 := range v1 {
			fmt.Printf("%v", v2)
		}
		fmt.Println()
	}
}

func assembleArr(a [3]int) {
	a[0] = 13
	a[1] = 15
	fmt.Printf("modify %v \n", len(a))
}
