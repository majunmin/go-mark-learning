/**
  @Author: majm@ushareit.com
  @date: 2020/12/21
  @note:
**/
package json_learn

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseTag(t *testing.T) {
	student := Student{
		Name:   "majm",
		Age:    18,
		Sex:    "sex",
		Status: 1,
	}
	tags := findJsonTags(student)
	fmt.Println(strings.Join(tags, ","))
}
