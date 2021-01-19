/**
  @Author: majm@ushareit.com
  @date: 2021/1/19
  @note:
**/
package tail_learn

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

var (
	fileName = "./a.log"
)

func initConfig() {
	config := tail.Config{
		Location:    &tail.SeekInfo{Offset: 0, Whence: 2}, // 从哪个位置开始读取
		ReOpen:      true,
		MustExist:   false,
		Poll:        false,
		RateLimiter: nil,
		Follow:      true, // 是否跟随
		MaxLineSize: 0,
	}

	file, _ := tail.TailFile(fileName, config)

	for true {
		line, ok := <-file.Lines
		if !ok {
			fmt.Printf("tail file close reopen,filename: %s", fileName)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("msg: ", line.Text)
	}
}
