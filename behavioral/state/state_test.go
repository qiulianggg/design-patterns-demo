package main

import (
	"strings"
	"testing"
)

// 正常流转：待支付 -> 已支付 -> 已发货。
func TestStateHappyPath(t *testing.T) {
	o := NewOrder()
	if o.State() != "待支付" {
		t.Fatalf("初始状态应为待支付, 实际 %q", o.State())
	}
	o.Pay()
	if o.State() != "已支付" {
		t.Fatalf("支付后应为已支付, 实际 %q", o.State())
	}
	o.Ship()
	if o.State() != "已发货" {
		t.Fatalf("发货后应为已发货, 实际 %q", o.State())
	}
}

// 未支付不能发货，状态不变。
func TestStateShipBeforePay(t *testing.T) {
	o := NewOrder()
	got := o.Ship()
	if !strings.Contains(got, "❌") || o.State() != "待支付" {
		t.Fatalf("未支付发货应被拒绝且状态不变, got=%q state=%q", got, o.State())
	}
}

// 重复支付应被拒绝。
func TestStateDoublePay(t *testing.T) {
	o := NewOrder()
	o.Pay()
	if got := o.Pay(); !strings.Contains(got, "❌") {
		t.Fatalf("重复支付应被拒绝, got=%q", got)
	}
}
