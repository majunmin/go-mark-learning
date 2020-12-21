/**
  @Author: majm@ushareit.com
  @date: 2020/12/4
  @note:
**/
package chan_select

import (
	"fmt"
	"strings"
	"testing"
)

func TestDemo1(t *testing.T) {

	fmt.Println(int64(1e6))
	Demo1()

	fmt.Println(fmt.Sprintf("%s/%v", "hello", 2006))

	fmt.Println(strings.Compare("123400", "14567"))

	split := strings.Split("_1234", "_")
	fmt.Println(split[0])

	res := make([]string, 1)
	GetSliceData(&res)
	fmt.Println(res)

}
