package main

import "fmt"

// 中介者模式：用一个中介对象封装一组对象之间的交互，
// 使各对象不必显式互相引用，从而松散耦合。
// 本例：聊天室(中介者)协调多个用户的消息收发，用户之间互不直接持有引用。

// Mediator 是中介者接口。
type Mediator interface {
	Register(u *User)
	Send(from *User, msg string)
}

// ChatRoom 是具体中介者：统一转发消息。
type ChatRoom struct {
	users map[string]*User
}

func NewChatRoom() *ChatRoom { return &ChatRoom{users: make(map[string]*User)} }

func (c *ChatRoom) Register(u *User) {
	c.users[u.name] = u
	u.room = c
}

// Send 由中介者决定如何投递（这里广播给除发送者外的所有人）。
func (c *ChatRoom) Send(from *User, msg string) {
	for name, u := range c.users {
		if name != from.name {
			u.Receive(from.name, msg)
		}
	}
}

// User 是同事(Colleague)：只依赖中介者，不直接引用其他 User。
type User struct {
	name string
	room Mediator
}

func NewUser(name string) *User { return &User{name: name} }

func (u *User) Send(msg string) {
	fmt.Printf("[%s 发送] %s\n", u.name, msg)
	u.room.Send(u, msg)
}

func (u *User) Receive(from, msg string) {
	fmt.Printf("    [%s 收到] 来自 %s: %s\n", u.name, from, msg)
}

func main() {
	room := NewChatRoom()
	alice, bob, carol := NewUser("Alice"), NewUser("Bob"), NewUser("Carol")
	room.Register(alice)
	room.Register(bob)
	room.Register(carol)

	// Alice 发消息，无需知道 Bob/Carol 的存在，交给中介者转发
	alice.Send("大家好！")
	bob.Send("你好 Alice~")
}
