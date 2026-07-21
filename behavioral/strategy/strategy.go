package main

import "fmt"

// 策略模式：定义一系列算法，把它们各自封装起来，并使它们可以互相替换。
// 本例：一个购物车，支持在结算时切换不同的折扣策略。

// DiscountStrategy 是策略接口：不同折扣算法。
type DiscountStrategy interface {
	Apply(amount float64) float64
	Name() string
}

// NoDiscount 无折扣。
type NoDiscount struct{}

func (NoDiscount) Name() string             { return "无折扣" }
func (NoDiscount) Apply(a float64) float64 { return a }

// PercentageDiscount 百分比折扣。
type PercentageDiscount struct{ percent float64 }

func (p PercentageDiscount) Name() string { return fmt.Sprintf("打%.0f折", (1-p.percent)*10) }
func (p PercentageDiscount) Apply(a float64) float64 {
	return a * (1 - p.percent)
}

// ThresholdDiscount 满减：满 threshold 减 minus。
type ThresholdDiscount struct{ threshold, minus float64 }

func (t ThresholdDiscount) Name() string {
	return fmt.Sprintf("满%.0f减%.0f", t.threshold, t.minus)
}
func (t ThresholdDiscount) Apply(a float64) float64 {
	if a >= t.threshold {
		return a - t.minus
	}
	return a
}

// Cart 是上下文：持有一个策略，把计算委托给它。
type Cart struct {
	amount   float64
	discount DiscountStrategy
}

func (c *Cart) SetDiscount(s DiscountStrategy) { c.discount = s }

func (c *Cart) Checkout() {
	final := c.discount.Apply(c.amount)
	fmt.Printf("原价 %.2f, 策略[%s] => 实付 %.2f\n", c.amount, c.discount.Name(), final)
}

func main() {
	cart := &Cart{amount: 200}

	// 运行时自由切换算法，Cart 代码不变
	for _, s := range []DiscountStrategy{
		NoDiscount{},
		PercentageDiscount{percent: 0.2},   // 打8折
		ThresholdDiscount{threshold: 150, minus: 30}, // 满150减30
	} {
		cart.SetDiscount(s)
		cart.Checkout()
	}
}
