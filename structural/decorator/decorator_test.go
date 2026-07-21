package main

import "testing"

// 装饰器叠加：大写 + 感叹号。
func TestDecoratorsStack(t *testing.T) {
	var ds DataSource = PlainSource{}
	ds = NewUpperCase(ds)
	ds = NewExclaim(ds)

	if got := ds.Write("hi"); got != "HI!!!" {
		t.Fatalf("期望 HI!!!, 实际 %q", got)
	}
}

// 单个装饰器只做自己的事。
func TestSingleDecorator(t *testing.T) {
	if got := NewExclaim(PlainSource{}).Write("ok"); got != "ok!!!" {
		t.Fatalf("期望 ok!!!, 实际 %q", got)
	}
	if got := NewUpperCase(PlainSource{}).Write("ok"); got != "OK" {
		t.Fatalf("期望 OK, 实际 %q", got)
	}
}
