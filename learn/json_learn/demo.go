/**
  @Author: majm@ushareit.com
  @date: 2020/11/20
  @note:
**/
package json_learn

// struct tag : `json:"name,omitempty"`
// omitempty: 如果该字段为nil或零值，则打包的JSON结果不会有这个字段
// -        : 忽略该字段
type Student struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age"`
	Sex    string `json:"sex"`
	Status int
}
