/**
  @Author: majm@ushareit.com
  @date: 2020/12/13
  @note:
**/
package interface_learn

import (
	"fmt"
	"os"
	"testing"
)

func TestDemo1(t *testing.T) {
	Name := MyString("qwerpvnavopqtqpeovba")
	var v VowelsFinder
	v = Name // valid only if MyString implements VowelsFinder
	fmt.Printf("Vowels are %c ", v.FindVowels())
}

func TestDemo2(t *testing.T) {
	emp1 := Permanent{empId: 1, basicPay: 8000, pf: 540}
	emp2 := Contract{empId: 2, basicPay: 8000}
	employees := []SalaryCalculater{emp1, emp2}
	totalExpense(employees)
}

func TestDemo3(t *testing.T) {
	findType("Naveen")
	findType(11)
	findType(89.98)
	findType(Person{name: "majm", age: 18})
}

func TestDemo4(t *testing.T) {
	fmt.Println("hello.xlsx world!")
	fmt.Fprintln(UpperWriter{os.Stdout}, "Hello, world!")

	// 2
	fmt.Fprintln(os.Stdout, UpperString("Hello, world!"))
}
