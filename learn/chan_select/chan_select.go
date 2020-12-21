/**
  @Author: majm@ushareit.com
  @date: 2020/12/4
  @note:
**/
package chan_select

import (
	"fmt"
	"net/http"
)

func Demo1() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go Getdata("https://www.baidu.com", ch1)
	go Getdata("https://www.baidu.com", ch2)
	go Getdata("https://www.baidu.com", ch3)

	select {
	case v := <-ch1:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	case v := <-ch3:
		fmt.Println(v)
	}
}

func Getdata(url string, ch chan int) {
	if req, err := http.Get(url); err == nil {
		ch <- req.StatusCode
	}
}

func GetSliceData(param *[]string) []string {
	*param = append(*param, "hello")
	*param = append(*param, "world")
	return *param
}
