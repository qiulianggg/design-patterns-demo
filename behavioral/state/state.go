package main

import "fmt"

// 状态模式：允许一个对象在其内部状态改变时改变它的行为，看起来像换了个类。
// 本例：一个订单在 待支付 -> 已支付 -> 已发货 -> 已完成 之间流转，
// 每个状态下对同一操作（付款、发货）的响应不同。

// State 是状态接口：定义在该状态下各操作的行为。
type State interface {
	Name() string
	Pay(o *Order) string
	Ship(o *Order) string
}

// Order 是上下文(Context)：持有当前状态，把操作委托给状态对象。
type Order struct {
	state State
}

func NewOrder() *Order { return &Order{state: &PendingState{}} }

func (o *Order) setState(s State) { o.state = s }
func (o *Order) State() string    { return o.state.Name() }

// 上下文的操作直接委托给当前状态
func (o *Order) Pay() string  { return o.state.Pay(o) }
func (o *Order) Ship() string { return o.state.Ship(o) }

// ---------- 各具体状态 ----------

type PendingState struct{} // 待支付

func (PendingState) Name() string { return "待支付" }
func (s PendingState) Pay(o *Order) string {
	o.setState(&PaidState{})
	return "支付成功，状态 -> 已支付"
}
func (PendingState) Ship(*Order) string { return "❌ 未支付，不能发货" }

type PaidState struct{} // 已支付

func (PaidState) Name() string          { return "已支付" }
func (PaidState) Pay(*Order) string     { return "❌ 已支付，请勿重复支付" }
func (s PaidState) Ship(o *Order) string {
	o.setState(&ShippedState{})
	return "已发货，状态 -> 已发货"
}

type ShippedState struct{} // 已发货

func (ShippedState) Name() string      { return "已发货" }
func (ShippedState) Pay(*Order) string { return "❌ 已支付" }
func (ShippedState) Ship(*Order) string {
	return "❌ 已发货，请勿重复发货"
}

func main() {
	o := NewOrder()
	fmt.Println("初始状态:", o.State())

	fmt.Println(o.Ship()) // 未支付不能发货
	fmt.Println(o.Pay())  // 支付
	fmt.Println(o.Pay())  // 重复支付被拒
	fmt.Println(o.Ship()) // 发货
	fmt.Println(o.Ship()) // 重复发货被拒
	fmt.Println("最终状态:", o.State())
}
