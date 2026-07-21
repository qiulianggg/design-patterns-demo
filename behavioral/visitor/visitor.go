package main

import "fmt"

// 访问者模式：在不修改元素类的前提下，为一组对象结构定义新的操作。
// 把「操作」从「元素」中分离出来，放进访问者里。
// 本例：几何图形(圆、矩形)接受不同访问者(计算面积、导出XML)。

// Visitor 声明对每种具体元素的访问方法。
type Visitor interface {
	VisitCircle(c *Circle)
	VisitRectangle(r *Rectangle)
}

// Shape 是元素接口：接受访问者（双分派的第一步）。
type Shape interface {
	Accept(v Visitor)
}

// Circle 具体元素。
type Circle struct{ Radius float64 }

// Accept 回调 visitor 对应的方法（双分派：由元素类型选择调用哪个 Visit）。
func (c *Circle) Accept(v Visitor) { v.VisitCircle(c) }

// Rectangle 具体元素。
type Rectangle struct{ Width, Height float64 }

func (r *Rectangle) Accept(v Visitor) { v.VisitRectangle(r) }

// ---------- 访问者1：计算面积 ----------

type AreaVisitor struct{ Total float64 }

func (a *AreaVisitor) VisitCircle(c *Circle) {
	area := 3.14159 * c.Radius * c.Radius
	fmt.Printf("  圆(r=%.1f) 面积=%.2f\n", c.Radius, area)
	a.Total += area
}
func (a *AreaVisitor) VisitRectangle(r *Rectangle) {
	area := r.Width * r.Height
	fmt.Printf("  矩形(%.1fx%.1f) 面积=%.2f\n", r.Width, r.Height, area)
	a.Total += area
}

// ---------- 访问者2：导出为 XML（新增操作，无需改元素类）----------

type XMLExportVisitor struct{}

func (XMLExportVisitor) VisitCircle(c *Circle) {
	fmt.Printf("  <circle radius=\"%.1f\"/>\n", c.Radius)
}
func (XMLExportVisitor) VisitRectangle(r *Rectangle) {
	fmt.Printf("  <rect w=\"%.1f\" h=\"%.1f\"/>\n", r.Width, r.Height)
}

func main() {
	shapes := []Shape{
		&Circle{Radius: 2},
		&Rectangle{Width: 3, Height: 4},
		&Circle{Radius: 1},
	}

	fmt.Println("== 访问者A: 计算面积 ==")
	area := &AreaVisitor{}
	for _, s := range shapes {
		s.Accept(area)
	}
	fmt.Printf("总面积 = %.2f\n", area.Total)

	fmt.Println("== 访问者B: 导出 XML ==")
	xml := XMLExportVisitor{}
	for _, s := range shapes {
		s.Accept(xml)
	}
}
