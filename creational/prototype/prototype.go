package main

import "fmt"

// 原型模式：通过「克隆」现有对象来创建新对象，而不是 new 后逐字段赋值。
// 适合创建成本高、或想基于一个模板生成许多副本的场景。

// Cloneable 声明克隆能力。
type Cloneable interface {
	Clone() Cloneable
}

// Document 是原型：包含一个切片（引用类型），演示深拷贝的重要性。
type Document struct {
	Title   string
	Content string
	Tags    []string // 引用类型，浅拷贝会共享底层数组
}

// Clone 返回一个深拷贝副本：切片要单独复制，否则副本与原型共享底层数组。
func (d *Document) Clone() Cloneable {
	tags := make([]string, len(d.Tags))
	copy(tags, d.Tags) // 关键：复制切片内容，实现深拷贝
	return &Document{
		Title:   d.Title,
		Content: d.Content,
		Tags:    tags,
	}
}

func (d *Document) String() string {
	return fmt.Sprintf("《%s》 tags=%v", d.Title, d.Tags)
}

func main() {
	// 准备一个「模板」原型
	tpl := &Document{
		Title:   "季度报告模板",
		Content: "……",
		Tags:    []string{"内部", "财务"},
	}

	// 克隆出一份并修改，不影响原型
	q1 := tpl.Clone().(*Document)
	q1.Title = "2026 Q1 报告"
	q1.Tags = append(q1.Tags, "Q1")

	fmt.Println("原型:", tpl)
	fmt.Println("副本:", q1)

	// 验证深拷贝：修改副本的 Tags 未污染原型
	fmt.Printf("原型 Tags 仍为 %v（未被副本的 append 影响）\n", tpl.Tags)
}
