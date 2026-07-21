package main

import (
	"strings"
	"testing"
)

// 适配器应完成「分 -> 元」的接口转换。
func TestAdapterConvertsCentsToYuan(t *testing.T) {
	got := NewThirdPartyAdapter().Pay(12999)
	if !strings.Contains(got, "129.99") {
		t.Fatalf("12999 分应转换为 129.99 元, 实际: %q", got)
	}
}

// 适配器应满足目标接口。
func TestAdapterImplementsTarget(t *testing.T) {
	var _ PaymentProcessor = NewThirdPartyAdapter()
}
