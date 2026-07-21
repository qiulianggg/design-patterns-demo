package main

import "fmt"

// 迭代器模式：提供一种方法顺序访问聚合对象中的元素，而不暴露其内部表示。
// 本例：自定义一个环形缓冲区，通过迭代器统一遍历，隐藏其内部实现细节。

// Iterator 是迭代器接口。
type Iterator interface {
	HasNext() bool
	Next() int
}

// Aggregate 是聚合接口：能创建自己的迭代器。
type Aggregate interface {
	CreateIterator() Iterator
}

// RingBuffer 是具体聚合：内部用切片 + 起始下标模拟环形存储。
type RingBuffer struct {
	data  []int
	start int // 逻辑起点
}

func NewRingBuffer(data []int, start int) *RingBuffer {
	return &RingBuffer{data: data, start: start}
}

func (r *RingBuffer) CreateIterator() Iterator {
	return &ringIterator{rb: r, count: 0}
}

// ringIterator 是具体迭代器：封装遍历逻辑，客户端无需知道 start/取模等细节。
type ringIterator struct {
	rb    *RingBuffer
	count int
}

func (it *ringIterator) HasNext() bool {
	return it.count < len(it.rb.data)
}

func (it *ringIterator) Next() int {
	idx := (it.rb.start + it.count) % len(it.rb.data)
	it.count++
	return it.rb.data[idx]
}

func main() {
	// 环形缓冲：物理顺序 [10,20,30,40]，但从下标2开始逻辑遍历
	rb := NewRingBuffer([]int{10, 20, 30, 40}, 2)

	// 客户端只用 Iterator 接口，不关心环形/取模等内部实现
	it := rb.CreateIterator()
	fmt.Print("遍历结果: ")
	for it.HasNext() {
		fmt.Printf("%d ", it.Next())
	}
	fmt.Println()

	fmt.Println("对比 Go 原生 range（仅作参考，顺序为物理顺序）:")
	for _, v := range rb.data {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
