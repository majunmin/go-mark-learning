/**
  @Author: majm@ushareit.com
  @date: 2020/11/21
  @note:
**/
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// UpperWriter 继承  Writer
type UpperWriter struct {
	io.Writer
}

func (p UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

func main() {
	fmt.Println("hello world!")
	fmt.Fprintln(UpperWriter{os.Stdout}, "Hello, world!")

	// 2
	fmt.Fprintln(os.Stdout, UpperString("Hello, world!"))

}

//------------------ 2 自定义打印格式
type UpperString string

func (s UpperString) String() string {
	return strings.ToUpper(string(s))
}
