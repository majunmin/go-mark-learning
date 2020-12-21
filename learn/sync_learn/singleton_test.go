/**
  @Author: majm@ushareit.com
  @date: 2020/12/18
  @note:
**/
package sync_learn

import (
	"fmt"
	"testing"
)

func TestSinngleton(t *testing.T) {
	// 将标准代码提取出来 就成了 标准库里的 sync.Once
	// custom:
	instance1 := Instance()
	fmt.Println(instance1 == Instance())

	fmt.Println("============")
	println(GetInstance())
	println(GetInstance())
	println(GetInstance())

}
