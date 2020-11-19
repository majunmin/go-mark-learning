/**
  @Author: majm@ushareit.com
  @date: 2020/11/20
  @note:
**/
package common

import "time"

//easyjson:json
type Student struct {
	Id       int       `json:"id"`
	Name     string    `json:"name,omitempty"`
	School   School    `json:"school"`
	Birthday time.Time `json:"birthday"`
}
