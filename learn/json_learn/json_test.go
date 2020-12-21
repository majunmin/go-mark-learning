/**
  @Author: majm@ushareit.com
  @date: 2020/12/21
  @note:
**/
package json_learn

import (
	"encoding/json"
	"fmt"
	"go-mark-learning/common"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	stu := &Student{Age: 18, Sex: "男"}

	//将结构体封装成json格式，并返回[]byte
	jsonData, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal error").Error())
		return
	}
	fmt.Println(string(jsonData)) // 会忽略 小写字母开头的field  {"name":"majm","age":18,"sex":"男"}
	fmt.Println(stu)

	intArr := []int{14, 15, 17, 19}
	jsonArrData, err := json.Marshal(intArr)
	if err != nil {
		fmt.Println(fmt.Errorf("Marshal error").Error())
		return
	}
	fmt.Println(string(jsonArrData)) // [14,15,17,19]
}

func TestApp(t *testing.T) {
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
