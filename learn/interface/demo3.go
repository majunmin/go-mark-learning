package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old!", p.name, p.age)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am string and my value is %s \n", i.(string))
	case int:
		fmt.Printf("I am int and my value is %d \n", i.(int))
	case Describer:
		i.(Describer).Describe()
	default:
		fmt.Printf("Unknown Type \n")
	}
}
func main() {
	findType("Naveen")
	findType(11)
	findType(89.98)
	findType(Person{name: "majm", age: 18})
}
