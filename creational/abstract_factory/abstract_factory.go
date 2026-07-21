package main

import "fmt"

// 抽象工厂模式：创建「一系列相关产品」的家族，且保证同一家族的产品搭配使用。
// 本例：不同操作系统的 GUI 组件族（按钮 + 复选框）。

// ---------- 抽象产品 ----------

type Button interface{ Paint() string }
type Checkbox interface{ Paint() string }

// ---------- Windows 产品族 ----------

type WinButton struct{}

func (WinButton) Paint() string { return "渲染 Windows 风格按钮" }

type WinCheckbox struct{}

func (WinCheckbox) Paint() string { return "渲染 Windows 风格复选框" }

// ---------- macOS 产品族 ----------

type MacButton struct{}

func (MacButton) Paint() string { return "渲染 macOS 风格按钮" }

type MacCheckbox struct{}

func (MacCheckbox) Paint() string { return "渲染 macOS 风格复选框" }

// ---------- 抽象工厂 ----------

// GUIFactory 声明创建整族产品的方法，保证产品搭配一致。
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WinFactory struct{}

func (WinFactory) CreateButton() Button     { return WinButton{} }
func (WinFactory) CreateCheckbox() Checkbox { return WinCheckbox{} }

type MacFactory struct{}

func (MacFactory) CreateButton() Button     { return MacButton{} }
func (MacFactory) CreateCheckbox() Checkbox { return MacCheckbox{} }

// renderUI 是客户端代码：只依赖抽象工厂与抽象产品，
// 不关心到底是 Windows 还是 macOS。
func renderUI(f GUIFactory) {
	fmt.Println(f.CreateButton().Paint())
	fmt.Println(f.CreateCheckbox().Paint())
}

func main() {
	var factory GUIFactory

	fmt.Println("== 运行在 Windows ==")
	factory = WinFactory{}
	renderUI(factory)

	fmt.Println("== 运行在 macOS ==")
	factory = MacFactory{}
	renderUI(factory)
}
