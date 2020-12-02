/**
  @Author: majm@ushareit.com
  @date: 2020/12/1
  @note:
**/
package panic_recover

import (
	"fmt"
	"testing"
)

func TestFoo(t *testing.T) {
	Foo()
	fmt.Printf("hello world")
}

func TestFullName(t *testing.T) {

	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	FullName(&firstName, nil)
	fmt.Println("returned normally from main")
}

func TestFullName_Reciver(t *testing.T) {
	defer RecoverName()
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	FullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
