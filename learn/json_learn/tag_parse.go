/**
  @Author: majm@ushareit.com
  @date: 2020/12/21
  @note:
**/
package json_learn

import (
	"fmt"
	"reflect"
	"strings"
)

func findJsonTags(i interface{}) []string {
	t := reflect.TypeOf(i)
	numField := t.NumField()
	var tags []string
	for idx := 0; idx < numField; idx++ {
		tag := t.Field(idx).Tag
		if _, ok := tag.Lookup("json"); !ok {
			// 没标记 json
			fmt.Printf("field %s not contains tag json \n", t.Field(idx).Name)
			// process
		}
		jsonTag := tag.Get("json")
		if len(jsonTag) == 0 {
			continue
		}
		split := strings.Split(jsonTag, ",")
		tags = append(tags, split[0])
	}
	return tags
}
