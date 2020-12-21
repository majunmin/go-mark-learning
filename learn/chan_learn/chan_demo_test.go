/**
  @Author: majm@ushareit.com
  @date: 2020/12/8
  @note:
**/
package chanchan

import (
	"fmt"
	"testing"
)

func TestAGoutinue(t *testing.T) {
	go AGoutinue()
	res := <-done
	fmt.Println(msg)
	fmt.Println(res)

	Operate()
}

func TestSelectOpe(t *testing.T) {
	SelectOpe()
}

func TestSelectForPend(t *testing.T) {
	SelectForPend()
}
