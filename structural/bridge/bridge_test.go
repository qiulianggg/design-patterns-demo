package main

import (
	"strings"
	"testing"
)

func TestNormalMessageOverSMS(t *testing.T) {
	m := NormalMessage{Message{sender: SMSSender{}}}
	got := m.Notify("hi")
	if !strings.Contains(got, "[短信]") || !strings.Contains(got, "hi") {
		t.Fatalf("普通短信内容错误: %q", got)
	}
}

// 加急消息应加上【加急】标记，且可搭配任意渠道。
func TestUrgentMessageAddsMark(t *testing.T) {
	m := UrgentMessage{Message{sender: EmailSender{}}}
	got := m.Notify("down")
	if !strings.Contains(got, "【加急】") || !strings.Contains(got, "[邮件]") {
		t.Fatalf("加急邮件内容错误: %q", got)
	}
}
