package main

import "testing"

// 保存快照后可逐步恢复到历史状态。
func TestMementoUndo(t *testing.T) {
	editor := &Editor{}
	history := &History{}

	editor.Type("Hello")
	history.Push(editor.Save())

	editor.Type(", World")
	history.Push(editor.Save())

	editor.Type("!!!")
	if editor.Content() != "Hello, World!!!" {
		t.Fatalf("当前内容错误: %q", editor.Content())
	}

	editor.Restore(history.Pop())
	if editor.Content() != "Hello, World" {
		t.Fatalf("撤销一次后错误: %q", editor.Content())
	}

	editor.Restore(history.Pop())
	if editor.Content() != "Hello" {
		t.Fatalf("撤销两次后错误: %q", editor.Content())
	}
}

// 空历史 Pop 应返回 nil。
func TestHistoryPopEmpty(t *testing.T) {
	if (&History{}).Pop() != nil {
		t.Fatal("空历史 Pop 应返回 nil")
	}
}
