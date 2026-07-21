package main

import (
	"strings"
	"testing"
)

func buildChain(rateQuota int) Handler {
	auth := &AuthHandler{}
	rate := &RateLimitHandler{remaining: rateQuota}
	biz := &BusinessHandler{}
	auth.SetNext(rate).SetNext(biz)
	return auth
}

// 正常请求应走到链尾并被处理。
func TestChainSuccess(t *testing.T) {
	r := &Request{User: "alice", Token: "abc"}
	got := buildChain(1).Handle(r)
	if !r.Handled || !strings.Contains(got, "业务处理完成") {
		t.Fatalf("正常请求应被处理, got=%q handled=%v", got, r.Handled)
	}
}

// 缺少 Token 应在认证环节中断。
func TestChainAuthInterrupt(t *testing.T) {
	r := &Request{User: "eve"}
	got := buildChain(1).Handle(r)
	if r.Handled || !strings.Contains(got, "认证失败") {
		t.Fatalf("无 Token 应被拦截, got=%q", got)
	}
}

// 配额为 0 应触发限流中断。
func TestChainRateLimit(t *testing.T) {
	got := buildChain(0).Handle(&Request{User: "bob", Token: "x"})
	if !strings.Contains(got, "限流") {
		t.Fatalf("应触发限流, got=%q", got)
	}
}
