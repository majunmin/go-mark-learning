/**
  @Author: majm@ushareit.com
  @date: 2020/11/26
  @note:
**/
package simple_factory

import "fmt"

type Api interface {
	Say(name string) string
}

type hiApi struct{}

type helloApi struct{}

// return Api instance by type
// @param typ  1: hiApi 2:helloApi
func NewApi(typ int) Api {
	if typ == 1 {
		return &hiApi{}
	} else if typ == 2 {
		return &helloApi{}
	}
	return nil
}

func (receiver *hiApi) Say(name string) string {
	return fmt.Sprintf("hi %v", name)
}

func (receiver *helloApi) Say(name string) string {
	return fmt.Sprintf("hello.xlsx %v", name)
}
