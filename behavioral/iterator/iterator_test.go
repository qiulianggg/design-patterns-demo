package main

import (
	"reflect"
	"testing"
)

// 环形缓冲从 start 开始、环绕遍历，顺序应正确。
func TestRingIteratorOrder(t *testing.T) {
	rb := NewRingBuffer([]int{10, 20, 30, 40}, 2)
	var got []int
	it := rb.CreateIterator()
	for it.HasNext() {
		got = append(got, it.Next())
	}
	want := []int{30, 40, 10, 20}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("遍历顺序错误: got=%v want=%v", got, want)
	}
}

// 遍历元素个数应等于底层长度。
func TestRingIteratorCount(t *testing.T) {
	rb := NewRingBuffer([]int{1, 2, 3}, 0)
	n := 0
	for it := rb.CreateIterator(); it.HasNext(); it.Next() {
		n++
	}
	if n != 3 {
		t.Fatalf("期望遍历 3 个元素, 实际 %d", n)
	}
}
