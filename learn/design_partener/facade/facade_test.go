/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
**/
package facade

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var expect = "A module running\nB module running"

func TestFacadeApi(t *testing.T) {

	Convey("should equals", t, func() {
		api := NewApi()
		ret := api.Test()
		So(ret, ShouldEqual, expect)
	})
}
