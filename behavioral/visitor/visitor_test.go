package main

import (
	"math"
	"testing"
)

// 面积访问者应累加所有图形的面积。
func TestAreaVisitor(t *testing.T) {
	shapes := []Shape{
		&Circle{Radius: 2},        // 3.14159*4 = 12.56636
		&Rectangle{Width: 3, Height: 4}, // 12
	}
	v := &AreaVisitor{}
	for _, s := range shapes {
		s.Accept(v)
	}
	want := 3.14159*4 + 12
	if math.Abs(v.Total-want) > 1e-6 {
		t.Fatalf("总面积错误: got=%v want=%v", v.Total, want)
	}
}

// 双分派：同一组元素可被不同访问者处理而不 panic。
func TestXMLExportVisitor(t *testing.T) {
	shapes := []Shape{&Circle{Radius: 1}, &Rectangle{Width: 2, Height: 3}}
	v := XMLExportVisitor{}
	for _, s := range shapes {
		s.Accept(v)
	}
}
