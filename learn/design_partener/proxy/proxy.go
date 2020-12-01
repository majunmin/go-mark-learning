/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
**/
package proxy

import "fmt"

type Seller interface {
	sell(name string) string
}

//火车站
type Station struct {
	stock int //库存
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理点中： %s 买了一张票，剩余： %d \n", name, station.stock)

	} else {
		fmt.Println("票已售空")
	}
}

// 火车代理点
type StationProxy struct {
	station *Station //持有一个火车站对象
}

func NewStationProxy(station *Station) *StationProxy {
	return &StationProxy{station: station}
}

func (sp *StationProxy) sell(name string) {
	if sp.station.stock > 0 {
		sp.station.stock--
		fmt.Printf("代理点中: %s 买了一张票， 剩余 %d \n", name, sp.station.stock)
	} else {
		fmt.Println("票已售空")
	}
}
