package main

import "testing"

// 容器的 Size 应递归累加所有子节点大小。
func TestDirectorySizeAggregates(t *testing.T) {
	root := NewDirectory("root")
	root.Add(&File{"a", 2})
	sub := NewDirectory("sub")
	sub.Add(&File{"b", 3}).Add(&File{"c", 5})
	root.Add(sub)

	if got := root.Size(); got != 10 {
		t.Fatalf("期望总大小 10, 实际 %d", got)
	}
}

func TestLeafSize(t *testing.T) {
	f := &File{"x", 7}
	if f.Size() != 7 {
		t.Fatalf("叶子大小错误: %d", f.Size())
	}
}
