/**
  @Author: majm@ushareit.com
  @date: 2020/12/1
  @note:
**/
package reflectreflect

import (
	"fmt"
	"go-mark-learning/common"
	"reflect"
)

func Demo1() {
	//var stu2 *common.Student
	stu2 := &common.Student{}
	v := reflect.ValueOf(stu2)
	fmt.Println(v)
	fmt.Println(v.Type().Name())
	fmt.Println(v.Type().Kind())
	fmt.Println(v.Type().Elem().Kind())

	typ := reflect.TypeOf(stu2)
	fmt.Println(typ.Name())
	fmt.Println(typ.Kind())
}

// Type : 表示 interface{} 的实际类型  --
// Kind : 表示 该类型的特定类别
func Func2(i interface{}) {
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(i).Kind())
}

func Func3(i interface{}) {
	if reflect.TypeOf(i).Kind() == reflect.Struct {
		value := reflect.ValueOf(i)
		fmt.Printf("Num of Fields %v, \n", value.NumField())
		fmt.Printf("Num of Method %v, \n", value.NumMethod())
		for i := 0; i < value.NumField(); i++ {
			fmt.Printf("Field : %v, type: %T, value: %v \n", i, value.Field(i), value.Field(i))
		}
	}
}

func Func4() {
	i := 56
	x := reflect.ValueOf(i).Int()
	fmt.Printf("type： %T,  value: %v \n", x, x)

	str := "majm"
	y := reflect.ValueOf(str).String()
	fmt.Printf("type： %T,  value: %v \n", y, y)
}

func DemoCase(i interface{}) {
	if reflect.ValueOf(i).Kind() == reflect.Struct {
		typ := reflect.TypeOf(i)
		val := reflect.ValueOf(i)
		name := typ.Name()
		fmt.Printf("typeName: %v \n", name)

		for j := 0; j < val.NumField(); j++ {
			f := val.Field(j)
			switch f.Kind() {
			case reflect.Int:
				// process
				fmt.Printf("%d \n", f.Int())
			case reflect.String:
				fmt.Printf("%s \n", f.String())
			default:
				fmt.Println("Unresolved Type!")
			}
		}
	}
}
