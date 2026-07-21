package main

import (
	"strings"
	"testing"
)

func TestRoadLogisticsUsesTruck(t *testing.T) {
	got := NewRoadLogistics().PlanDelivery("电视机")
	if !strings.Contains(got, "卡车") {
		t.Fatalf("陆运应使用卡车, 实际: %q", got)
	}
}

func TestSeaLogisticsUsesShip(t *testing.T) {
	got := NewSeaLogistics().PlanDelivery("集装箱")
	if !strings.Contains(got, "轮船") {
		t.Fatalf("海运应使用轮船, 实际: %q", got)
	}
}

// 工厂返回的产品应满足 Transport 接口。
func TestFactoryProductType(t *testing.T) {
	var _ Transport = NewRoadLogistics().CreateTransport()
	var _ Transport = NewSeaLogistics().CreateTransport()
}
