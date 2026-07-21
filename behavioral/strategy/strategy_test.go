package main

import "testing"

func TestNoDiscount(t *testing.T) {
	if got := (NoDiscount{}).Apply(200); got != 200 {
		t.Fatalf("无折扣应原价, 实际 %v", got)
	}
}

func TestPercentageDiscount(t *testing.T) {
	// 打 8 折 -> 200 * 0.8 = 160
	if got := (PercentageDiscount{percent: 0.2}).Apply(200); got != 160 {
		t.Fatalf("8折应为 160, 实际 %v", got)
	}
}

func TestThresholdDiscount(t *testing.T) {
	s := ThresholdDiscount{threshold: 150, minus: 30}
	if got := s.Apply(200); got != 170 { // 满150减30
		t.Fatalf("满减后应为 170, 实际 %v", got)
	}
	if got := s.Apply(100); got != 100 { // 未达门槛不减
		t.Fatalf("未满门槛应原价, 实际 %v", got)
	}
}

// 上下文切换策略后行为随之改变。
func TestCartSwitchStrategy(t *testing.T) {
	cart := &Cart{amount: 100}
	cart.SetDiscount(PercentageDiscount{percent: 0.1})
	if cart.discount.Apply(cart.amount) != 90 {
		t.Fatal("切换策略后计算错误")
	}
}
