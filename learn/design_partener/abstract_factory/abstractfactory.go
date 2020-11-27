/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
  order orderDetail 分别存储在 xml 存储 和 数据库存储
**/
package abstract_factory

import "fmt"

// 订单表
type OrderMainDao interface {
	SaveOrderMain()
}

// 订单详情表
type OrderDetailDao interface {
	SaveOrderDetail()
}

// 抽象工厂🏭
type DaoFactory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

// 对应数据库存储
type RDBMainDao struct {
}

type RDBDetailDao struct {
}

func (rdb *RDBMainDao) SaveOrderMain() {
	fmt.Println("RDBMainDao SaveOrderMain !")
}

func (rdb *RDBDetailDao) SaveOrderDetail() {
	fmt.Println("RDBDetailDao SaveOrderDetail !")
}

type RDBFactory struct {
}

func (receiver *RDBFactory) CreateOrderMainDao() OrderMainDao {
	return &RDBMainDao{}
}

func (receiver *RDBFactory) CreateOrderDetailDao() OrderDetailDao {
	return &RDBDetailDao{}
}

// ----------- 对应 XML存储
type XMLMainDao struct {
}

type XMLDetailDao struct {
}

func (xml *XMLMainDao) SaveOrderMain() {
	fmt.Println("XMLMainDao SaveOrderMain !")
}

func (xml *XMLDetailDao) SaveOrderDetail() {
	fmt.Println("XMLDetailDao SaveOrderDetail !")
}

type XMLFactory struct {
}

func (receiver *XMLFactory) CreateOrderMainDao() OrderMainDao {
	return &XMLMainDao{}
}

func (receiver *XMLFactory) CreateOrderDetailDao() OrderDetailDao {
	return &XMLDetailDao{}
}
