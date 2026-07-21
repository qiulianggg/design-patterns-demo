package main

import (
	"strings"
	"testing"
)

// 子系统各自行为正确（外观正是编排这些调用）。
func TestSubsystems(t *testing.T) {
	if !strings.Contains(Projector{}.On(), "投影仪") {
		t.Error("投影仪开机文案错误")
	}
	if !strings.Contains(Amplifier{}.SetVolume(8), "8") {
		t.Error("音量设置错误")
	}
	if !strings.Contains(Player{}.Play("片名"), "片名") {
		t.Error("播放文案错误")
	}
}

// 外观方法应能顺利编排整套流程而不 panic。
func TestFacadeWatchMovie(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("WatchMovie 不应 panic: %v", r)
		}
	}()
	NewHomeTheater().WatchMovie("盗梦空间")
}
