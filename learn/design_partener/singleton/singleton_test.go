/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
**/
package singleton

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetSingleton(t *testing.T) {

	instance := GetSingleton()
	fmt.Println(&instance)

	Convey("for test", t, func() {
		for i := 0; i < 50; i++ {
			So(GetSingleton(), ShouldEqual, instance)
		}
	})
}
