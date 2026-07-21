package main

import "testing"

// 克隆应为深拷贝：修改副本不应影响原型。
func TestCloneIsDeepCopy(t *testing.T) {
	tpl := &Document{Title: "模板", Tags: []string{"a", "b"}}

	clone := tpl.Clone().(*Document)
	clone.Title = "副本"
	clone.Tags = append(clone.Tags, "c")
	clone.Tags[0] = "changed"

	if tpl.Title != "模板" {
		t.Errorf("原型 Title 被污染: %q", tpl.Title)
	}

	if len(tpl.Tags) != 2 || tpl.Tags[0] != "a" {
		t.Errorf("原型 Tags 被污染: %v", tpl.Tags)
	}

}

// 克隆内容应与原型一致。
func TestCloneCopiesValues(t *testing.T) {
	tpl := &Document{Title: "T", Content: "C", Tags: []string{"x"}}
	clone := tpl.Clone().(*Document)
	if clone.Title != "T" || clone.Content != "C" || clone.Tags[0] != "x" {
		t.Fatalf("克隆内容不一致: %+v", clone)
	}
}
