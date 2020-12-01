/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
**/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// Note: https://colobu.com/2015/10/09/Linux-Signals/
// go run stack_obtain.go
// Enter `Ctrl + c` quit, will print stackTrace
func main() {

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for true {
			<-sigs
			buf := make([]byte, 1<<20)
			// 获取 **所有** goroutine 的 stacktrace
			// 如果需要获取 **当前** goroutine 的 stacktrace, 第二个参数需要为 `false`
			runtime.Stack(buf, true)
			fmt.Printf("=== goroutine stack trace...\n%s\n=== end\n", buf)
			done <- true
		}
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
