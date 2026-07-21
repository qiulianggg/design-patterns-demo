package main

import "testing"

// 注册后中介者应持有所有同事，且同事持有中介者引用。
func TestMediatorRegister(t *testing.T) {
	room := NewChatRoom()
	alice := NewUser("Alice")
	bob := NewUser("Bob")
	room.Register(alice)
	room.Register(bob)

	if len(room.users) != 2 {
		t.Fatalf("期望 2 个用户, 实际 %d", len(room.users))
	}
	if alice.room != room {
		t.Fatal("同事应持有中介者引用")
	}
}

// 通过中介者发送消息不应 panic（发送者不直接引用其他同事）。
func TestMediatorSend(t *testing.T) {
	room := NewChatRoom()
	a, b := NewUser("A"), NewUser("B")
	room.Register(a)
	room.Register(b)
	a.Send("hello") // 广播给 b，不应 panic
}
