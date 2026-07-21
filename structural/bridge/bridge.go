package main

import "fmt"

// 桥接模式：把「抽象」和「实现」拆成两个独立变化的维度，用组合连接。
// 本例两个维度：
//   维度1（抽象）：消息类型   —— 普通消息 / 加急消息
//   维度2（实现）：发送渠道   —— 短信 / 邮件
// 若用继承，会产生 2x2=4 个类；用桥接只需 2+2 个，且可自由组合。

// ---------- 实现维度：发送渠道 ----------

type MessageSender interface {
	Send(text string) string
}

type SMSSender struct{}

func (SMSSender) Send(text string) string { return "[短信] " + text }

type EmailSender struct{}

func (EmailSender) Send(text string) string { return "[邮件] " + text }

// ---------- 抽象维度：消息类型 ----------

// Message 持有一个 MessageSender（桥），把「怎么发」委托出去。
type Message struct {
	sender MessageSender
}

// NormalMessage 普通消息。
type NormalMessage struct {
	Message
}

func (m NormalMessage) Notify(content string) string {
	return m.sender.Send(content)
}

// UrgentMessage 加急消息：在内容上加标记，复用同样的 sender。
type UrgentMessage struct {
	Message
}

func (m UrgentMessage) Notify(content string) string {
	return m.sender.Send("【加急】" + content)
}

func main() {
	// 自由组合两个维度
	n := NormalMessage{Message{sender: SMSSender{}}}
	fmt.Println(n.Notify("您有一条系统通知"))

	u := UrgentMessage{Message{sender: EmailSender{}}}
	fmt.Println(u.Notify("服务器 CPU 超过阈值"))

	// 换一种渠道，抽象部分完全不用改
	u2 := UrgentMessage{Message{sender: SMSSender{}}}
	fmt.Println(u2.Notify("数据库连接失败"))
}
