/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
**/
package abstract_factory

import "testing"

func TestRDBFactory(t *testing.T) {
	factory := &RDBFactory{}
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
}

func TestXMLFactory(t *testing.T) {
	factory := &XMLFactory{}
	factory.CreateOrderMainDao().SaveOrderMain()
	factory.CreateOrderDetailDao().SaveOrderDetail()
}
