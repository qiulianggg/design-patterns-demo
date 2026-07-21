package main

import "testing"

func TestBuilderSetsFields(t *testing.T) {
	pc := NewComputerBuilder().
		WithCPU("M3").
		WithRAM(24).
		WithStorage("1TB").
		WithGPU("RTX").
		Build()

	if pc.CPU != "M3" || pc.RAM != 24 || pc.Storage != "1TB" || pc.GPU != "RTX" {
		t.Fatalf("字段设置错误: %+v", pc)
	}
}

// 未设置 RAM 时应使用默认值 8。
func TestBuilderDefaultRAM(t *testing.T) {
	pc := NewComputerBuilder().WithCPU("x").Build()
	if pc.RAM != 8 {
		t.Fatalf("期望默认 RAM=8, 实际 %d", pc.RAM)
	}
}

func TestDirectorGamingPC(t *testing.T) {
	pc := Director{}.BuildGamingPC(NewComputerBuilder())
	if pc.GPU == "" {
		t.Fatal("游戏机应配置独立显卡")
	}
}
