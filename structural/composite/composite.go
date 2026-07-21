package main

import (
	"fmt"
	"strings"
)

// 组合模式：把对象组织成树形结构，让「单个对象（叶子）」和「组合对象（容器）」
// 对客户端表现出一致的接口。本例：文件系统的文件与目录。

// Node 是统一接口：文件和目录都实现它。
type Node interface {
	Name() string
	Size() int             // 叶子返回自身大小；容器返回子节点大小之和
	Print(indent string)   // 递归打印
}

// File 是叶子节点。
type File struct {
	name string
	size int
}

func (f *File) Name() string { return f.name }
func (f *File) Size() int    { return f.size }
func (f *File) Print(indent string) {
	fmt.Printf("%s- %s (%dKB)\n", indent, f.name, f.size)
}

// Directory 是容器节点：内部持有一组 Node（可能是文件，也可能是子目录）。
type Directory struct {
	name     string
	children []Node
}

func NewDirectory(name string) *Directory { return &Directory{name: name} }

func (d *Directory) Add(n Node) *Directory {
	d.children = append(d.children, n)
	return d
}

func (d *Directory) Name() string { return d.name }

// Size 递归累加：容器的行为是「委托给所有子节点」。
func (d *Directory) Size() int {
	total := 0
	for _, c := range d.children {
		total += c.Size()
	}
	return total
}

func (d *Directory) Print(indent string) {
	fmt.Printf("%s+ %s/ (%dKB)\n", indent, d.name, d.Size())
	for _, c := range d.children {
		c.Print(indent + "  ")
	}
}

func main() {
	root := NewDirectory("root")
	root.Add(&File{"readme.md", 2})

	src := NewDirectory("src")
	src.Add(&File{"main.go", 10}).Add(&File{"util.go", 5})

	root.Add(src)
	root.Add(&File{"go.mod", 1})

	fmt.Println(strings.Repeat("=", 30))
	// 客户端对叶子和容器一视同仁地调用 Print / Size
	root.Print("")
	fmt.Printf("整棵树总大小: %dKB\n", root.Size())
}
