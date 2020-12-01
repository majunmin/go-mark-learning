/**
  @Author: majm@ushareit.com
  @date: 2020/11/25
  @note:
**/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", TestDo)
	log.Fatalf("%v", http.ListenAndServe("localhost:8080", nil))
}

func TestDo(w http.ResponseWriter, req *http.Request) {
	// 休眠 2 秒再处理，一会来测试超时
	time.Sleep(2e9)

	// csrf
	if req.Header.Get("_csrf") != "12345" {
		http.Error(w, "无效的 csrf token", 400)
		return
	}

	// 获取 Cookie
	cookie := req.Header.Get("Cookie")
	if cookie == "" {
		http.Error(w, "请登陆后再操作", 401)
		return
	}

	fmt.Println(cookie)

	index := strings.Index(cookie, "=")
	name := cookie[index+1:]
	fmt.Println(name)

	if name != "majm" {
		http.Error(w, "当前用户没有权限操作", 401)
		return
	}

	io.WriteString(w, "Hello "+name)

}
