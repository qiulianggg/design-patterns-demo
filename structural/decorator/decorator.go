package main

import (
	"fmt"
	"strings"
)

// 装饰器模式：在不修改原对象、也不用子类的前提下，动态地给对象「层层包裹」新增职责。
// 本例：给数据流叠加「大写转换」和「加感叹号」两种处理。

// DataSource 是被装饰的核心接口。
type DataSource interface {
	Write(data string) string
}

// PlainSource 是具体组件（最内层，真正干活的）。
type PlainSource struct{}

func (PlainSource) Write(data string) string {
	return data
}

// decorator 是装饰器基类：持有被包裹的 DataSource（组合）。
type decorator struct {
	wrapped DataSource
}

// UpperCaseDecorator 装饰器：把内容转成大写。
type UpperCaseDecorator struct {
	decorator
}

func NewUpperCase(ds DataSource) *UpperCaseDecorator {
	return &UpperCaseDecorator{decorator{ds}}
}
func (d *UpperCaseDecorator) Write(data string) string {
	// 先让被包裹对象处理，再叠加自己的行为
	return strings.ToUpper(d.wrapped.Write(data))
}

// ExclaimDecorator 装饰器：在末尾加感叹号。
type ExclaimDecorator struct {
	decorator
}

func NewExclaim(ds DataSource) *ExclaimDecorator {
	return &ExclaimDecorator{decorator{ds}}
}
func (d *ExclaimDecorator) Write(data string) string {
	return d.wrapped.Write(data) + "!!!"
}

func main() {
	// 从内到外层层包裹：PlainSource -> 大写 -> 加感叹号
	var ds DataSource = PlainSource{}
	ds = NewUpperCase(ds)
	ds = NewExclaim(ds)

	fmt.Println(ds.Write("hello world"))

	// 调整装饰顺序 / 组合，得到不同效果，且核心组件不变
	only := NewExclaim(PlainSource{})
	fmt.Println(only.Write("nice"))
}
