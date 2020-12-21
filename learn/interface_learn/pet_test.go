/**
  @Author: majm@ushareit.com
  @date: 2020/12/13
  @note:
**/
package interface_learn

import (
	"fmt"
	"testing"
)

func TestFirst(t *testing.T) {
	cat := Cat{name: "cat"} // 如果 cat := &Cat{name: "cat"}， 结构又是怎样的呢？
	var pet = cat
	cat.SetName("wangcai")
	fmt.Println(pet.Name()) //? wangcai
	fmt.Println(cat.Name()) //? cat
}

func TestSecond(t *testing.T) {
	cat := Cat{name: "cat"} // 如果 cat := &Cat{name: "cat"}， 结构又是怎样的呢？
	pet := cat
	cat.SetName("wangcai")
	fmt.Println(pet.Name()) //? wangcai
	fmt.Println(cat.Name()) //? cat

	//NewCat("mofan").SetName("mifan") //为啥编译不通过? 因为 NewCat("mofan")是不可寻址的
}

func TestNil(t *testing.T) {
	var cat *Cat
	fmt.Println("the first cat is nil. [wrap1]")
	cat2 := cat
	fmt.Println("the second cat is nil. [wrap2]")

	var pet Pet = cat2
	if pet != nil {
		fmt.Println("pet is not nil. [wrap1]") // ?
	} else {
		fmt.Println("pet is nil. [wrap1]")
	}
}
