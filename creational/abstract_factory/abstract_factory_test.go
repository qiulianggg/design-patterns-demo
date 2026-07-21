package main

import (
	"strings"
	"testing"
)

// 同一工厂产出的产品应属于同一风格族。
func TestWinFactoryFamily(t *testing.T) {
	f := WinFactory{}
	if !strings.Contains(f.CreateButton().Paint(), "Windows") {
		t.Error("Win 按钮风格错误")
	}
	if !strings.Contains(f.CreateCheckbox().Paint(), "Windows") {
		t.Error("Win 复选框风格错误")
	}
}

func TestMacFactoryFamily(t *testing.T) {
	f := MacFactory{}
	if !strings.Contains(f.CreateButton().Paint(), "macOS") {
		t.Error("Mac 按钮风格错误")
	}
	if !strings.Contains(f.CreateCheckbox().Paint(), "macOS") {
		t.Error("Mac 复选框风格错误")
	}
}

// 具体工厂应满足抽象工厂接口。
func TestFactoriesImplementInterface(t *testing.T) {
	var _ GUIFactory = WinFactory{}
	var _ GUIFactory = MacFactory{}
}
