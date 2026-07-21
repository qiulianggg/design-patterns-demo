package main

import "fmt"

// 责任链模式：把多个处理器串成一条链，请求沿链传递，
// 每个处理器决定「自己处理」还是「交给下一个」。
// 本例：一个 HTTP 请求依次经过 认证 -> 限流 -> 业务处理。

// Request 表示流经链条的请求。
type Request struct {
	User    string
	Token   string
	Handled bool
}

// Handler 是处理器接口。
type Handler interface {
	SetNext(Handler) Handler // 返回参数便于链式串联
	Handle(*Request) string
}

// baseHandler 提供「调用下一个」的公共逻辑，被各处理器内嵌复用。
type baseHandler struct {
	next Handler
}

func (b *baseHandler) SetNext(h Handler) Handler {
	b.next = h
	return h
}

// callNext 若存在下一个处理器则继续传递，否则结束。
func (b *baseHandler) callNext(r *Request) string {
	if b.next != nil {
		return b.next.Handle(r)
	}
	return "链路结束"
}

// AuthHandler 认证处理器。
type AuthHandler struct{ baseHandler }

func (h *AuthHandler) Handle(r *Request) string {
	if r.Token == "" {
		return "❌ 认证失败：缺少 Token（中断链路）"
	}
	fmt.Println("  ✓ 认证通过:", r.User)
	return h.callNext(r)
}

// RateLimitHandler 限流处理器。
type RateLimitHandler struct {
	baseHandler
	remaining int
}

func (h *RateLimitHandler) Handle(r *Request) string {
	if h.remaining <= 0 {
		return "❌ 限流：请求过多（中断链路）"
	}
	h.remaining--
	fmt.Printf("  ✓ 限流检查通过，剩余配额 %d\n", h.remaining)
	return h.callNext(r)
}

// BusinessHandler 业务处理器（链尾）。
type BusinessHandler struct{ baseHandler }

func (h *BusinessHandler) Handle(r *Request) string {
	r.Handled = true
	return "✅ 业务处理完成，返回数据给 " + r.User
}

func main() {
	// 组装链条：auth -> rateLimit -> business
	auth := &AuthHandler{}
	rate := &RateLimitHandler{remaining: 1}
	biz := &BusinessHandler{}
	auth.SetNext(rate).SetNext(biz)

	fmt.Println("请求1（正常）:")
	fmt.Println(" =>", auth.Handle(&Request{User: "alice", Token: "abc"}))

	fmt.Println("请求2（触发限流）:")
	fmt.Println(" =>", auth.Handle(&Request{User: "bob", Token: "xyz"}))

	fmt.Println("请求3（无 Token）:")
	fmt.Println(" =>", auth.Handle(&Request{User: "eve"}))
}
