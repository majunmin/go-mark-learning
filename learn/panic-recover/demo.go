/**
  @Author: majm@ushareit.com
  @date: 2020/12/1
  @note:
**/
package panic_recover

import (
	"fmt"
	"runtime/debug"
)

func Foo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var bar *int
	fmt.Println(*bar)
}

func FullName(firstName, lastName *string) {
	defer fmt.Println("deferred call in FullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

// recover 必须在 defer 的 下一层才会生效
func RecoverName() {
	//func(){
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
		debug.PrintStack()
	}
	//}()
}
