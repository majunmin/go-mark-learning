/**
  @Author: majm@ushareit.com
  @date: 2020/12/1
  @note:
**/
package reflectreflect

import (
	"go-mark-learning/common"
	"testing"
)

func TestDemo1(t *testing.T) {
	Demo1()
}

func TestFunc2(t *testing.T) {
	stu := common.Student{}
	stu2 := &stu
	Func2(stu)  // common.Student    struct
	Func2(stu2) // *common.Student    ptr
}

func TestFunc3(t *testing.T) {
	stu := common.Student{}
	Func3(stu)
}

func TestFunc4(t *testing.T) {
	Func4()
}

func TestDemoCase(t *testing.T) {
	stu := common.Student{}
	DemoCase(stu)
}
