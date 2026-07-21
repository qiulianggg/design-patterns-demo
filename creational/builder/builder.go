package main

import (
	"fmt"
	"strings"
)

// 建造者模式：分步骤构建复杂对象。本例构建一台电脑配置。

// Computer 是最终要构建的复杂产品。
type Computer struct {
	CPU     string
	RAM     int
	Storage string
	GPU     string
}

func (c Computer) String() string {
	gpu := c.GPU
	if gpu == "" {
		gpu = "集成显卡"
	}
	return fmt.Sprintf("电脑{CPU:%s, RAM:%dGB, 存储:%s, 显卡:%s}", c.CPU, c.RAM, c.Storage, gpu)
}

// ComputerBuilder 使用「链式调用（fluent interface）」分步设置各部件。
// 每个 With 方法返回 *ComputerBuilder，便于连缀。
type ComputerBuilder struct {
	c Computer
}

func NewComputerBuilder() *ComputerBuilder { return &ComputerBuilder{} }

func (b *ComputerBuilder) WithCPU(cpu string) *ComputerBuilder {
	b.c.CPU = cpu
	return b
}
func (b *ComputerBuilder) WithRAM(gb int) *ComputerBuilder {
	b.c.RAM = gb
	return b
}
func (b *ComputerBuilder) WithStorage(s string) *ComputerBuilder {
	b.c.Storage = s
	return b
}
func (b *ComputerBuilder) WithGPU(gpu string) *ComputerBuilder {
	b.c.GPU = gpu
	return b
}

// Build 收尾并返回构建好的产品（可在此做校验/默认值）。
func (b *ComputerBuilder) Build() Computer {
	if b.c.RAM == 0 {
		b.c.RAM = 8 // 默认值
	}
	return b.c
}

// Director（指挥者，可选角色）：封装常见的构建流程，复用建造步骤。
type Director struct{}

func (Director) BuildOfficePC(b *ComputerBuilder) Computer {
	return b.WithCPU("Intel i5").WithRAM(16).WithStorage("512GB SSD").Build()
}
func (Director) BuildGamingPC(b *ComputerBuilder) Computer {
	return b.WithCPU("AMD Ryzen 9").WithRAM(64).WithStorage("2TB SSD").WithGPU("RTX 4090").Build()
}

func main() {
	// 方式一：直接用建造者链式构建（最灵活）
	pc := NewComputerBuilder().
		WithCPU("Apple M3").
		WithRAM(24).
		WithStorage("1TB SSD").
		Build()
	fmt.Println("自定义:", pc)

	// 方式二：交给指挥者按预设流程构建
	d := Director{}
	fmt.Println("办公机:", d.BuildOfficePC(NewComputerBuilder()))
	fmt.Println("游戏机:", d.BuildGamingPC(NewComputerBuilder()))

	fmt.Println(strings.Repeat("-", 30))
	// 只设置部分字段，其余用默认值
	fmt.Println("最小配置:", NewComputerBuilder().WithCPU("入门 CPU").WithStorage("256GB").Build())
}
