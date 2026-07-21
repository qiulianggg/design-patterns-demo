package main

import "testing"

func TestInterpretBoolean(t *testing.T) {
	ctx := &Context{vars: map[string]bool{"x": true, "y": false}}
	cases := []struct {
		tokens []string
		want   bool
	}{
		{[]string{"true", "AND", "false"}, false},
		{[]string{"false", "OR", "true"}, true},
		{[]string{"x", "AND", "y"}, false},
		{[]string{"x", "OR", "y"}, true},
	}
	for _, c := range cases {
		got := parse(c.tokens, ctx).Interpret(ctx)
		if got != c.want {
			t.Errorf("%v => %v, 期望 %v", c.tokens, got, c.want)
		}
	}
}

// 终结符表达式直接求值。
func TestTerminals(t *testing.T) {
	ctx := &Context{vars: map[string]bool{"v": true}}
	if !(Constant{true}).Interpret(ctx) {
		t.Error("Constant{true} 应为 true")
	}
	if !(Variable{"v"}).Interpret(ctx) {
		t.Error("Variable{v} 应取到 true")
	}
}
