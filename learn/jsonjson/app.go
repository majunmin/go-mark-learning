/**
  @Author: majm@ushareit.com
  @date: 2020/11/20
  @note:
**/
package main

import (
	"fmt"
	"go-mark-learning/common"
	"time"
)

func main() {

	student := common.Student{
		Id:   11,
		Name: "majm",
		School: common.School{
			Name: "NJUPT",
			Addr: "NJ",
		},
		Birthday: time.Now(),
	}

	json, err := student.MarshalJSON()
	fmt.Println(string(json), err)

	jsonStr := `{"id":11,"s_name":"qq","s_chool":{"name":"CUMT","addr":"xz"},"birthday":"2017-08-04T20:58:07.9894603+08:00"}`
	ss := common.Student{}
	err = ss.UnmarshalJSON([]byte(jsonStr))
	if err != nil {
		fmt.Println("UnmarshalJSON error", err)
		return
	}
	fmt.Println("UnmarshalJSON ", ss)
}
