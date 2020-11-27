/**
  @Author: majm@ushareit.com
  @date: 2020/11/28
  @note:
**/
package chain

import "testing"

func TestHandlerChain(t *testing.T) {
	adHandler := &AdHandler{}
	yellowHandler := &YellowHandler{}
	sensitiveHandler := &SensitiveHandler{}

	adHandler.handler = yellowHandler
	yellowHandler.handler = sensitiveHandler

	adHandler.Do("我是正常内容，我是广告，我是涉黄，我是敏感词，我是正常内容")
}
