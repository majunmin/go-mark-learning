/**
  @Author: majm@ushareit.com
  @date: 2020/11/28
  @note:
**/
package template

import (
	"fmt"
	"go-mark-learning/common/tool"
)

const (
	// ConstActTypeTime 按时间抽奖类型
	ConstActTypeTime int32 = 1
	// ConstActTypeTimes 按抽奖次数抽奖
	ConstActTypeTimes int32 = 2
	// ConstActTypeAmount 按数额范围区间抽奖
	ConstActTypeAmount int32 = 3
)

type ActInfo struct {
	// 抽奖类型
	ActiveType int32
}

// Context 上下文
type Context struct {
	ActInfo *ActInfo
}

type BehaviorInterface interface {
	// 参数校验(不同活动类型实现不同)
	checkParams(ctx *Context) error
	// 获取node奖品信息(不同活动类型实现不同)
	getPrizesByNode(ctx *Context) error
}

// TimeDraw 具体抽奖行为
//按时间类型抽奖 比如红包雨
type TimeDraw struct {
}

func (td TimeDraw) checkParams(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "按时间类型抽奖:特殊参数校验...")
	return
}

func (td TimeDraw) getPrizesByNode(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "do nothing(抽取该场次的奖品即可，无需其他逻辑)...")
	return
}

// TimesDraw 具体抽奖行为
//按次数类型抽奖 比如答题闯关
type TimesDraw struct {
}

func (tsd TimesDraw) checkParams(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "按抽奖次数抽奖类型:特殊参数校验...")
	return
}

func (tsd TimesDraw) getPrizesByNode(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "1. 判断是该用户第几次抽奖...")
	fmt.Println(tool.RunFuncName(), "2. 获取对应node的奖品信息...")
	fmt.Println(tool.RunFuncName(), "3. 复写原所有奖品信息(抽取该node节点的奖品)...")
	return
}

// AmountDraw 具体抽奖行为
// 按数额范围区间抽奖 比如订单金额刮奖
type AmountDraw struct {
}

// checkParams 其他参数校验(不同活动类型实现不同)
func (draw *AmountDraw) checkParams(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "按数额范围区间抽奖:特殊参数校验...")
	return
}

// getPrizesByNode 获取node奖品信息(不同活动类型实现不同)
func (draw *AmountDraw) getPrizesByNode(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "1. 判断属于哪个数额区间...")
	fmt.Println(tool.RunFuncName(), "2. 获取对应node的奖品信息...")
	fmt.Println(tool.RunFuncName(), "3. 复写原所有奖品信息(抽取该node节点的奖品)...")
	return
}

// 抽奖模板 Lottery
type Lottery struct {
	concreteBehavior BehaviorInterface
}

func (lottery *Lottery) Run(ctx *Context) (err error) {
	// 检验活动编号(seq_nno)是否存在 并获取活动信息
	if err = lottery.checkSerialNo(ctx); err != nil {
		return err
	}

	// 具体方法：校验活动、场次是否正在进行
	if err = lottery.checkStatus(ctx); err != nil {
		return err
	}

	// ”抽象方法“：其他参数校验
	if err = lottery.concreteBehavior.checkParams(ctx); err != nil {
		return err
	}

	// 具体方法：活动抽奖次数校验(同时扣减)
	if err = lottery.checkTimesByAct(ctx); err != nil {
		return err
	}

	// 具体方法：活动是否需要消费积分
	if err = lottery.consumePointsByAct(ctx); err != nil {
		return err
	}

	// 具体方法：场次抽奖次数校验(同时扣减)
	if err = lottery.checkTimesBySession(ctx); err != nil {
		return err
	}

	// 具体方法：获取场次奖品信息
	if err = lottery.getPrizesBySession(ctx); err != nil {
		return err
	}

	// ”抽象方法“：获取node奖品信息
	if err = lottery.getPrizesByNode(ctx); err != nil {
		return err
	}

	// 具体方法：抽奖
	if err = lottery.drawPrizes(ctx); err != nil {
		return err
	}

	// 具体方法：奖品数量判断
	if err = lottery.checkPrizesStock(ctx); err != nil {
		return err
	}

	// 具体方法：组装奖品信息
	if err = lottery.packagePrizeInfo(ctx); err != nil {
		return err
	}
	return

}

// checkSerialNo 校验活动编号(serial_no)是否存在
func (lottery *Lottery) checkSerialNo(ctx *Context) (err error) {
	fmt.Printf(tool.RunFuncName(), "校验活动编号(serial_no)是否存在、并获取活动信息...")
	ctx.ActInfo = &ActInfo{
		ActiveType: ConstActTypeTimes,
	}

	switch ctx.ActInfo.ActiveType {
	case ConstActTypeTime:
		lottery.concreteBehavior = &TimeDraw{}
	case ConstActTypeTimes:
		lottery.concreteBehavior = &TimesDraw{}
	case ConstActTypeAmount:
		lottery.concreteBehavior = &AmountDraw{}
	default:
		return fmt.Errorf("不存在的活动类型")
	}
	return

}

func (lottery *Lottery) checkStatus(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "检查活动场次是否正在进行...")
	return
}

// checkParams 其他参数校验(不同活动类型实现不同)
// 不同场景变化的算法 转化为依赖抽象
func (lottery *Lottery) checkParams(ctx *Context) (err error) {
	// 实际依赖的接口的抽象方法
	return lottery.concreteBehavior.checkParams(ctx)
}

// checkTimesByAct 活动抽奖次数校验
func (lottery *Lottery) checkTimesByAct(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "活动抽奖次数校验...")
	return
}

// consumePointsByAct 活动是否需要消费积分
func (lottery *Lottery) consumePointsByAct(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "活动是否需要消费积分...")
	return
}

// checkTimesBySession 活动抽奖次数校验
func (lottery *Lottery) checkTimesBySession(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "活动抽奖次数校验...")
	return
}

// getPrizesBySession 获取场次奖品信息
func (lottery *Lottery) getPrizesBySession(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "获取场次奖品信息...")
	return
}

// getPrizesByNode 获取node奖品信息(不同活动类型实现不同)
// 不同场景变化的算法 转化为依赖抽象
func (lottery *Lottery) getPrizesByNode(ctx *Context) (err error) {
	// 实际依赖的接口的抽象方法
	return lottery.concreteBehavior.getPrizesByNode(ctx)
}

// drawPrizes 抽奖
func (lottery *Lottery) drawPrizes(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "抽奖...")
	return
}

// checkPrizesStock 奖品数量判断
func (lottery *Lottery) checkPrizesStock(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "奖品数量判断...")
	return
}

// packagePrizeInfo 组装奖品信息
func (lottery *Lottery) packagePrizeInfo(ctx *Context) (err error) {
	fmt.Println(tool.RunFuncName(), "组装奖品信息...")
	return
}
