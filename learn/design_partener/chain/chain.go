/**
  @Author: majm@ushareit.com
  @date: 2020/11/28
  @note:
**/
package chain

import (
	"fmt"
	"strings"
)

type Handler interface {
	Do(content string)
	next(handler Handler, content string)
}

// 广告过滤
type AdHandler struct {
	handler Handler
}

// 涉黄过滤
type YellowHandler struct {
	handler Handler
}

// 敏感词过滤
type SensitiveHandler struct {
	hadnler Handler
}

func (ad *AdHandler) Do(content string) {
	fmt.Println("执行广告过滤")
	newContent := strings.ReplaceAll(content, "广告", "**")
	fmt.Println(newContent)
	ad.next(ad.handler, newContent)
}

func (ad *AdHandler) next(handler Handler, content string) {
	if handler != nil {
		handler.Do(content)
	}
}

func (yellow *YellowHandler) Do(content string) {
	fmt.Println("执行涉黄过滤")
	newContent := strings.ReplaceAll(content, "涉黄", "**")
	fmt.Println(newContent)
	yellow.next(yellow.handler, newContent)
}

func (yellow *YellowHandler) next(handler Handler, content string) {
	if handler != nil {
		handler.Do(content)
	}
}

func (sensitive *SensitiveHandler) Do(content string) {
	fmt.Println("执行敏感词过滤")
	newContent := strings.ReplaceAll(content, "敏感", "**")
	fmt.Println(newContent)
	sensitive.next(sensitive.hadnler, newContent)
}

func (sensitive *SensitiveHandler) next(handler Handler, content string) {
	if handler != nil {
		handler.Do(content)
	}
}

//过滤器链
type FilterChain interface {
	//添加一个
	AddFilter(h Handler)
	//删除一个
	RemoveFilter(f Handler)
	//迭代
	Iterator() Handler
	//调整顺序
	AdjustmentFilter(f Handler, index int)
}
