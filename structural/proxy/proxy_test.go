package main

import "testing"

// 虚拟代理：创建时不加载真实对象，首次使用才加载。
func TestProxyLazyLoad(t *testing.T) {
	p := NewProxyImage("photo.png")
	if p.real != nil {
		t.Fatal("代理创建时不应加载真实对象")
	}

	p.Display()
	if p.real == nil {
		t.Fatal("首次 Display 后应已加载真实对象")
	}

	// 再次调用应复用同一真实对象
	before := p.real
	p.Display()
	if p.real != before {
		t.Fatal("重复调用不应重新加载真实对象")
	}
}
