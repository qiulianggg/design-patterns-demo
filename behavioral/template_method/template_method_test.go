package main

import "testing"

// 各实现的可变步骤行为正确。
func TestCSVProcessorSteps(t *testing.T) {
	p := CSVProcessor{}
	if p.Read() != "a,b,c" {
		t.Fatalf("CSV Read 错误: %q", p.Read())
	}
	if got := p.Process("a,b,c"); got != "a | b | c" {
		t.Fatalf("CSV Process 错误: %q", got)
	}
}

func TestJSONProcessorSteps(t *testing.T) {
	p := JSONProcessor{}
	if got := p.Process(`{"a":1}`); got != `{"A":1}` {
		t.Fatalf("JSON Process 应转大写: %q", got)
	}
}

// 模板方法骨架应能编排任意实现而不 panic。
func TestRunSkeleton(t *testing.T) {
	Run(CSVProcessor{})
	Run(JSONProcessor{})
}
