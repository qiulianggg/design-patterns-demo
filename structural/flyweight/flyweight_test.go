package main

import "testing"

// 相同 (char,font) 应复用同一享元实例（指针相等）。
func TestFlyweightShared(t *testing.T) {
	f := NewGlyphFactory()
	g1 := f.Get('a', "Arial")
	g2 := f.Get('a', "Arial")
	if g1 != g2 {
		t.Fatal("相同内部状态应返回同一享元实例")
	}
}

// 不同内部状态应是不同实例，且池大小正确。
func TestFlyweightCount(t *testing.T) {
	f := NewGlyphFactory()
	for _, ch := range "hello" { // h e l l o -> 4 个不同字符
		f.Get(ch, "Arial")
	}
	if f.Count() != 4 {
		t.Fatalf("期望 4 个享元, 实际 %d", f.Count())
	}
}
