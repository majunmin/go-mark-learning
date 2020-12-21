/**
  @Author: majm@ushareit.com
  @date: 2020/11/21
  @note:
**/
package interface_learn

import (
	"bytes"
	"io"
	"strings"
)

// UpperWriter 继承  Writer
type UpperWriter struct {
	io.Writer
}

func (p UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

//------------------ 2 自定义打印格式
type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}
