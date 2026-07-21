package main

import "testing"

// 单例：多次获取应为同一实例，且状态共享。
func TestGetConfigReturnsSameInstance(t *testing.T) {
	c1 := GetConfig()
	c2 := GetConfig()
	if c1 != c2 {
		t.Fatalf("期望同一实例，得到不同指针: %p vs %p", c1, c2)
	}
}

func TestConfigStateShared(t *testing.T) {
	GetConfig().Set("k", "v")
	if got := GetConfig().Get("k"); got != "v" {
		t.Fatalf("期望通过另一引用读到 v，实际 %q", got)
	}
}
