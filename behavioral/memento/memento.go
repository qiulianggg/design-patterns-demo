package main

import "fmt"

// 备忘录模式：在不破坏封装的前提下，捕获对象的内部状态并保存到外部，
// 以便日后将对象恢复到该状态。本例：文本编辑器的撤销功能。

// Memento 是备忘录：保存某一时刻的状态快照。
// 它对外只是不透明的「令牌」，只有 Originator 知道如何读写它。
type Memento struct {
	content string
}

// Editor 是原发器(Originator)：拥有需要保存/恢复的状态。
type Editor struct {
	content string
}

func (e *Editor) Type(text string) {
	e.content += text
}

func (e *Editor) Content() string { return e.content }

// Save 生成当前状态的备忘录。
func (e *Editor) Save() *Memento {
	return &Memento{content: e.content}
}

// Restore 从备忘录恢复状态。
func (e *Editor) Restore(m *Memento) {
	e.content = m.content
}

// History 是管理者(Caretaker)：保存备忘录栈，但不查看/修改其内容。
type History struct {
	stack []*Memento
}

func (h *History) Push(m *Memento) { h.stack = append(h.stack, m) }

func (h *History) Pop() *Memento {
	if len(h.stack) == 0 {
		return nil
	}
	m := h.stack[len(h.stack)-1]
	h.stack = h.stack[:len(h.stack)-1]
	return m
}

func main() {
	editor := &Editor{}
	history := &History{}

	editor.Type("Hello")
	history.Push(editor.Save()) // 快照1: "Hello"

	editor.Type(", World")
	history.Push(editor.Save()) // 快照2: "Hello, World"

	editor.Type("!!!")
	fmt.Println("当前内容:", editor.Content()) // Hello, World!!!

	// 撤销一次 -> 回到快照2
	editor.Restore(history.Pop())
	fmt.Println("撤销后:", editor.Content()) // Hello, World

	// 再撤销 -> 回到快照1
	editor.Restore(history.Pop())
	fmt.Println("再撤销:", editor.Content()) // Hello
}
