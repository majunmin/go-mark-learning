/**
  @Author: majm@ushareit.com
  @date: 2020/12/1
  @note:
**/
package robfig_cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func Democase1() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("*/1 * * * * *", func() {
		format := time.Now().Format("20060102 15:04:05")
		fmt.Println(format + "hello world!")
	})
	c.Start()
}
