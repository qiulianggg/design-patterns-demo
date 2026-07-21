package main

import "fmt"

// ---------- 产品接口 ----------

// Transport 是所有交通工具的抽象产品。
type Transport interface {
	Deliver(cargo string) string
}

// Truck 具体产品：卡车（陆运）。
type Truck struct{}

func (Truck) Deliver(cargo string) string {
	return fmt.Sprintf("用卡车经陆路运输【%s】", cargo)
}

// Ship 具体产品：轮船（海运）。
type Ship struct{}

func (Ship) Deliver(cargo string) string {
	return fmt.Sprintf("用轮船经海路运输【%s】", cargo)
}

// ---------- 工厂接口 ----------

// Logistics 是抽象工厂：声明「工厂方法」CreateTransport，
// 由具体工厂决定创建哪种 Transport。
type Logistics interface {
	CreateTransport() Transport
	// PlanDelivery 是复用工厂方法的通用业务逻辑，
	// 它不关心具体产品类型，面向 Transport 接口编程。
	PlanDelivery(cargo string) string
}

// baseLogistics 提供 PlanDelivery 的公共实现，
// 通过内嵌被具体工厂复用（Go 用组合替代继承）。
type baseLogistics struct {
	factory func() Transport
}

func (b baseLogistics) PlanDelivery(cargo string) string {
	t := b.factory() // 调用工厂方法拿到产品
	return "规划配送 -> " + t.Deliver(cargo)
}

// RoadLogistics 具体工厂：生产卡车。
type RoadLogistics struct{ baseLogistics }

func NewRoadLogistics() *RoadLogistics {
	l := &RoadLogistics{}
	l.factory = l.CreateTransport
	return l
}
func (RoadLogistics) CreateTransport() Transport { return Truck{} }

// SeaLogistics 具体工厂：生产轮船。
type SeaLogistics struct{ baseLogistics }

func NewSeaLogistics() *SeaLogistics {
	l := &SeaLogistics{}
	l.factory = l.CreateTransport
	return l
}
func (SeaLogistics) CreateTransport() Transport { return Ship{} }

func main() {
	var logistics Logistics

	logistics = NewRoadLogistics()
	fmt.Println(logistics.PlanDelivery("电视机"))

	logistics = NewSeaLogistics()
	fmt.Println(logistics.PlanDelivery("集装箱"))
}
