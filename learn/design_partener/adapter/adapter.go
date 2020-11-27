/**
  @Author: majm@ushareit.com
  @date: 2020/11/27
  @note:
**/
package adapter

// 适配后的目标接口
type Target interface {
	Request() string
}

// 被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// 被适配类的 工厂方法
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// implements Adaptee
type adapteeImpl struct {
}

func (receiver *adapteeImpl) SpecificRequest() string {
	return "adaptee method !"
}

//Adapter 是转换Adaptee为Target接口的适配器
type Adapter struct {
	Adaptee
}

// 适配类 工厂方法
func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{adaptee}
}

// 适配器类需要 实现 目标接口
func (receiver *Adapter) Request() string {
	return receiver.SpecificRequest()
}
