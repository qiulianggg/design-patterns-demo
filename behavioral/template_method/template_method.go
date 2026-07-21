package main

import (
	"fmt"
	"strings"
)

// 模板方法模式：在父类中定义算法的骨架，把某些步骤延迟到子类实现。
// 本例：数据处理流程固定为 读取 -> 处理 -> 保存，
// 但「读取」和「处理」的具体做法由不同数据源实现。
//
// Go 没有继承，用「接口定义可变步骤 + 一个模板函数编排骨架」来实现。

// DataProcessor 定义算法中「可变的步骤」（钩子）。
type DataProcessor interface {
	Read() string           // 变化点：从哪读
	Process(raw string) string // 变化点：怎么处理
	Save(result string)     // 变化点：存到哪
}

// Run 是模板方法：固定算法骨架（步骤顺序不可变），只把各步骤委托给具体实现。
func Run(p DataProcessor) {
	fmt.Println("=== 开始处理流程（骨架固定）===")
	raw := p.Read()
	fmt.Println("  1) 读取:", raw)
	result := p.Process(raw)
	fmt.Println("  2) 处理:", result)
	p.Save(result)
	fmt.Println("=== 处理完成 ===")
}

// CSVProcessor 具体实现之一。
type CSVProcessor struct{}

func (CSVProcessor) Read() string { return "a,b,c" }
func (CSVProcessor) Process(raw string) string {
	return strings.ReplaceAll(raw, ",", " | ")
}
func (CSVProcessor) Save(result string) { fmt.Println("  3) 保存到 CSV 文件:", result) }

// JSONProcessor 具体实现之二。
type JSONProcessor struct{}

func (JSONProcessor) Read() string              { return `{"name":"tom"}` }
func (JSONProcessor) Process(raw string) string { return strings.ToUpper(raw) }
func (JSONProcessor) Save(result string)        { fmt.Println("  3) 保存到数据库:", result) }

func main() {
	// 同一套骨架 Run()，不同的步骤实现
	Run(CSVProcessor{})
	fmt.Println()
	Run(JSONProcessor{})
}
