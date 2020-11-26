/**
  @Author: majm@ushareit.com
  @date: 2020/11/25
  @note:
**/
package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.DefaultTimeLoc = time.UTC

	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Profile))
	orm.RegisterModel(new(Post))
	orm.RegisterModel(new(Tag))
	orm.RegisterModel(new(PostTag))
}
