/**
  @Author: majm@ushareit.com
  @date: 2020/11/20
  @note:
**/
package main

// struct tag : `json:"name,omitempty"`
// omitempty: 如果该字段为nil或零值，则打包的JSON结果不会有这个字段
// -        : 忽略该字段
type Student struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age"`
	Sex    string `json:"sex"`
	status int
}

//func main() {
//    stu := &Student{Age: 18, Sex: "男"}
//
//    //将结构体封装成json格式，并返回[]byte
//    jsonData, err := json.Marshal(stu)
//    if err != nil {
//        fmt.Println(fmt.Errorf("Marshal error").Error())
//        return
//    }
//    fmt.Println(string(jsonData)) // 会忽略 小写字母开头的field  {"name":"majm","age":18,"sex":"男"}
//    fmt.Println(stu)
//
//    intArr := []int{14, 15, 17, 19}
//    jsonArrData, err := json.Marshal(intArr)
//    if err != nil {
//        fmt.Println(fmt.Errorf("Marshal error").Error())
//        return
//    }
//    fmt.Println(string(jsonArrData)) // [14,15,17,19]
//
//}
